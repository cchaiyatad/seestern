package alias

type waitForValueToCompleteBeforeGoToKey struct {
	parser *parser
}

func (s *waitForValueToCompleteBeforeGoToKey) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *waitForValueToCompleteBeforeGoToKey) isFoundKey(line string) error {
	panic("not implement")
}
func (s *waitForValueToCompleteBeforeGoToKey) isFoundValue(line string) error {
	value := strip(line)

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}
func (s *waitForValueToCompleteBeforeGoToKey) isValueComplete() error {
	panic("not implement")
}
