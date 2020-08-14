package customerror

import (
	"fmt"
	"strings"
)

type Error struct {
	statusCode int
	stackTrace []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v", strings.Join(e.stackTrace, " | "))
}

func (e *Error) StatusCode() int {
	return e.statusCode
}

func (e *Error) PrependToStackTrace(message string) {
	e.stackTrace = append([]string{message}, e.stackTrace...)
}

func Wrap(err error, message string) error {
	if customError, ok := err.(*Error); ok {
		return WrapWithStatusCode(err, customError.statusCode, message)
	}

	return WrapWithStatusCode(err, 0, message)
}

func WrapWithStatusCode(err error, statusCode int, message string) error {
	if customError, ok := err.(*Error); ok {
		customError.statusCode = statusCode
		customError.PrependToStackTrace(message)
		return customError
	}

	var stackTrace []string
	if err != nil {
		stackTrace = []string{message, err.Error()}
	} else {
		stackTrace = []string{message}
	}

	return &Error{stackTrace: stackTrace, statusCode: statusCode}
}
