package httprequester

import (
	"context"
	"net/http"
	"testing"

	"github.com/ttanik/http-client/httperror"
	"github.com/ttanik/http-client/httprequester/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServerHttpRequester(t *testing.T) {
	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)

	httpRequester := NewHTTPRequester(requester, decoder)
	assert.Equal(t, requester, httpRequester.requestExecutioner)
}

func TestServerHttpRequester_ExecuteRequest(t *testing.T) {
	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)
	requester.On("Do", &http.Request{}).Return(&http.Response{}, nil)

	httpRequester := NewHTTPRequester(requester, decoder)
	response, httpError := httpRequester.ExecuteRequest(&http.Request{})

	assert.Nil(t, httpError)
	assert.Equal(t, &http.Response{}, response)
}

func TestServerHttpRequester_ExecuteRequest_DoError(t *testing.T) {
	httpError := httperror.HTTPError{
		Status:  http.StatusBadGateway,
		Message: "fatal error",
	}

	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)
	requester.On("Do", &http.Request{}).Return(nil, &httpError)

	httpRequester := NewHTTPRequester(requester, decoder)
	response, err := httpRequester.ExecuteRequest(&http.Request{})

	assert.Nil(t, response)
	assert.Equal(t, "error executing request", err.Message)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
}

func TestServerHttpRequester_ExecuteRequest_InternalServerError(t *testing.T) {
	mockedResponse := http.Response{
		StatusCode: http.StatusInternalServerError,
	}

	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)
	requester.On("Do", &http.Request{}).Return(&mockedResponse, nil)
	decoder.On("DecodeErrorBody", mock.Anything, mock.Anything).Return(nil)

	httpRequester := NewHTTPRequester(requester, decoder)
	response, err := httpRequester.ExecuteRequest(&http.Request{})

	assert.Nil(t, response)
	assert.Equal(t, "dependency failed", err.Message)
	assert.Equal(t, http.StatusFailedDependency, err.Status)
}

func TestServerHttpRequester_ExecuteRequest_DeadlineExceededError(t *testing.T) {
	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)
	requester.On("Do", &http.Request{}).Return(nil, context.DeadlineExceeded)

	httpRequester := NewHTTPRequester(requester, decoder)
	response, err := httpRequester.ExecuteRequest(&http.Request{})

	assert.Nil(t, response)
	assert.Equal(t, "request timed out", err.Message)
	assert.Equal(t, http.StatusGatewayTimeout, err.Status)
}

func TestHttpRequester_Do_Success(t *testing.T) {
	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)
	requester.On("Do", &http.Request{}).Return(&http.Response{}, nil)

	httpRequester := NewHTTPRequester(requester, decoder)
	response, httpError := httpRequester.Do(&http.Request{})

	assert.NoError(t, httpError)
	assert.Equal(t, &http.Response{}, response)
}

func TestHttpRequester_Do_Error(t *testing.T) {
	requester := new(mocks.Requester)
	decoder := new(mocks.Decoder)
	requester.On("Do", &http.Request{}).Return(&http.Response{}, context.DeadlineExceeded)

	httpRequester := NewHTTPRequester(requester, decoder)
	response, httpError := httpRequester.Do(&http.Request{})

	assert.Nil(t, response)
	assert.Error(t, httpError)
}
