package cf

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/cchaiyatad/seestern/pkg/gen"
)

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

func (ssconfig *SSConfig) Gen() {
	for _, db := range ssconfig.Databases {
		// find order
		ssconfig.genDB(&db)
	}
}

func (*SSConfig) genDB(db *Database) error {
	dbName := db.D_name
	collName := db.Collection.C_name

	count := db.Collection.Count
	if count <= 0 {
		return &ErrCollectionCountIsInvalid{dbName: dbName, collName: collName, count: count}
	}

	for i := 0; i < count; i++ {
		// create document
		document := map[string]interface{}{}

		for _, field := range db.Collection.Fields {
			// has set?

			// should omit?

			// random from constraint
			constraint := getRandomConstraint(field.Constraints)
			value := getValueFromConstraint(constraint)
			document[field.F_name] = value
		}
		fmt.Println(document)
	}

	return nil
}

func getRandomConstraint(constraints []Constraint) Constraint {
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

	if max == min {
		max = math.MaxInt
	}

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
