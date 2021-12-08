package db

import (
	"fmt"
	"strings"
)

type DBController struct {
	worker dbWorker
}

type dbWorker interface {
	ping() error
	ps(string) (databaseCollectionInfo, error)
	initConfigFile(*InitParam) (string, error)
}

type databaseCollectionInfo map[string]map[string]struct{}

func createDBController(cntStr string, vendor string) (*DBController, error) {
	controller := DBController{}
	switch vendor {
	case "mongo":
		controller.worker = createMongoDBWorker(cntStr)
	default:
		return nil, &ErrVendorNotSupport{vendor}
	}

	if err := controller.worker.ping(); err != nil {
		return nil, &ErrConnectionStringIsInvalid{err}
	}

	return &controller, nil
}

func PS(param *PSParam) (databaseCollectionInfo, error) {
	controller, err := createDBController(param.CntStr, param.Vendor)
	if err != nil {
		return nil, err
	}
	return controller.worker.ps(param.DBName)
}

func Init(param *InitParam) (string, error) {
	controller, err := createDBController(param.CntStr, param.Vendor)
	if err != nil {
		return "", err
	}
	return controller.worker.initConfigFile(param)

}

func (info databaseCollectionInfo) String() string {
	if len(info) == 0 {
		return "database does not exists\n"
	}

	var strBuilder strings.Builder

	for dbName, collNames := range info {
		fmt.Fprintf(&strBuilder, "database: %s\n", dbName)

		if len(collNames) == 0 {
			strBuilder.WriteString(" None\n")
			continue
		}

		idx := 1
		for collName := range collNames {
			fmt.Fprintf(&strBuilder, " %-2d: %s\n", idx, collName)
			idx += 1
		}
	}
	return strBuilder.String()
}
