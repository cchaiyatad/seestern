package dataformat

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

func TestEncoderEncode(t *testing.T) {
	t.Parallel()
	givenData := struct {
		Name       string   `json:"name" toml:"name" yaml:"name"`
		Age        int      `json:"age" toml:"age" yaml:"age"`
		FriendList []string `json:"friend_list" toml:"friend_list" yaml:"friend_list"`
	}{
		"Bobby",
		35,
		[]string{"Tom", "Alice", "Ted"},
	}

	cases := []struct {
		fileType string
		expected string
	}{
		{"json", "{\"name\":\"Bobby\",\"age\":35,\"friend_list\":[\"Tom\",\"Alice\",\"Ted\"]}\n"},
		{"yaml", "name: Bobby\nage: 35\nfriend_list:\n    - Tom\n    - Alice\n    - Ted\n"},
		{"toml", "name = \"Bobby\"\nage = 35\nfriend_list = [\"Tom\", \"Alice\", \"Ted\"]\n"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Encode with file type %s", tc.fileType), func(t *testing.T) {
			t.Parallel()
			gotEncoder := NewEncoder(tc.fileType)

			gotEncoder.Encode(givenData)
			assert.Equal(t, tc.expected, gotEncoder.Buf.String())
		})
	}
}
