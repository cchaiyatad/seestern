package alias

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	noAliasPath              = "./test/alias/00_no_alias.ss.toml"
	simpleAliasPath          = "./test/alias/01_simple_alias.ss.toml"
	complexAliasPath         = "./test/alias/02_complex_alias.ss.toml"
	simpleWithOtherValuePath = "./test/alias/03_simple_with_other_value.ss.toml"
	twoAliasPath             = "./test/alias/04_two_alias.ss.toml"
	noCoresspondAliasPath    = "./test/alias/05_no_coresspond_alias.ss.toml"
)

func TestGetAlias(t *testing.T) {
	t.Parallel()

	t.Run("getAlias with no-exist file", func(t *testing.T) {
		cases := []struct {
			filePath    string
			expectedErr string
		}{
			{"./test/alias/no-exist", "stat ./test/alias/no-exist: no such file or directory"},
			{"./test/alias", "a given path is not a file: ./test/alias"},
			{"./test/alias/config.ss.json", "error: alias only support in toml"},
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
		// for k, v := range gotAliases {
		// 	fmt.Printf("%s: %s\n", k, string(v))
		// }
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
}
