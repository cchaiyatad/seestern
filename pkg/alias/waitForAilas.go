package alias

type waitForAilas struct {
	parser *parser
}

func (s *waitForAilas) isFoundAilas(line string) error {
	if strip(line) == "[[alias]]" {
		s.parser.setState(s.parser.foundAilas)
	}
	return nil
}
func (s *waitForAilas) isFoundKey(line string) error {
	panic("not implement")
}
func (s *waitForAilas) isFoundValue(line string) error {
	panic("not implement")
}
func (s *waitForAilas) isValueComplete() error {
	panic("not implement")
}
