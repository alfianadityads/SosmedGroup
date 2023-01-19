// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	user "sosmedapps/features/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userToken
func (_m *UserService) Delete(userToken interface{}) error {
	ret := _m.Called(userToken)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(userToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: username, password
func (_m *UserService) Login(username string, password string) (string, user.Core, error) {
	ret := _m.Called(username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 user.Core
	if rf, ok := ret.Get(1).(func(string, string) user.Core); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Get(1).(user.Core)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Logout provides a mock function with given fields:
func (_m *UserService) Logout() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
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

// Profile provides a mock function with given fields: userToken
func (_m *UserService) Profile(userToken interface{}) (interface{}, error) {
	ret := _m.Called(userToken)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(userToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(userToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUser
func (_m *UserService) Register(newUser user.Core) (user.Core, error) {
	ret := _m.Called(newUser)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Searching provides a mock function with given fields: quote
func (_m *UserService) Searching(quote string) ([]user.Core, error) {
	ret := _m.Called(quote)

	var r0 []user.Core
	if rf, ok := ret.Get(0).(func(string) []user.Core); ok {
		r0 = rf(quote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(quote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: formHeader, userToken, updateData
func (_m *UserService) Update(formHeader multipart.FileHeader, userToken interface{}, updateData user.Core) (user.Core, error) {
	ret := _m.Called(formHeader, userToken, updateData)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(multipart.FileHeader, interface{}, user.Core) user.Core); ok {
		r0 = rf(formHeader, userToken, updateData)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(multipart.FileHeader, interface{}, user.Core) error); ok {
		r1 = rf(formHeader, userToken, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
