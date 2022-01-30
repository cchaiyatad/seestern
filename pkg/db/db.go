package db

import (
	"github.com/cchaiyatad/seestern/pkg/cf"
)

type DBController interface {
	ping() error
	PS(string) (NameRecord, error)
	InitConfigFile([]string, *cf.ConfigFileGenerator) error
	Insert(string, string, ...interface{}) error
	Drop(string, string) error
}

func CreateAppController(cntStr string, vendor string) (DBController, error) {
	var controller DBController
	switch vendor {
	case "mongo":
		controller = createMongoDBController(cntStr)
	default:
		return nil, &ErrVendorNotSupport{vendor}
	}

	if err := controller.ping(); err != nil {
		return nil, &ErrConnectionStringIsInvalid{err}
	}

	return controller, nil
}
