package db

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run("ping with valid connection string", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr}

		gotErr := givenController.ping()

		assert.Nil(t, gotErr)
	})

	t.Run("ping with invalid connection string (random)", func(t *testing.T) {
		givenController := mongoDBWorker{"random"}

		expectErr := `error parsing uri: scheme must be "mongodb" or "mongodb+srv"`
		gotErr := givenController.ping()

		assert.Equal(t, expectErr, gotErr.Error())
	})

	t.Run("ping with invalid connection string (informat)", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr + "123"}

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`
		gotErr := givenController.ping()

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
	})
}

func TestGetDatabaseCollectionInfo(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"
	t.Run("getDatabaseCollectionInfo with valid connection string", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr}

		expected := databaseCollectionInfo{"admin": map[string]struct{}{}, "local": map[string]struct{}{"clustermanager": {}, "oplog.rs": {}, "replset.election": {}, "replset.initialSyncId": {}, "replset.minvalid": {}, "replset.oplogTruncateAfterPoint": {}, "startup_log": {}}, "sample_airbnb": map[string]struct{}{"listingsAndReviews": {}}, "sample_analytics": map[string]struct{}{"accounts": {}, "customers": {}, "transactions": {}}, "sample_geospatial": map[string]struct{}{"shipwrecks": {}}, "sample_mflix": map[string]struct{}{"comments": {}, "movies": {}, "sessions": {}, "theaters": {}, "users": {}}, "sample_restaurants": map[string]struct{}{"neighborhoods": {}, "restaurants": {}}, "sample_supplies": map[string]struct{}{"sales": {}}, "sample_training": map[string]struct{}{"companies": {}, "grades": {}, "inspections": {}, "posts": {}, "routes": {}, "trips": {}, "zips": {}}, "sample_weatherdata": map[string]struct{}{"data": {}}}

		gotInfo, gotErr := givenController.getDatabaseCollectionInfo()

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotInfo)
	})

	t.Run("getDatabaseCollectionInfo with invalid connection string", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr + "123"}

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`

		gotInfo, gotErr := givenController.getDatabaseCollectionInfo()

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotInfo)
	})
}

func TestPS(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run(`ps with valid connection string dbNameFilter = ""`, func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr}
		givenDBNameFilter := ""

		expected := databaseCollectionInfo{"admin": map[string]struct{}{}, "local": map[string]struct{}{"clustermanager": {}, "oplog.rs": {}, "replset.election": {}, "replset.initialSyncId": {}, "replset.minvalid": {}, "replset.oplogTruncateAfterPoint": {}, "startup_log": {}}, "sample_airbnb": map[string]struct{}{"listingsAndReviews": {}}, "sample_analytics": map[string]struct{}{"accounts": {}, "customers": {}, "transactions": {}}, "sample_geospatial": map[string]struct{}{"shipwrecks": {}}, "sample_mflix": map[string]struct{}{"comments": {}, "movies": {}, "sessions": {}, "theaters": {}, "users": {}}, "sample_restaurants": map[string]struct{}{"neighborhoods": {}, "restaurants": {}}, "sample_supplies": map[string]struct{}{"sales": {}}, "sample_training": map[string]struct{}{"companies": {}, "grades": {}, "inspections": {}, "posts": {}, "routes": {}, "trips": {}, "zips": {}}, "sample_weatherdata": map[string]struct{}{"data": {}}}

		gotInfo, gotErr := givenController.ps(givenDBNameFilter)

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotInfo)
	})

	t.Run("ps with valid connection string dbNameFilter is the database that exist", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr}
		givenDBNameFilter := "local"

		expected := databaseCollectionInfo{"local": map[string]struct{}{"clustermanager": {}, "oplog.rs": {}, "replset.election": {}, "replset.initialSyncId": {}, "replset.minvalid": {}, "replset.oplogTruncateAfterPoint": {}, "startup_log": {}}}

		gotInfo, gotErr := givenController.ps(givenDBNameFilter)

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotInfo)
	})

	t.Run("ps with valid connection string dbNameFilter is the database that not exist", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr}
		givenDBNameFilter := "not-exists"

		expected := databaseCollectionInfo{}

		gotInfo, gotErr := givenController.ps(givenDBNameFilter)

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotInfo)
	})

	t.Run("ps with invalid connection string ", func(t *testing.T) {
		givenController := mongoDBWorker{valid_mongo_cntStr + "123"}
		givenDBNameFilter := ""

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`

		gotInfo, gotErr := givenController.ps(givenDBNameFilter)

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotInfo)
	})

}

func TestGetCursor(t *testing.T) {
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"
	givenValidController := mongoDBWorker{valid_mongo_cntStr}
	givenValidClient, err := givenValidController.connect()

	assert.Nil(t, err)
	t.Parallel()

	t.Run("with valid client, valid database and valid collection", func(t *testing.T) {
		givenDB := "sample_mflix"
		givenColl := "users"

		gotCursor, gotErr := givenValidController.getCursor(givenValidClient, givenDB, givenColl)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotCursor)
	})

	t.Run("with valid client, valid database and invalid collection (cursor with zero document)", func(t *testing.T) {
		givenDB := "sample_mflix"
		givenColl := "not-exist"

		gotCursor, gotErr := givenValidController.getCursor(givenValidClient, givenDB, givenColl)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotCursor)
	})

	t.Run("with nil client, valid database and valid collection", func(t *testing.T) {
		givenDB := "sample_mflix"
		givenColl := "users"

		expectedErr := ErrClientIsNil
		gotCursor, gotErr := givenValidController.getCursor(nil, givenDB, givenColl)

		assert.Equal(t, expectedErr, gotErr)
		assert.Nil(t, gotCursor)
	})

	t.Run("with invalid client, valid database and valid collection", func(t *testing.T) {
		givenInValidController := mongoDBWorker{valid_mongo_cntStr + "123"}
		givenInValidClient, err := givenInValidController.connect()
		assert.NotNil(t, err)
		assert.Nil(t, givenInValidClient)

		givenDB := "sample_mflix"
		givenColl := "users"

		expectedErr := ErrClientIsNil
		gotCursor, gotErr := givenInValidController.getCursor(givenInValidClient, givenDB, givenColl)

		assert.Equal(t, expectedErr, gotErr)
		assert.Nil(t, gotCursor)
	})
}

func TestInit(t *testing.T) {
}
