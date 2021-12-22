package alias

type waitForValueToCompleteBeforeGoToWaitAlias struct {
	parser *parser
}

func (s *waitForValueToCompleteBeforeGoToWaitAlias) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *waitForValueToCompleteBeforeGoToWaitAlias) isFoundKey(line string) error {
	panic("not implement")
}
func (s *waitForValueToCompleteBeforeGoToWaitAlias) isFoundValue(line string) error {
	value := strip(line)

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}
func (s *waitForValueToCompleteBeforeGoToWaitAlias) isValueComplete() error {
	panic("not implement")
}
