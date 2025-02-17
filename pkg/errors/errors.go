package errors

import (
	"errors"
	"fmt"
	"runtime"
)

type AppError struct {
	Code    int   
	Message string 
	Err     error  
	Caller  string 
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Caller:  getCaller(),
	}
}

func Wrap(err error, message string) *AppError {
	if err == nil {
		return nil
	}

	return &AppError{
		Code:    InternalError, 
		Message: message,
		Err:     err,
		Caller:  getCaller(),
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %s (caller: %s)", e.Code, e.Message, e.Err.Error(), e.Caller)
	}
	return fmt.Sprintf("[%d] %s (caller: %s)", e.Code, e.Message, e.Caller)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func Is(err error, target *AppError) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == target.Code
	}
	return false
}

func getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", file, line)
}
