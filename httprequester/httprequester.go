package httprequester

import (
	"context"
	"errors"
	"net/http"

	"github.com/ttanik/http-client/httperror"
)

// Requester ...
type Requester interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHTTPRequester ...
func NewHTTPRequester(
	requester Requester,
	decoder Decoder,
) *HTTPRequester {
	return &HTTPRequester{requester, decoder}
}

// Decoder ...
type Decoder interface {
	DecodeErrorBody(ctx context.Context, response *http.Response) *httperror.HTTPError
}

// HTTPRequester ...
type HTTPRequester struct {
	requestExecutioner Requester
	decoder            Decoder
}

// ExecuteRequest ...
func (requester *HTTPRequester) ExecuteRequest(request *http.Request) (*http.Response, *httperror.HTTPError) {

	response, err := requester.requestExecutioner.Do(request)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, &httperror.HTTPError{
				Status:  http.StatusGatewayTimeout,
				Message: "request timed out",
				Err:     err,
			}
		}

		return nil, &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "error executing request",
			Err:     err,
		}
	}

	if http.StatusInternalServerError == response.StatusCode {
		responseError := requester.decoder.DecodeErrorBody(request.Context(), response)
		return nil, &httperror.HTTPError{
			Status:  http.StatusFailedDependency,
			Message: "dependency failed",
			Err:     responseError,
		}
	}

	return response, nil
}

// Do ...
func (requester *HTTPRequester) Do(request *http.Request) (*http.Response, error) {
	response, err := requester.ExecuteRequest(request)
	if err != nil {
		return response, err
	}
	return response, nil
}
