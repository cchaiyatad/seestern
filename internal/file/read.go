package file

import (
	"io/ioutil"
	"os"
)

// read toml
func getBytesFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, &ErrIOFile{"read", err}
	}

	defer file.Close()

	return ioutil.ReadAll(file)
}
