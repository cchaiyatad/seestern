package cf

import "github.com/cchaiyatad/seestern/internal/random"

type fieldGenerator struct {
	fields        []Field
	constraintGen constraintRandomizer
	// setGen setGenerator
	vendor string
}

func newFieldGenerator(fields []Field, vendor string) *fieldGenerator {
	return &fieldGenerator{
		fields:        fields,
		constraintGen: genConstraintRandomizer(fields),
		vendor:        vendor,
	}
}

func (gen *fieldGenerator) getConstraint(field string) *Constraint {
	return gen.constraintGen.getConstraint(field)
}

func (gen *fieldGenerator) genFromSet(sets []Set, idx int) (interface{}, bool) {
	// Big O performance
	for _, set := range sets {
		for _, at := range set.At {
			if at == idx {
				return getValueFromItem(set.Value.Value, set.Enum.Enum, set.Type, gen.vendor), true
			}
		}
	}
	return nil, false
}

func (gen *fieldGenerator) genFromConstraint(fieldName string) interface{} {
	constraint := gen.getConstraint(fieldName)
	if constraint == nil {
		return nil
	}
	return getValueFromItem(constraint.Value.Value, constraint.Enum.Enum, constraint.Type, gen.vendor)
}

type constraintRandomizer map[string]*random.RandomWeightTree

func genConstraintRandomizer(fields []Field) constraintRandomizer {
	randomizer := make(map[string]*random.RandomWeightTree)

	for _, field := range fields {
		tree := random.NewRandomWeightTree()
		for _, con := range field.Constraints {
			tree.Insert(con.Weight, con)
		}
		randomizer[field.F_name] = tree
	}

	return randomizer
}

func (randomizer constraintRandomizer) getConstraint(field string) *Constraint {
	if tree, ok := randomizer[field]; ok {
		if con, ok := tree.GetRandom().(Constraint); ok {
			return &con
		}
		return nil
	}
	return nil
}
