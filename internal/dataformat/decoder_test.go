package dataformat

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecoder(t *testing.T) {
	t.Parallel()

	t.Run("NewDecoder with valid file", func(t *testing.T) {
		cases := []struct {
			filePath            string
			expectedDecoderType string
			expectedData        string
		}{
			{"./test/decoder/file.json", "*dataformat.jsonUnmarshaler", "{\n    \"Age\": 25,\n    \"Cats\": [\n        \"Cauchy\",\n        \"Plato\"\n    ],\n    \"Perfection\": [\n        6,\n        28,\n        496,\n        8128\n    ],\n    \"Pi\": 3.14\n}"},
			{"./test/decoder/file.yaml", "*dataformat.yamlUnmarshaler", "age: 25\ncats:\n  - Cauchy\n  - Plato\npi: 3.14\nperfection:\n  - 6\n  - 28\n  - 496\n  - 8128"},
			{"./test/decoder/file.toml", "*dataformat.tomlUnmarshaler", "Age = 25\nCats = [ \"Cauchy\", \"Plato\" ]\nPi = 3.14\nPerfection = [ 6, 28, 496, 8128 ]"},
			{"./test/decoder/file.ss.json", "*dataformat.jsonUnmarshaler", "{\n    \"Age\": 25,\n    \"Cats\": [\n        \"Cauchy\",\n        \"Plato\"\n    ],\n    \"Perfection\": [\n        6,\n        28,\n        496,\n        8128\n    ],\n    \"Pi\": 3.14\n}"},
			{"./test/decoder/file.ss.yaml", "*dataformat.yamlUnmarshaler", "age: 25\ncats:\n  - Cauchy\n  - Plato\npi: 3.14\nperfection:\n  - 6\n  - 28\n  - 496\n  - 8128"},
			{"./test/decoder/file.ss.toml", "*dataformat.tomlUnmarshaler", "Age = 25\nCats = [ \"Cauchy\", \"Plato\" ]\nPi = 3.14\nPerfection = [ 6, 28, 496, 8128 ]"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("NewDecoder with file type %s expected to get dec with type %s", tc.filePath, tc.expectedDecoderType), func(t *testing.T) {
				t.Parallel()
				gotDecoder, gotErr := NewDecoder(tc.filePath)

				assert.Nil(t, gotErr)

				gotDecType := reflect.TypeOf(gotDecoder.dec)
				assert.Equal(t, tc.expectedDecoderType, gotDecType.String())
				assert.Equal(t, tc.expectedData, string(gotDecoder.data))
			})
		}
	})

	t.Run("NewDecoder with invalid file", func(t *testing.T) {
		cases := []struct {
			filePath    string
			expectedErr string
		}{
			{"./test/decoder/dir", "a given path is not a file: ./test/decoder/dir"},
			{"./test/decoder/file", "error: not support file type (support only .json .toml .yaml)"},
			{"./test/decoder/filejson", "error: not support file type (support only .json .toml .yaml)"},
			{"./test/decoder/file-not-exist", "stat ./test/decoder/file-not-exist: no such file or directory"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("NewDecoder with file type %s expected to get error", tc.filePath), func(t *testing.T) {
				t.Parallel()
				gotDecoder, gotErr := NewDecoder(tc.filePath)

				assert.Nil(t, gotDecoder)
				assert.Equal(t, tc.expectedErr, gotErr.Error())
			})
		}
	})
}

func TestDecoderDecode(t *testing.T) {
	t.Parallel()

	type data struct {
		Age        int
		Cats       []string
		Pi         float64
		Perfection []int
	}

	t.Run("Decode without option", func(t *testing.T) {
		expectedData := data{Age: 25, Cats: []string{"Cauchy", "Plato"}, Pi: 3.14, Perfection: []int{6, 28, 496, 8128}}
		cases := []struct {
			filePath string
		}{
			{"./test/decoder/file.json"},
			{"./test/decoder/file.toml"},
			{"./test/decoder/file.yaml"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("Decode with file %s", tc.filePath), func(t *testing.T) {
				gotDecoder, gotErr := NewDecoder(tc.filePath)
				assert.Nil(t, gotErr)

				var data data

				gotErr = gotDecoder.Decode(&data)
				assert.Nil(t, gotErr)
				assert.Equal(t, expectedData, data)
			})
		}
	})

	t.Run("Decode with option", func(t *testing.T) {})
}
