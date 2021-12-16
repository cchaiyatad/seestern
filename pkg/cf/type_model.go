package cf

import "fmt"

type Type struct {
	Type DataType `json:"type,omitempty" toml:"type,omitzero" yaml:"type,omitempty"`

	// array and object
	P_ElementType []interface{} `json:"element_type,omitempty" toml:"element_type,omitzero" yaml:"element_type,omitempty"`

	// only constraint
	P_Ref string `json:"ref,omitempty" toml:"ref,omitzero" yaml:"ref,omitempty"`

	// string
	P_Prefix string `json:"prefix,omitempty" toml:"prefix,omitzero" yaml:"prefix,omitempty"`
	P_Suffix string `json:"suffix,omitempty" toml:"suffix,omitzero" yaml:"suffix,omitempty"`
	P_Length int    `json:"length,omitempty" toml:"length,omitzero" yaml:"length,omitempty"`

	// int, double
	P_Min interface{} `json:"min,omitempty" toml:"min,omitzero" yaml:"min,omitempty"`
	P_Max interface{} `json:"max,omitempty" toml:"max,omitzero" yaml:"max,omitempty"`

	// array
	P_Sets []Set `json:"sets,omitempty" toml:"sets,omitzero" yaml:"sets,omitempty"`
}

func (t Type) String() string {
	return fmt.Sprintf("type: %s element_type: %s", t.Type, t.P_ElementType)
}

// TODO: refactor with generic
func (t Type) Ref() string {
	return t.P_Ref
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

	if min, ok := t.P_Min.(int); ok {
		return min
	}
	return 0
}

func (t Type) MaxInt() int {
	if t.Type != Integer {
		return 0
	}

	if max, ok := t.P_Max.(int); ok {
		return max
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

func (t Type) ElementTypeArray() []Constraint {
	constraints := []Constraint{}
	if t.Type != Array {
		return constraints
	}

	for _, value := range t.P_ElementType {
		if constraint, ok := value.(Constraint); ok {
			constraints = append(constraints, constraint)
		}
	}

	return constraints
}

func (t Type) ElementTypeObject() []Field {
	fields := []Field{}
	if t.Type != Object {
		return fields
	}

	for _, value := range t.P_ElementType {
		if field, ok := value.(Field); ok {
			fields = append(fields, field)
		}
	}

	return fields
}
