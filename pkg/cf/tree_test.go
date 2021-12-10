package cf

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSchemaTree(t *testing.T) {
	t.Run("ParseSchemaTree", func(t *testing.T) {
		cases := []struct {
			givenData map[string]interface{}
			expected  *SchemaTree
		}{
			{map[string]interface{}{}, &SchemaTree{Root: &Node{DataType: Object}}},
			{map[string]interface{}{"height": 500}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "height", DataType: Integer}}}}},
			{map[string]interface{}{"pi": 3.14}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "pi", DataType: Double}}}}},
			{map[string]interface{}{"name": "john"}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "name", DataType: String}}}}},
			{map[string]interface{}{"_id": "123"}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "_id", DataType: ObjectID}}}}},

			{map[string]interface{}{"height": 500, "weight": 500}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "height", DataType: Integer}, {Name: "weight", DataType: Integer}}}}},
			{map[string]interface{}{"height": 500, "pi": 3.14}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "height", DataType: Integer}, {Name: "pi", DataType: Double}}}}},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("ParseSchemaTree on with %v", tc.givenData), func(t *testing.T) {
				t.Parallel()

				got := ParseSchemaTree(tc.givenData)
				assert.Equal(t, true, reflect.DeepEqual(tc.expected, got), fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))
			})
		}
	})

	t.Run("ParseSchemaTree Array", func(t *testing.T) {
		cases := []struct {
			givenData map[string]interface{}
			expected  *SchemaTree
		}{
			{map[string]interface{}{"names": []interface{}{}}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "names", DataType: Array}}}}},
			{map[string]interface{}{"names": []interface{}{"john"}}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "names", DataType: Array, Payload: []*Node{{Name: "", DataType: String}}}}}}},
			{map[string]interface{}{"names": []interface{}{"john", "johnson"}}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "names", DataType: Array, Payload: []*Node{{Name: "", DataType: String}}}}}}},
			{map[string]interface{}{"names": []interface{}{"john", 3.14}}, &SchemaTree{Root: &Node{DataType: Object, Payload: []*Node{{Name: "names", DataType: Array, Payload: []*Node{{Name: "", DataType: String}, {Name: "", DataType: Double}}}}}}},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("ParseSchemaTree on with %v", tc.givenData), func(t *testing.T) {
				t.Parallel()

				got := ParseSchemaTree(tc.givenData)
				assert.Equal(t, true, reflect.DeepEqual(tc.expected, got), fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))
			})
		}
	})
}
