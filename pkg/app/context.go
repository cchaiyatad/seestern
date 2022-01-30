package app

import (
	"context"
	"time"
)

func getCtxForConnect() (context.Context, context.CancelFunc) {
	return getCtxWithTimeout(10)
}
func getCtxForTransaction() (context.Context, context.CancelFunc) {
	return getCtxWithTimeout(5)
}
func getCtxWithTimeout(second int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(second)*time.Second)
}
