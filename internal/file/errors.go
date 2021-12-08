package file

import "fmt"

type ErrIsNotDir struct {
	path string
}

func (e *ErrIsNotDir) Error() string {
	return fmt.Sprintf("a given path is not a directory: %s", e.path)
}

type ErrIOFile struct {
	ops    string
	reason error
}

func (e *ErrIOFile) Error() string {
	return fmt.Sprintf("cannot %s a file: %s", e.ops, e.reason.Error())
}
