package alias

import (
	"fmt"
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

var replaceAliasKeyPattern = `"#{{%s}}"`

func getReplaceAliasFuncFromKey(key string, value []byte) (func([]byte) []byte, error) {
	pattern := fmt.Sprintf(replaceAliasKeyPattern, key)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return func(b []byte) []byte {
		return re.ReplaceAll(b, value)
	}, nil
}
