package file

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetBytesFromFile(path string) ([]byte, error) {
	file, err := OpenFile(path)
	if err != nil {
		return []byte{}, err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, &ErrIOFile{"read", err}
	}

	return b, nil
}

func OpenFile(path string) (*os.File, error) {
	if err := isFile(path); err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, &ErrIOFile{"read", err}
	}
	return file, nil
}

func CheckIsFileExistAndGetFileType(path string) (string, error) {
	if err := isFile(path); err != nil {
		return "", err
	}

	var fileType string
	if fileInfo, err := os.Stat(path); err == nil {
		fileNames := strings.Split(fileInfo.Name(), ".")
		fileType = fileNames[len(fileNames)-1]
	}
	return fileType, nil
}
