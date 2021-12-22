package alias

type foundValueAfterKey struct {
	parser *parser
}

func (s *foundValueAfterKey) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *foundValueAfterKey) isFoundKey(line string) error {
	panic("not implement")
}
func (s *foundValueAfterKey) isFoundValue(line string) error {
	value := strip(line)

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}
func (s *foundValueAfterKey) isValueComplete() error {
	panic("not implement")
}
