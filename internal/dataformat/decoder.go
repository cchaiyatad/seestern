package dataformat

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/cchaiyatad/seestern/internal/file"
	"gopkg.in/yaml.v3"
)

var ErrFileTypeNotSupport = errors.New("error: not support file type (support only .json .toml .yaml)")

type Decoder struct {
	dec  unmarshaler
	data []byte
}

type unmarshaler interface {
	unmarshal([]byte, interface{}) error
}

type DecodeOption func([]byte) []byte

func (e *Decoder) Decode(v interface{}, opts ...DecodeOption) error {
	newData := e.data
	if len(opts) > 0 {
		newData = make([]byte, len(e.data))
		copy(newData, e.data)

		for _, opt := range opts {
			newData = opt(newData)
		}
	}

	return e.dec.unmarshal(newData, v)
}

func NewDecoder(filePath string) (*Decoder, error) {
	fileType, err := file.GetFileType(filePath)
	if err != nil {
		return nil, err
	}

	decoder := &Decoder{}

	switch strings.ToLower(fileType) {
	case "json":
		decoder.dec = &jsonUnmarshaler{}
	case "toml":
		decoder.dec = &tomlUnmarshaler{}
	case "yaml":
		decoder.dec = &yamlUnmarshaler{}
	default:
		return nil, ErrFileTypeNotSupport
	}

	data, err := file.GetBytesFromFile(filePath)
	if err != nil {
		return nil, err
	}
	decoder.data = data
	return decoder, nil
}

type jsonUnmarshaler struct{}

func (d *jsonUnmarshaler) unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type yamlUnmarshaler struct{}

func (d *yamlUnmarshaler) unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

type tomlUnmarshaler struct{}

func (d *tomlUnmarshaler) unmarshal(data []byte, v interface{}) error {
	_, err := toml.Decode(string(data), v)
	return err
}
