package cf

import (
	"fmt"
	"strings"
	"time"
)

const file_sub_type = ".ss."

func GetFilename(fileType string) string {
	return fmt.Sprintf("%d%s%s", time.Now().Unix(), file_sub_type, fileTypeValidator(fileType))
}

func fileTypeValidator(fileType string) string {
	fileType = strings.ToLower(fileType)
	if fileType == "json" || fileType == "yaml" || fileType == "toml" {
		return fileType
	}
	return "yaml"
}
