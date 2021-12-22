package alias

type waitForKeyAfterFoundValue struct {
	parser *parser
}

func (s *waitForKeyAfterFoundValue) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *waitForKeyAfterFoundValue) isFoundKey(line string) error {
	key, ok := findKey(line)
	if !ok {
		return nil
	}

	s.parser.currentKey = key
	s.parser.insertCurrentAlias()
	s.parser.setState(s.parser.waitForAilas)
	return nil
}
func (s *waitForKeyAfterFoundValue) isFoundValue(line string) error {
	panic("not implement")
}
func (s *waitForKeyAfterFoundValue) isValueComplete() error {
	panic("not implement")
}
