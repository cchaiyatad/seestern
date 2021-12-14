package db

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cchaiyatad/seestern/internal/file"
	"github.com/cchaiyatad/seestern/pkg/cf"
)

type DBController struct {
	worker dbWorker
}

type dbWorker interface {
	ping() error
	ps(string) (databaseCollectionInfo, error)
	initConfigFile(*InitParam, *cf.ConfigFileGenerator) error
	insert()
	drop()
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
	if param.Outpath != "" {
		if err := file.PrepareDir(param.Outpath); err != nil {
			return "", err
		}
	}

	controller, err := createDBController(param.CntStr, param.Vendor)
	if err != nil {
		return "", err
	}

	configGen := cf.NewConfigFileGenerator(param.FileType)
	if err := controller.worker.initConfigFile(param, configGen); err != nil {
		return "", err
	}

	go func() {
		for range configGen.OutChan {
			configGen.Done()
		}

	}()

	configGen.Wait()
	configGen.Close()

	configByte, err := configGen.Bytes()
	if err != nil {
		return "", err
	}

	if param.Verbose {
		fmt.Print(string(configByte))
	}

	if param.Outpath == "" {
		return "", nil
	}

	fileName := cf.GetFilename(param.FileType)
	if err := file.WriteFile(param.Outpath, fileName, configByte); err != nil {
		return "", nil
	}

	return filepath.Join(param.Outpath, fileName), nil
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
