package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaTreeToSSConfig(t *testing.T) {
	t.Parallel()
	t.Run("ToSSConfig for one name one field (string)", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{stringNode}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []*Database{
			{D_name: "sample_supplies", Collection: &Collection{C_name: "sales", Fields: []*Field{
				{F_name: "storeLocation", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_String}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})

	t.Run("ToSSConfig for one name two field (string, int)", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{stringNode, integerNode}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []*Database{
			{D_name: "sample_supplies", Collection: &Collection{C_name: "sales", Fields: []*Field{
				{F_name: "storeLocation", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_String}}}, {Item: &Item{Type: &Type{Type: SS_Integer}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})
	t.Run("ToSSConfig for array", func(t *testing.T) {
		givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "storeLocation", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{Name: "", NodeTypes: []*NodeType{stringNode}}}}}}}}}}, Collection: "sales", Database: "sample_supplies"}
		expected := &SSConfig{Databases: []*Database{
			{D_name: "sample_supplies", Collection: &Collection{C_name: "sales", Fields: []*Field{
				{F_name: "storeLocation", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_Array}}}}},
			}}},
		}}
		got := givenTree.ToSSConfig()

		assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	})

	// t.Run("ToSSConfig for whole object", func(t *testing.T) {
	// 	givenTree := &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "_id", NodeTypes: []*NodeType{objectIDNode}}, {Name: "couponUsed", NodeTypes: []*NodeType{booleanNode}}, {Name: "customer", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "age", NodeTypes: []*NodeType{integerNode}}, {Name: "email", NodeTypes: []*NodeType{stringNode}}, {Name: "gender", NodeTypes: []*NodeType{stringNode}}, {Name: "satisfaction", NodeTypes: []*NodeType{integerNode}}}}}}, {Name: "items", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "name", NodeTypes: []*NodeType{stringNode}}, {Name: "price", NodeTypes: []*NodeType{doubleNode}}, {Name: "quantity", NodeTypes: []*NodeType{integerNode}}, {Name: "tags", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{Name: "", NodeTypes: []*NodeType{stringNode}}}}}}}}}}}}}}, {Name: "purchaseMethod", NodeTypes: []*NodeType{stringNode}}, {Name: "saleDate", NodeTypes: []*NodeType{integerNode}}, {Name: "storeLocation", NodeTypes: []*NodeType{stringNode}}}}}}, Collection: "sales", Database: "sample_supplies"}
	// 	expected := &SSConfig{Databases: []*Database{
	// 		{D_name: "sample_supplies", Collection: &Collection{C_name: "sales", Fields: []*Field{
	// 			{F_name: "_id", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_ObjectID}}}}},
	// 			{F_name: "couponUsed", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_Boolean}}}}},
	// 			{F_name: "customer", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_Object}}}}},
	// 			{F_name: "items", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_Array}}}}},
	// 			{F_name: "purchaseMethod", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_String}}}}},
	// 			{F_name: "saleDate", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_Integer}}}}},
	// 			{F_name: "storeLocation", Constraints: []*Constraint{{Item: &Item{Type: &Type{Type: SS_String}}}}},
	// 		}}},
	// 	}}
	// 	got := givenTree.ToSSConfig()

	// 	assert.Equal(t, expected, got, fmt.Sprintf("expected: %v\ngot: %v", expected, got))
	// })

}

// map[_id:ObjectID("5bd761deae323e45a93ce364")
// couponUsed:false
// customer:map[age:48 email:laah@lawef.ae gender:F satisfaction:5]
// items:[
// 	map[name:pens price:59.95 quantity:2 tags:[writing office school stationary]]
// 	map[name:binder price:23.5 quantity:5 tags:[school general organization]]
// 	map[name:laptop price:635.41 quantity:4 tags:[electronics school office]]
// 	map[name:printer paper price:37.03 quantity:5 tags:[office stationary]]
// 	map[name:backpack price:96.6 quantity:5 tags:[school travel kids]]
// 	map[name:binder price:28.79 quantity:1 tags:[school general organization]]]
// purchaseMethod:In store
// saleDate:1485894409273
// storeLocation:Denver]
