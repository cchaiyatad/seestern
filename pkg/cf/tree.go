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
	Name string
	DataType
	Payload []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("name: %s type: %v payload: %s", n.Name, n.DataType, n.Payload)
}

type SchemaTree struct {
	Root *Node
}

func ParseSchemaTree(data map[string]interface{}) *SchemaTree {
	root := &Node{DataType: Object}

	for key, value := range data {
		if node := parse(key, reflect.ValueOf(value)); node != nil {
			root.Payload = append(root.Payload, node)
		}
	}

	return &SchemaTree{Root: root}
}

func parse(key string, value reflect.Value) *Node {

	if key == "_id" {
		return &Node{Name: key, DataType: ObjectID}
	}

	switch value.Kind() {
	case reflect.Invalid:
		return nil
	case reflect.Bool:
		return &Node{Name: key, DataType: Boolean}
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return &Node{Name: key, DataType: Integer}
	case reflect.Float32, reflect.Float64:
		return &Node{Name: key, DataType: Double}
	case reflect.String:
		return &Node{Name: key, DataType: String}
	case reflect.Array, reflect.Slice:
		node := &Node{Name: key, DataType: Array}
		// countKey := 0
		for i := 0; i < value.Len(); i++ {
			gotNode := parse("", value.Index(i))

			if gotNode == nil {
				continue
			}

			isDup := false
			for _, payloadNode := range node.Payload {
				// Deep Equal?
				if reflect.DeepEqual(gotNode, payloadNode) {
					isDup = true
					break
				}
			}

			if !isDup {
				node.Payload = append(node.Payload, gotNode)
			}

		}
		return node
	case reflect.Map:
		node := &Node{Name: key, DataType: Object}
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
