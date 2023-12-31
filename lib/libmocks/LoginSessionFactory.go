// Code generated by mockery v2.36.1. DO NOT EDIT.

package libmocks

import (
	entgenerated "api/ent/entgenerated"
	lib "api/lib"

	mock "github.com/stretchr/testify/mock"
)

// LoginSessionFactory is an autogenerated mock type for the LoginSessionFactory type
type LoginSessionFactory struct {
	mock.Mock
}

type LoginSessionFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *LoginSessionFactory) EXPECT() *LoginSessionFactory_Expecter {
	return &LoginSessionFactory_Expecter{mock: &_m.Mock}
}

// FromRefreshTokenCookie provides a mock function with given fields: cookieReader, userTokenFactory
func (_m *LoginSessionFactory) FromRefreshTokenCookie(cookieReader lib.CookieReader, userTokenFactory lib.UserTokenFactory) (*entgenerated.LoginSession, error) {
	ret := _m.Called(cookieReader, userTokenFactory)

	var r0 *entgenerated.LoginSession
	var r1 error
	if rf, ok := ret.Get(0).(func(lib.CookieReader, lib.UserTokenFactory) (*entgenerated.LoginSession, error)); ok {
		return rf(cookieReader, userTokenFactory)
	}
	if rf, ok := ret.Get(0).(func(lib.CookieReader, lib.UserTokenFactory) *entgenerated.LoginSession); ok {
		r0 = rf(cookieReader, userTokenFactory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entgenerated.LoginSession)
		}
	}

	if rf, ok := ret.Get(1).(func(lib.CookieReader, lib.UserTokenFactory) error); ok {
		r1 = rf(cookieReader, userTokenFactory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginSessionFactory_FromRefreshTokenCookie_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FromRefreshTokenCookie'
type LoginSessionFactory_FromRefreshTokenCookie_Call struct {
	*mock.Call
}

// FromRefreshTokenCookie is a helper method to define mock.On call
//   - cookieReader lib.CookieReader
//   - userTokenFactory lib.UserTokenFactory
func (_e *LoginSessionFactory_Expecter) FromRefreshTokenCookie(cookieReader interface{}, userTokenFactory interface{}) *LoginSessionFactory_FromRefreshTokenCookie_Call {
	return &LoginSessionFactory_FromRefreshTokenCookie_Call{Call: _e.mock.On("FromRefreshTokenCookie", cookieReader, userTokenFactory)}
}

func (_c *LoginSessionFactory_FromRefreshTokenCookie_Call) Run(run func(cookieReader lib.CookieReader, userTokenFactory lib.UserTokenFactory)) *LoginSessionFactory_FromRefreshTokenCookie_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(lib.CookieReader), args[1].(lib.UserTokenFactory))
	})
	return _c
}

func (_c *LoginSessionFactory_FromRefreshTokenCookie_Call) Return(_a0 *entgenerated.LoginSession, _a1 error) *LoginSessionFactory_FromRefreshTokenCookie_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LoginSessionFactory_FromRefreshTokenCookie_Call) RunAndReturn(run func(lib.CookieReader, lib.UserTokenFactory) (*entgenerated.LoginSession, error)) *LoginSessionFactory_FromRefreshTokenCookie_Call {
	_c.Call.Return(run)
	return _c
}

// FromRefreshTokenCookieBypassPrivacy provides a mock function with given fields: cookieReader, userTokenFactory
func (_m *LoginSessionFactory) FromRefreshTokenCookieBypassPrivacy(cookieReader lib.CookieReader, userTokenFactory lib.UserTokenFactory) (*entgenerated.LoginSession, error) {
	ret := _m.Called(cookieReader, userTokenFactory)

	var r0 *entgenerated.LoginSession
	var r1 error
	if rf, ok := ret.Get(0).(func(lib.CookieReader, lib.UserTokenFactory) (*entgenerated.LoginSession, error)); ok {
		return rf(cookieReader, userTokenFactory)
	}
	if rf, ok := ret.Get(0).(func(lib.CookieReader, lib.UserTokenFactory) *entgenerated.LoginSession); ok {
		r0 = rf(cookieReader, userTokenFactory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entgenerated.LoginSession)
		}
	}

	if rf, ok := ret.Get(1).(func(lib.CookieReader, lib.UserTokenFactory) error); ok {
		r1 = rf(cookieReader, userTokenFactory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FromRefreshTokenCookieBypassPrivacy'
type LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call struct {
	*mock.Call
}

// FromRefreshTokenCookieBypassPrivacy is a helper method to define mock.On call
//   - cookieReader lib.CookieReader
//   - userTokenFactory lib.UserTokenFactory
func (_e *LoginSessionFactory_Expecter) FromRefreshTokenCookieBypassPrivacy(cookieReader interface{}, userTokenFactory interface{}) *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call {
	return &LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call{Call: _e.mock.On("FromRefreshTokenCookieBypassPrivacy", cookieReader, userTokenFactory)}
}

func (_c *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call) Run(run func(cookieReader lib.CookieReader, userTokenFactory lib.UserTokenFactory)) *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(lib.CookieReader), args[1].(lib.UserTokenFactory))
	})
	return _c
}

func (_c *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call) Return(_a0 *entgenerated.LoginSession, _a1 error) *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call) RunAndReturn(run func(lib.CookieReader, lib.UserTokenFactory) (*entgenerated.LoginSession, error)) *LoginSessionFactory_FromRefreshTokenCookieBypassPrivacy_Call {
	_c.Call.Return(run)
	return _c
}

// NewLoginSessionFactory creates a new instance of LoginSessionFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoginSessionFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoginSessionFactory {
	mock := &LoginSessionFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
