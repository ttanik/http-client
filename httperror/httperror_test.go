package httperror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpError_Error(t *testing.T) {
	httpError := HTTPError{
		Status:  http.StatusInternalServerError,
		Message: "error message",
		Err:     errors.New("internal server error"),
	}
	assert.Equal(t, "status: 500 message: error message error: internal server error time: 0001-01-01 00:00:00 +0000 UTC", httpError.Error())
}

func TestHttpError_Error_WithoutError(t *testing.T) {
	httpError := HTTPError{
		Status:  http.StatusInternalServerError,
		Message: "error message",
	}
	assert.Equal(t, "status: 500 message: error message", httpError.Error())
}
