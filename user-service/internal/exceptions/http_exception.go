package exceptions

import "fmt"

// Custom error type
type HttpException struct {
	Code    int    // HTTP status code
	Message string // Error message
}

// Implement the error interface
func (e *HttpException) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

// Constructor for the error
func NewHttpException(code int, message string) *HttpException {
	return &HttpException{
		Code:    code,
		Message: message,
	}
}
