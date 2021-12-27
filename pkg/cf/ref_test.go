package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseGetRef(t *testing.T) {

	cases := []struct {
		givenPath     string
		expectedRefs1 []string
		expectedRefs2 []string
	}{
		{"./test/01_simple_ref.ss.toml", []string{}, []string{"bookstore.publisher._id"}},
		{"./test/02_two_simple_ref.ss.toml", []string{}, []string{"bookstore.publisher._id", "bookstore.publisher.name"}},
		{"./test/03_array_ref.ss.toml", []string{}, []string{"bookstore.publisher._id"}},
		{"./test/04_object_ref.ss.toml", []string{}, []string{"bookstore.publisher._id", "bookstore.publisher.name"}},
		{"./test/05_with_invalid_ref.ss.toml", []string{}, []string{"bookstore.publisher._id"}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("GetRef for file %s", tc.givenPath), func(t *testing.T) {
			t.Parallel()
			gotSSConfig, gotErr := NewConfigFileReader(tc.givenPath, "").GetSSConfig()
			assert.Nil(t, gotErr)

			gotRefs1 := gotSSConfig.Databases[0].GetRef()
			gotRefs2 := gotSSConfig.Databases[1].GetRef()

			assert.Equal(t, tc.expectedRefs1, gotRefs1)
			assert.Equal(t, tc.expectedRefs2, gotRefs2)
		})
	}

}
