package cf

import (
	"math/rand"

	"github.com/cchaiyatad/seestern/pkg/gen"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getValueFromItem(item *Item, fieldGen *fieldGenerator) interface{} {
	value := item.Value.Value
	enum := item.Enum.Enum
	t := item.Type

	if value != nil {
		return value
	}
	if len(enum) > 0 {
		return enum[rand.Intn(len(enum))]
	}
	return genType(t, fieldGen)
}

func genType(t Type, gen *fieldGenerator) interface{} {
	if shouldGenFromRef(t) {
		refs := t.Ref()
		return genFromRef(gen, refs)
	}

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
		return genObjectID(t, gen)
	case Array:
		return genArray(t, gen)
	case Object:
		return genObject(t, gen)
	}
	return nil
}
func shouldGenFromRef(t Type) bool {
	return t.P_Ref != ""
}

func genFromRef(gen *fieldGenerator, refs []string) interface{} {
	ref := refs[rand.Intn(len(refs))]
	nym, nymRef, ok := SplitRef(ref)
	if !ok {
		return nil
	}

	db, coll, ok := SplitNym(nym)
	if !ok {
		return nil
	}

	docs, err := gen.result.GetDocuments(db, coll)
	if err != nil {
		return nil
	}

	json, err := docs.ToJson()
	if err != nil {
		return nil
	}

	nymRef = "#." + nymRef
	refResult := gjson.Get(string(json), nymRef).Array()
	refValue := refResult[rand.Intn(len(refResult))]
	return refValue.Value()
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

func genObjectID(_ Type, fieldGen *fieldGenerator) interface{} {
	if fieldGen.vendor == "mongo" {
		return primitive.NewObjectID()
	}
	return gen.GenString(20, "", "")
}

func genArray(t Type, fieldGen *fieldGenerator) interface{} {
	data := []interface{}{}
	minItem := t.MinItem()
	maxItem := t.MaxItem()

	if minItem > maxItem {
		return data
	}

	constraintRandomTree := NewConstraintRandomTree(t.ElementType())
	setMap := newSetMap(t.Sets())

	maxItem = rand.Intn(maxItem-minItem) + minItem

	for i := minItem; i < maxItem; i++ {
		if item, ok := setMap.getItem(i); ok {
			value := getValueFromItem(item, fieldGen)
			data = append(data, value)
		}

		item := constraintRandomTree.getItem()
		value := getValueFromItem(item, fieldGen)
		data = append(data, value)
	}
	return data
}

func genObject(t Type, fieldGen *fieldGenerator) interface{} {
	fields := t.Fields()

	newFieldGen := newFieldGenerator(fields, fieldGen.vendor, fieldGen.result)
	document := genDocument(0, newFieldGen)

	return document
}
