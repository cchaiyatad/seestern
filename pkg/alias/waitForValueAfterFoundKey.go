package alias

type waitForValueAfterFoundKey struct {
	parser *parser
}

func (s *waitForValueAfterFoundKey) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *waitForValueAfterFoundKey) isFoundKey(line string) error {
	panic("not implement")
}
func (s *waitForValueAfterFoundKey) isFoundValue(line string) error {
	value, ok := findValue(line)
	if !ok {
		return nil
	}

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}

func (s *waitForValueAfterFoundKey) isValueComplete() error {
	panic("not implement")
}
