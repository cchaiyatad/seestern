package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func GetBytesFromFile(path string) ([]byte, error) {
	file, err := openFile(path)
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
func IterateLineFromFile(path string, callback func(string)) error {
	file, err := openFile(path)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}
	return scanner.Err()
}

func openFile(path string) (*os.File, error) {
	if err := isFile(path); err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, &ErrIOFile{"read", err}
	}
	return file, nil
}

func GetFileType(path string) (string, error) {
	if err := isFile(path); err != nil {
		return "", err
	}

	var fileType string
	var fileNames []string
	if fileInfo, err := os.Stat(path); err == nil {
		fileNames = strings.Split(fileInfo.Name(), ".")
	}

	if len(fileNames) <= 1 {
		return "", nil
	}

	fileType = fileNames[len(fileNames)-1]
	return fileType, nil
}
