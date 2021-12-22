package alias

import (
	"regexp"
	"strings"
)

var isValueReg = regexp.MustCompile("value( *)=( *)")
var isKeyReg = regexp.MustCompile("key( *)=( *)")

func strip(s string) string {
	return strings.TrimSpace(s)
}

func isMatchRegex(re *regexp.Regexp, line string) bool {
	return re.MatchString(line)
}

func removeRegex(re *regexp.Regexp, source string) string {
	return re.ReplaceAllString(source, "")
}
