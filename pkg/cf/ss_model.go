package cf

import "fmt"

type SSConfig struct {
	Aliases   []*Alias    `json:"aliases" toml:"aliases" yaml:"aliases"`
	Databases []*Database `json:"databases" toml:"databases" yaml:"databases"`
}

func (s *SSConfig) String() string {
	return fmt.Sprintf("aliases: %s databases: %s", s.Aliases, s.Databases)
}

type Alias struct {
	Key   string `json:"key" toml:"key" yaml:"key"`
	Value string `json:"value" toml:"value" yaml:"value"`
}

func (a *Alias) String() string {
	return fmt.Sprintf("key: %s value: %s", a.Key, a.Value)
}

type Database struct {
	D_name     string      `json:"d_name" toml:"d_name" yaml:"d_name"`
	Collection *Collection `json:"collection" toml:"collection" yaml:"collection"`
}

func (d *Database) String() string {
	return fmt.Sprintf("d_name: %s collection: %s", d.D_name, d.Collection)
}

type Collection struct {
	C_name string   `json:"c_name" toml:"c_name" yaml:"c_name"`
	Count  int      `json:"count,omitempty" toml:"count,omitzero" yaml:"count,omitzero"`
	Fields []*Field `json:"fields" toml:"fields" yaml:"fields"`
}

func (c *Collection) String() string {
	return fmt.Sprintf("c_name: %s count: %d fields: %s", c.C_name, c.Count, c.Fields)
}

type Field struct {
	F_name      string        `json:"f_name" toml:"f_name" yaml:"f_name"`
	Omit_weight float64       `json:"omit_weight,omitempty" toml:"omit_weight,omitzero" yaml:"omit_weight,omitzero"`
	Constraints []*Constraint `json:"constraints" toml:"constraints" yaml:"constraints"`
	Sets        []*Set        `json:"sets" toml:"sets" yaml:"sets"`
}

func (f *Field) String() string {
	return fmt.Sprintf("f_name: %s constraints: %s sets: %s", f.F_name, f.Constraints, f.Sets)
}

type Constraint struct {
	Weight int `json:"weight,omitempty" toml:"weight,omitzero" yaml:"weight,omitzero"`
	*Item
}

func (c *Constraint) String() string {
	return fmt.Sprintf("item: %s", c.Item)
}

type Set struct {
	At []int `json:"at" toml:"at" yaml:"at"`
	*Item
}

func (s *Set) String() string {
	return fmt.Sprintf("item: %s", s.Item)
}

type Item struct {
	*Value
	*Enum
	*Type
}

func (i *Item) String() string {
	return fmt.Sprintf("value: %s: enum: %s type: %s", i.Value, i.Enum, i.Type)
}

type Value struct {
	Value string `json:"value,omitempty" toml:"value,omitzero" yaml:"value,omitzero"`
}

func (v *Value) String() string {
	return fmt.Sprintf("value: %s", v.Value)
}

type Enum struct {
	Enum string `json:"enum,omitempty" toml:"enum,omitzero" yaml:"enum,omitzero"`
}

func (e *Enum) String() string {
	return fmt.Sprintf("enum: %s", e.Enum)
}

type Type struct {
	Type        SS_DataType   `json:"type,omitempty" toml:"type,omitzero" yaml:"type,omitzero"`
	ElementType []interface{} `json:"element_type,omitempty" toml:"element_type,omitzero" yaml:"element_type,omitzero"`
	Params      `json:"params,omitempty" toml:"params,omitzero" yaml:"params,omitzero"`
}

func (t *Type) String() string {
	return fmt.Sprintf("type: %s element_type: %s", t.Type, t.ElementType)
}

type Params map[string]interface{}

type SS_DataType string

const (
	SS_Null     SS_DataType = "null"
	SS_String   SS_DataType = "string"
	SS_Integer  SS_DataType = "integer"
	SS_Double   SS_DataType = "double"
	SS_Boolean  SS_DataType = "boolean"
	SS_ObjectID SS_DataType = "objectID"
	SS_Array    SS_DataType = "array"
	SS_Object   SS_DataType = "object"
)

func (d DataType) toSS_DataType() SS_DataType {
	switch d {
	case Null:
		return SS_Null
	case String:
		return SS_String
	case Integer:
		return SS_Integer
	case Double:
		return SS_Double
	case Boolean:
		return SS_Boolean
	case ObjectID:
		return SS_ObjectID
	case Array:
		return SS_Array
	case Object:
		return SS_Object
	default:
		return SS_Null
	}
}
