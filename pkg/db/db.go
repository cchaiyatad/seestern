package db

type DBController struct {
	worker dbWorker
}

type dbWorker interface {
	ping() error
	ps(dbName string) ([]string, error)
}

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

func PS(cntStr string, vendor string, dbName string) ([]string, error) {
	controller, err := createDBController(cntStr, vendor)
	if err != nil {
		return nil, err
	}
	return controller.worker.ps(dbName)
}
