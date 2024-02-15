package errors

import "fmt"

type WSUpgradeError struct {
	msg string
	err error
}

func NewWSUpgradeError(err error) *WSUpgradeError {
	return &WSUpgradeError{
		msg: "error upgrading connection to websocket",
		err: err,
	}
}

func (e *WSUpgradeError) Error() string {
	return fmt.Sprintf("%s: %v", e.msg, e.err)
}
