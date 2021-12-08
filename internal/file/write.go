package file

import (
	"os"
	"path/filepath"
)

func WriteFile(path string, fileName string, data []byte) error {
	f, err := createFile(path, fileName)
	if err != nil {
		return &ErrIOFile{"write", err}
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return &ErrIOFile{"write", err}
	}

	return nil
}

func createFile(path string, fileName string) (*os.File, error) {
	filePath := filepath.Join(path, fileName)
	return os.Create(filePath)
}
