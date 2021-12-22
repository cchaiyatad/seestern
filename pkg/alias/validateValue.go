package alias

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type validateValue struct {
	parser *parser
}

func (s *validateValue) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *validateValue) isFoundKey(line string) error {
	panic("not implement")
}
func (s *validateValue) isFoundValue(line string) error {
	panic("not implement")
}
func (s *validateValue) isValueComplete() error {
	value := s.parser.currentValue.String()

	if err := isTomlValid(value); err != nil {
		fmt.Println(value)
		fmt.Println(err)

		if s.parser.currentKey == "" {
			s.parser.setState(s.parser.foundValueBeforeKey)
		} else {
			s.parser.setState(s.parser.foundValueAfterKey)
		}
		return nil
	}

	s.parser.insertCurrentAlias()
	s.parser.setState(s.parser.waitForAilas)
	return nil
}

func isTomlValid(tomlData string) error {
	var tmp interface{}
	_, err := toml.Decode("tmp="+tomlData, &tmp)
	return err
}
