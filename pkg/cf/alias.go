package cf

import (
	"errors"
	"fmt"

	"github.com/cchaiyatad/seestern/internal/dataformat"
	"github.com/cchaiyatad/seestern/internal/file"
)

type Alias struct {
	Key   string      `json:"key" toml:"key" yaml:"key"`
	Value interface{} `json:"value" toml:"value" yaml:"value"`
}

func (a Alias) String() string {
	return fmt.Sprintf("key: %s value: %s", a.Key, a.Value)
}

var ErrAliasNotSupport = errors.New("error: alias only support in toml")
var ErrDoesnotHaveAlias = errors.New("error: this file does not have alias")

func getAlias(filepath string) ([]Alias, error) {
	var aliases []Alias

	fileType, err := file.GetFileType(filepath)
	if err != nil {
		return aliases, err
	}

	if fileType != "toml" {
		return aliases, ErrAliasNotSupport
	}

	decoder, err := dataformat.NewDecoder(filepath)
	if err != nil {
		return aliases, err
	}

	if err = decoder.Decode(&aliases); err != nil {
		return aliases, err
	}

	return aliases, nil
}

func getParseAliasFunc(filepath string) (dataformat.DecodeOption, error) {
	aliases, err := getAlias(filepath)
	if err != nil {
		return nil, err
	}

	if len(aliases) == 0 {
		return nil, ErrDoesnotHaveAlias
	}

	parseFunc := func(data []byte) []byte {
		return data
	}

	return parseFunc, nil

}
