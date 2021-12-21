package cf

import (
	"fmt"
	"strings"
	"time"
)

const file_sub_type = ".ss."

func GetInitFilename(fileType string) string {
	return getFilename(file_sub_type, fileType)
}

func GetGenFilename(database, collection string) string {
	midName := fmt.Sprintf("_%s_%s.", database, collection)
	return getFilename(midName, "json")
}

func getFilename(midName, fileType string) string {
	return fmt.Sprintf("%d%s%s", time.Now().Unix(), midName, fileTypeValidator(fileType))
}

func fileTypeValidator(fileType string) string {
	fileType = strings.ToLower(fileType)
	if fileType == "json" || fileType == "yaml" || fileType == "toml" {
		return fileType
	}
	return "yaml"
}
