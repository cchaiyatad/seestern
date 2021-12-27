package cf

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseGetRef(t *testing.T) {
	cases := []struct {
		givenPath     string
		expectedRefs1 []string
		expectedRefs2 []string
	}{
		{filepath.FromSlash("./test/01_simple_ref.ss.toml"), []string{}, []string{"bookstore.publisher._id"}},
		{filepath.FromSlash("./test/02_two_simple_ref.ss.toml"), []string{}, []string{"bookstore.publisher._id", "bookstore.publisher.name"}},
		{filepath.FromSlash("./test/03_array_ref.ss.toml"), []string{}, []string{"bookstore.publisher._id"}},
		{filepath.FromSlash("./test/04_object_ref.ss.toml"), []string{}, []string{"bookstore.publisher._id", "bookstore.publisher.name"}},
		{filepath.FromSlash("./test/05_with_invalid_ref.ss.toml"), []string{}, []string{"bookstore.publisher._id"}},
		{filepath.FromSlash("./test/06_self_ref.ss.toml"), []string{}, []string{"bookstore.book._id"}},
		{filepath.FromSlash("./test/07_not_exist_ref.ss.toml"), []string{}, []string{"bookstore.customer._id"}},
		{filepath.FromSlash("./test/08_cyclic_ref.ss.toml"), []string{"bookstore.book._id"}, []string{"bookstore.publisher._id"}},
		{filepath.FromSlash("./test/09_simple_ref_ref_come_first.ss.toml"), []string{"bookstore.publisher._id"}, []string{}},
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
