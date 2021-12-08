package cf

import (
	"fmt"
	"time"
)

const file_type = ".ss.toml"

// gen function
// take path to config file
// read file

// validate config file

func getFilename(path string) string {
	return fmt.Sprintf("%d%s", time.Now().Unix(), file_type)
}
