package db

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run("connect with valid connection string", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr}

		gotClinet, gotErr := givenController.connect()

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotClinet)

		gotSecondClient, gotErr := givenController.connect()
		assert.Nil(t, gotErr)
		assert.Equal(t, gotClinet, gotSecondClient)
	})

	t.Run("connect with invalid connection string (random)", func(t *testing.T) {
		givenController := mongoDBController{cntStr: "random"}

		expectErr := `error parsing uri: scheme must be "mongodb" or "mongodb+srv"`
		gotClinet, gotErr := givenController.connect()

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotClinet)
	})

	t.Run("connect with invalid connection string (in format)", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr + "123"}

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`
		gotClinet, gotErr := givenController.connect()

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotClinet)
	})

}
func TestPing(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run("ping with valid connection string", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr}

		gotErr := givenController.ping()

		assert.Nil(t, gotErr)
	})

	t.Run("ping with invalid connection string (random)", func(t *testing.T) {
		givenController := mongoDBController{cntStr: "random"}

		expectErr := `error parsing uri: scheme must be "mongodb" or "mongodb+srv"`
		gotErr := givenController.ping()

		assert.Equal(t, expectErr, gotErr.Error())
	})

	t.Run("ping with invalid connection string (in format)", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr + "123"}

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`
		gotErr := givenController.ping()

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
	})
}

func TestGetNameRecord(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"
	t.Run("getNameRecord with valid connection string", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr}

		expected := NameRecord{"admin": map[string]struct{}{}, "local": map[string]struct{}{"clustermanager": {}, "oplog.rs": {}, "replset.election": {}, "replset.initialSyncId": {}, "replset.minvalid": {}, "replset.oplogTruncateAfterPoint": {}, "startup_log": {}}, "sample_airbnb": map[string]struct{}{"listingsAndReviews": {}}, "sample_analytics": map[string]struct{}{"accounts": {}, "customers": {}, "transactions": {}}, "sample_geospatial": map[string]struct{}{"shipwrecks": {}}, "sample_mflix": map[string]struct{}{"comments": {}, "movies": {}, "sessions": {}, "theaters": {}, "users": {}}, "sample_restaurants": map[string]struct{}{"neighborhoods": {}, "restaurants": {}}, "sample_supplies": map[string]struct{}{"sales": {}}, "sample_training": map[string]struct{}{"companies": {}, "grades": {}, "inspections": {}, "posts": {}, "routes": {}, "trips": {}, "zips": {}}, "sample_weatherdata": map[string]struct{}{"data": {}}}

		gotRecord, gotErr := givenController.getNameRecord()

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotRecord)
	})

	t.Run("getNameRecord with invalid connection string", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr + "123"}

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`

		gotRecord, gotErr := givenController.getNameRecord()

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotRecord)
	})
}

func TestPS(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run(`ps with valid connection string dbNameFilter = ""`, func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr}
		givenDBNameFilter := ""

		expected := NameRecord{"admin": map[string]struct{}{}, "local": map[string]struct{}{"clustermanager": {}, "oplog.rs": {}, "replset.election": {}, "replset.initialSyncId": {}, "replset.minvalid": {}, "replset.oplogTruncateAfterPoint": {}, "startup_log": {}}, "sample_airbnb": map[string]struct{}{"listingsAndReviews": {}}, "sample_analytics": map[string]struct{}{"accounts": {}, "customers": {}, "transactions": {}}, "sample_geospatial": map[string]struct{}{"shipwrecks": {}}, "sample_mflix": map[string]struct{}{"comments": {}, "movies": {}, "sessions": {}, "theaters": {}, "users": {}}, "sample_restaurants": map[string]struct{}{"neighborhoods": {}, "restaurants": {}}, "sample_supplies": map[string]struct{}{"sales": {}}, "sample_training": map[string]struct{}{"companies": {}, "grades": {}, "inspections": {}, "posts": {}, "routes": {}, "trips": {}, "zips": {}}, "sample_weatherdata": map[string]struct{}{"data": {}}}

		gotRecord, gotErr := givenController.PS(givenDBNameFilter)

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotRecord)
	})

	t.Run("ps with valid connection string dbNameFilter is the database that exist", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr}
		givenDBNameFilter := "local"

		expected := NameRecord{"local": map[string]struct{}{"clustermanager": {}, "oplog.rs": {}, "replset.election": {}, "replset.initialSyncId": {}, "replset.minvalid": {}, "replset.oplogTruncateAfterPoint": {}, "startup_log": {}}}

		gotRecord, gotErr := givenController.PS(givenDBNameFilter)

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotRecord)
	})

	t.Run("ps with valid connection string dbNameFilter is the database that not exist", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr}
		givenDBNameFilter := "not-exists"

		expected := NameRecord{}

		gotRecord, gotErr := givenController.PS(givenDBNameFilter)

		assert.Nil(t, gotErr)
		assert.Equal(t, expected, gotRecord)
	})

	t.Run("ps with invalid connection string ", func(t *testing.T) {
		givenController := mongoDBController{cntStr: valid_mongo_cntStr + "123"}
		givenDBNameFilter := ""

		expectPrefixRegexErr := `error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`

		gotRecord, gotErr := givenController.PS(givenDBNameFilter)

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotRecord)
	})

}

func TestGetCursor(t *testing.T) {
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"
	givenValidController := mongoDBController{cntStr: valid_mongo_cntStr}

	t.Parallel()

	t.Run("with valid client, valid database and valid collection", func(t *testing.T) {
		givenDB := "sample_mflix"
		givenColl := "users"

		gotCursor, gotErr := givenValidController.getCursor(givenDB, givenColl)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotCursor)
	})

	t.Run("with valid client, valid database and invalid collection (cursor with zero document)", func(t *testing.T) {
		givenDB := "sample_mflix"
		givenColl := "not-exist"

		gotCursor, gotErr := givenValidController.getCursor(givenDB, givenColl)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotCursor)
	})

	t.Run("with nil client, valid database and valid collection", func(t *testing.T) {
		givenDB := "sample_mflix"
		givenColl := "users"

		givenNewValidController := mongoDBController{cntStr: valid_mongo_cntStr}
		_, err := givenNewValidController.connect()
		assert.Nil(t, err)
		givenNewValidController.client = nil
		assert.Nil(t, givenNewValidController.client)

		expectedErr := ErrClientIsNil
		gotCursor, gotErr := givenNewValidController.getCursor(givenDB, givenColl)

		assert.Equal(t, expectedErr, gotErr)
		assert.Nil(t, gotCursor)
	})

	t.Run("with invalid client, valid database and valid collection", func(t *testing.T) {
		givenInValidController := mongoDBController{cntStr: valid_mongo_cntStr + "123"}

		givenDB := "sample_mflix"
		givenColl := "users"

		expectPrefixRegexErr := "error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123 on 192.168.1.1:53: no such host"
		gotCursor, gotErr := givenInValidController.getCursor(givenDB, givenColl)

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotCursor)
	})
}

func TestInit(t *testing.T) {
}
