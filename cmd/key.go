package cmd

import "errors"

const (
	connectionStringKey = "connectionString"
	verboseKey          = "verbose"

	databaseKey   = "database"
	collectionKey = "collection"

	// init
	outputKey   = "output"
	fileTypeKey = "type"

	// gen
	fileKey   = "file"
	dropKey   = "drop"
	insertKey = "insert"
)

var verbose bool
var collections []string

var isDrop bool
var isInsert bool

func isEitherVerboseOrOutSet(out string, verbose bool) error {
	if out == "" && !verbose {
		return errors.New("if verbose is not set, output has to be set")
	}
	return nil
}

func isCntStrSetWhenEitherDropOrInsertSet(cntStr string, isDrop bool, isInsert bool) error {
	if (isDrop || isInsert) && cntStr == "" {
		return errors.New("if set drop or insert, connection string has to be provided")
	}
	return nil
}
