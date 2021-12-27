package random

import (
	"math/rand"
	"time"
)

func init() {
	setSeed()
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())
}

// AVL
type RandomWeightTree struct {
	sumOfWeight int
	tree        *SplayTree
}

func NewRandomWeightTree() *RandomWeightTree {
	return &RandomWeightTree{tree: NewSplayTree(less, inRandomRange)}
}

func (t *RandomWeightTree) Insert(weight int, payload interface{}) {
	if weight <= 0 {
		weight = 1
	}
	node := &randomWeightNode{
		accumulateWeight: t.sumOfWeight,
		weight:           weight,
		payload:          payload,
	}
	t.tree.Add(node)
	t.sumOfWeight += weight
}

func (t *RandomWeightTree) GetRandom() interface{} {
	if t.sumOfWeight == 0 {
		return nil
	}
	randomWeight := rand.Intn(t.sumOfWeight)
	value := t.tree.Get(randomWeight)
	if value == nil {
		return nil
	}

	if node, ok := value.(*randomWeightNode); ok {
		return node.payload
	}
	return nil
}

func less(i1, i2 interface{}) bool {
	weightIdx := 0
	var weight int

	if gotWeight, ok := i1.(int); ok {
		weight = gotWeight
		weightIdx = 1
	}

	if gotWeight, ok := i2.(int); ok {
		weight = gotWeight
		weightIdx = 2
	}

	// both i1, i2 is node
	if weightIdx == 0 {
		return lessNode(i1, i2)
	}

	if node, ok := i1.(*randomWeightNode); ok {
		return node.accumulateWeight < weight
	}

	if node, ok := i2.(*randomWeightNode); ok {
		return weight < node.accumulateWeight
	}

	return false
}

func lessNode(n1, n2 interface{}) bool {
	return n1.(*randomWeightNode).accumulateWeight < n2.(*randomWeightNode).accumulateWeight
}

// i1 and i2 will be int or *randomWeightNode but not the same time
func inRandomRange(i1, i2 interface{}) bool {
	if weight, ok := i1.(int); ok {
		if node, ok := i2.(*randomWeightNode); ok {
			return node.accumulateWeight <= weight && weight < node.accumulateWeight+node.weight
		}
	}

	return false
}

// [accumulateWeight, accumulateWeight + weight)
type randomWeightNode struct {
	accumulateWeight int
	weight           int
	payload          interface{}
}
