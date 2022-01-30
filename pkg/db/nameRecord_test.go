package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameRecordString(t *testing.T) {
	t.Parallel()
	t.Run("String on len == 0", func(t *testing.T) {
		givenRecord := make(NameRecord)

		expect := "database does not exists\n"

		assert.Equal(t, expect, givenRecord.String())
	})

	t.Run("String on database that doesn't not have any collection", func(t *testing.T) {
		givenRecord := make(NameRecord)
		givenDB := "dbName"

		givenRecord[givenDB] = make(map[string]struct{})

		expect := "database: dbName\n None\n"
		assert.Equal(t, expect, givenRecord.String())
	})

	t.Run("String on database that have some collection", func(t *testing.T) {
		givenRecord := make(NameRecord)
		givenDB := "dbName"

		givenColl1 := "coll1"
		givenColl2 := "coll2"

		givenRecord[givenDB] = make(map[string]struct{})
		givenRecord[givenDB][givenColl1] = struct{}{}
		givenRecord[givenDB][givenColl2] = struct{}{}

		expectOne := "database: dbName\n 1 : coll1\n 2 : coll2\n"
		expectTwo := "database: dbName\n 1 : coll2\n 2 : coll1\n"
		expect := []string{expectOne, expectTwo}

		assert.Subset(t, expect, []string{givenRecord.String()})
	})
}
