package app

import (
	"fmt"
	"path/filepath"

	"github.com/cchaiyatad/seestern/internal/file"
	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/cf"
	"github.com/cchaiyatad/seestern/pkg/db"
)

type AppController struct {
	dbController db.DBController
}

func createAppController(cntStr string, vendor string) (*AppController, error) {
	controller := AppController{}
	dbController, err := db.CreateAppController(cntStr, vendor)
	if err != nil {
		return nil, err
	}
	controller.dbController = dbController

	return &controller, nil
}

func PS(param *PSParam) (db.NameRecord, error) {
	controller, err := createAppController(param.CntStr, param.Vendor)
	if err != nil {
		return nil, err
	}
	return controller.dbController.PS(param.DBName)
}

func Init(param *InitParam) (string, error) {
	if param.isWriteFile() {
		if err := file.PrepareDir(param.Outpath); err != nil {
			return "", err
		}
	}

	controller, err := createAppController(param.CntStr, param.Vendor)
	if err != nil {
		return "", err
	}

	configGen := cf.NewConfigFileGenerator(param.FileType)
	if err := controller.dbController.InitConfigFile(param.TargetColls, configGen); err != nil {
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

	fileName := cf.GetInitFilename(param.FileType)
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

	var controller *AppController
	var err error

	if param.shouldConnectDB() {
		controller, err = createAppController(param.CntStr, param.Vendor)
		if err != nil {
			return err
		}
	}

	ssConfig, err := cf.NewConfigFileReader(param.File, param.Vendor).GetSSConfig()
	if err != nil {
		return err
	}

	result, err := ssConfig.Gen()
	if err != nil {
		return err
	}

	for db, colls := range result {
		for coll := range colls {
			if param.IsDrop {
				if err := controller.dbController.Drop(db, coll); err == nil {
					log.Logf(log.Info, "drop database %s collection %s\n", db, coll)
				} else {
					log.Logf(log.Warning, "drop database %s collection %s fail with reason %s\n", db, coll, err.Error())
				}
			}

			documents, err := result.GetDocuments(db, coll)
			if err != nil {
				log.Logf(log.Warning, "can not generate database %s collection %s with reason %s \n", db, coll, err.Error())
			}

			if param.IsInsert {
				if err := controller.dbController.Insert(db, coll, documents.ToInterfaceSlice()...); err == nil {
					log.Logf(log.Info, "insert database %s collection %s\n", db, coll)
				} else {
					log.Logf(log.Warning, "insert database %s collection %s fail with reason\n %s", db, coll, err.Error())
				}
			}

			if !param.shouldGenJson() {
				continue
			}

			outByte, err := documents.ToJson()
			if err != nil {
				log.Logf(log.Warning, "fail to display generate data of database %s collection %s\n", db, coll)
				continue
			}

			if param.Verbose {
				fmt.Printf("// database %s collection %s\n", db, coll)
				fmt.Println(string(outByte))
			}

			if param.isWriteFile() {
				fileName := cf.GetGenFilename(db, coll)
				if err := file.WriteFile(param.Outpath, fileName, outByte); err != nil {
					log.Logf(log.Warning, "can not save database %s collection %s to %s \n", db, coll, param.Outpath)
				} else {
					log.Logf(log.Info, "save database %s collection %s to %s \n", db, coll, param.Outpath)
				}
			}
		}
	}

	return nil
}
