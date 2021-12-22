package alias

import "regexp"

type state interface {
	isFoundAilas(string) error
	isFoundKey(string) error
	isFoundValue(string) error
	isValueComplete(string) error
}

func findValue(line string) (string, bool) {
	return find(line, isValueReg)
}

func findKey(line string) (string, bool) {
	return find(line, isKeyReg)
}

func find(line string, Reg *regexp.Regexp) (string, bool) {
	line = strip(line)
	if !isMatchRegex(Reg, line) {
		return "", false
	}

	value := removeRegex(Reg, line)
	return value, true
}
