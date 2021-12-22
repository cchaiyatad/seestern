package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomWeightTree(t *testing.T) {
	tree := NewRandomWeightTree()

	tree.Insert(1, "A")
	tree.Insert(2, "B")
	tree.Insert(3, "C")
	tree.Insert(4, "D")
	tree.Insert(4, "E")
	tree.Insert(5, "F")
	tree.Insert(0, "G")

	count := map[string]int{}

	for i := 0; i < 1000000; i++ {
		got := tree.GetRandom().(string)
		count[got]++
	}

	assert.InDelta(t, float64(count["A"])/1000000.0, 1.0/20.0, 0.05)
	assert.InDelta(t, float64(count["B"])/1000000.0, 2.0/20.0, 0.05)
	assert.InDelta(t, float64(count["C"])/1000000.0, 3.0/20.0, 0.05)
	assert.InDelta(t, float64(count["D"])/1000000.0, 4.0/20.0, 0.05)
	assert.InDelta(t, float64(count["E"])/1000000.0, 4.0/20.0, 0.05)
	assert.InDelta(t, float64(count["F"])/1000000.0, 5.0/20.0, 0.05)
	assert.InDelta(t, float64(count["G"])/1000000.0, 1.0/20.0, 0.05)
}
