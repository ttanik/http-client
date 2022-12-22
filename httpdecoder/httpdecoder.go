package httpdecoder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ttanik/http-client/httperror"
)

// NewHTTPDecoder ...
func NewHTTPDecoder() *Decoder {
	return &Decoder{}
}

// Decoder ...
type Decoder struct {
}

// DecodeResponseBody ...
func (d *Decoder) DecodeResponseBody(ctx context.Context, response *http.Response, target interface{}) *httperror.HTTPError {
	if response == nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "response cannot be nil",
		}
	}

	if response.Body == nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "response body cannot be nil",
		}
	}

	if target == nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "target cannot be nil",
		}
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("error closing response body %s", err)
		}
	}()

	err := json.NewDecoder(response.Body).Decode(target)
	if err != nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "error decoding response body",
			Err:     err,
		}
	}

	return nil
}

// DecodeErrorBody ...
func (d *Decoder) DecodeErrorBody(ctx context.Context, response *http.Response) *httperror.HTTPError {
	var requestError httperror.HTTPError
	err := d.DecodeResponseBody(ctx, response, &requestError)
	if err != nil {
		return err
	}

	if requestError.Status == 0 {
		requestError.Status = response.StatusCode
	}

	return &requestError
}

// DecodeRequestBody ...
func (d *Decoder) DecodeRequestBody(request *http.Request, target interface{}) *httperror.HTTPError {
	if request == nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "request cannot be nil",
		}
	}

	if request.Body == nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "request body cannot be nil",
		}
	}

	if target == nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "target cannot be nil",
		}
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		httpError := &httperror.HTTPError{
			Status:  http.StatusBadRequest,
			Message: "error reading request body",
			Err:     err,
		}
		return httpError
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		httpError := &httperror.HTTPError{
			Status:  http.StatusBadRequest,
			Message: "error unmarshalling request body",
			Err:     err,
		}
		return httpError
	}

	return nil
}
