package cf

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type ConfigFileDecoder struct {
	enc Decoder
	Buf *bytes.Buffer
}

func (e *ConfigFileDecoder) Encode(v interface{}) error {
	return e.enc.Encode(v)
}

type Decoder interface {
	Encode(v interface{}) error
}

func NewDecoder(fileType string) *ConfigFileDecoder {
	buf := new(bytes.Buffer)

	switch strings.ToLower(fileType) {
	case "json":
		return &ConfigFileDecoder{enc: json.NewEncoder(buf), Buf: buf}
	case "toml":
		return &ConfigFileDecoder{enc: toml.NewEncoder(buf), Buf: buf}
	case "yaml":
		return &ConfigFileDecoder{enc: yaml.NewEncoder(buf), Buf: buf}
	default:
		return &ConfigFileDecoder{enc: yaml.NewEncoder(buf), Buf: buf}
	}
}
