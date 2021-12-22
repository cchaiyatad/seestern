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
	// check is complete
	s.parser.setState(s.parser.waitForAilas)
	return nil
}

func (s *foundKeyBeforeValue) isValueComplete(line string) error {
	panic("not implement")
}
