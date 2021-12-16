package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBytesFromFile(t *testing.T) {
	t.Parallel()

	t.Run("test GetBytesFromFile with exist file", func(t *testing.T) {
		givenPath := "./test/testRead"
		expected := "Hello, World!"

		gotByte, gotErr := GetBytesFromFile(givenPath)
		assert.Nil(t, gotErr)
		assert.Equal(t, expected, string(gotByte))
	})

	t.Run("test GetBytesFromFile with non-exist file", func(t *testing.T) {
		givenPath := "./test/no-exist"
		expected := []byte{}
		expectedErr := "stat ./test/no-exist: no such file or directory"

		gotByte, gotErr := GetBytesFromFile(givenPath)

		assert.Equal(t, expectedErr, gotErr.Error())
		assert.Equal(t, expected, gotByte)
	})

	t.Run("test GetBytesFromFile with directory", func(t *testing.T) {
		givenPath := "./test/"
		expected := []byte{}
		expectedErr := "a given path is not a file: ./test/"

		gotByte, gotErr := GetBytesFromFile(givenPath)

		assert.Equal(t, expectedErr, gotErr.Error())
		assert.Equal(t, expected, gotByte)
	})

	// t.Run("test GetBytesFromFile with configuration file", func(t *testing.T) {
	// 	givenPath := "./../../test/config/04_configSpec_alias.ss.toml"

	// 	gotByte, _ := GetBytesFromFile(givenPath)
	// 	var ssConfig cf.SSConfig

	// 	// f, _ := os.Open(givenPath)
	// 	// dec := toml.NewDecoder(f)

	// 	// _, _ = dec.Decode(&ssConfig)
	// 	_, err := toml.Decode(string(gotByte), &ssConfig)

	// 	fmt.Printf("%s\n", err)
	// 	fmt.Printf("%#v\n", ssConfig)

	// })
}
func TestGetFileType(t *testing.T) {
	t.Run("GetFileType when is not file", func(t *testing.T) {
		cases := []struct {
			path     string
			expected string
		}{
			{"./test/fileType/not-exist-file", "stat ./test/fileType/not-exist-file: no such file or directory"},
			{"./test/fileType/dir", "a given path is not a file: ./test/fileType/dir"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("GetFileType on path %s expected to get error %s", tc.path, tc.expected), func(t *testing.T) {
				t.Parallel()

				gotFileType, gotErr := GetFileType(tc.path)

				assert.Equal(t, "", gotFileType)
				assert.Equal(t, tc.expected, gotErr.Error())
			})
		}
	})

	t.Run("GetFileType when is file", func(t *testing.T) {
		cases := []struct {
			path     string
			expected string
		}{
			{"./test/fileType/file", ""},
			{"./test/fileType/file.filetype", "filetype"},
			{"./test/fileType/file.subfiletype.filetype", "filetype"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("GetFileType on path %s expected to get %s", tc.path, tc.expected), func(t *testing.T) {
				t.Parallel()

				gotFileType, gotErr := GetFileType(tc.path)

				assert.Nil(t, gotErr)
				assert.Equal(t, tc.expected, gotFileType)
			})
		}
	})
}
