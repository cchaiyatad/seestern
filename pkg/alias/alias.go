package alias

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/cchaiyatad/seestern/internal/dataformat"
	"github.com/cchaiyatad/seestern/internal/file"
)

type Alias struct {
	dict  map[string][]byte
	order []string
}

var ErrAliasNotSupport = errors.New("error: alias only support in toml")
var ErrDoesnotHaveAlias = errors.New("error: this file does not have alias")

type ErrTomlFileIsInvalid struct {
	reason error
}

func (err *ErrTomlFileIsInvalid) Error() string {
	return fmt.Sprintf("error: this .toml file is invalid reason: %s", err.reason)
}

func (a *Alias) len() int {
	return len(a.order)
}

func GetParseAliasFunc(filepath string) ([]dataformat.DecodeOption, error) {
	ailas, err := getAlias(filepath)
	if err != nil {
		return nil, err
	}
	parseFuncs := make([]dataformat.DecodeOption, 0, ailas.len())

	for _, key := range ailas.order {
		value := ailas.dict[key]
		if funcs, err := getReplaceAliasFuncFromKey(key, value); err == nil {
			parseFuncs = append(parseFuncs, funcs)
		}
	}

	return parseFuncs, nil
}

func getAlias(filepath string) (*Alias, error) {
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

	alias := &Alias{
		dict:  make(map[string][]byte),
		order: []string{},
	}

	if err := file.IterateLineFromFile(filepath, alias.getCreateAliasByLineFunc()); err != nil {
		return nil, err
	}

	return alias, nil
}

func (alias *Alias) getCreateAliasByLineFunc() func(string) {
	parser := alias.newParser()
	return parser.parse
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
