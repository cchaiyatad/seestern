package cf

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cchaiyatad/seestern/pkg/gen"
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
	// TODO: find order

	// TODO: iterate though order
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

func (*SSConfig) genDB(db *Database) ([]document, error) {
	dbName := db.D_name
	collName := db.Collection.C_name

	count := db.Collection.Count
	if count <= 0 {
		return nil, &ErrCollectionCountIsInvalid{dbName: dbName, collName: collName, count: count}
	}

	documents := make([]document, 0, count)
	for i := 0; i < count; i++ {
		doc := genDocument(i, db.Collection.Fields)
		documents = append(documents, doc)
	}

	return documents, nil
}

func genDocument(idx int, fields []Field) document {
	document := make(document)

	for _, field := range fields {
		// has set?

		if field.Omit_weight != 0 && shouldOmit(field.Omit_weight) {
			continue
		}

		// random from constraint
		constraint := getRandomConstraint(field.Constraints)
		value := getValueFromConstraint(constraint)
		document[field.F_name] = value
	}
	return document
}

func getRandomConstraint(constraints []Constraint) Constraint {
	// TODO: Use weight?
	return constraints[rand.Intn(len(constraints))]
}

func getValueFromConstraint(constraint Constraint) interface{} {
	if constraint.Value.Value != nil {
		return constraint.Value.Value
	}
	if constraint.Enum.Enum != nil && len(constraint.Enum.Enum) > 0 {
		return constraint.Enum.Enum[rand.Intn(len(constraint.Enum.Enum))]
	}
	return genType(constraint.Type, true)
}

func shouldOmit(omitWeight float64) bool {
	return rand.Float64() < omitWeight
}

func genType(t Type, isConstraint bool) interface{} {
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
		return genString(t)
	case Array:
		return genArray(t)
	case Object:
		return genObject(t)
	}
	return nil
}

func genNull() interface{} {
	return nil // is it work?
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

func genArray(t Type) interface{} {
	return false
}

func genObject(t Type) interface{} {
	return false
}
