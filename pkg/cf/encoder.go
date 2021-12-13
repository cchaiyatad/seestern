package cf

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type ConfigFileEncoder struct {
	enc Encoder
	Buf *bytes.Buffer
}

func (e *ConfigFileEncoder) Encode(v interface{}) error {
	return e.enc.Encode(v)
}

type Encoder interface {
	Encode(v interface{}) error
}

func NewEncoder(fileType string) *ConfigFileEncoder {
	buf := new(bytes.Buffer)

	switch strings.ToLower(fileType) {
	case "json":
		return &ConfigFileEncoder{enc: json.NewEncoder(buf), Buf: buf}
	case "toml":
		return &ConfigFileEncoder{enc: toml.NewEncoder(buf), Buf: buf}
	case "yaml":
		return &ConfigFileEncoder{enc: yaml.NewEncoder(buf), Buf: buf}
	default:
		return &ConfigFileEncoder{enc: yaml.NewEncoder(buf), Buf: buf}
	}
}
