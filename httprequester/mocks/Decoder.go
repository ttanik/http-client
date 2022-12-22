// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	"github.com/ttanik/http-client/httperror"

	mock "github.com/stretchr/testify/mock"
)

// Decoder is an autogenerated mock type for the Decoder type
type Decoder struct {
	mock.Mock
}

// DecodeErrorBody provides a mock function with given fields: ctx, response
func (_m *Decoder) DecodeErrorBody(ctx context.Context, response *http.Response) *httperror.HTTPError {
	ret := _m.Called(ctx, response)

	var r0 *httperror.HTTPError
	if rf, ok := ret.Get(0).(func(context.Context, *http.Response) *httperror.HTTPError); ok {
		r0 = rf(ctx, response)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*httperror.HTTPError)
		}
	}

	return r0
}