package db

import "fmt"

type DBController struct {
	worker dbWorker
}

type dbWorker interface {
	ping() error
	ps(string) (databaseCollectionInfo, error)
}

type databaseCollectionInfo map[string][]string

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

func PS(cntStr string, vendor string, dbName string) (databaseCollectionInfo, error) {
	controller, err := createDBController(cntStr, vendor)
	if err != nil {
		return nil, err
	}
	return controller.worker.ps(dbName)
}

func (info databaseCollectionInfo) String() string {
	if len(info) == 0 {
		return "there are no exist database\n"
	}

	for dbName, collNames := range info {
		fmt.Printf("database: %s\n", dbName)
		if len(collNames) == 0 {
			fmt.Println(" None")
		}

		for idx, collName := range collNames {
			fmt.Printf(" %-2d: %s\n", idx+1, collName)
		}
	}
	return ""
}
