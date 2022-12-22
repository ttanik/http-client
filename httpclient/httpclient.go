package httpclient

import (
	"context"
	"net/http"

	"github.com/ttanik/http-client/httperror"
)

// Marshaller ...
type Marshaller interface {
	MarshalBody(body interface{}) ([]byte, *httperror.HTTPError)
}

// Requester ...
type Requester interface {
	ExecuteRequest(request *http.Request) (*http.Response, *httperror.HTTPError)
}

// HTTPRequestBuilder ...
type HTTPRequestBuilder interface {
	WithEndpoint(endpoint string) HTTPRequestBuilder
	WithMethod(method string) HTTPRequestBuilder
	WithHeaders(headers map[string]string) HTTPRequestBuilder
	WithBody(body interface{}) HTTPRequestBuilder
	Build() (*http.Request, *httperror.HTTPError)
}

// NewHTTPClient ...
func NewHTTPClient(requester Requester, marshaller Marshaller) *Client {
	builderConfigs := RequestBuilderConfigs{
		Marshaller: marshaller,
	}

	return &Client{
		requester:      requester,
		builderConfigs: builderConfigs,
	}
}

// Client ...
type Client struct {
	requester      Requester
	builderConfigs RequestBuilderConfigs
}

// Get ...
func (client *Client) Get(ctx context.Context, endpoint string) (*http.Response, *httperror.HTTPError) {
	request, err := NewRequestBuilder(ctx, client.builderConfigs).
		WithEndpoint(endpoint).
		WithMethod(http.MethodGet).
		Build()
	if err != nil {
		return nil, err
	}

	return client.ExecuteRequest(request)
}

// Patch ...
func (client *Client) Patch(ctx context.Context, endpoint string, body interface{}, headers ...map[string]string) (*http.Response, *httperror.HTTPError) {
	requestHeaders := makeHeader(headers...)
	request, err := NewRequestBuilder(ctx, client.builderConfigs).
		WithEndpoint(endpoint).
		WithMethod(http.MethodPatch).
		WithBody(body).
		WithHeaders(requestHeaders).
		Build()
	if err != nil {
		return nil, err
	}

	return client.ExecuteRequest(request)
}

// Put ...
func (client *Client) Put(ctx context.Context, endpoint string, body interface{}, headers ...map[string]string) (*http.Response, *httperror.HTTPError) {
	requestHeaders := makeHeader(headers...)
	request, err := NewRequestBuilder(ctx, client.builderConfigs).
		WithEndpoint(endpoint).
		WithMethod(http.MethodPut).
		WithBody(body).
		WithHeaders(requestHeaders).
		Build()
	if err != nil {
		return nil, err
	}

	return client.ExecuteRequest(request)
}

// Post ...
func (client *Client) Post(ctx context.Context, endpoint string, body interface{}, headers ...map[string]string) (*http.Response, *httperror.HTTPError) {
	requestHeaders := makeHeader(headers...)
	request, err := NewRequestBuilder(ctx, client.builderConfigs).
		WithEndpoint(endpoint).
		WithMethod(http.MethodPost).
		WithBody(body).
		WithHeaders(requestHeaders).
		Build()
	if err != nil {
		return nil, err
	}

	return client.ExecuteRequest(request)
}

// ExecuteRequest ...
func (client *Client) ExecuteRequest(request *http.Request) (*http.Response, *httperror.HTTPError) {
	return client.requester.ExecuteRequest(request)
}

// NewRequestBuilder ...
func (client *Client) NewRequestBuilder(ctx context.Context) HTTPRequestBuilder {
	return NewRequestBuilder(ctx, client.builderConfigs)
}

func makeHeader(headers ...map[string]string) map[string]string {
	var requestHeaders = make(map[string]string)
	for _, headerMap := range headers {
		for key, value := range headerMap {
			requestHeaders[key] = value
		}
	}
	return requestHeaders
}
