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

func isFlagValid(out string, verbose bool) error {
	if out == "" && !verbose {
		return errors.New("if verbose is not set, output has to be set")
	}
	return nil
}
