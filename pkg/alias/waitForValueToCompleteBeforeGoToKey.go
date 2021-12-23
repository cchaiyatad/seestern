package alias

type waitForValueToCompleteBeforeGoToKey struct {
	parser *parser
}

func (*waitForValueToCompleteBeforeGoToKey) String() string {
	return "waitForValueToCompleteBeforeGoToKey"
}

func (s *waitForValueToCompleteBeforeGoToKey) isFoundAilas(line string) error {
	return ErrIllegalMethod
}
func (s *waitForValueToCompleteBeforeGoToKey) isFoundKey(line string) error {
	return ErrIllegalMethod
}
func (s *waitForValueToCompleteBeforeGoToKey) isFoundValue(line string) error {
	value := strip(line)

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}
func (s *waitForValueToCompleteBeforeGoToKey) isValueComplete() error {
	return ErrIllegalMethod
}
