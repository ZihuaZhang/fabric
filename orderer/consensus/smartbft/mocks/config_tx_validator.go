// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ZihuaZhang/fabric-protos-go/common"
	mock "github.com/stretchr/testify/mock"
)

// ConfigTxValidator is an autogenerated mock type for the ConfigTxValidator type
type ConfigTxValidator struct {
	mock.Mock
}

// ChannelID provides a mock function with given fields:
func (_m *ConfigTxValidator) ChannelID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ConfigProto provides a mock function with given fields:
func (_m *ConfigTxValidator) ConfigProto() *common.Config {
	ret := _m.Called()

	var r0 *common.Config
	if rf, ok := ret.Get(0).(func() *common.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Config)
		}
	}

	return r0
}

// ProposeConfigUpdate provides a mock function with given fields: configtx
func (_m *ConfigTxValidator) ProposeConfigUpdate(configtx *common.Envelope) (*common.ConfigEnvelope, error) {
	ret := _m.Called(configtx)

	var r0 *common.ConfigEnvelope
	if rf, ok := ret.Get(0).(func(*common.Envelope) *common.ConfigEnvelope); ok {
		r0 = rf(configtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.ConfigEnvelope)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Envelope) error); ok {
		r1 = rf(configtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sequence provides a mock function with given fields:
func (_m *ConfigTxValidator) Sequence() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Validate provides a mock function with given fields: configEnv
func (_m *ConfigTxValidator) Validate(configEnv *common.ConfigEnvelope) error {
	ret := _m.Called(configEnv)

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.ConfigEnvelope) error); ok {
		r0 = rf(configEnv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConfigTxValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfigTxValidator creates a new instance of ConfigTxValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigTxValidator(t mockConstructorTestingTNewConfigTxValidator) *ConfigTxValidator {
	mock := &ConfigTxValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
