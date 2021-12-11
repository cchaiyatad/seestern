package cf

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/jinzhu/copier"
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

	keyList := getKeyList(reflect.ValueOf(data).MapKeys())

	for _, key := range keyList {
		if node := parse(key, reflect.ValueOf(data[key])); node != nil {
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
		keyList := getKeyList(value.MapKeys())

		for _, key := range keyList {
			if gotNode := parse(key, value.MapIndex(reflect.ValueOf(key))); gotNode != nil {
				node.NodeTypes[0].Payload = append(node.NodeTypes[0].Payload, gotNode)
			}
		}

		return node
	case reflect.Interface:
		if value.IsNil() {
			return nil
		}
		return parse(key, value.Elem())
	default:
		return nil
	}
}

func (t *SchemaTree) isValid() bool {
	return t.Root.Name == "_root" && t.Root.NodeTypes[0].DataType == Object
}

func MergeSchemaTree(t1, t2 *SchemaTree) (*SchemaTree, error) {
	if !(t1.isValid() && t2.isValid()) {
		return nil, ErrInvalidSchemaTree
	}

	mergedTree := &SchemaTree{}
	copier.Copy(mergedTree, t1)

	mergedPayloads := mergedTree.Root.NodeTypes[0].Payload
	t2Payloads := t2.Root.NodeTypes[0].Payload

	for _, node := range t2Payloads {
		keyName := node.Name
		hasKey := false
		keyIdx := -1

		for idx, searchedNode := range mergedPayloads {
			if searchedNode.Name == keyName {
				hasKey = true
				keyIdx = idx
				break
			}
		}

		if hasKey {
			isDup := false
			for _, searchedNodeType := range mergedPayloads[keyIdx].NodeTypes {
				if reflect.DeepEqual(searchedNodeType, node.NodeTypes[0]) { //[0] ??
					isDup = true
					break
				}
			}

			if !isDup {
				mergedTree.Root.NodeTypes[0].Payload[keyIdx].NodeTypes = append(mergedPayloads[keyIdx].NodeTypes, node.NodeTypes[0]) //[0] ??
			}
		} else {
			mergedTree.Root.NodeTypes[0].Payload = append(mergedPayloads, node)
		}

	}

	return mergedTree, nil
}

func getKeyList(keys []reflect.Value) []string {
	keyList := make([]string, 0, len(keys))
	for _, key := range keys {
		keyList = append(keyList, key.Interface().(string))
	}
	sort.Strings(keyList)
	return keyList
}
