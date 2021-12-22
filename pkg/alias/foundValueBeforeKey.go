package alias

type foundValueBeforeKey struct {
	parser *parser
}

func (s *foundValueBeforeKey) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *foundValueBeforeKey) isFoundKey(line string) error {
	panic("not implement")
}
func (s *foundValueBeforeKey) isFoundValue(line string) error {
	value := strip(line)

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}
func (s *foundValueBeforeKey) isValueComplete() error {
	panic("not implement")
}
