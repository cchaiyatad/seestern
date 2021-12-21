package cf

import "fmt"

type SSConfig struct {
	Databases []Database `json:"databases" toml:"databases" yaml:"databases"`

	vendor string
}

func (s SSConfig) String() string {
	return fmt.Sprintf("databases: %s", s.Databases)
}

type Database struct {
	D_name     string     `json:"d_name" toml:"d_name" yaml:"d_name"`
	Collection Collection `json:"collection" toml:"collection" yaml:"collection"`
}

func (d Database) String() string {
	return fmt.Sprintf("d_name: %s collection: %s", d.D_name, d.Collection)
}

type Collection struct {
	C_name string  `json:"c_name" toml:"c_name" yaml:"c_name"`
	Count  int     `json:"count,omitempty" toml:"count,omitzero" yaml:"count,omitempty"`
	Fields []Field `json:"fields" toml:"fields" yaml:"fields"`
}

func (c Collection) String() string {
	return fmt.Sprintf("c_name: %s count: %d fields: %s", c.C_name, c.Count, c.Fields)
}

type Field struct {
	F_name      string       `json:"f_name" toml:"f_name" yaml:"f_name"`
	Omit_weight float64      `json:"omit_weight,omitempty" toml:"omit_weight,omitzero" yaml:"omit_weight,omitempty"`
	Constraints []Constraint `json:"constraints" toml:"constraints" yaml:"constraints"`
	Sets        []Set        `json:"sets,omitempty" toml:"sets,omitzero" yaml:"sets,omitempty"`
}

func (f Field) String() string {
	return fmt.Sprintf("f_name: %s constraints: %s sets: %s", f.F_name, f.Constraints, f.Sets)
}

type Constraint struct {
	Weight int `json:"weight,omitempty" toml:"weight,omitzero" yaml:"weight,omitempty"`
	Item   `json:",omitempty" yaml:",inline,omitempty"`
}

func (c Constraint) String() string {
	return fmt.Sprintf("item: %s", c.Item)
}

type Set struct {
	At   []int `json:"at,omitempty" toml:"at,omitzero" yaml:"at,omitempty"`
	Item `json:",omitempty" toml:",omitzero"  yaml:",inline,omitempty"`
}

func (s Set) String() string {
	return fmt.Sprintf("item: %s", s.Item)
}

type Item struct {
	Value `json:",omitempty" toml:",omitzero" yaml:",inline,omitempty"`
	Enum  `json:",omitempty" toml:",omitzero" yaml:",inline,omitempty"`
	Type  `json:",omitempty" toml:",omitzero" yaml:",inline,omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf("value: %s: enum: %s type: %s", i.Value, i.Enum, i.Type)
}

type Value struct {
	Value interface{} `json:"value,omitempty" toml:"value,omitzero" yaml:"value,omitempty"`
}

func (v Value) String() string {
	return fmt.Sprintf("value: %s", v.Value)
}

type Enum struct {
	Enum []interface{} `json:"enum,omitempty" toml:"enum,omitzero" yaml:"enum,omitempty"`
}

func (e Enum) String() string {
	return fmt.Sprintf("enum: %s", e.Enum)
}
