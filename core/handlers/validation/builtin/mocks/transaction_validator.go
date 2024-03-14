// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ZihuaZhang/fabric-protos-go/common"
	errors "github.com/ZihuaZhang/fabric/common/errors"

	mock "github.com/stretchr/testify/mock"
)

// TransactionValidator is an autogenerated mock type for the TransactionValidator type
type TransactionValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: block, namespace, txPosition, actionPosition, policy
func (_m *TransactionValidator) Validate(block *common.Block, namespace string, txPosition int, actionPosition int, policy []byte) errors.TxValidationError {
	ret := _m.Called(block, namespace, txPosition, actionPosition, policy)

	var r0 errors.TxValidationError
	if rf, ok := ret.Get(0).(func(*common.Block, string, int, int, []byte) errors.TxValidationError); ok {
		r0 = rf(block, namespace, txPosition, actionPosition, policy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.TxValidationError)
		}
	}

	return r0
}
