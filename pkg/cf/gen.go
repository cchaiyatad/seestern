package cf

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cchaiyatad/seestern/internal/random"
)

type document map[string]interface{}

type ErrCollectionCountIsInvalid struct {
	dbName   string
	collName string
	count    int
}

func (e *ErrCollectionCountIsInvalid) Error() string {
	return fmt.Sprintf("count have to be more than zero got: %d (db: %s, coll: %s)", e.count, e.dbName, e.collName)
}

type fieldGenerator struct {
	fields        []Field
	constraintGen constraintRandomizer
	setGen        setGenerator
	vendor        string
}

type constraintRandomizer map[string]*constraintRandomTree
type constraintRandomTree struct {
	tree *random.RandomWeightTree
}

type setGenerator map[string]setMap
type setMap map[int]*Item

func init() {
	setSeed()
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())
}

func (ssconfig *SSConfig) Gen() result {
	result := ssconfig.NewResult()
	order := ssconfig.NewGenOrder()

	order.IterateDB(func(db *Database) {
		documents, err := ssconfig.genDB(db)
		if err != nil {
			result[db.D_name][db.Collection.C_name] = err
		} else {
			result[db.D_name][db.Collection.C_name] = documents
		}
	})
	return result
}

func (ssconfig *SSConfig) genDB(db *Database) (documents, error) {
	dbName := db.D_name
	collName := db.Collection.C_name

	count := db.Collection.Count
	if count <= 0 {
		return nil, &ErrCollectionCountIsInvalid{dbName: dbName, collName: collName, count: count}
	}

	documents := make(documents, 0, count)
	fieldGen := newFieldGenerator(db.Collection.Fields, ssconfig.vendor)

	for i := 0; i < count; i++ {
		doc := genDocument(i, fieldGen)
		documents = append(documents, doc)
	}

	return documents, nil
}

func genDocument(idx int, fieldGen *fieldGenerator) document {
	document := make(document)

	for _, field := range fieldGen.fields {
		if value, ok := fieldGen.genFromSet(field.F_name, idx); ok {
			document[field.F_name] = value
			continue
		}

		if field.Omit_weight != 0 && shouldOmit(field.Omit_weight) {
			continue
		}

		document[field.F_name] = fieldGen.genFromConstraint(field.F_name)
	}
	return document
}

func shouldOmit(omitWeight float64) bool {
	return rand.Float64() < omitWeight
}

func newFieldGenerator(fields []Field, vendor string) *fieldGenerator {
	return &fieldGenerator{
		fields:        fields,
		constraintGen: newConstraintRandomizer(fields),
		setGen:        newSetGenerator(fields),
		vendor:        vendor,
	}
}

func (gen *fieldGenerator) getItemFromConstraint(fieldName string) *Item {
	return gen.constraintGen.getItemFromConstraint(fieldName)
}

func (gen *fieldGenerator) genFromConstraint(fieldName string) interface{} {
	item := gen.getItemFromConstraint(fieldName)
	if item == nil {
		return nil
	}
	return getValueFromItemFromConstraint(item, gen.vendor)
}

func (gen *fieldGenerator) getItemFromSet(fieldName string, idx int) (*Item, bool) {
	return gen.setGen.getItemFromSet(fieldName, idx)
}

func (gen *fieldGenerator) genFromSet(fieldName string, idx int) (interface{}, bool) {
	if item, ok := gen.getItemFromSet(fieldName, idx); ok {
		return getValueFromItemFromSet(item, gen.vendor), true
	}
	return nil, false
}

func newConstraintRandomizer(fields []Field) constraintRandomizer {
	randomizer := make(constraintRandomizer)

	for _, field := range fields {
		randomizer[field.F_name] = NewConstraintRandomTree(field.Constraints)
	}

	return randomizer
}

func (randomizer constraintRandomizer) getItemFromConstraint(field string) *Item {
	if tree, ok := randomizer[field]; ok {
		return tree.getItem()
	}
	return nil
}

func NewConstraintRandomTree(constraints []Constraint) *constraintRandomTree {
	tree := random.NewRandomWeightTree()
	for _, con := range constraints {
		tree.Insert(con.Weight, con)
	}
	return &constraintRandomTree{tree: tree}
}

func (tree *constraintRandomTree) getItem() *Item {
	if con, ok := tree.tree.GetRandom().(Constraint); ok {
		return &Item{
			Value: con.Value,
			Enum:  con.Enum,
			Type:  con.Type,
		}
	}
	return nil
}

func newSetGenerator(fields []Field) setGenerator {
	generator := make(setGenerator)

	for _, field := range fields {
		generator[field.F_name] = newSetMap(field.Sets)
	}

	return generator
}

func (gen setGenerator) getItemFromSet(field string, idx int) (*Item, bool) {
	if set, ok := gen[field]; ok {
		return set.getItem(idx)
	}
	return nil, false
}

func newSetMap(sets []Set) setMap {
	setMap := make(setMap)
	for _, set := range sets {
		for _, at := range set.At {
			setMap[at] = &Item{
				Value: set.Value,
				Enum:  set.Enum,
				Type:  set.Type,
			}
		}
	}
	return setMap
}

func (setMap setMap) getItem(idx int) (*Item, bool) {
	if item, ok := setMap[idx]; ok {
		return item, true
	}
	return nil, false
}
