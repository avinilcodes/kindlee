// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import db "kindlee/db"
import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// addUser provides a mock function with given fields: ctx, _a1
func (_m *Service) addUser(ctx context.Context, _a1 db.User) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// listUsers provides a mock function with given fields: ctx
func (_m *Service) listUsers(ctx context.Context) ([]db.User, error) {
	ret := _m.Called(ctx)

	var r0 []db.User
	if rf, ok := ret.Get(0).(func(context.Context) []db.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
