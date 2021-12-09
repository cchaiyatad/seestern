package cf

type Alias struct {
	Key   string `json:"key" toml:"key"`
	Value string `json:"value" toml:"value"`
}

type Database struct {
	D_name      string     `json:"d_name" toml:"d_name"`
	Collections Collection `json:"collections" toml:"collections"`
}

type Collection struct {
	C_name string  `json:"c_name" toml:"c_name"`
	Count  int     `json:"count" toml:"count"`
	Fileds []Field `json:"fileds" toml:"fileds"`
}

type Field struct {
	F_name      string       `json:"f_name" toml:"f_name"`
	Omit_weight float64      `json:"omit_weight" toml:"omit_weight"`
	Constraints []Constraint `json:"constraints" toml:"constraints"`
	Sets        []Set        `json:"sets" toml:"sets"`
}

type Constraint struct {
	Weight int `json:"weight" toml:"weight"`
}

type Set struct {
	At    []int  `json:"at" toml:"at"`
	Items []Item `json:"items" toml:"items"`
}

type Item struct {
	Value Value `json:"value" toml:"value"`
	Enum  Enum  `json:"enum" toml:"enum"`
	Type  Type  `json:"type" toml:"type"`
}

type Value struct {
	Value string `json:"value" toml:"value"`
}

type Enum struct {
	Enum string `json:"enum" toml:"enum"`
}

type Type struct {
	Type   string `json:"type" toml:"type"`
	Params `json:"params" toml:"params"`
}

type Params map[string]interface{}
