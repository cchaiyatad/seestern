package alias

import (
	"strings"
)

type parser struct {
	waitForAilas state
	foundAilas   state

	waitForValueAfterFoundKey                 state
	waitForValueToCompleteBeforeGoToWaitAlias state

	waitForValueToCompleteBeforeGoToKey state
	waitForKeyAfterFoundValue           state

	validateValue state
	currentState  state

	alias        Alias
	currentKey   string
	currentValue *strings.Builder
}

func (alias Alias) newParser() *parser {
	parser := &parser{alias: alias}

	parser.waitForAilas = &waitForAilas{parser: parser}
	parser.foundAilas = &foundAilas{parser: parser}

	parser.waitForValueAfterFoundKey = &waitForValueAfterFoundKey{parser: parser}
	parser.waitForValueToCompleteBeforeGoToWaitAlias = &waitForValueToCompleteBeforeGoToWaitAlias{parser: parser}

	parser.waitForValueToCompleteBeforeGoToKey = &waitForValueToCompleteBeforeGoToKey{parser: parser}
	parser.waitForKeyAfterFoundValue = &waitForKeyAfterFoundValue{parser: parser}

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
	key := strings.Trim(parser.currentKey, `"`)
	parser.alias[key] = []byte(parser.currentValue.String())
	parser.clearCurrentData()
}

func (parser *parser) parse(line string) {
	var err error
	switch parser.currentState {
	case parser.waitForAilas:
		err = parser.isFoundAlias(line)
	case parser.foundAilas:
		if keyErr := parser.isFoundKey(line); keyErr == ErrKeyNotFound {
			err = parser.isFoundValue(line)
		} else {
			err = keyErr
		}
	case parser.waitForValueAfterFoundKey:
		err = parser.isFoundValue(line)
	case parser.waitForValueToCompleteBeforeGoToWaitAlias:
		err = parser.isFoundValue(line)
	case parser.waitForValueToCompleteBeforeGoToKey:
		err = parser.isFoundValue(line)
	case parser.waitForKeyAfterFoundValue:
		err = parser.isFoundKey(line)

	}
	parser.checkIllegalState(err)

	if parser.currentState == parser.validateValue {
		err = parser.isValueComplete()
		parser.checkIllegalState(err)
	}
}
