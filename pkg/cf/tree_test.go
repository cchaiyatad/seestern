package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestParseSchemaTree(t *testing.T) {
	t.Run("parseSchemaTree", func(t *testing.T) {
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
								{DataType: String}, // is string because "123" is not primitive.ObjectId
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
			t.Run(fmt.Sprintf("parseSchemaTree on with %v", tc.givenData), func(t *testing.T) {
				t.Parallel()

				got := parseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

	t.Run("parseSchemaTree for Array", func(t *testing.T) {
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
			t.Run(fmt.Sprintf("parseSchemaTree on with %v", tc.givenData), func(t *testing.T) {
				t.Parallel()

				got := parseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

	t.Run("parseSchemaTree for Object", func(t *testing.T) {
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
			t.Run(fmt.Sprintf("parseSchemaTree on with %v", tc.givenData), func(t *testing.T) {
				t.Parallel()

				got := parseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

	t.Run("parseSchemaTree for complex item", func(t *testing.T) {
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
			t.Run(fmt.Sprintf("parseSchemaTree on with %v", tc.givenData), func(t *testing.T) {
				t.Parallel()

				got := parseSchemaTree("", "", tc.givenData)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))

			})
		}
	})

	t.Run("parseSchemaTree for mongo", func(t *testing.T) {
		given := map[string]interface{}{"_id": primitive.ObjectID{0x5b, 0xd7, 0x61, 0xde, 0xae, 0x32, 0x3e, 0x45, 0xa9, 0x3c, 0xe3, 0x69}, "couponUsed": false, "customer": map[string]interface{}{"age": 45, "email": "fili@vu.so", "gender": "F", "satisfaction": 5}, "items": primitive.A{map[string]interface{}{"name": "backpack", "price": primitive.Decimal128{}, "quantity": 1, "tags": primitive.A{"school", "travel", "kids"}}, map[string]interface{}{"name": "printer paper", "price": primitive.Decimal128{}, "quantity": 3, "tags": primitive.A{"office", "stationary"}}}, "purchaseMethod": "Online", "saleDate": 1363803439463, "storeLocation": "Denver"}
		expected := &SchemaTree{
			Root: &Node{Name: "_root", NodeTypes: []*NodeType{
				{DataType: Object, Payload: []*Node{
					{Name: "_id", NodeTypes: []*NodeType{objectIDNode}},
					{Name: "couponUsed", NodeTypes: []*NodeType{booleanNode}},
					{Name: "customer", NodeTypes: []*NodeType{{DataType: Object,
						Payload: []*Node{
							{Name: "age", NodeTypes: []*NodeType{integerNode}},
							{Name: "email", NodeTypes: []*NodeType{stringNode}},
							{Name: "gender", NodeTypes: []*NodeType{stringNode}},
							{Name: "satisfaction", NodeTypes: []*NodeType{integerNode}},
						},
					}}},
					{Name: "items", NodeTypes: []*NodeType{{DataType: Array,
						Payload: []*Node{
							{NodeTypes: []*NodeType{{DataType: Object,
								Payload: []*Node{
									{Name: "name", NodeTypes: []*NodeType{stringNode}},
									{Name: "price", NodeTypes: []*NodeType{doubleNode}},
									{Name: "quantity", NodeTypes: []*NodeType{integerNode}},
									{Name: "tags", NodeTypes: []*NodeType{{DataType: Array,
										Payload: []*Node{
											{Name: "", NodeTypes: []*NodeType{stringNode}},
										},
									}}},
								},
							}}},
						},
					}}},
					{Name: "purchaseMethod", NodeTypes: []*NodeType{stringNode}},
					{Name: "saleDate", NodeTypes: []*NodeType{integerNode}},
					{Name: "storeLocation", NodeTypes: []*NodeType{stringNode}},
				}},
			}},
			Collection: "sales",
			Database:   "sample_supplies",
		}

		got := parseSchemaTree("sample_supplies", "sales", given)
		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})
}

func TestMergeSchemaTree(t *testing.T) {
	t.Run("MergeSchemaTree", func(t *testing.T) {
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

				got, err := mergeSchemaTree(tc.givenTreeOne, tc.givenTreeTwo)
				assert.Nil(t, err)
				assert.Equal(t, tc.expected, got, fmt.Sprintf("expected: %v\ngot: %v", tc.expected, got))
			})
		}
	})

}

func TestSchemaTreeToSSConfig(t *testing.T) {
	t.Parallel()
	t.Run("ToSSConfig for one name one field (string)", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{stringNode}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []Database{
			{D_name: "sample_supplies", Collection: Collection{C_name: "sales", Fields: []Field{
				{F_name: "storeLocation", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})

	t.Run("ToSSConfig for one name two field (string, int)", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{stringNode, integerNode}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []Database{
			{D_name: "sample_supplies", Collection: Collection{C_name: "sales", Fields: []Field{
				{F_name: "storeLocation", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}, {Item: Item{Type: Type{Type: Integer}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})
	t.Run("ToSSConfig for array", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{Name: "", NodeTypes: []*NodeType{stringNode}}}}}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []Database{
			{D_name: "sample_supplies", Collection: Collection{C_name: "sales", Fields: []Field{
				{F_name: "storeLocation", Constraints: []Constraint{{Item: Item{Type: Type{Type: Array, P_ElementType: []interface{}{Item{Type: Type{Type: String}}}}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})
	t.Run("ToSSConfig for Object", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "location", NodeTypes: []*NodeType{stringNode}}}}}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []Database{
			{D_name: "sample_supplies", Collection: Collection{C_name: "sales", Fields: []Field{
				{F_name: "storeLocation", Constraints: []Constraint{{Item: Item{Type: Type{Type: Object, P_Fields: []Field{
					{F_name: "location", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
				}}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})
	t.Run("ToSSConfig for whole object", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "_id", NodeTypes: []*NodeType{objectIDNode}}, {Name: "couponUsed", NodeTypes: []*NodeType{booleanNode}}, {Name: "customer", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "age", NodeTypes: []*NodeType{integerNode}}, {Name: "email", NodeTypes: []*NodeType{stringNode}}, {Name: "gender", NodeTypes: []*NodeType{stringNode}}, {Name: "satisfaction", NodeTypes: []*NodeType{integerNode}}}}}}, {Name: "items", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "name", NodeTypes: []*NodeType{stringNode}}, {Name: "price", NodeTypes: []*NodeType{doubleNode}}, {Name: "quantity", NodeTypes: []*NodeType{integerNode}}, {Name: "tags", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{Name: "", NodeTypes: []*NodeType{stringNode}}}}}}}}}}}}}}, {Name: "purchaseMethod", NodeTypes: []*NodeType{stringNode}}, {Name: "saleDate", NodeTypes: []*NodeType{integerNode}}, {Name: "storeLocation", NodeTypes: []*NodeType{stringNode}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []Database{
			{D_name: "sample_supplies", Collection: Collection{C_name: "sales", Fields: []Field{
				{F_name: "_id", Constraints: []Constraint{{Item: Item{Type: Type{Type: ObjectID}}}}},
				{F_name: "couponUsed", Constraints: []Constraint{{Item: Item{Type: Type{Type: Boolean}}}}},
				{F_name: "customer", Constraints: []Constraint{{Item: Item{Type: Type{Type: Object, P_Fields: []Field{
					{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: Integer}}}}},
					{F_name: "email", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
					{F_name: "gender", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
					{F_name: "satisfaction", Constraints: []Constraint{{Item: Item{Type: Type{Type: Integer}}}}},
				}}}}}},
				{F_name: "items", Constraints: []Constraint{{Item: Item{Type: Type{Type: Array, P_ElementType: []interface{}{
					Item{Type: Type{Type: Object, P_Fields: []Field{
						{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
						{F_name: "price", Constraints: []Constraint{{Item: Item{Type: Type{Type: Double}}}}},
						{F_name: "quantity", Constraints: []Constraint{{Item: Item{Type: Type{Type: Integer}}}}},
						{F_name: "tags", Constraints: []Constraint{{Item: Item{Type: Type{Type: Array, P_ElementType: []interface{}{Item{Type: Type{Type: String}}}}}}}},
					}}},
				}}}}}},
				{F_name: "purchaseMethod", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
				{F_name: "saleDate", Constraints: []Constraint{{Item: Item{Type: Type{Type: Integer}}}}},
				{F_name: "storeLocation", Constraints: []Constraint{{Item: Item{Type: Type{Type: String}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})
}
