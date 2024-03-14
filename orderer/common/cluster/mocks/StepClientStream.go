// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	orderer "github.com/ZihuaZhang/fabric-protos-go/orderer"
	mock "github.com/stretchr/testify/mock"
)

// StepClientStream is an autogenerated mock type for the StepClientStream type
type StepClientStream struct {
	mock.Mock
}

// Auth provides a mock function with given fields:
func (_m *StepClientStream) Auth() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *StepClientStream) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Recv provides a mock function with given fields:
func (_m *StepClientStream) Recv() (*orderer.StepResponse, error) {
	ret := _m.Called()

	var r0 *orderer.StepResponse
	if rf, ok := ret.Get(0).(func() *orderer.StepResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orderer.StepResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: request
func (_m *StepClientStream) Send(request *orderer.StepRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*orderer.StepRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStepClientStream interface {
	mock.TestingT
	Cleanup(func())
}

// NewStepClientStream creates a new instance of StepClientStream. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStepClientStream(t mockConstructorTestingTNewStepClientStream) *StepClientStream {
	mock := &StepClientStream{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
