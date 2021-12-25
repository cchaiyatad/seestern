package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitRef(t *testing.T) {
	cases := []struct {
		nym string

		expectedNym    string
		expectedNymRef string
		expectedOK     bool
	}{
		{"school.student.name", "school.student", "name", true},
		{"school.teacher._id", "school.teacher", "_id", true},

		{"school.student", "", "", false},
		{"school.teacher", "", "", false},

		{"", "", "", false},
		{"absedse", "", "", false},
		{"abc.def.gh", "abc.def", "gh", true},

		{"school.student.name.firstname", "school.student", "name.firstname", true},
		{"school.student.name.firstname.middlename", "school.student", "name.firstname.middlename", true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("SplitRef %s", tc.nym), func(t *testing.T) {
			t.Parallel()

			gotNym, gotNymRef, gotOk := SplitRef(tc.nym)
			assert.Equal(t, tc.expectedNym, gotNym)
			assert.Equal(t, tc.expectedNymRef, gotNymRef)
			assert.Equal(t, tc.expectedOK, gotOk)
		})
	}
}

func TestSplitNymRef(t *testing.T) {
	cases := []struct {
		nym string

		expected   []string
		expectedOK bool
	}{
		{"school.student", []string{"school", "student"}, true},
		{"school.teacher", []string{"school", "teacher"}, true},

		{"", []string{}, false},
		{"absedse", []string{"absedse"}, true},
		{"abc.def.gh", []string{"abc", "def", "gh"}, true},
		{"abc.def.gh.ij", []string{"abc", "def", "gh", "ij"}, true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("SplitNymRef %s", tc.nym), func(t *testing.T) {
			t.Parallel()

			got, gotOk := SplitNymRef(tc.nym)
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.expectedOK, gotOk)
		})
	}
}

func TestSplitNym(t *testing.T) {
	cases := []struct {
		nym string

		expectedDBName    string
		exptectedCollName string
		expectedOK        bool
	}{
		{"school.student", "school", "student", true},
		{"school.teacher", "school", "teacher", true},

		{"", "", "", false},
		{"absedse", "", "", false},
		{"abc.def.gh", "abc", "def", true},
		{"abc.def.gh.ij", "abc", "def", true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("SplitNym %s", tc.nym), func(t *testing.T) {
			t.Parallel()

			gotDBName, gotCollName, gotOk := SplitNym(tc.nym)
			assert.Equal(t, tc.expectedDBName, gotDBName)
			assert.Equal(t, tc.exptectedCollName, gotCollName)
			assert.Equal(t, tc.expectedOK, gotOk)
		})
	}
}

func TestCreateNym(t *testing.T) {
	cases := []struct {
		dbName, collName, expected string
	}{
		{"school", "student", "school.student"},
		{"school", "teacher", "school.teacher"},

		{"", "teacher", "teacher"},
		{"school", "", "school"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("CreateNym %s %s show got %s", tc.dbName, tc.collName, tc.expected), func(t *testing.T) {
			t.Parallel()

			got := CreateNym(tc.dbName, tc.collName)
			assert.Equal(t, tc.expected, got)
		})
	}
}
func TestCreateNymRef(t *testing.T) {
	cases := []struct {
		dbName, collName, expected string
	}{
		{"school", "student", "school.student"},
		{"school", "teacher", "school.teacher"},

		{"", "teacher", "teacher"},
		{"school", "", "school"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("CreateNymRef %s %s show got %s", tc.dbName, tc.collName, tc.expected), func(t *testing.T) {
			t.Parallel()

			got := CreateNymRef(tc.dbName, tc.collName)
			assert.Equal(t, tc.expected, got)
		})
	}
}
