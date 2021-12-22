package cf

import "github.com/cchaiyatad/seestern/internal/random"

type constraintRandomizer struct {
	trees map[string]*random.RandomWeightTree
}

func genConstraintRandomizer(fields []Field) *constraintRandomizer {
	randomizer := &constraintRandomizer{trees: make(map[string]*random.RandomWeightTree)}

	for _, field := range fields {
		tree := random.NewRandomWeightTree()
		for _, con := range field.Constraints {
			tree.Insert(con.Weight, con)
		}
		randomizer.trees[field.F_name] = tree
	}

	return randomizer
}

func (randomizer *constraintRandomizer) getConstraint(field string) *Constraint {
	if tree, ok := randomizer.trees[field]; ok {
		if con, ok := tree.GetRandom().(Constraint); ok {
			return &con
		}
		return nil
	}
	return nil
}
