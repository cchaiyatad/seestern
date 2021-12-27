package cf

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenOrderSimple(t *testing.T) {
	cases := []struct {
		givenPath      string
		exptectedOrder []int
	}{
		{filepath.FromSlash("./test/01_simple_ref.ss.toml"), []int{0, 1}},
		{filepath.FromSlash("./test/02_two_simple_ref.ss.toml"), []int{0, 1}},
		{filepath.FromSlash("./test/03_array_ref.ss.toml"), []int{0, 1}},
		{filepath.FromSlash("./test/04_object_ref.ss.toml"), []int{0, 1}},
		{filepath.FromSlash("./test/09_simple_ref_ref_come_first.ss.toml"), []int{1, 0}},
		{filepath.FromSlash("./test/10_complex_ref.ss.toml"), []int{0, 3, 1, 2}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("NewGenOrder for file %s", tc.givenPath), func(t *testing.T) {
			t.Parallel()
			gotSSConfig, gotErr := NewConfigFileReader(tc.givenPath, "").GetSSConfig()
			assert.Nil(t, gotErr)
			gotGenOrder, gotErr := gotSSConfig.NewGenOrder()
			assert.Nil(t, gotErr)
			assert.Equal(t, tc.exptectedOrder, gotGenOrder.order)
		})
	}

}
func TestNewGenOrderWithErr(t *testing.T) {
	cases := []struct {
		givenPath   string
		expectedErr error
	}{
		{filepath.FromSlash("./test/05_with_invalid_ref.ss.toml"), nil}, // it is filltered out before find order
		{filepath.FromSlash("./test/06_self_ref.ss.toml"), ErrSelfReference},
		{filepath.FromSlash("./test/07_not_exist_ref.ss.toml"), ErrRefToNotExist},
		{filepath.FromSlash("./test/08_cyclic_ref.ss.toml"), ErrRefIsCyclic},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("NewGenOrder for file %s", tc.givenPath), func(t *testing.T) {
			t.Parallel()
			gotSSConfig, gotErr := NewConfigFileReader(tc.givenPath, "").GetSSConfig()
			assert.Nil(t, gotErr)
			_, gotErr = gotSSConfig.NewGenOrder()
			assert.Equal(t, tc.expectedErr, gotErr)
		})
	}

}

func TestGenOrderIterate(t *testing.T) {
	cases := []struct {
		givenPath     string
		expectedOrder []string
	}{
		{filepath.FromSlash("./test/01_simple_ref.ss.toml"), []string{"bookstore.publisher", "bookstore.book"}},
		{filepath.FromSlash("./test/02_two_simple_ref.ss.toml"), []string{"bookstore.publisher", "bookstore.book"}},
		{filepath.FromSlash("./test/03_array_ref.ss.toml"), []string{"bookstore.publisher", "bookstore.book"}},
		{filepath.FromSlash("./test/04_object_ref.ss.toml"), []string{"bookstore.publisher", "bookstore.book"}},
		{filepath.FromSlash("./test/05_with_invalid_ref.ss.toml"), []string{"bookstore.publisher", "bookstore.book"}},
		{filepath.FromSlash("./test/06_self_ref.ss.toml"), []string{}},
		{filepath.FromSlash("./test/07_not_exist_ref.ss.toml"), []string{}},
		{filepath.FromSlash("./test/08_cyclic_ref.ss.toml"), []string{}},
		{filepath.FromSlash("./test/09_simple_ref_ref_come_first.ss.toml"), []string{"bookstore.publisher", "bookstore.book"}},
		{filepath.FromSlash("./test/10_complex_ref.ss.toml"), []string{"db.A", "db.D", "db.B", "db.C"}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Iterate for file %s", tc.givenPath), func(t *testing.T) {
			t.Parallel()
			gotSSConfig, gotErr := NewConfigFileReader(tc.givenPath, "").GetSSConfig()
			assert.Nil(t, gotErr)
			gotGenOrder, _ := gotSSConfig.NewGenOrder()

			givenTestIterate := make([]string, 0)
			givenFunc := func(db *Database) {
				givenTestIterate = append(givenTestIterate, db.CreateNym())
			}

			gotGenOrder.IterateDB(givenFunc)
			assert.Equal(t, tc.expectedOrder, givenTestIterate)
		})
	}
}
