package alias

import (
	"github.com/BurntSushi/toml"
)

type validateValue struct {
	parser *parser
}

func (*validateValue) String() string {
	return "validateValue"
}

func (s *validateValue) isFoundAilas(line string) error {
	return ErrIllegalMethod
}
func (s *validateValue) isFoundKey(line string) error {
	return ErrIllegalMethod
}
func (s *validateValue) isFoundValue(line string) error {
	return ErrIllegalMethod
}
func (s *validateValue) isValueComplete() error {
	value := s.parser.currentValue.String()

	if err := isTomlValid(value); err != nil {
		s.setStateByKey(s.parser.waitForValueToCompleteBeforeGoToWaitAlias, s.parser.waitForValueToCompleteBeforeGoToKey)
		return nil
	}

	s.setStateByKey(s.parser.waitForAilas, s.parser.waitForKeyAfterFoundValue)
	s.callBackByKey(s.parser.insertCurrentAlias, func() {})
	return nil
}

func (s *validateValue) setStateByKey(foundKeyFirst, foundValueFirst state) {
	s.callBackByKey(s.getSetStateFunc(foundKeyFirst), s.getSetStateFunc(foundValueFirst))
}

func (s *validateValue) getSetStateFunc(state state) func() {
	return func() {
		s.parser.setState(state)
	}
}
func (s *validateValue) callBackByKey(foundKeyFirstFuncs, foundValueFirstFuncs func()) {
	if s.parser.currentKey == "" {
		foundValueFirstFuncs()
	} else {
		foundKeyFirstFuncs()
	}
}

func isTomlValid(tomlData string) error {
	var tmp interface{}
	_, err := toml.Decode("tmp="+tomlData, &tmp)
	return err
}
