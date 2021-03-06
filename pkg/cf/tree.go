package cf

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataType string

const (
	Null     DataType = "null"
	String   DataType = "string"
	Integer  DataType = "integer"
	Double   DataType = "double"
	Boolean  DataType = "boolean"
	ObjectID DataType = "objectID"
	Array    DataType = "array"
	Object   DataType = "object"
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
	nullNode     *NodeType = &NodeType{DataType: Null}
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
	Root       *Node
	Database   string
	Collection string
}

func (t *SchemaTree) getRootPayload() []*Node {
	if t == nil || !t.isValid() {
		return nil
	}
	return t.Root.NodeTypes[0].Payload
}

func (t *SchemaTree) setRootPayload(payload []*Node) {
	if t == nil || !t.isValid() {
		return
	}
	t.Root.NodeTypes[0].Payload = payload
}

func newSchemaTree(dbName, collName string) *SchemaTree {
	root := &Node{Name: "_root", NodeTypes: []*NodeType{{DataType: Object}}}

	return &SchemaTree{Root: root, Database: dbName, Collection: collName}
}

func parseSchemaTree(dbName, collName string, data map[string]interface{}) *SchemaTree {
	tree := newSchemaTree(dbName, collName)
	keyList := getKeyList(reflect.ValueOf(data).MapKeys())

	for _, key := range keyList {
		if node := parse(key, reflect.ValueOf(data[key])); node != nil {
			tree.setRootPayload(append(tree.getRootPayload(), node))
		}
	}

	return tree
}

func parse(key string, value reflect.Value) *Node {
	// bson Null, primitive.Decimal128 and primitive.ObjectID (refactor to be a option param?)
	switch value.Interface().(type) {
	case primitive.Null:
		return &Node{key, []*NodeType{nullNode}}
	case primitive.ObjectID:
		return &Node{key, []*NodeType{objectIDNode}}
	case primitive.Decimal128:
		return &Node{key, []*NodeType{doubleNode}}
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

func mergeSchemaTree(t1, t2 *SchemaTree) (*SchemaTree, error) {
	if !(t1.isValid() && t2.isValid()) {
		return nil, ErrInvalidSchemaTree
	}

	if t1.Database != t2.Database && t1.Collection != t2.Collection {
		return nil, &ErrMergeDiffTree{t1, t2}
	}

	mergedTree := &SchemaTree{}
	if err := copier.Copy(mergedTree, t1); err != nil {
		return nil, err
	}

	mergedPayload := mergedTree.getRootPayload()
	t2Payloads := t2.getRootPayload()

	for _, t2Node := range t2Payloads {
		keyName := t2Node.Name
		hasKey := false
		keyIdx := -1

		for idx, searchedNode := range mergedPayload {
			if searchedNode.Name == keyName {
				hasKey = true
				keyIdx = idx
				break
			}
		}

		if hasKey {
			isDup := false
			for _, searchedNodeType := range mergedPayload[keyIdx].NodeTypes {
				if reflect.DeepEqual(searchedNodeType, t2Node.NodeTypes[0]) {
					isDup = true
					break
				}
			}

			if !isDup {
				mergedPayload[keyIdx].NodeTypes = append(mergedPayload[keyIdx].NodeTypes, t2Node.NodeTypes[0])
			}
		} else {
			mergedPayload = append(mergedPayload, t2Node)
		}

		mergedTree.setRootPayload(mergedPayload)
	}

	return mergedTree, nil
}

func (t *SchemaTree) ToSSConfig() *SSConfig {
	if t == nil || !t.isValid() {
		return nil
	}

	coll := Collection{
		C_name: t.Collection,
	}

	for _, node := range t.getRootPayload() {
		field := node.toField()
		coll.Fields = append(coll.Fields, field)
	}

	db := Database{
		D_name:     t.Database,
		Collection: coll,
	}

	return &SSConfig{
		Databases: []Database{db},
	}
}

func (n *Node) toField() Field {
	if n == nil {
		return Field{}
	}

	field := Field{
		F_name: n.Name,
	}

	for _, nodeType := range n.NodeTypes {
		field.Constraints = append(field.Constraints, nodeType.toConstraint())
	}

	return field
}

func (n *NodeType) toConstraint() Constraint {
	if n == nil {
		return Constraint{}
	}

	item := Item{Type: Type{Type: n.DataType}}

	// Param for array and obj
	switch n.DataType {
	case Array:
		for _, payload := range n.Payload {
			for _, payloadNodeType := range payload.NodeTypes {
				con := payloadNodeType.toConstraint()
				item.P_ElementType = append(item.P_ElementType, con)
			}
		}
	case Object:
		for _, payload := range n.Payload {
			field := payload.toField()
			item.P_Fields = append(item.P_Fields, field)
		}
	}

	return Constraint{Item: item}
}

func getKeyList(keys []reflect.Value) []string {
	keyList := make([]string, 0, len(keys))
	for _, key := range keys {
		keyList = append(keyList, key.Interface().(string))
	}
	sort.Strings(keyList)
	return keyList
}
