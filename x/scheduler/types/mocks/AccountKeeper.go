// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper is an autogenerated mock type for the AccountKeeper type
type AccountKeeper struct {
	mock.Mock
}

// GetAccount provides a mock function with given fields: ctx, addr
func (_m *AccountKeeper) GetAccount(ctx context.Context, addr types.AccAddress) types.AccountI {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for GetAccount")
	}

	var r0 types.AccountI
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) types.AccountI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AccountI)
		}
	}

	return r0
}

// HasAccount provides a mock function with given fields: ctx, addr
func (_m *AccountKeeper) HasAccount(ctx context.Context, addr types.AccAddress) bool {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for HasAccount")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) bool); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewAccountWithAddress provides a mock function with given fields: ctx, addr
func (_m *AccountKeeper) NewAccountWithAddress(ctx context.Context, addr types.AccAddress) types.AccountI {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for NewAccountWithAddress")
	}

	var r0 types.AccountI
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) types.AccountI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AccountI)
		}
	}

	return r0
}

// SetAccount provides a mock function with given fields: ctx, acc
func (_m *AccountKeeper) SetAccount(ctx context.Context, acc types.AccountI) {
	_m.Called(ctx, acc)
}

// NewAccountKeeper creates a new instance of AccountKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *AccountKeeper {
	mock := &AccountKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
