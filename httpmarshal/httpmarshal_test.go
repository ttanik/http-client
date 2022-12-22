package httpmarshal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServerHttpMarshal(t *testing.T) {
	serverHTTPMarshal := NewHTTPMarshal()
	assert.Equal(t, &HTTPMarshal{}, serverHTTPMarshal)
}

func TestServerHttpMarshal_MarshalBody(t *testing.T) {
	var body struct {
		Name string
	}
	body.Name = "test"

	serverHTTPMarshal := NewHTTPMarshal()
	jsonBody, httpError := serverHTTPMarshal.MarshalBody(body)

	assert.Nil(t, httpError)
	assert.Equal(t, `{"Name":"test"}`, string(jsonBody))
}

func TestServerHttpMarshal_MarshalBody_Error(t *testing.T) {
	var body complex64
	serverHTTPMarshal := NewHTTPMarshal()
	jsonBody, httpError := serverHTTPMarshal.MarshalBody(&body)

	assert.Nil(t, jsonBody)
	assert.Equal(t, "error marshalling body", httpError.Message)
	assert.Equal(t, http.StatusInternalServerError, httpError.Status)
}

func TestServerHttpMarshal_Unmarshal_Error(t *testing.T) {
	serverHTTPMarshal := NewHTTPMarshal()
	httpError := serverHTTPMarshal.Unmarshal(make([]byte, 1), nil)

	assert.Equal(t, "error unmarshalling source", httpError.Message)
	assert.Equal(t, http.StatusInternalServerError, httpError.Status)
}

func TestServerHttpMarshal_Unmarshal(t *testing.T) {
	var body struct {
		Name string
	}

	serverHTTPMarshal := NewHTTPMarshal()
	httpError := serverHTTPMarshal.Unmarshal([]byte(`{"Name":"test"}`), &body)

	assert.Nil(t, httpError)
	assert.Equal(t, "test", body.Name)
}
