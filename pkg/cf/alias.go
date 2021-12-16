package cf

import "fmt"

type Alias struct {
	Key   string `json:"key" toml:"key" yaml:"key"`
	Value string `json:"value" toml:"value" yaml:"value"`
}

func (a Alias) String() string {
	return fmt.Sprintf("key: %s value: %s", a.Key, a.Value)
}

// Aliases   []Alias    `json:"aliases,omitempty" toml:"aliases" yaml:"aliases,omitempty"`
