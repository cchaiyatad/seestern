package alias

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found")

type foundAilas struct {
	parser *parser
}

func (*foundAilas) String() string {
	return "foundAilas"
}

func (s *foundAilas) isFoundAilas(line string) error {
	return ErrIllegalMethod
}
func (s *foundAilas) isFoundKey(line string) error {
	key, ok := findKey(line)
	if !ok {
		return ErrKeyNotFound
	}

	s.parser.currentKey = key
	s.parser.setState(s.parser.waitForValueAfterFoundKey)
	return nil
}
func (s *foundAilas) isFoundValue(line string) error {
	value, ok := findValue(line)
	if !ok {
		return nil
	}

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}

func (s *foundAilas) isValueComplete() error {
	return ErrIllegalMethod
}
