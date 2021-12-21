package gen

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenInt(t *testing.T) {
	cases := []struct {
		min         int
		max         int
		expectedMin int
		expectedMax int
	}{
		{0, 0, 0, 100},
		{10, 0, 10, 110},

		{10, 15, 10, 14},
		{0, 100, 0, 99},

		{-10, 40, -10, 39},
		{-50, -10, -50, -9},

		{0, 1, 0, 0},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("GenInt with min: %d max: %d", tc.min, tc.max), func(t *testing.T) {
			t.Parallel()
			got := GenInt(tc.min, tc.max)

			assert.GreaterOrEqual(t, got, tc.expectedMin)
			assert.LessOrEqual(t, got, tc.expectedMax)
		})
	}
}

func TestGenDouble(t *testing.T) {
	cases := []struct {
		min         float64
		max         float64
		expectedMin float64
		expectedMax float64
	}{
		{0, 0, 0, 100},
		{10, 0, 0, 110},

		{10, 15, 10, 15},
		{0, 100, 0, 100},

		{-10, 40, -10, 40},
		{-50, -10, -50, -10},

		{0, 1, 0, 1},

		{0.5, 1.0, 0.5, 1.0},
		{-0.5, 1.0, -0.5, 1.0},
		{-1.5, -1.0, -1.5, -1.0},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("GenDouble with min: %f max: %f", tc.min, tc.max), func(t *testing.T) {
			t.Parallel()

			got := GenDouble(tc.min, tc.max)

			assert.GreaterOrEqual(t, got, tc.expectedMin)
			assert.LessOrEqual(t, got, tc.expectedMax)
		})
	}
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

			expectPrefixRegexErr := fmt.Sprintf(`^%s[a-zA-Z0-9]{0,%d}%s$`, tc.prefix, tc.exptectedMaxLength, tc.suffix)
			assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), got, fmt.Sprintf("expect prefix %s suffix: %s length %d: got %s", tc.prefix, tc.suffix, tc.length, got))
		})
	}
}

// func FuzzGenInt(f *testing.F) {

// }
// func FuzzGenDouble(f *testing.F) {

// }
