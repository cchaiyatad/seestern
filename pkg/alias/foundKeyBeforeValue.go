package alias

import "fmt"

type foundKeyBeforeValue struct {
	parser *parser
}

func (s *foundKeyBeforeValue) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *foundKeyBeforeValue) isFoundKey(line string) error {
	panic("not implement")
}
func (s *foundKeyBeforeValue) isFoundValue(line string) error {
	value, ok := findValue(line)
	if !ok {
		return nil
	}

	fmt.Printf("value: %s\n", value)
	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}

func (s *foundKeyBeforeValue) isValueComplete() error {
	panic("not implement")
}
