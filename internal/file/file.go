package file

import (
	"os"
)

func PrepareDir(path string) error {
	_ = createDir(path)
	return isDirectory(path)
}

func createDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func isDirectory(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return &ErrIsNotDir{path: path}
	}
	return nil
}
