package httpclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ttanik/http-client/httpclient/mocks"
)

func TestClient_ExecuteRequest(t *testing.T) {
	requester := new(mocks.Requester)
	requester.On("ExecuteRequest", mock.Anything).Return(&http.Response{}, nil)

	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	response, err := client.ExecuteRequest(&http.Request{})

	assert.Nil(t, err)
	assert.Equal(t, &http.Response{}, response)

	requester.AssertExpectations(t)
}

func TestClient_Get_CreateRequestError(t *testing.T) {
	requester := new(mocks.Requester)
	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	response, httpError := client.Get(context.Background(), "\n / /")

	assert.Nil(t, response)
	assert.Equal(t, "error creating request", httpError.Message)
	assert.Equal(t, http.StatusInternalServerError, httpError.Status)

	requester.AssertNotCalled(t, "ExecuteRequest", mock.Anything)
}

func TestClient_Get(t *testing.T) {
	requester := new(mocks.Requester)
	requester.On("ExecuteRequest", mock.Anything).Return(&http.Response{}, nil)

	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	getResponse, httpError := client.Get(context.Background(), "/test/get")

	assert.Nil(t, httpError)
	assert.Equal(t, &http.Response{}, getResponse)

	requester.AssertExpectations(t)
}

func TestClient_Post_CreateRequestError(t *testing.T) {
	requester := new(mocks.Requester)
	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	response, httpError := client.Post(context.Background(), " \n / a", nil)

	assert.Nil(t, response)
	assert.Equal(t, "error creating request", httpError.Message)
	assert.Equal(t, http.StatusInternalServerError, httpError.Status)

	requester.AssertNotCalled(t, "ExecuteRequest", mock.Anything)
}

func TestClient_Post(t *testing.T) {
	requester := new(mocks.Requester)
	requester.On("ExecuteRequest", mock.Anything).Return(&http.Response{}, nil)

	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	getResponse, httpError := client.Post(context.Background(), "/test/post", nil)

	assert.Nil(t, httpError)
	assert.Equal(t, &http.Response{}, getResponse)

	requester.AssertExpectations(t)
}

func TestClient_Put_CreateRequestError(t *testing.T) {
	requester := new(mocks.Requester)
	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	response, httpError := client.Put(context.Background(), "  \n a", nil, nil)

	assert.Nil(t, response)
	assert.Equal(t, "error creating request", httpError.Message)
	assert.Equal(t, http.StatusInternalServerError, httpError.Status)

	requester.AssertNotCalled(t, "ExecuteRequest", mock.Anything)
}

func TestClient_Patch(t *testing.T) {
	requester := new(mocks.Requester)
	requester.On("ExecuteRequest", mock.Anything).Return(&http.Response{}, nil)

	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	getResponse, httpError := client.Patch(context.Background(), "/test/patch", nil)

	assert.Nil(t, httpError)
	assert.Equal(t, &http.Response{}, getResponse)

	requester.AssertExpectations(t)
}

func TestClient_Put(t *testing.T) {
	requester := new(mocks.Requester)
	requester.On("ExecuteRequest", mock.Anything).Return(&http.Response{}, nil)

	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	getResponse, httpError := client.Put(context.Background(), "/test/put", nil)

	assert.Nil(t, httpError)
	assert.Equal(t, &http.Response{}, getResponse)

	requester.AssertExpectations(t)
}

func TestClient_NewRequestBuilder(t *testing.T) {

	requester := new(mocks.Requester)
	marshaller := new(mocks.Marshaller)
	client := NewHTTPClient(requester, marshaller)

	newContextBuilder := client.NewRequestBuilder(context.Background())

	assert.IsType(t, &RequestBuilder{}, newContextBuilder)
}
