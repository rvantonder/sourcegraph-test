package httpapi

import "fmt"

type ClientError struct { /* all structs must go */ }

func (e *ClientError) Error() string {
	return e.err.Error()
}

func clientError(message string, vals ...interface{}) error {
	return &ClientError{err: fmt.Errorf(message, vals...)}
}
