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
			{map[string]interface{}{}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object},
				},
				},
			},
			},
			{map[string]interface{}{"height": 500}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "height", NodeTypes: []*NodeType{
								{DataType: Integer},
							}},
						},
					},
				},
				},
			},
			},
			{map[string]interface{}{"pi": 3.14}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "pi", NodeTypes: []*NodeType{
								{DataType: Double},
							}},
						},
					},
				},
				},
			},
			},
			{map[string]interface{}{"name": "john"}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "name", NodeTypes: []*NodeType{
								{DataType: String},
							}},
						},
					},
				},
				},
			},
			},
			{map[string]interface{}{"_id": "123"}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "_id", NodeTypes: []*NodeType{
								{DataType: ObjectID},
							}},
						},
					},
				},
				},
			},
			},
			{map[string]interface{}{"height": 500, "weight": 30}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "height", NodeTypes: []*NodeType{
								{DataType: Integer},
							}},
							{Name: "weight", NodeTypes: []*NodeType{
								{DataType: Integer},
							}},
						},
					},
				},
				},
			},
			},
			{map[string]interface{}{"height": 500, "pi": 3.14}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "height", NodeTypes: []*NodeType{
								{DataType: Integer},
							}},
							{Name: "pi", NodeTypes: []*NodeType{
								{DataType: Double},
							}},
						},
					},
				},
				},
			},
			},
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

	t.Run("ParseSchemaTree for Array", func(t *testing.T) {
		cases := []struct {
			givenData map[string]interface{}
			expected  *SchemaTree
		}{
			{map[string]interface{}{"names": []interface{}{}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "names", NodeTypes: []*NodeType{
								{DataType: Array},
							}},
						},
					},
				},
				},
			},
			},

			{map[string]interface{}{"names": []interface{}{"john"}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "names", NodeTypes: []*NodeType{
								{DataType: Array, Payload: []*Node{
									{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
								}},
							}},
						},
					},
				},
				},
			},
			},

			{map[string]interface{}{"names": []interface{}{"john", "jane"}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "names", NodeTypes: []*NodeType{
								{DataType: Array, Payload: []*Node{
									{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
								}},
							}},
						},
					},
				},
				},
			},
			},

			{map[string]interface{}{"names": []interface{}{"john", 3.14}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "names", NodeTypes: []*NodeType{
								{DataType: Array, Payload: []*Node{
									{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
									{Name: "", NodeTypes: []*NodeType{{DataType: Double}}},
								}},
							}},
						},
					},
				},
				},
			},
			},
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
