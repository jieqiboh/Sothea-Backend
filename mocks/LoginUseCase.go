// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/jieqiboh/sothea_backend/entities"
	mock "github.com/stretchr/testify/mock"
)

// LoginUseCase is an autogenerated mock type for the LoginUseCase type
type LoginUseCase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, user
func (_m *LoginUseCase) Login(ctx context.Context, user entities.User) (string, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.User) (string, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLoginUseCase creates a new instance of LoginUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoginUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoginUseCase {
	mock := &LoginUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
