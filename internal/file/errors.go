package file

import "fmt"

type ErrIOTypeNotCorrectType struct {
	path string
	want string
}

func (e *ErrIOTypeNotCorrectType) Error() string {
	return fmt.Sprintf("a given path is not a %s: %s", e.want, e.path)
}

type ErrIOFile struct {
	ops    string
	reason error
}

func (e *ErrIOFile) Error() string {
	return fmt.Sprintf("cannot %s a file: %s", e.ops, e.reason.Error())
}
