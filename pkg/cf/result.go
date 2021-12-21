package cf

import (
	"errors"
	"fmt"
)

// db[coll]
type result map[string]map[string]interface{}

func (ssconfig *SSConfig) NewResult() result {
	info := make(result)

	for _, db := range ssconfig.Databases {
		dbName := db.D_name
		collName := db.Collection.C_name
		if _, ok := info[dbName]; !ok {
			info[dbName] = make(map[string]interface{})
		}
		info[dbName][collName] = struct{}{}
	}

	return info
}

var ErrDBCollNotExist = errors.New("database or collection name doesn't exist in result")
var ErrUnknownError = errors.New("got an unknown error")

func (info result) GetDocuments(dbName, collName string) ([]document, error) {
	var value interface{}
	var ok bool
	if value, ok = info[dbName][collName]; !ok {
		return nil, ErrDBCollNotExist
	}

	switch castedValue := value.(type) {
	case []document:
		return castedValue, nil
	case *ErrCollectionCountIsInvalid:
		return nil, castedValue
	default:
		return nil, ErrUnknownError
	}
}

// db.coll
type dbcollInfo map[string]interface{}

func (ssconfig *SSConfig) GetdbcollInfo() dbcollInfo {
	info := make(dbcollInfo)

	for _, db := range ssconfig.Databases {
		dbName := db.D_name
		collName := db.Collection.C_name
		key := fmt.Sprintf("%s.%s", dbName, collName)

		info[key] = struct{}{}
	}

	return info
}
