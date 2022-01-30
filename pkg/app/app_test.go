package app

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAppController(t *testing.T) {
	t.Parallel()
	const valid_mongo_cntStr = "mongodb+srv://testReadOnly:testSeeStern@ps-command-cluster.h0n2k.mongodb.net"

	t.Run("createAppController with valid connection string and correct vendor (mongo)", func(t *testing.T) {
		givenCntStr := valid_mongo_cntStr
		givenVendor := "mongo"

		gotController, gotErr := createAppController(givenCntStr, givenVendor)

		assert.Nil(t, gotErr)
		assert.NotNil(t, gotController)
	})

	t.Run("createAppController with valid connection string but incorrect vendor (mongo)", func(t *testing.T) {
		givenCntStr := valid_mongo_cntStr
		givenVendor := "mysql"

		expectErr := "database vendor is not support: mysql"
		gotController, gotErr := createAppController(givenCntStr, givenVendor)

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotController)
	})

	t.Run("createAppController with invalid connection string (completely random) but correct vendor (mongo)", func(t *testing.T) {
		givenCntStr := "random_value"
		givenVendor := "mongo"

		expectErr := `can not connect to database with connection string: error parsing uri: scheme must be "mongodb" or "mongodb+srv"`
		gotController, gotErr := createAppController(givenCntStr, givenVendor)

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotController)
	})

	t.Run("createAppController with invalid connection string (in the format) but correct vendor (mongo)", func(t *testing.T) {
		givenCntStr := valid_mongo_cntStr + "123"
		givenVendor := "mongo"

		expectPrefixRegexErr := `can not connect to database with connection string: error parsing uri: lookup _mongodb._tcp.ps-command-cluster.h0n2k.mongodb.net123`
		gotController, gotErr := createAppController(givenCntStr, givenVendor)

		assert.Regexp(t, regexp.MustCompile(expectPrefixRegexErr), gotErr.Error())
		assert.Nil(t, gotController)
	})

	t.Run("createAppController with not supported vendor", func(t *testing.T) {
		givenCntStr := "random_value"
		givenVendor := "mysql"

		expectErr := "database vendor is not support: mysql"
		gotController, gotErr := createAppController(givenCntStr, givenVendor)

		assert.Equal(t, expectErr, gotErr.Error())
		assert.Nil(t, gotController)
	})
}
