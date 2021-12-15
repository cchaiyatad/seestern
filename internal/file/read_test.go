package file

import (
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
		expectedErr := "cannot read a file: open ./test/no-exist: no such file or directory"

		gotByte, gotErr := GetBytesFromFile(givenPath)

		assert.Equal(t, expectedErr, gotErr.Error())
		assert.Equal(t, expected, gotByte)
	})

	t.Run("test GetBytesFromFile with directory", func(t *testing.T) {
		givenPath := "./test/"
		expected := []byte{}
		expectedErr := "read ./test/: is a directory"

		gotByte, gotErr := GetBytesFromFile(givenPath)

		assert.Equal(t, expectedErr, gotErr.Error())
		assert.Equal(t, expected, gotByte)
	})
}
