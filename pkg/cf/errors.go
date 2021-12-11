package cf

import (
	"errors"
	"fmt"
)

var ErrInvalidSchemaTree = errors.New(`invalid schema tree: root node of schmea tree have to be { Name: "_root", Type: Object }`)

type ErrMergeDiffTree struct {
	t1 *SchemaTree
	t2 *SchemaTree
}

func (e *ErrMergeDiffTree) Error() string {
	return fmt.Sprintf("can not merge tree %s.%s with %s.%s", e.t1.Database, e.t1.Collection, e.t2.Database, e.t2.Collection)
}
