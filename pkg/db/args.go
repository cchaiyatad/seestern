package db

import (
	"strings"
)

func parseCollectionInputFromArgs(args []string) map[string][]string {
	output := make(map[string][]string)

	for _, arg := range args {
		spilted_arg := strings.SplitN(arg, ".", 2)
		if len(spilted_arg) != 2 {
			continue
		}

		db := spilted_arg[0]
		coll := spilted_arg[1]

		if !contains(output[db], coll) {
			output[db] = append(output[db], coll)
		}
	}
	return output
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
