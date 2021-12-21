package db

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cchaiyatad/seestern/internal/file"
	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/cf"
)

type DBController struct {
	worker dbWorker
}

type dbWorker interface {
	ping() error
	ps(string) (nameRecord, error)
	initConfigFile(*InitParam, *cf.ConfigFileGenerator) error
	insert(string, string, ...interface{}) error
	drop(string, string) error
}

type nameRecord map[string]map[string]struct{}

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

func PS(param *PSParam) (nameRecord, error) {
	controller, err := createDBController(param.CntStr, param.Vendor)
	if err != nil {
		return nil, err
	}
	return controller.worker.ps(param.DBName)
}

func Init(param *InitParam) (string, error) {
	if param.isWriteFile() {
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

	if !param.isWriteFile() {
		return "", nil
	}

	fileName := cf.GetFilename(param.FileType)
	if err := file.WriteFile(param.Outpath, fileName, configByte); err != nil {
		return "", nil
	}

	return filepath.Join(param.Outpath, fileName), nil
}

func Gen(param *GenParam) error {
	if param.isWriteFile() {
		if err := file.PrepareDir(param.Outpath); err != nil {
			return err
		}
	}

	var controller *DBController
	var err error

	if param.shouldConnectDB() {
		controller, err = createDBController(param.CntStr, param.Vendor)
		if err != nil {
			return err
		}
	}

	ssConfig, err := cf.NewConfigFileReader(param.File).GetSSConfig()
	if err != nil {
		return err
	}

	result := ssConfig.Gen()

	for db, colls := range result {
		for coll := range colls {
			if param.IsDrop {
				controller.worker.drop(db, coll)
				log.Logf(log.Info, "drop database %s collection %s\n", db, coll)
			}

			documents, err := result.GetDocuments(db, coll)
			if err != nil {
				log.Logf(log.Warning, "can not generate database %s collection %s with reason %s \n", db, coll, err.Error())
			}

			if param.IsInsert {
				controller.worker.insert(db, coll, documents.ToInterfaceSlice()...)
				log.Logf(log.Info, "insert database %s collection %s\n", db, coll)
			}

			if !param.shouldGenJson() {
				continue
			}

			json, err := documents.ToJson()
			if err != nil {
				log.Logf(log.Warning, "fail to display generate data of database %s collection %s\n", db, coll)
				continue
			}

			if param.Verbose {
				fmt.Printf("database %s collection %s\n", db, coll)
				fmt.Println(string(json))
			}

			if param.isWriteFile() {
				// TODO: save documents to json
				log.Logf(log.Info, "save database %s collection %s to %s \n", db, coll, param.Outpath)
			}
		}
	}

	return nil
}

func (record nameRecord) empty() bool {
	return len(record) == 0
}

func (record nameRecord) String() string {
	if record.empty() {
		return "database does not exists\n"
	}

	var strBuilder strings.Builder

	for dbName, collNames := range record {
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
