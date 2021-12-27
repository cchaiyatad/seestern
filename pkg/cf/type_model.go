package cf

import (
	"fmt"
)

type Type struct {
	Type DataType `json:"type,omitempty" toml:"type,omitzero" yaml:"type,omitempty" mapstructure:"type,omitempty"`

	// only constraint
	P_Ref string `json:"ref,omitempty" toml:"ref,omitzero" yaml:"ref,omitempty" mapstructure:"ref,omitempty"`

	// string
	P_Prefix string `json:"prefix,omitempty" toml:"prefix,omitzero" yaml:"prefix,omitempty" mapstructure:"prefix,omitempty"`
	P_Suffix string `json:"suffix,omitempty" toml:"suffix,omitzero" yaml:"suffix,omitempty" mapstructure:"suffix,omitempty"`
	P_Length int    `json:"length,omitempty" toml:"length,omitzero" yaml:"length,omitempty" mapstructure:"length,omitempty"`

	// int, double
	P_Min interface{} `json:"min,omitempty" toml:"min,omitzero" yaml:"min,omitempty" mapstructure:"min,omitempty"`
	P_Max interface{} `json:"max,omitempty" toml:"max,omitzero" yaml:"max,omitempty" mapstructure:"max,omitempty"`

	// array
	P_Sets        []Set        `json:"sets,omitempty" toml:"sets,omitzero" yaml:"sets,omitempty" mapstructure:"sets,omitempty"`
	P_MaxItem     int          `json:"max_item,omitempty" toml:"max_item,omitzero" yaml:"max_item,omitempty" mapstructure:"max_item,omitempty"`
	P_MinItem     int          `json:"min_item,omitempty" toml:"min_item,omitzero" yaml:"min_item,omitempty" mapstructure:"min_item,omitempty"`
	P_ElementType []Constraint `json:"element_type,omitempty" toml:"element_type,omitzero" yaml:"element_type,omitempty" mapstructure:"element_type,omitempty"`

	// object
	P_Fields []Field `json:"fields,omitempty" toml:"fields,omitzero" yaml:"fields,omitempty" mapstructure:"fields,omitempty"`
}

func (t Type) String() string {
	return fmt.Sprintf("type: %s element_type: %s", t.Type, t.P_ElementType)
}

// TODO: refactor with generic
func (t Type) Ref() []string {
	if t.Type == Array {
		refs := make([]string, 0)
		for _, con := range t.ElementType() {
			refs = append(refs, con.Ref()...)
		}
		return refs
	}

	if t.Type == Object {
		refs := make([]string, 0)
		for _, field := range t.Fields() {
			for _, con := range field.Constraints {
				refs = append(refs, con.Ref()...)
			}
		}
		return refs
	}

	return []string{t.P_Ref}
}

func (t Type) Prefix() string {
	if t.Type != String {
		return ""
	}
	return t.P_Prefix
}

func (t Type) Suffix() string {
	if t.Type != String {
		return ""
	}
	return t.P_Suffix
}

func (t Type) Length() int {
	if t.Type != String {
		return 0
	}
	return t.P_Length
}

func (t Type) MinInt() int {
	if t.Type != Integer {
		return 0
	}

	if min, ok := t.P_Min.(int64); ok {
		return int(min)
	}
	return 0
}

func (t Type) MaxInt() int {
	if t.Type != Integer {
		return 0
	}

	if max, ok := t.P_Max.(int64); ok {
		return int(max)
	}
	return 0
}

func (t Type) MinDouble() float64 {
	if t.Type != Double {
		return 0
	}

	if min, ok := t.P_Min.(float64); ok {
		return min
	}
	return 0
}

func (t Type) MaxDouble() float64 {
	if t.Type != Double {
		return 0
	}

	if max, ok := t.P_Max.(float64); ok {
		return max
	}
	return 0
}

func (t Type) Sets() []Set {
	if t.Type != Array {
		return []Set{}
	}

	return t.P_Sets
}

func (t Type) MinItem() int {
	if t.Type != Array {
		return 0
	}

	if t.P_MinItem > 0 {
		return t.P_MinItem
	}
	return 0
}

func (t Type) MaxItem() int {
	if t.Type != Array {
		return 0
	}

	if t.P_MaxItem > 0 {
		return t.P_MaxItem
	}
	return 10
}

func (t Type) ElementType() []Constraint {
	if t.Type != Array {
		return []Constraint{}
	}

	return t.P_ElementType
}

func (t Type) Fields() []Field {
	if t.Type != Object {
		return []Field{}
	}
	return t.P_Fields
}
