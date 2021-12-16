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
	return checkFileType(path, true, "directory")
}

func isFile(path string) error {
	return checkFileType(path, false, "file")
}

func checkFileType(path string, isDir bool, want string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() != isDir {
		return &ErrIOTypeNotCorrectType{want: want, path: path}
	}
	return nil
}
