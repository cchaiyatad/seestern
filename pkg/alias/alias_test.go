package alias

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	noAliasPath                    = "./test/valid/00_no_alias.ss.toml"
	simpleAliasPath                = "./test/valid/01_simple_alias.ss.toml"
	complexAliasPath               = "./test/valid/02_complex_alias.ss.toml"
	simpleWithOtherValuePath       = "./test/valid/03_simple_with_other_value.ss.toml"
	twoAliasPath                   = "./test/valid/04_two_alias.ss.toml"
	noCoresspondAliasPath          = "./test/valid/05_no_coresspond_alias.ss.toml"
	complexAliasValueBeforeKeyPath = "./test/valid/06_complex_value_before_key.ss.toml"

	// Invalid
	noValuePath        = "./test/invalid/01_no_value.ss.toml"
	noKeyPath          = "./test/invalid/02_no_key.ss.toml"
	noKeyInvalidPath   = "./test/invalid/03_key_invalid.ss.toml"
	noValueInvalidPath = "./test/invalid/04_value_invalid.ss.toml"
)

func TestGetAlias(t *testing.T) {
	t.Parallel()

	t.Run("getAlias with no-exist file", func(t *testing.T) {
		cases := []struct {
			filePath    string
			expectedErr string
		}{
			{"./test/valid/no-exist", "stat ./test/valid/no-exist: no such file or directory"},
			{"./test/valid", "a given path is not a file: ./test/valid"},
			{"./test/valid/config.ss.json", "error: alias only support in toml"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("getAlias from %s ", tc.filePath), func(t *testing.T) {
				t.Parallel()
				gotAliases, gotErr := getAlias(tc.filePath)
				assert.Nil(t, gotAliases)

				assert.Equal(t, tc.expectedErr, gotErr.Error())
			})
		}
	})

	t.Run("getAlias with noAlias file", func(t *testing.T) {
		givenPath := noAliasPath
		gotAliases, gotErr := getAlias(givenPath)
		expectedErr := "error: this file does not have alias"
		assert.Nil(t, gotAliases)
		assert.Equal(t, expectedErr, gotErr.Error())

	})

	t.Run("getAlias with simpleAliasPath file", func(t *testing.T) {
		givenPath := simpleAliasPath

		expectedAlias := Alias{
			"stringLenTen": []byte(`[{type = "string", length = 10}]`),
		}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})

	t.Run("getAlias with complexAlias file", func(t *testing.T) {
		givenPath := complexAliasPath

		expectedAlias := Alias{
			"class_constraint": []byte(`[{type = "array", element_type = [{type = "object", fields = [{f_name = "class_name", constraint = [{type = "string"}]},{f_name = "instructor", constraint = [{type = "string"}]},]},]},]`),
		}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})

	t.Run("getAlias with simpleWithOtherValue file", func(t *testing.T) {
		givenPath := simpleWithOtherValuePath

		expectedAlias := Alias{
			"stringLenTen": []byte(`[{type = "string", length = 10}]`),
		}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})
	t.Run("getAlias with twoAlias file", func(t *testing.T) {
		givenPath := twoAliasPath

		expectedAlias := Alias{
			"stringLenFifteen": []byte(`[{type = "string", length = 15}]`),
			"stringLenTen":     []byte(`[{type = "string", length = 10}]`),
		}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})
	t.Run("getAlias with noCoresspondAlias file", func(t *testing.T) {
		givenPath := noCoresspondAliasPath

		expectedAlias := Alias{
			"stringLenFive": []byte(`[{type = "string", length = 5}]`),
		}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})
	t.Run("getAlias with complexAlias value before key file", func(t *testing.T) {
		givenPath := complexAliasValueBeforeKeyPath

		expectedAlias := Alias{
			"class_constraint": []byte(`[{type = "array", element_type = [{type = "object", fields = [{f_name = "class_name", constraint = [{type = "string"}]},{f_name = "instructor", constraint = [{type = "string"}]},]},]},]`),
		}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})
}

func TestGetAliasForInvalidFile(t *testing.T) {
	t.Run("getAlias with noValuePath file", func(t *testing.T) {
		givenPath := noValuePath

		expectedAlias := Alias{}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})
	t.Run("getAlias with noKeyPath file", func(t *testing.T) {
		givenPath := noKeyPath

		expectedAlias := Alias{}

		gotAliases, gotErr := getAlias(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expectedAlias, gotAliases)
	})
	t.Run("getAlias with noKeyInvalidPath file", func(t *testing.T) {
		givenPath := noKeyInvalidPath

		expectPrefixRegexErr := "error: this .toml file is invalid reason: "

		gotAliases, gotErr := getAlias(givenPath)
		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotAliases)
	})
	t.Run("getAlias with noValueInvalidPath file", func(t *testing.T) {
		givenPath := noValueInvalidPath

		expectPrefixRegexErr := "error: this .toml file is invalid reason: "

		gotAliases, gotErr := getAlias(givenPath)
		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotAliases)
	})

}
