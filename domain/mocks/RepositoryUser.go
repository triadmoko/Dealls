// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "app/model"
)

// RepositoryUser is an autogenerated mock type for the RepositoryUser type
type RepositoryUser struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *RepositoryUser) Create(ctx context.Context, user model.User) (model.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) (model.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User) model.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetailByID provides a mock function with given fields: ctx, id
func (_m *RepositoryUser) DetailByID(ctx context.Context, id string) (model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DetailByID")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetailByUsername provides a mock function with given fields: ctx, username
func (_m *RepositoryUser) DetailByUsername(ctx context.Context, username string) (model.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for DetailByUsername")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPartner provides a mock function with given fields: filter
func (_m *RepositoryUser) SearchPartner(filter model.FilterInterest) ([]model.User, int, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for SearchPartner")
	}

	var r0 []model.User
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(model.FilterInterest) ([]model.User, int, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.FilterInterest) []model.User); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(model.FilterInterest) int); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(model.FilterInterest) error); ok {
		r2 = rf(filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePurchasePremium provides a mock function with given fields: ctx, userID, premium
func (_m *RepositoryUser) UpdatePurchasePremium(ctx context.Context, userID string, premium bool) error {
	ret := _m.Called(ctx, userID, premium)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePurchasePremium")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, userID, premium)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryUser creates a new instance of RepositoryUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryUser {
	mock := &RepositoryUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
