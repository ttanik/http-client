package httperror

import (
	"fmt"
	"time"
)

const (
	// ErrorStatus ...
	ErrorStatus = "error_status"
	// ErrorMessage ...
	ErrorMessage = "error_message"
)

// HTTPError ...
type HTTPError struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
	Time    time.Time `json:"time"`
}

// Error prints the error struct
func (e *HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("status: %d message: %s error: %s time: %s", e.Status, e.Message, e.Err.Error(), e.Time)
	}
	return fmt.Sprintf("status: %d message: %s", e.Status, e.Message)
}

// Unwrap returns the Error
func (e *HTTPError) Unwrap() error {
	return e.Err
}
