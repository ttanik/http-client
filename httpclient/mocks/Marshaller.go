// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/ttanik/http-client/httperror"

)

// Marshaller is an autogenerated mock type for the Marshaller type
type Marshaller struct {
	mock.Mock
}

// MarshalBody provides a mock function with given fields: body
func (_m *Marshaller) MarshalBody(body interface{}) ([]byte, *httperror.HTTPError) {
	ret := _m.Called(body)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}) []byte); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 *httperror.HTTPError
	if rf, ok := ret.Get(1).(func(interface{}) *httperror.HTTPError); ok {
		r1 = rf(body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*httperror.HTTPError)
		}
	}

	return r0, r1
}
