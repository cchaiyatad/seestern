package alias

type foundKeyAfterValue struct {
	parser *parser
}

func (s *foundKeyAfterValue) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *foundKeyAfterValue) isFoundKey(line string) error {
	key, ok := findKey(line)
	if !ok {
		return nil
	}

	s.parser.currentKey = key
	s.parser.insertCurrentAlias()
	s.parser.setState(s.parser.waitForAilas)
	return nil
}
func (s *foundKeyAfterValue) isFoundValue(line string) error {
	panic("not implement")
}
func (s *foundKeyAfterValue) isValueComplete() error {
	panic("not implement")
}
