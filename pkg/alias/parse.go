package alias

import "fmt"

type parser struct {
	alias Alias
}

func (alias Alias) newParser() *parser {
	return &parser{alias: alias}
}

func (parser *parser) getParseFunc() func(string) {
	funcs := func(line string) {

		if strip(line) == "[[alias]]" {
			fmt.Printf("line: %s\n", line)
		}

	}
	return funcs
}
