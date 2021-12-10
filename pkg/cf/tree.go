package cf

import (
	"fmt"
	"reflect"
)

type DataType int

const (
	Null DataType = iota
	String
	Integer
	Double
	Boolean
	ObjectID
	Array
	Object
)

type Node struct {
	Name      string
	NodeTypes []*NodeType
}

type NodeType struct {
	DataType
	Payload []*Node
}

var (
	// nullNode     *NodeType = &NodeType{DataType: Null}
	stringNode   *NodeType = &NodeType{DataType: String}
	integerNode  *NodeType = &NodeType{DataType: Integer}
	doubleNode   *NodeType = &NodeType{DataType: Double}
	booleanNode  *NodeType = &NodeType{DataType: Boolean}
	objectIDNode *NodeType = &NodeType{DataType: ObjectID}
)

func (n *Node) String() string {
	return fmt.Sprintf("name: %s type: %v", n.Name, n.NodeTypes)
}

func (n *NodeType) String() string {
	return fmt.Sprintf("type: %v payload: %s", n.DataType, n.Payload)
}

type SchemaTree struct {
	Root *Node
}

func ParseSchemaTree(data map[string]interface{}) *SchemaTree {
	root := &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object}}}

	for key, value := range data {
		if node := parse(key, reflect.ValueOf(value)); node != nil {
			root.NodeTypes[0].Payload = append(root.NodeTypes[0].Payload, node)
		}
	}

	return &SchemaTree{Root: root}
}

func parse(key string, value reflect.Value) *Node {
	if key == "_id" {
		return &Node{key, []*NodeType{objectIDNode}}
	}

	switch value.Kind() {
	case reflect.Invalid:
		return nil
	case reflect.Bool:
		return &Node{key, []*NodeType{booleanNode}}
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return &Node{key, []*NodeType{integerNode}}
	case reflect.Float32, reflect.Float64:
		return &Node{key, []*NodeType{doubleNode}}
	case reflect.String:
		return &Node{key, []*NodeType{stringNode}}

	case reflect.Array, reflect.Slice:
		node := &Node{Name: key, NodeTypes: []*NodeType{{DataType: Array}}}
		for i := 0; i < value.Len(); i++ {
			gotNode := parse("", value.Index(i))

			if gotNode == nil {
				continue
			}

			isDup := false
			for _, payloadNode := range node.NodeTypes[0].Payload {
				// Deep Equal?
				if reflect.DeepEqual(gotNode, payloadNode) {
					isDup = true
					break
				}
			}

			if !isDup {
				node.NodeTypes[0].Payload = append(node.NodeTypes[0].Payload, gotNode)
			}

		}
		return node
	case reflect.Map:
		node := &Node{Name: key, NodeTypes: []*NodeType{{DataType: Object}}}
		for _, key := range value.MapKeys() {
			if gotNode := parse(key.String(), value.MapIndex(key)); gotNode != nil {
				node.NodeTypes[0].Payload = append(node.NodeTypes[0].Payload, gotNode)
			}
		}
		return node
	case reflect.Struct:
		return nil
	case reflect.Interface:
		if value.IsNil() {
			return nil
		}
		return parse(key, value.Elem())
	default:
		return nil
	}
}
