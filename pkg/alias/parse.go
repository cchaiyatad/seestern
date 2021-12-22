package alias

import (
	"fmt"
	"strings"
)

type parser struct {
	waitForAilas state
	foundAilas   state

	foundKeyBeforeValue state
	foundValueAfterKey  state

	foundValueBeforeKey state
	foundKeyAfterValue  state

	currentState state

	alias        Alias
	currentKey   string
	currentValue *strings.Builder
}

func (alias Alias) newParser() *parser {
	parser := &parser{alias: alias}

	parser.waitForAilas = &waitForAilas{parser: parser}
	parser.foundAilas = &foundAilas{parser: parser}

	parser.foundKeyBeforeValue = &foundKeyBeforeValue{parser: parser}
	parser.foundValueAfterKey = &foundValueAfterKey{parser: parser}

	parser.foundValueBeforeKey = &foundValueBeforeKey{parser: parser}
	parser.foundKeyAfterValue = &foundKeyAfterValue{parser: parser}

	parser.setState(parser.waitForAilas)

	parser.clearCurrentData()
	return parser
}

func (parser *parser) isFoundAlias(line string) error {
	return parser.currentState.isFoundAilas(line)
}
func (parser *parser) isFoundKey(line string) error {
	return parser.currentState.isFoundKey(line)
}
func (parser *parser) isFoundValue(line string) error {
	return parser.currentState.isFoundValue(line)
}
func (parser *parser) isValueComplete(line string) error {
	return parser.currentState.isValueComplete(line)
}
func (parser *parser) setState(s state) {
	parser.currentState = s
}
func (parser *parser) checkIllegalState(err error) {
	if err != nil {
		fmt.Println("illegal")
		parser.currentState = parser.waitForAilas
		parser.clearCurrentData()
	}
}

func (parser *parser) clearCurrentData() {
	parser.currentKey = ""
	parser.currentValue = &strings.Builder{}
}

func (parser *parser) parse(line string) {
	var err error
	switch parser.currentState {
	case parser.waitForAilas:
		err = parser.isFoundAlias(line)
	case parser.foundAilas:
		err = parser.isFoundKey(line)
	case parser.foundKeyBeforeValue:
		err = parser.isFoundValue(line)
	case parser.foundValueAfterKey:

	case parser.foundValueBeforeKey:

	case parser.foundKeyAfterValue:

	}

	parser.checkIllegalState(err)
}

func (parser *parser) getParseFunc() func(string) {

	return parser.parse
}
