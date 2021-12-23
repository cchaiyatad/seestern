package alias

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrIllegalMethod = errors.New("this method is not implemented for this state")

type state interface {
	isFoundAilas(string) error
	isFoundKey(string) error
	isFoundValue(string) error
	isValueComplete() error
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
func (parser *parser) isValueComplete() error {
	return parser.currentState.isValueComplete()
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

func findValue(line string) (string, bool) {
	return find(line, isValueReg)
}

func findKey(line string) (string, bool) {
	return find(line, isKeyReg)
}

func find(line string, Reg *regexp.Regexp) (string, bool) {
	line = strip(line)
	if !isMatchRegex(Reg, line) {
		return "", false
	}

	value := removeRegex(Reg, line)
	return value, true
}
