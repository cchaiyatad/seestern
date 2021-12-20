package gen

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenInt(t *testing.T) {

}

func TestGenDouble(t *testing.T) {

}

func TestGenString(t *testing.T) {
	cases := []struct {
		prefix             string
		suffix             string
		length             int
		exptectedMaxLength int
	}{
		{"", "", 15, 15},
		{"Pre", "", 15, 15},
		{"", "ly", 10, 10},
		{"Pre", "ly", 5, 5},

		{"", "", 0, 0},
		{"", "", -5, 0},
		{"pre", "ly", -5, 0},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("GenString with prefix: %s suffix: %s length %d", tc.prefix, tc.suffix, tc.length), func(t *testing.T) {
			t.Parallel()

			got := GenString(tc.length, tc.prefix, tc.suffix)

			expectPrefixRegexErr := fmt.Sprintf(`^%s[a-zA-Z]{0,%d}%s$`, tc.prefix, tc.exptectedMaxLength, tc.suffix)
			assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), got, fmt.Sprintf("expect prefix %s suffix: %s length %d: got %s", tc.prefix, tc.suffix, tc.length, got))
		})
	}
}

// func FuzzGenInt(f *testing.F) {

// }
// func FuzzGenDouble(f *testing.F) {

// }
