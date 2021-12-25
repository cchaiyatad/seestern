package nym

import "strings"

// nym is <databaseName>.<collectionName>

func SplitNym(nym string) (string, string, bool) {
	spilted_arg := strings.SplitN(nym, ".", 2)
	if len(spilted_arg) != 2 {
		return "", "", false
	}

	db := spilted_arg[0]
	coll := spilted_arg[1]
	return db, coll, true

}
