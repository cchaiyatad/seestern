package alias

type waitForValueAfterFoundKey struct {
	parser *parser
}

func (*waitForValueAfterFoundKey) String() string {
	return "waitForValueAfterFoundKey"
}

func (s *waitForValueAfterFoundKey) isFoundAilas(line string) error {
	return ErrIllegalMethod
}
func (s *waitForValueAfterFoundKey) isFoundKey(line string) error {
	return ErrIllegalMethod
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
	return ErrIllegalMethod
}
