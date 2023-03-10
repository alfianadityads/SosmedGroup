// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	comment "sosmedapps/features/comment"

	mock "github.com/stretchr/testify/mock"
)

// CommentData is an autogenerated mock type for the CommentData type
type CommentData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userID, commentID
func (_m *CommentData) Delete(userID uint, commentID uint) error {
	ret := _m.Called(userID, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCom provides a mock function with given fields:
func (_m *CommentData) GetCom() ([]comment.Core, error) {
	ret := _m.Called()

	var r0 []comment.Core
	if rf, ok := ret.Get(0).(func() []comment.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comment.Core)
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

// NewComment provides a mock function with given fields: userID, contentID, newComment
func (_m *CommentData) NewComment(userID int, contentID uint, newComment string) (comment.Core, error) {
	ret := _m.Called(userID, contentID, newComment)

	var r0 comment.Core
	if rf, ok := ret.Get(0).(func(int, uint, string) comment.Core); ok {
		r0 = rf(userID, contentID, newComment)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, uint, string) error); ok {
		r1 = rf(userID, contentID, newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentData creates a new instance of CommentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentData(t mockConstructorTestingTNewCommentData) *CommentData {
	mock := &CommentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
