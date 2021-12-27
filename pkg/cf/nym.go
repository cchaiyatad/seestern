package cf

import (
	"fmt"
	"regexp"
	"strings"
)

const delimiter = "."

var nymReg = regexp.MustCompile("[^.\n ]+[.][^.\n ]+")
var nymRefReg = regexp.MustCompile("[^.\n ]+([.][^.\n ]+)*")

// nym is <databaseName>.<collectionName>
// nymRef is <f_name>[.<f_name>....]
// ref is <nym>.<nymRef>
func SplitRef(ref string) (string, string, bool) {
	ref = strings.TrimSpace(ref)
	if !isRefValid(ref) {
		return "", "", false
	}

	nym := nymReg.FindString(ref)
	ref = splitNymFromRef(ref, nym)

	nymRef := nymRefReg.FindString(ref)
	return nym, nymRef, true
}

func isRefValid(ref string) bool {
	if !nymReg.MatchString(ref) {
		return false
	}
	nym := nymReg.FindString(ref)
	ref = splitNymFromRef(ref, nym)

	return nymRefReg.MatchString(ref)
}

func splitNymFromRef(ref, nym string) string {
	ref = ref[len(nym):]
	return strings.Trim(ref, delimiter)
}

func SplitNymRef(nymRef string) ([]string, bool) {
	if !nymRefReg.MatchString(nymRef) {
		return []string{}, false
	}
	nymRef = nymRefReg.FindString(nymRef)
	return strings.Split(nymRef, delimiter), true
}

func SplitNym(nym string) (string, string, bool) {
	if !nymReg.MatchString(nym) {
		return "", "", false
	}

	nym = nymReg.FindString(nym)
	spilted_arg := strings.SplitN(nym, delimiter, 2)

	db := spilted_arg[0]
	coll := spilted_arg[1]
	return db, coll, true
}

func CreateNym(dbName, collName string) string {
	return joinDelimiter(dbName, collName)
}

func CreateNymRef(nymRef, fName string) string {
	return joinDelimiter(nymRef, fName)
}

func joinDelimiter(s1, s2 string) string {
	if s1 == "" {
		return s2
	}

	if s2 == "" {
		return s1
	}
	return fmt.Sprintf("%s%s%s", s1, delimiter, s2)
}
