package cf

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncoder(t *testing.T) {
	cases := []struct {
		fileType        string
		expectedEncType string
	}{
		{"json", "*json.Encoder"},
		{"yaml", "*yaml.Encoder"},
		{"toml", "*toml.Encoder"},
		{"jSOn", "*json.Encoder"},
		{"YAML", "*yaml.Encoder"},
		{"YAML2", "*yaml.Encoder"},
		{"toMl", "*toml.Encoder"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("NewEncoder with file type %s expected to get enc with type %s", tc.fileType, tc.expectedEncType), func(t *testing.T) {
			t.Parallel()
			gotEncoder := NewEncoder(tc.fileType)
			gotEncType := reflect.TypeOf(gotEncoder.enc)
			assert.Equal(t, tc.expectedEncType, gotEncType.String())
		})
	}
}
