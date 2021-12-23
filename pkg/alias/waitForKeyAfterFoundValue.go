package alias

type waitForKeyAfterFoundValue struct {
	parser *parser
}

func (*waitForKeyAfterFoundValue) String() string {
	return "waitForKeyAfterFoundValue"
}

func (s *waitForKeyAfterFoundValue) isFoundAilas(line string) error {
	return ErrIllegalMethod
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
	return ErrIllegalMethod
}
func (s *waitForKeyAfterFoundValue) isValueComplete() error {
	return ErrIllegalMethod
}
