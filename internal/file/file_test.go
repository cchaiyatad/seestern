package file

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	// clear all create dir

	_ = os.Remove("./test/createDir/new1")
	_ = os.Remove("./test/createDir/new2/new2.1")
	_ = os.Remove("./test/createDir/new2")

	_ = os.Remove("./test/prepare/new1")
	_ = os.Remove("./test/prepare/new2/new2.1")
	_ = os.Remove("./test/prepare/new2")

	os.Exit(exitVal)
}

func TestPrepareDir(t *testing.T) {
	t.Run("PrepareDir when path is file", func(t *testing.T) {
		cases := []struct {
			path     string
			expected string
		}{
			{"./test/prepare/file1", "a given path is not a directory: ./test/prepare/file1"},
			{"./test/prepare/dir1/file1.1", "a given path is not a directory: ./test/prepare/dir1/file1.1"},
			{"./test/prepare/dir1/dir2/file2.1", "a given path is not a directory: ./test/prepare/dir1/dir2/file2.1"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("PrepareDir on path %s expected an error", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Equal(t, tc.expected, PrepareDir(tc.path).Error())
			})
		}
	})

	t.Run("PrepareDir when path is dir", func(t *testing.T) {
		cases := []struct {
			path string
		}{
			{"./test/prepare/"},
			{"./test/prepare/dir1"},
			{"./test/prepare/dir1/dir2"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("PrepareDir on path %s expected to get nill", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Nil(t, PrepareDir(tc.path))
			})
		}
	})
	t.Run("PrepareDir when path is new", func(t *testing.T) {
		cases := []struct {
			path string
		}{
			{"./test/prepare/new1"},
			{"./test/prepare/new2/new2.1"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("PrepareDir on path %s expected to get nill", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Nil(t, PrepareDir(tc.path))
			})
		}
	})
}

func TestCreateDir(t *testing.T) {
	t.Run("createDir when path is file", func(t *testing.T) {
		cases := []struct {
			path string
		}{
			{"./test/createDir/file1"},
			{"./test/createDir/dir1/file1.1"},
			{"./test/createDir/dir1/dir2/file2.1"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("createDir on path %s expected an error", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.NotNil(t, createDir(tc.path).Error())
			})
		}
	})

	t.Run("createDir when path is dir", func(t *testing.T) {
		cases := []struct {
			path string
		}{
			{"./test/createDir/"},
			{"./test/createDir/dir1"},
			{"./test/createDir/dir1/dir2"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("createDir on path %s expected to get nill", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Nil(t, createDir(tc.path))
			})
		}
	})
	t.Run("createDir when path is new", func(t *testing.T) {
		cases := []struct {
			path string
		}{
			{"./test/createDir/new1"},
			{"./test/createDir/new2/new2.1"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("createDir on path %s expected to get nill", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Nil(t, createDir(tc.path))
			})
		}
	})
}

func TestIsDirectory(t *testing.T) {
	t.Run("isDirectory when is not dir", func(t *testing.T) {
		cases := []struct {
			path     string
			expected string
		}{
			{"./test/is-not-exist", "stat ./test/is-not-exist: no such file or directory"},
			{"./test/isDir/file1", "a given path is not a directory: ./test/isDir/file1"},
			{"./test/isDir/dir1/file1.1", "a given path is not a directory: ./test/isDir/dir1/file1.1"},
			{"./test/isDir/dir1/dir2/file2.1", "a given path is not a directory: ./test/isDir/dir1/dir2/file2.1"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("isDirectory on path %s expected an error", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Equal(t, tc.expected, isDirectory(tc.path).Error())
			})
		}
	})

	t.Run("isDirectory when is dir", func(t *testing.T) {
		cases := []struct {
			path string
		}{
			{"./test/isDir/"},
			{"./test/isDir/dir1"},
			{"./test/isDir/dir1/dir2"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("isDirectory on path %s expected to get nill", tc.path), func(t *testing.T) {
				t.Parallel()

				assert.Nil(t, isDirectory(tc.path))
			})
		}
	})
}
