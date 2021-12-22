package cf

import (
	"math/rand"

	"github.com/cchaiyatad/seestern/pkg/gen"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getValueFromItemFromSet(item *Item, vendor string) interface{} {
	return getValueFromItem(item, vendor, false)
}
func getValueFromItemFromConstraint(item *Item, vendor string) interface{} {
	return getValueFromItem(item, vendor, true)
}

func getValueFromItem(item *Item, vendor string, isConstraint bool) interface{} {
	value := item.Value.Value
	enum := item.Enum.Enum
	t := item.Type

	if value != nil {
		return value
	}
	if len(enum) > 0 {
		return enum[rand.Intn(len(enum))]
	}
	return genType(t, vendor, isConstraint)
}

// TODO: ref (using isConstraint value)
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
