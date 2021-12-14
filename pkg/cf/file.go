package cf

import (
	"fmt"
	"strings"
	"time"
)

const file_sub_type = ".ss."

// gen function
// take path to config file
// read file

// validate config file

func GetFilename(fileType string) string {
	return fmt.Sprintf("%d%s%s", time.Now().Unix(), file_sub_type, fileTypeValid(fileType))
}

func fileTypeValid(fileType string) string {
	fileType = strings.ToLower(fileType)
	if fileType == "json" || fileType == "yaml" || fileType == "toml" {
		return fileType
	}
	return "yaml"
}
