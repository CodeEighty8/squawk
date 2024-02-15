package errors

import "fmt"

type UUIDParseError struct {
	msg string
	err error
}

func NewUUIDParseError(err error) *UUIDParseError {
	return &UUIDParseError{
		msg: "error parsing uuid",
		err: err,
	}
}

func (e *UUIDParseError) Error() string {
	return fmt.Sprintf("%s: %v", e.msg, e.err)
}
