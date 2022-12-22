package httpmarshal

import (
	"encoding/json"
	"net/http"

	"github.com/ttanik/http-client/httperror"
)

// NewHTTPMarshal ...
func NewHTTPMarshal() *HTTPMarshal {
	return &HTTPMarshal{}
}

// HTTPMarshal ...
type HTTPMarshal struct {
}

// MarshalBody ...
func (httpMarshal *HTTPMarshal) MarshalBody(body interface{}) ([]byte, *httperror.HTTPError) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "error marshalling body",
			Err:     err,
		}
	}
	return jsonBody, nil
}

// Unmarshal ...
func (httpMarshal *HTTPMarshal) Unmarshal(source []byte, target interface{}) *httperror.HTTPError {
	err := json.Unmarshal(source, &target)
	if err != nil {
		return &httperror.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "error unmarshalling source",
			Err:     err,
		}
	}
	return nil
}
