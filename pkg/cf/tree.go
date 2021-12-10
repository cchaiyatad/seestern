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
		if node := parse(root, key, value); node != nil {
			root.Payload = append(root.Payload, node)
		}
	}

	return &SchemaTree{Root: root}
}

// Complex64
// Complex128
// Array
// Chan
// Func
// Interface
// Map
// Ptr
// Slice
// Struct
// UnsafePointer
func parse(parent *Node, key string, value interface{}) *Node {

	if key == "_id" {
		return &Node{Name: key, DataType: ObjectID}
	}

	reflectedValue := reflect.ValueOf(value)

	switch reflectedValue.Kind() {
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
	default:
		return nil

	}
}
