package cf

import (
	"fmt"
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

				got := ParseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

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

				got := ParseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

	t.Run("ParseSchemaTree for Object", func(t *testing.T) {
		cases := []struct {
			givenData map[string]interface{}
			expected  *SchemaTree
		}{
			{map[string]interface{}{"data": map[string]interface{}{}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "data", NodeTypes: []*NodeType{
								{DataType: Object},
							}},
						},
					},
				},
				},
			},
			},

			{map[string]interface{}{"data": map[string]interface{}{"name": "john"}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "data", NodeTypes: []*NodeType{
								{DataType: Object, Payload: []*Node{
									{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
								}},
							}},
						},
					},
				},
				},
			},
			},

			{map[string]interface{}{"data": map[string]interface{}{"name": "john", "surname": "son"}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "data", NodeTypes: []*NodeType{
								{DataType: Object, Payload: []*Node{
									{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
									{Name: "surname", NodeTypes: []*NodeType{{DataType: String}}},
								}},
							}},
						},
					},
				},
				},
			},
			},

			{map[string]interface{}{"data": map[string]interface{}{"name": "john", "age": 320}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "data", NodeTypes: []*NodeType{
								{DataType: Object, Payload: []*Node{
									{Name: "age", NodeTypes: []*NodeType{{DataType: Integer}}},
									{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
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

				got := ParseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

	t.Run("ParseSchemaTree for complex item", func(t *testing.T) {
		cases := []struct {
			givenData map[string]interface{}
			expected  *SchemaTree
		}{

			{map[string]interface{}{"data": map[string]interface{}{"arrayOfInt": []interface{}{1, 2, 3}, "arrayOfStr": []interface{}{"1", "123"}}}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "data", NodeTypes: []*NodeType{
								{DataType: Object, Payload: []*Node{
									{Name: "arrayOfInt", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{
										{Name: "", NodeTypes: []*NodeType{{DataType: Integer}}},
									}}}},
									{Name: "arrayOfStr", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{
										{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
									}}}},
								}},
							}},
						},
					},
				},
				},
			},
			},
			{map[string]interface{}{"data": map[string]interface{}{
				"arrayOfObj": []interface{}{
					map[string]interface{}{"name": "some name", "age": 123},
					map[string]interface{}{"product": "some name", "price": 123.00},
				}},
			}, &SchemaTree{
				Root: &Node{Name: "_root", NodeTypes: []*NodeType{
					{DataType: Object,
						Payload: []*Node{
							{Name: "data", NodeTypes: []*NodeType{
								{DataType: Object, Payload: []*Node{
									{Name: "arrayOfObj", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{
										{Name: "", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{
											{Name: "age", NodeTypes: []*NodeType{{DataType: Integer}}},
											{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
										}}}},
										{Name: "", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{
											{Name: "price", NodeTypes: []*NodeType{{DataType: Double}}},
											{Name: "product", NodeTypes: []*NodeType{{DataType: String}}},
										}}}},
									}}}},
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

				got := ParseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

}

func TestMergeTree(t *testing.T) {
	t.Run("MergeTree", func(t *testing.T) {
		cases := []struct {
			givenTreeOne *SchemaTree
			givenTreeTwo *SchemaTree
			expected     *SchemaTree
		}{
			{
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Integer},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Integer},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Integer},
								}},
							},
						},
					},
					},
				},
			},
			{
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Integer},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Double},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Integer}, {DataType: Double},
								}},
							},
						},
					},
					},
				},
			},
			{
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data0", NodeTypes: []*NodeType{
									{DataType: Integer},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data1", NodeTypes: []*NodeType{
									{DataType: Double},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data0", NodeTypes: []*NodeType{
									{DataType: Integer},
								}},
								{Name: "data1", NodeTypes: []*NodeType{
									{DataType: Double},
								}},
							},
						},
					},
					},
				},
			},
			{
				&SchemaTree{
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
				&SchemaTree{
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
				&SchemaTree{
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
			{
				&SchemaTree{
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
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "names", NodeTypes: []*NodeType{
									{DataType: Array, Payload: []*Node{
										{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "", NodeTypes: []*NodeType{{DataType: Integer}}},
									}},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "names", NodeTypes: []*NodeType{
									{DataType: Array, Payload: []*Node{
										{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
									}},
									{DataType: Array, Payload: []*Node{
										{Name: "", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "", NodeTypes: []*NodeType{{DataType: Integer}}},
									}},
								}},
							},
						},
					},
					},
				},
			},
			{
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "surname", NodeTypes: []*NodeType{{DataType: String}}},
									}},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "surname", NodeTypes: []*NodeType{{DataType: String}}},
									}},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "surname", NodeTypes: []*NodeType{{DataType: String}}},
									}},
								}},
							},
						},
					},
					},
				},
			},
			{
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "surname", NodeTypes: []*NodeType{{DataType: String}}},
									}},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
									}},
								}},
							},
						},
					},
					},
				},
				&SchemaTree{
					Root: &Node{Name: "_root", NodeTypes: []*NodeType{
						{DataType: Object,
							Payload: []*Node{
								{Name: "data", NodeTypes: []*NodeType{
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
										{Name: "surname", NodeTypes: []*NodeType{{DataType: String}}},
									}},
									{DataType: Object, Payload: []*Node{
										{Name: "name", NodeTypes: []*NodeType{{DataType: String}}},
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
			t.Run(fmt.Sprintf("MergeTree expected %v", tc.expected), func(t *testing.T) {
				t.Parallel()

				got, err := MergeSchemaTree(tc.givenTreeOne, tc.givenTreeTwo)
				assert.Nil(t, err)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))
			})
		}
	})

}
