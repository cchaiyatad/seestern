package db

import (
	"fmt"
)

type ErrVendorNotSupport struct {
	connectionString string
}

func (e *ErrVendorNotSupport) Error() string {
	return fmt.Sprintf("database vendor is not support: %s", e.connectionString)
}

type ErrConnectionStringIsInvalid struct {
	reason error
}

func (e *ErrConnectionStringIsInvalid) Error() string {
	return fmt.Sprintf("can not connect to database with connection string: %s", e.reason)
}
