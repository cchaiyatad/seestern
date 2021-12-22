package alias

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/cchaiyatad/seestern/internal/dataformat"
	"github.com/cchaiyatad/seestern/internal/file"
)

type Alias map[string]string

var ErrAliasNotSupport = errors.New("error: alias only support in toml")
var ErrDoesnotHaveAlias = errors.New("error: this file does not have alias")

type ErrTomlFileIsInvalid struct {
	reason error
}

func (err *ErrTomlFileIsInvalid) Error() string {
	return fmt.Sprintf("error: this .toml file is invalid reason: %s", err.reason)
}

func getAlias(filepath string) (Alias, error) {
	fileType, err := file.GetFileType(filepath)
	if err != nil {
		return nil, err
	}

	if fileType != "toml" {
		return nil, ErrAliasNotSupport
	}

	if err := isTomlFileValid(filepath); err != nil {
		return nil, &ErrTomlFileIsInvalid{err}
	}

	alias := make(Alias)
	if err := file.IterateLineFromFile(filepath, alias.getCreateAliasByLineFunc()); err != nil {
		return nil, err
	}

	return alias, nil
}

func (alias Alias) getCreateAliasByLineFunc() func(string) {
	parser := alias.newParser()
	return parser.getParseFunc()
}

func strip(s string) string {
	return strings.TrimSpace(s)
}

func getParseAliasFunc(filepath string) (dataformat.DecodeOption, error) {
	alias, err := getAlias(filepath)
	if err != nil {
		return nil, err
	}

	if len(alias) == 0 {
		return nil, ErrDoesnotHaveAlias
	}

	parseFunc := func(data []byte) []byte {
		return data
	}

	return parseFunc, nil

}

func isTomlFileValid(filepath string) error {
	var tmp interface{}
	_, err := toml.DecodeFile(filepath, &tmp)
	return err
}
