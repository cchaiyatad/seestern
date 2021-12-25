package db

import "github.com/cchaiyatad/seestern/pkg/cf"

func parseCollectionInputFromArgs(args []string) map[string][]string {
	output := make(map[string][]string)

	for _, arg := range args {
		db, coll, ok := cf.SplitNym(arg)
		if !ok {
			continue
		}

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
