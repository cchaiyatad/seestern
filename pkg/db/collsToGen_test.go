package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCollectionInputFromArgs(t *testing.T) {
	toString := func(c collsToGen) string {
		return fmt.Sprintf("%T", c)
	}

	t.Parallel()

	t.Run("ParseCollectionInputFromArgs on empty slice", func(t *testing.T) {
		givenArgs := []string{}

		expected := map[string][]string{}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})

	t.Run("ParseCollectionInputFromArgs on one correct item", func(t *testing.T) {
		givenArgs := []string{"db1.coll1"}

		expected := map[string][]string{"db1": {"coll1"}}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})

	t.Run("ParseCollectionInputFromArgs on two correct item with same db", func(t *testing.T) {
		givenArgs := []string{"db1.coll1", "db1.coll2"}

		expected := map[string][]string{"db1": {"coll1", "coll2"}}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})

	t.Run("ParseCollectionInputFromArgs on two correct item with same db same coll", func(t *testing.T) {
		givenArgs := []string{"db1.coll1", "db1.coll1"}

		expected := map[string][]string{"db1": {"coll1"}}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})

	t.Run("ParseCollectionInputFromArgs on two correct item with same coll", func(t *testing.T) {
		givenArgs := []string{"db1.coll1", "db2.coll1"}

		expected := map[string][]string{"db1": {"coll1"}, "db2": {"coll1"}}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})

	t.Run("ParseCollectionInputFromArgs on incorrect format", func(t *testing.T) {
		givenArgs := []string{"db1coll1"}

		expected := map[string][]string{}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})
	t.Run("ParseCollectionInputFromArgs on two dot", func(t *testing.T) {
		// collection name should not be two dot to avoid ambigious
		givenArgs := []string{"db1.coll1.sub"}

		expected := map[string][]string{"db1": {"coll1"}}
		got := parseCollectionInputFromArgs(givenArgs)

		assert.Equal(t, toString(expected), toString(got), fmt.Sprintf("expected: %s got: %s", expected, got))
	})

}
