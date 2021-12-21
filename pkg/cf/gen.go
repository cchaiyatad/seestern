package cf

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cchaiyatad/seestern/pkg/gen"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func init() {
	setSeed()
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())
}

func (ssconfig *SSConfig) Gen() result {
	info := ssconfig.NewResult()
	// dbcollInfo := ssconfig.GetdbcollInfo()
	// TODO 6: find order

	// TODO 6: iterate though order
	for _, db := range ssconfig.Databases {

		documents, err := ssconfig.genDB(&db)
		if err != nil {
			info[db.D_name][db.Collection.C_name] = err
		} else {
			info[db.D_name][db.Collection.C_name] = documents
		}
	}
	return info
}

func (ssconfig *SSConfig) genDB(db *Database) (documents, error) {
	dbName := db.D_name
	collName := db.Collection.C_name

	count := db.Collection.Count
	if count <= 0 {
		return nil, &ErrCollectionCountIsInvalid{dbName: dbName, collName: collName, count: count}
	}

	documents := make(documents, 0, count)
	for i := 0; i < count; i++ {
		doc := ssconfig.genDocument(i, db.Collection.Fields)
		documents = append(documents, doc)
	}

	return documents, nil
}

func (ssconfig *SSConfig) genDocument(idx int, fields []Field) document {
	document := make(document)

	for _, field := range fields {
		if value, ok := genFromSet(field.Sets, idx, ssconfig.vendor); ok {
			document[field.F_name] = value
			continue
		}

		if field.Omit_weight != 0 && shouldOmit(field.Omit_weight) {
			continue
		}

		document[field.F_name] = genFromConstraint(field.Constraints, ssconfig.vendor)
	}
	return document
}

func genFromSet(sets []Set, idx int, vendor string) (interface{}, bool) {
	for _, set := range sets {
		for _, at := range set.At {
			if at == idx {
				return getValueFromItem(set.Value.Value, set.Enum.Enum, set.Type, vendor), true
			}
		}
	}
	return nil, false
}

func genFromConstraint(constraints []Constraint, vendor string) interface{} {
	constraint := getRandomConstraint(constraints)
	return getValueFromItem(constraint.Value.Value, constraint.Enum.Enum, constraint.Type, vendor)
}

func getRandomConstraint(constraints []Constraint) Constraint {
	// TODO 1: Use weight?
	return constraints[rand.Intn(len(constraints))]
}

func getValueFromItem(value interface{}, enum []interface{}, t Type, vendor string) interface{} {
	if value != nil {
		return value
	}
	if len(enum) > 0 {
		return enum[rand.Intn(len(enum))]
	}
	return genType(t, vendor, true)
}

func shouldOmit(omitWeight float64) bool {
	return rand.Float64() < omitWeight
}

func genType(t Type, vendor string, isConstraint bool) interface{} {
	switch t.Type {
	case Null:
		return genNull()
	case String:
		return genString(t)
	case Integer:
		return genInt(t)
	case Double:
		return genDouble(t)
	case Boolean:
		return genBoolean(t)
	case ObjectID:
		return genObjectID(t, vendor)
	case Array:
		return genArray(t)
	case Object:
		return genObject(t)
	}
	return nil
}

func genNull() interface{} {
	return nil
}

func genString(t Type) interface{} {
	prefix := t.Prefix()
	suffix := t.Suffix()
	lenght := t.Length()

	if lenght == 0 {
		lenght = 20
	}

	return gen.GenString(lenght, prefix, suffix)
}

func genInt(t Type) interface{} {
	min := t.MinInt()
	max := t.MaxInt()

	return gen.GenInt(min, max)
}

func genDouble(t Type) interface{} {
	min := t.MinDouble()
	max := t.MaxDouble()

	return gen.GenDouble(min, max)
}

func genBoolean(t Type) interface{} {
	return gen.GenBoolean()
}

func genObjectID(t Type, vendor string) interface{} {
	if vendor == "mongo" {
		return primitive.NewObjectID()
	}
	return gen.GenString(20, "", "")
}

func genArray(t Type) interface{} {
	// TODO 3: Array?
	return false
}

func genObject(t Type) interface{} {
	// TODO 4: Object?
	return false
}
