package cf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type ConfigFileEncoder struct {
	enc Encoder
}

func (e *ConfigFileEncoder) Encode(v interface{}) error {
	return e.enc.Encode(v)
}

type Encoder interface {
	Encode(v interface{}) error
}

type ErrInvalidEncoderFileType struct {
	fileType string
}

func (e *ErrInvalidEncoderFileType) Error() string {
	return fmt.Sprintf("invalid encoder filetype : got %s", e.fileType)
}

func NewEncoder(fileType string) (*ConfigFileEncoder, error) {

	switch strings.ToLower(fileType) {
	case "json":
		return &ConfigFileEncoder{enc: json.NewEncoder(new(bytes.Buffer))}, nil
	case "yaml":
		return &ConfigFileEncoder{enc: yaml.NewEncoder(new(bytes.Buffer))}, nil
	case "toml":
		return &ConfigFileEncoder{enc: toml.NewEncoder(new(bytes.Buffer))}, nil
	default:
		return nil, &ErrInvalidEncoderFileType{fileType: fileType}
	}
}
