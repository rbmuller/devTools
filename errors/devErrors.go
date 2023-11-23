// devErrors/devErrors.go

package devErrors

import (
	"fmt"
	"runtime"
)

// CustomError represents a custom error type with file and line information.
type errorFinder struct {
	File string
	Line int
	Err  error
}

// Error implements the error interface for CustomError.
func (e *errorFinder) Error() string {
	return fmt.Sprintf("%s:%d - %s", e.File, e.Line, e.Err.Error())
}

func NewError(message error) error {
	_, file, line, _ := runtime.Caller(1)
	return &errorFinder{
		Err:  message,
		File: file,
		Line: line,
	}
}
