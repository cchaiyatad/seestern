package cf

import (
	"encoding/json"
	"errors"
)

// db[coll] -> documents
type result map[string]map[string]interface{}
type documents []document

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

func (info result) GetDocuments(dbName, collName string) (documents, error) {
	var value interface{}
	var ok bool
	if value, ok = info[dbName][collName]; !ok {
		return nil, ErrDBCollNotExist
	}

	switch castedValue := value.(type) {
	case documents:
		return castedValue, nil
	case *ErrCollectionCountIsInvalid:
		return nil, castedValue
	default:
		return nil, ErrUnknownError
	}
}
func (documents documents) ToInterfaceSlice() []interface{} {
	result := make([]interface{}, len(documents))
	for i := range documents {
		result[i] = documents[i]
	}
	return result
}

func (documents documents) ToJson() ([]byte, error) {
	return json.MarshalIndent(documents, "", "    ")
}
