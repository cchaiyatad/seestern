package alias

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/cchaiyatad/seestern/internal/dataformat"
	"github.com/cchaiyatad/seestern/internal/file"
)

type Alias map[string][]byte

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

	if err := isTomlFileValidAndHasAliase(filepath); err != nil {
		return nil, err
	}

	alias := make(Alias)
	if err := file.IterateLineFromFile(filepath, alias.getCreateAliasByLineFunc()); err != nil {
		return nil, err
	}

	return alias, nil
}

func (alias Alias) getCreateAliasByLineFunc() func(string) {
	parser := alias.newParser()
	return parser.parse
}

func getParseAliasFunc(filepath string) (dataformat.DecodeOption, error) {
	_, err := getAlias(filepath)
	if err != nil {
		return nil, err
	}

	parseFunc := func(data []byte) []byte {
		return data
	}

	return parseFunc, nil
}

func isTomlFileValidAndHasAliase(filepath string) error {
	type T_Alias struct {
		Key   string      `toml:"key"`
		Value interface{} `toml:"value"`
	}

	type T_Aliases struct {
		Aliases []T_Alias `toml:"alias"`
	}

	var aliases T_Aliases

	if _, err := toml.DecodeFile(filepath, &aliases); err != nil {
		return &ErrTomlFileIsInvalid{err}
	}

	if len(aliases.Aliases) == 0 {
		return ErrDoesnotHaveAlias
	}

	return nil
}
