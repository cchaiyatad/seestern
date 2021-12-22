package alias

type foundAilas struct {
	parser *parser
}

func (s *foundAilas) isFoundAilas(line string) error {
	panic("not implement")
}
func (s *foundAilas) isFoundKey(line string) error {
	key, ok := findKey(line)
	if !ok {
		return nil
	}

	s.parser.currentKey = key
	s.parser.setState(s.parser.waitForValueAfterFoundKey)
	return nil
}
func (s *foundAilas) isFoundValue(line string) error {
	value, ok := findKey(line)
	if !ok {
		return nil
	}

	s.parser.currentValue.WriteString(value)
	s.parser.setState(s.parser.validateValue)
	return nil
}

func (s *foundAilas) isValueComplete() error {
	panic("not implement")
}
