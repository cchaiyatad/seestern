package cf

// var tree = &SchemaTree{Root: &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "_id", NodeTypes: []*NodeType{objectIDNode}}, {Name: "couponUsed", NodeTypes: []*NodeType{booleanNode}}, {Name: "customer", NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "age", NodeTypes: []*NodeType{integerNode}}, {Name: "email", NodeTypes: []*NodeType{stringNode}}, {Name: "gender", NodeTypes: []*NodeType{stringNode}}, {Name: "satisfaction", NodeTypes: []*NodeType{integerNode}}}}}}, {Name: "items", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{NodeTypes: []*NodeType{{DataType: Object, Payload: []*Node{{Name: "name", NodeTypes: []*NodeType{stringNode}}, {Name: "price", NodeTypes: []*NodeType{doubleNode}}, {Name: "quantity", NodeTypes: []*NodeType{integerNode}}, {Name: "tags", NodeTypes: []*NodeType{{DataType: Array, Payload: []*Node{{Name: "", NodeTypes: []*NodeType{stringNode}}}}}}}}}}}}}}, {Name: "purchaseMethod", NodeTypes: []*NodeType{stringNode}}, {Name: "saleDate", NodeTypes: []*NodeType{integerNode}}, {Name: "storeLocation", NodeTypes: []*NodeType{stringNode}}}}}}, Collection: "sales", Database: "sample_supplies"}

// &{Root:name: _root type: [type: 7 payload: [name: storeLocation type: [type: 1 payload: []] name: saleDate type: [type: 2 payload: []] name: purchaseMethod type: [type: 1 payload: []] name: items type: [type: 6 payload: [name:  type: [type: 7 payload: [name: name type: [type: 1 payload: []] name: quantity type: [type: 2 payload: []] name: tags type: [type: 6 payload: [name:  type: [type: 1 payload: []]]]]]]] name: customer type: [type: 7 payload: [name: age type: [type: 2 payload: []] name: email type: [type: 1 payload: []] name: gender type: [type: 1 payload: []] name: satisfaction type: [type: 2 payload: []]]] name: couponUsed type: [type: 4 payload: []] name: _id type: [type: 5 payload: []]]] Database:sample_supplies Collection:sales}

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

// map[_id:ObjectID("5bd761deae323e45a93ce369") couponUsed:false customer:map[age:45 email:fili@vu.so gender:F satisfaction:5] items:[map[name:backpack price:148.77 quantity:1 tags:[school travel kids]] map[name:printer paper price:42.74 quantity:3 tags:[office stationary]]] purchaseMethod:Online saleDate:1363803439463 storeLocation:Denver]
// map[string]interface {}{"_id":primitive.ObjectID{0x5b, 0xd7, 0x61, 0xde, 0xae, 0x32, 0x3e, 0x45, 0xa9, 0x3c, 0xe3, 0x69}, "couponUsed":false, "customer":map[string]interface {}{"age":45, "email":"fili@vu.so", "gender":"F", "satisfaction":5}, "items":primitive.A{map[string]interface {}{"name":"backpack", "price":primitive.Decimal128{h:0x303c000000000000, l:0x3a1d}, "quantity":1, "tags":primitive.A{"school", "travel", "kids"}}, map[string]interface {}{"name":"printer paper", "price":primitive.Decimal128{h:0x303c000000000000, l:0x10b2}, "quantity":3, "tags":primitive.A{"office", "stationary"}}}, "purchaseMethod":"Online", "saleDate":1363803439463, "storeLocation":"Denver"}

// &{Root:name: _root type: [type: 7 payload: [name:
// storeLocation type: [type: 1 payload: []]
// name: saleDate type: [type: 2 payload: []]
// name: purchaseMethod type: [type: 1 payload: []]
// name: items type: [type: 6 payload: [name:  type: [type: 7 payload: [name: name type: [type: 1 payload: []] name: price type: [type: 3 payload: []] name: quantity type: [type: 2 payload: []] name: tags type: [type: 6 payload: [name:  type: [type: 1 payload: []]]]]]]] name: customer type: [type: 7 payload: [name: age type: [type: 2 payload: []] name: email type: [type: 1 payload: []] name: gender type: [type: 1 payload: []] name: satisfaction type: [type: 2 payload: []]]] name: couponUsed type: [type: 4 payload: []] name: _id type: [type: 5 payload: []]]] Database:sample_supplies Collection:sales}
