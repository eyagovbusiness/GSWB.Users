package errors

import "fmt"

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Msg)
}

func NewValidationError(msg string) error {
	return &ValidationError{Msg: msg}
}

type UnexpectedError struct {
	Msg string
	Err error
}

func (e *UnexpectedError) Error() string {
	return fmt.Sprintf("unexpected error: %s - %v", e.Msg, e.Err)
}

func NewUnexpectedError(msg string, err error) error {
	return &UnexpectedError{Msg: msg, Err: err}
}
