package alias

import (
	"strings"
)

type parser struct {
	waitForAilas state
	foundAilas   state

	foundKeyBeforeValue state
	foundValueAfterKey  state

	foundValueBeforeKey state
	foundKeyAfterValue  state
	validateValue       state

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

	parser.validateValue = &validateValue{parser: parser}

	parser.setState(parser.waitForAilas)

	parser.clearCurrentData()
	return parser
}

func (parser *parser) clearCurrentData() {
	parser.currentKey = ""
	parser.currentValue = &strings.Builder{}
}

func (parser *parser) insertCurrentAlias() {
	parser.alias[parser.currentKey] = []byte(parser.currentValue.String())
	parser.clearCurrentData()
}

func (parser *parser) parse(line string) {
	var err error
	switch parser.currentState {
	case parser.waitForAilas:
		err = parser.isFoundAlias(line)
	case parser.foundAilas:
		err = parser.isFoundKey(line)
		// err = parser.isFoundValue(line)

	case parser.foundKeyBeforeValue:
		err = parser.isFoundValue(line)
	case parser.foundValueAfterKey:
		err = parser.isFoundValue(line)

	case parser.foundValueBeforeKey:
		err = parser.isFoundValue(line)

	case parser.foundKeyAfterValue:
		err = parser.isFoundKey(line)

	}
	parser.checkIllegalState(err)

	if parser.currentState == parser.validateValue {
		err = parser.isValueComplete()
		parser.checkIllegalState(err)
	}
}

func (parser *parser) getParseFunc() func(string) {

	return parser.parse
}
