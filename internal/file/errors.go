package file

import "fmt"

type ErrIsNotDir struct {
	path string
}

func (e *ErrIsNotDir) Error() string {
	return fmt.Sprintf("a given path is not a directory: %s", e.path)
}
