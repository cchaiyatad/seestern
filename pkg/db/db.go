package db

import "fmt"

type DBController struct {
	worker dbWorker
}

type dbWorker interface {
	ping() error
	ps(string) (databaseCollectionInfo, error)
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

func PS(cntStr string, vendor string, dbName string) (databaseCollectionInfo, error) {
	controller, err := createDBController(cntStr, vendor)
	if err != nil {
		return nil, err
	}
	return controller.worker.ps(dbName)
}

func (info databaseCollectionInfo) String() string {
	if len(info) == 0 {
		return "database not exists\n"
	}

	for dbName, collNames := range info {
		fmt.Printf("database: %s\n", dbName)
		if len(collNames) == 0 {
			fmt.Println(" None")
			continue
		}

		idx := 1
		for collName := range collNames {
			fmt.Printf(" %-2d: %s\n", idx, collName)
			idx += 1
		}
	}
	return ""
}
