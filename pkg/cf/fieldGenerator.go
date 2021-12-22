package cf

import "github.com/cchaiyatad/seestern/internal/random"

type fieldGenerator struct {
	fields        []Field
	constraintGen constraintRandomizer
	setGen        setGenerator
	vendor        string
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
	return getValueFromItem(item, gen.vendor)
}

func (gen *fieldGenerator) getItemFromSet(fieldName string, idx int) (*Item, bool) {
	return gen.setGen.getItemFromSet(fieldName, idx)
}

func (gen *fieldGenerator) genFromSet(fieldName string, idx int) (interface{}, bool) {
	if item, ok := gen.getItemFromSet(fieldName, idx); ok {
		return getValueFromItem(item, gen.vendor), true
	}
	return nil, false
}

type constraintRandomizer map[string]*random.RandomWeightTree

func newConstraintRandomizer(fields []Field) constraintRandomizer {
	randomizer := make(constraintRandomizer)

	for _, field := range fields {
		tree := random.NewRandomWeightTree()
		for _, con := range field.Constraints {
			tree.Insert(con.Weight, con)
		}
		randomizer[field.F_name] = tree
	}

	return randomizer
}

func (randomizer constraintRandomizer) getItemFromConstraint(field string) *Item {
	if tree, ok := randomizer[field]; ok {
		if con, ok := tree.GetRandom().(Constraint); ok {
			return &Item{
				Value: con.Value,
				Enum:  con.Enum,
				Type:  con.Type,
			}
		}
		return nil
	}
	return nil
}

type setGenerator map[string]map[int]*Item

func newSetGenerator(fields []Field) setGenerator {
	generator := make(setGenerator)

	for _, field := range fields {
		setMap := make(map[int]*Item)

		for _, set := range field.Sets {
			for _, at := range set.At {
				setMap[at] = &Item{
					Value: set.Value,
					Enum:  set.Enum,
					Type:  set.Type,
				}
			}
		}

		generator[field.F_name] = setMap
	}

	return generator
}

func (gen setGenerator) getItemFromSet(field string, idx int) (*Item, bool) {
	if set, ok := gen[field]; ok {
		if item, ok := set[idx]; ok {
			return item, true
		}
		return nil, false
	}
	return nil, false
}
