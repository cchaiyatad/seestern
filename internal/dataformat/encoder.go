package dataformat

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Encoder struct {
	enc encoder
	Buf *bytes.Buffer
}

func (e *Encoder) Encode(v interface{}) error {
	return e.enc.Encode(v)
}

type encoder interface {
	Encode(v interface{}) error
}

func NewEncoder(fileType string) *Encoder {
	buf := new(bytes.Buffer)

	switch strings.ToLower(fileType) {
	case "json":
		return &Encoder{enc: json.NewEncoder(buf), Buf: buf}
	case "toml":
		return &Encoder{enc: toml.NewEncoder(buf), Buf: buf}
	case "yaml":
		return &Encoder{enc: yaml.NewEncoder(buf), Buf: buf}
	default:
		return &Encoder{enc: yaml.NewEncoder(buf), Buf: buf}
	}
}
