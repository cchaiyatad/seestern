package app

import (
	"errors"
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

type ErrSkipCreateConfigfile struct {
	database   string
	collection string
	reason     string
}

func (e *ErrSkipCreateConfigfile) Error() string {
	return fmt.Sprintf("skip: database %s collection %s: reason: %s", e.database, e.collection, e.reason)
}

var ErrClientIsNil = errors.New("error: client cannot be nil")
