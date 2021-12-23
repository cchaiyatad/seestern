package alias

type waitForValueToCompleteBeforeGoToWaitAlias struct {
	parser *parser
}

func (*waitForValueToCompleteBeforeGoToWaitAlias) String() string {
	return "waitForValueToCompleteBeforeGoToWaitAlias"
}

func (s *waitForValueToCompleteBeforeGoToWaitAlias) isFoundAilas(line string) error {
	return ErrIllegalMethod
}
func (s *waitForValueToCompleteBeforeGoToWaitAlias) isFoundKey(line string) error {
	return ErrIllegalMethod
}
func (s *waitForValueToCompleteBeforeGoToWaitAlias) isFoundValue(line string) error {
	value := strip(line)

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}
func (s *waitForValueToCompleteBeforeGoToWaitAlias) isValueComplete() error {
	return ErrIllegalMethod
}
