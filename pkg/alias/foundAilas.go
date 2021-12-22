package alias

import "fmt"

type foundAilas struct {
	parser *parser
}

func (s *foundAilas) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *foundAilas) isFoundKey(line string) error {
	key, ok := findKey(line)
	if !ok {
		return nil
	}

	fmt.Printf("key: %s\n", key)
	s.parser.currentKey = key
	s.parser.setState(s.parser.foundKeyBeforeValue)
	return nil
}
func (s *foundAilas) isFoundValue(line string) error {
	value, ok := findKey(line)
	if !ok {
		return nil
	}

	fmt.Printf("value: %s\n", value)
	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.foundValueBeforeKey)
	return nil
}

func (s *foundAilas) isValueComplete() error {
	panic("not implement")
}
