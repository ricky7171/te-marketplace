// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// GetByFields provides a mock function with given fields: account, fields
func (_m *AccountRepository) GetByFields(account accountdom.Account, fields []string) (interface{}, error) {
	ret := _m.Called(account, fields)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(accountdom.Account, []string) interface{}); ok {
		r0 = rf(account, fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accountdom.Account, []string) error); ok {
		r1 = rf(account, fields)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAccountRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountRepository(t mockConstructorTestingTNewAccountRepository) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
