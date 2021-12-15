package file

import (
	"io/ioutil"
	"os"
)

func GetBytesFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, &ErrIOFile{"read", err}
	}

	defer file.Close()

	return ioutil.ReadAll(file)
}
