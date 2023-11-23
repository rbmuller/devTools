// devErrors/devErrors.go

package devErrors

import (
	"fmt"
	"path/filepath"
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

// NewError creates a new CustomError with file and line information.
func NewError(message string) error {
	return &errorFinder{
		File: getFileName(),
		Line: getLineNumber(),
		Err:  fmt.Errorf(message),
	}
}

// getFileName returns the name of the file where the function is called.
func getFileName() string {
	_, file, _, _ := runtime.Caller(1) // 1 indicates the caller of getFileName
	return filepath.Base(file)
}

// getLineNumber returns the line number where the function is called.
func getLineNumber() int {
	_, _, line, _ := runtime.Caller(1) // 1 indicates the caller of getLineNumber
	return line
}
