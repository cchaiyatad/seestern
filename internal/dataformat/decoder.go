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

func (e *Decoder) Decode(v interface{}, opts ...func([]byte) []byte) error {
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
		decoder.dec = &yamlUnmarshaler{}
	case "yaml":
		decoder.dec = &tomlUnmarshaler{}
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

// t.Run("test GetBytesFromFile with configuration file", func(t *testing.T) {
// 	givenPath := "./../../test/config/04_configSpec_alias.ss.toml"

// 	gotByte, _ := GetBytesFromFile(givenPath)
// 	var ssConfig cf.SSConfig

// 	// f, _ := os.Open(givenPath)
// 	// dec := toml.NewDecoder(f)

// 	// _, _ = dec.Decode(&ssConfig)
// 	_, err := toml.Decode(string(gotByte), &ssConfig)

// 	fmt.Printf("%s\n", err)
// 	fmt.Printf("%#v\n", ssConfig)

// })
