package db

import "github.com/cchaiyatad/seestern/pkg/cf"

type collsToGen map[string][]string

func parseCollectionInputFromArgs(args []string) collsToGen {
	output := make(collsToGen)

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
