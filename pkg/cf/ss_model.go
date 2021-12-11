package cf

import "fmt"

type SSConfig struct {
	Aliases   []*Alias    `json:"aliases" toml:"alias"`
	Databases []*Database `json:"databases" toml:"database"`
}

func (s *SSConfig) String() string {
	return fmt.Sprintf("aliases: %s databases: %s", s.Aliases, s.Databases)
}

type Alias struct {
	Key   string `json:"key" toml:"key"`
	Value string `json:"value" toml:"value"`
}

func (a *Alias) String() string {
	return fmt.Sprintf("key: %s value: %s", a.Key, a.Value)
}

type Database struct {
	D_name     string      `json:"d_name" toml:"d_name"`
	Collection *Collection `json:"collection" toml:"collection"`
}

func (d *Database) String() string {
	return fmt.Sprintf("d_name: %s collection: %s", d.D_name, d.Collection)
}

type Collection struct {
	C_name string   `json:"c_name" toml:"c_name"`
	Count  int      `json:"count" toml:"count,omitzero"`
	Fields []*Field `json:"fields" toml:"fields"`
}

func (c *Collection) String() string {
	return fmt.Sprintf("c_name: %s count: %d fields: %s", c.C_name, c.Count, c.Fields)
}

type Field struct {
	F_name      string        `json:"f_name" toml:"f_name"`
	Omit_weight float64       `json:"omit_weight" toml:"omit_weight,omitzero"`
	Constraints []*Constraint `json:"constraints" toml:"constraint"`
	Sets        []*Set        `json:"sets" toml:"set"`
}

func (f *Field) String() string {
	return fmt.Sprintf("f_name: %s constraints: %s sets: %s", f.F_name, f.Constraints, f.Sets)
}

type Constraint struct {
	Weight int         `json:"weight" toml:"weight,omitzero"`
	Value  string      `json:"value" toml:"value,omitzero"`
	Enum   string      `json:"enum" toml:"enum,omitzero"`
	Type   SS_DataType `json:"type" toml:"type,omitzero"`
	Params Params      `json:"params" toml:"param"`
}

func (c *Constraint) String() string {
	return fmt.Sprintf("constraint: value: %s enum: %s type: %s", c.Value, c.Enum, c.Type)
}

type Set struct {
	At     []int       `json:"at" toml:"at"`
	Value  string      `json:"value" toml:"value,omitzero"`
	Enum   string      `json:"enum" toml:"enum,omitzero"`
	Type   SS_DataType `json:"type" toml:"type,omitzero"`
	Params Params      `json:"params" toml:"param"`
}

func (s *Set) String() string {
	return fmt.Sprintf("set: value: %s enum: %s type: %s", s.Value, s.Enum, s.Type)
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
