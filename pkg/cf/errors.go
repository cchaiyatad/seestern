package cf

import "errors"

var ErrInvalidSchemaTree = errors.New(`invalid schema tree: root node of schmea tree have to be { Name: "_root", Type: Object }`)
