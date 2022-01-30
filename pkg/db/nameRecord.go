package db

import (
	"fmt"
	"strings"
)

type NameRecord map[string]map[string]struct{}

func (record NameRecord) empty() bool {
	return len(record) == 0
}

func (record NameRecord) String() string {
	if record.empty() {
		return "database does not exists\n"
	}

	var strBuilder strings.Builder

	for dbName, collNames := range record {
		fmt.Fprintf(&strBuilder, "database: %s\n", dbName)

		if len(collNames) == 0 {
			strBuilder.WriteString(" None\n")
			continue
		}

		idx := 1
		for collName := range collNames {
			fmt.Fprintf(&strBuilder, " %-2d: %s\n", idx, collName)
			idx += 1
		}
	}
	return strBuilder.String()
}
