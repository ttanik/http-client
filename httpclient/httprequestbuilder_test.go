package httpclient

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/go-chi/chi/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ttanik/http-client/httpclient/mocks"
	"github.com/ttanik/http-client/httperror"
)

func TestRequestBuilder_Build_MarshalError(t *testing.T) {
	marshallerError := httperror.HTTPError{
		Status:  http.StatusInternalServerError,
		Message: "Marshaller error",
	}
	marshaller := new(mocks.Marshaller)
	marshaller.On("MarshalBody", mock.Anything).Return(nil, &marshallerError)

	requestBuilder := &RequestBuilder{
		Marshaller: marshaller,
		Body:       123,
	}

	request, err := requestBuilder.Build()

	assert.Nil(t, request)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "Marshaller error", err.Message)

	marshaller.AssertExpectations(t)
}

func TestRequestBuilder_Build_NewRequestError(t *testing.T) {
	requestBuilder := &RequestBuilder{
		Method: "whatever",
	}

	request, err := requestBuilder.Build()

	assert.Nil(t, request)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "error creating request", err.Message)
}

func TestRequestBuilder_Build(t *testing.T) {
	marshaller := new(mocks.Marshaller)
	marshaller.On("MarshalBody", 123).Return([]byte(`{"test":"test"}`), nil)

	defaultHeaders := map[string]string{
		"test": "test",
	}

	requestBuilder := &RequestBuilder{
		Body:       123,
		Endpoint:   "http://rain.us/test",
		Marshaller: marshaller,
		Headers:    defaultHeaders,
		Method:     http.MethodPost,
		Ctx:        context.Background(),
	}

	request, err := requestBuilder.Build()

	requestBody, readErr := io.ReadAll(request.Body)
	if readErr != nil {
		t.Error(readErr)
	}
	assert.JSONEq(t, `{"test":"test"}`, string(requestBody))

	assert.Nil(t, err)
	assert.Equal(t, http.MethodPost, request.Method)
	assert.Equal(t, "http://rain.us/test", request.URL.String())
	assert.Equal(t, "test", request.Header.Get("test"))

	marshaller.AssertExpectations(t)
}

func TestRequestBuilder_Build_WithRequestID(t *testing.T) {
	marshaller := new(mocks.Marshaller)

	requestCtx := context.WithValue(context.Background(), middleware.RequestIDKey, "123123")
	requestBuilder := &RequestBuilder{
		Body:       nil,
		Endpoint:   "/test",
		Marshaller: marshaller,
		Method:     http.MethodPost,
		Ctx:        requestCtx,
	}

	request, err := requestBuilder.Build()

	assert.Nil(t, err)
	assert.Equal(t, "123123", request.Header.Get(middleware.RequestIDHeader))
}

func TestRequestBuilder_WithBody(t *testing.T) {
	requestBuilder := &RequestBuilder{}
	requestBuilder.WithBody(123)
	assert.Equal(t, 123, requestBuilder.Body)
}

func TestRequestBuilder_WithEndpoint(t *testing.T) {
	requestBuilder := &RequestBuilder{}
	requestBuilder.WithEndpoint("/test")
	assert.Equal(t, "/test", requestBuilder.Endpoint)
}

func TestRequestBuilder_WithHeader(t *testing.T) {
	defaultHeaders := map[string]string{
		"test":         "test",
		"content-type": "ninjas",
	}
	requestBuilder := &RequestBuilder{
		Headers: defaultHeaders,
	}
	var headerMap = make(map[string]string)
	headerMap["abc"] = "header-three"
	headerMap["new-header"] = "the-new-header"
	requestBuilder.WithHeaders(headerMap)
	assert.Len(t, requestBuilder.Headers, 4)

	assert.Equal(t, "test", requestBuilder.Headers["test"])
	assert.Equal(t, "ninjas", requestBuilder.Headers["content-type"])
	assert.Equal(t, "header-three", requestBuilder.Headers["abc"])
	assert.Equal(t, "the-new-header", requestBuilder.Headers["new-header"])
}

func TestRequestBuilder_WithMethod(t *testing.T) {
	requestBuilder := &RequestBuilder{}
	requestBuilder.WithMethod(http.MethodPost)
	assert.Equal(t, http.MethodPost, requestBuilder.Method)
}
