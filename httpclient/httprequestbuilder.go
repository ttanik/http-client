package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/ttanik/http-client/httperror"
)

// RequestBuilder ...
type RequestBuilder struct {
	Endpoint   string
	Method     string
	Marshaller Marshaller
	Body       interface{}
	Ctx        context.Context
	Headers    map[string]string
}

// RequestBuilderConfigs ...
type RequestBuilderConfigs struct {
	Marshaller Marshaller
	Headers    map[string]string
}

// NewRequestBuilder ...
func NewRequestBuilder(ctx context.Context, configs RequestBuilderConfigs) HTTPRequestBuilder {
	return &RequestBuilder{
		Ctx:        ctx,
		Marshaller: configs.Marshaller,
		Headers:    getHeaders(configs.Headers),
	}
}
func getHeaders(defaultHeaders map[string]string) map[string]string {
	headers := map[string]string{}

	for key, value := range defaultHeaders {
		headers[key] = value
	}

	return headers
}

// WithEndpoint ...
func (requestBuilder *RequestBuilder) WithEndpoint(endpoint string) HTTPRequestBuilder {
	requestBuilder.Endpoint = endpoint
	return requestBuilder
}

// WithMethod ...
func (requestBuilder *RequestBuilder) WithMethod(method string) HTTPRequestBuilder {
	requestBuilder.Method = method
	return requestBuilder
}

// WithHeaders ...
func (requestBuilder *RequestBuilder) WithHeaders(headers map[string]string) HTTPRequestBuilder {
	for key, value := range headers {
		requestBuilder.Headers[key] = value
	}
	return requestBuilder
}

// WithBody ...
func (requestBuilder *RequestBuilder) WithBody(body interface{}) HTTPRequestBuilder {
	requestBuilder.Body = body
	return requestBuilder
}

// Build ...
func (requestBuilder *RequestBuilder) Build() (*http.Request, *httperror.HTTPError) {
	requestBody, httpErr := requestBuilder.getRequestBody()
	if httpErr != nil {
		return nil, httpErr
	}

	url := requestBuilder.Endpoint

	request, err := http.NewRequestWithContext(requestBuilder.Ctx, requestBuilder.Method, url, requestBody)
	if err != nil {
		return nil, &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "error creating request",
			Err:     err,
			Time:    time.Now(),
		}
	}

	requestBuilder.applyHeaders(request)

	return request, nil
}

func (requestBuilder *RequestBuilder) getRequestBody() (io.Reader, *httperror.HTTPError) {
	if requestBuilder.Body != nil {
		jsonBody, httpError := requestBuilder.Marshaller.MarshalBody(requestBuilder.Body)
		if httpError != nil {
			return nil, httpError
		}
		return bytes.NewReader(jsonBody), nil
	}

	return nil, nil
}

func (requestBuilder *RequestBuilder) applyHeaders(request *http.Request) {
	for key, value := range requestBuilder.Headers {
		request.Header.Set(key, value)
	}

	requestID := middleware.GetReqID(request.Context())
	request.Header.Set(middleware.RequestIDHeader, requestID)
}
