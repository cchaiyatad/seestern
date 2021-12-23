package alias

type waitForAilas struct {
	parser *parser
}

func (*waitForAilas) String() string {
	return "waitForAilas"
}

func (s *waitForAilas) isFoundAilas(line string) error {
	if strip(line) == "[[alias]]" {
		s.parser.setState(s.parser.foundAilas)
	}
	return nil
}
func (s *waitForAilas) isFoundKey(line string) error {
	return ErrIllegalMethod
}
func (s *waitForAilas) isFoundValue(line string) error {
	return ErrIllegalMethod
}
func (s *waitForAilas) isValueComplete() error {
	return ErrIllegalMethod
}
