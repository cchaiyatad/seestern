package db

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDBController(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run("createDBController with valid connection string and correct vendor (mongo)", func(t *testing.T) {
		givenCntStr := valid_mongo_cntStr
		givenVendor := "mongo"

		gotController, gotErr := createDBController(givenCntStr, givenVendor)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotController)
	})

	t.Run("createDBController with valid connection string but incorrect vendor (mongo)", func(t *testing.T) {
		givenCntStr := valid_mongo_cntStr
		givenVendor := "mysql"

		expectErr := "database vendor is not support: mysql"
		gotController, gotErr := createDBController(givenCntStr, givenVendor)

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotController)
	})

	t.Run("createDBController with invalid connection string (completely random) but correct vendor (mongo)", func(t *testing.T) {
		givenCntStr := "random_value"
		givenVendor := "mongo"

		expectErr := `can not connect to database with connection string: error parsing uri: scheme must be "mongodb" or "mongodb+srv"`
		gotController, gotErr := createDBController(givenCntStr, givenVendor)

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotController)
	})

	t.Run("createDBController with invalid connection string (in the format) but correct vendor (mongo)", func(t *testing.T) {
		givenCntStr := valid_mongo_cntStr + "123"
		givenVendor := "mongo"

		expectPrefixRegexErr := `can not connect to database with connection string: error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`
		gotController, gotErr := createDBController(givenCntStr, givenVendor)

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotController)
	})

	t.Run("createDBController with not supported vendor", func(t *testing.T) {
		givenCntStr := "random_value"
		givenVendor := "mysql"

		expectErr := "database vendor is not support: mysql"
		gotController, gotErr := createDBController(givenCntStr, givenVendor)

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotController)
	})
}

func TestDatabaseCollectionInfoString(t *testing.T) {
	t.Parallel()
	t.Run("String on len == 0", func(t *testing.T) {
		givenInfo := make(databaseCollectionInfo)

		expect := "database does not exists\n"

		assert.Equal(t, expect, givenInfo.String())
	})

	t.Run("String on database that doesn't not have any collection", func(t *testing.T) {
		givenInfo := make(databaseCollectionInfo)
		givenDB := "dbName"

		givenInfo[givenDB] = make(map[string]struct{})

		expect := "database: dbName\n None\n"
		assert.Equal(t, expect, givenInfo.String())
	})

	t.Run("String on database that have some collection", func(t *testing.T) {
		givenInfo := make(databaseCollectionInfo)
		givenDB := "dbName"

		givenColl1 := "coll1"
		givenColl2 := "coll2"

		givenInfo[givenDB] = make(map[string]struct{})
		givenInfo[givenDB][givenColl1] = struct{}{}
		givenInfo[givenDB][givenColl2] = struct{}{}

		expect := "database: dbName\n 1 : coll1\n 2 : coll2\n"
		assert.Equal(t, expect, givenInfo.String())
	})
}
