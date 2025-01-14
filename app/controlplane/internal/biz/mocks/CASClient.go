// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// CASClient is an autogenerated mock type for the CASClient type
type CASClient struct {
	mock.Mock
}

// Configured provides a mock function with given fields:
func (_m *CASClient) Configured() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Download provides a mock function with given fields: ctx, secretID, w, digest
func (_m *CASClient) Download(ctx context.Context, secretID string, w io.Writer, digest string) error {
	ret := _m.Called(ctx, secretID, w, digest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Writer, string) error); ok {
		r0 = rf(ctx, secretID, w, digest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upload provides a mock function with given fields: ctx, secretID, content, filename, digest
func (_m *CASClient) Upload(ctx context.Context, secretID string, content io.Reader, filename string, digest string) error {
	ret := _m.Called(ctx, secretID, content, filename, digest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, string, string) error); ok {
		r0 = rf(ctx, secretID, content, filename, digest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCASClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewCASClient creates a new instance of CASClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCASClient(t mockConstructorTestingTNewCASClient) *CASClient {
	mock := &CASClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
