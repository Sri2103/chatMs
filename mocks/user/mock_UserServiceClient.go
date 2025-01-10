// Code generated by mockery v2.50.4. DO NOT EDIT.

package user

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	user "github.com/sri2103/chat_me/protos/user"
)

// MockUserServiceClient is an autogenerated mock type for the UserServiceClient type
type MockUserServiceClient struct {
	mock.Mock
}

type MockUserServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserServiceClient) EXPECT() *MockUserServiceClient_Expecter {
	return &MockUserServiceClient_Expecter{mock: &_m.Mock}
}

// AuthenticateUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) AuthenticateUser(ctx context.Context, in *user.AuthenticateRequest, opts ...grpc.CallOption) (*user.AuthenticateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthenticateUser")
	}

	var r0 *user.AuthenticateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.AuthenticateRequest, ...grpc.CallOption) (*user.AuthenticateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.AuthenticateRequest, ...grpc.CallOption) *user.AuthenticateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.AuthenticateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.AuthenticateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserServiceClient_AuthenticateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthenticateUser'
type MockUserServiceClient_AuthenticateUser_Call struct {
	*mock.Call
}

// AuthenticateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - in *user.AuthenticateRequest
//   - opts ...grpc.CallOption
func (_e *MockUserServiceClient_Expecter) AuthenticateUser(ctx interface{}, in interface{}, opts ...interface{}) *MockUserServiceClient_AuthenticateUser_Call {
	return &MockUserServiceClient_AuthenticateUser_Call{Call: _e.mock.On("AuthenticateUser",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockUserServiceClient_AuthenticateUser_Call) Run(run func(ctx context.Context, in *user.AuthenticateRequest, opts ...grpc.CallOption)) *MockUserServiceClient_AuthenticateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*user.AuthenticateRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockUserServiceClient_AuthenticateUser_Call) Return(_a0 *user.AuthenticateResponse, _a1 error) *MockUserServiceClient_AuthenticateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserServiceClient_AuthenticateUser_Call) RunAndReturn(run func(context.Context, *user.AuthenticateRequest, ...grpc.CallOption) (*user.AuthenticateResponse, error)) *MockUserServiceClient_AuthenticateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserDetails provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) GetUserDetails(ctx context.Context, in *user.GetUserRequest, opts ...grpc.CallOption) (*user.GetUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetUserDetails")
	}

	var r0 *user.GetUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserRequest, ...grpc.CallOption) (*user.GetUserResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserRequest, ...grpc.CallOption) *user.GetUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserServiceClient_GetUserDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserDetails'
type MockUserServiceClient_GetUserDetails_Call struct {
	*mock.Call
}

// GetUserDetails is a helper method to define mock.On call
//   - ctx context.Context
//   - in *user.GetUserRequest
//   - opts ...grpc.CallOption
func (_e *MockUserServiceClient_Expecter) GetUserDetails(ctx interface{}, in interface{}, opts ...interface{}) *MockUserServiceClient_GetUserDetails_Call {
	return &MockUserServiceClient_GetUserDetails_Call{Call: _e.mock.On("GetUserDetails",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockUserServiceClient_GetUserDetails_Call) Run(run func(ctx context.Context, in *user.GetUserRequest, opts ...grpc.CallOption)) *MockUserServiceClient_GetUserDetails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*user.GetUserRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockUserServiceClient_GetUserDetails_Call) Return(_a0 *user.GetUserResponse, _a1 error) *MockUserServiceClient_GetUserDetails_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserServiceClient_GetUserDetails_Call) RunAndReturn(run func(context.Context, *user.GetUserRequest, ...grpc.CallOption) (*user.GetUserResponse, error)) *MockUserServiceClient_GetUserDetails_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserDetails provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) UpdateUserDetails(ctx context.Context, in *user.UpdateUserRequest, opts ...grpc.CallOption) (*user.UpdateUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserDetails")
	}

	var r0 *user.UpdateUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUserRequest, ...grpc.CallOption) (*user.UpdateUserResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUserRequest, ...grpc.CallOption) *user.UpdateUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UpdateUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.UpdateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserServiceClient_UpdateUserDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserDetails'
type MockUserServiceClient_UpdateUserDetails_Call struct {
	*mock.Call
}

// UpdateUserDetails is a helper method to define mock.On call
//   - ctx context.Context
//   - in *user.UpdateUserRequest
//   - opts ...grpc.CallOption
func (_e *MockUserServiceClient_Expecter) UpdateUserDetails(ctx interface{}, in interface{}, opts ...interface{}) *MockUserServiceClient_UpdateUserDetails_Call {
	return &MockUserServiceClient_UpdateUserDetails_Call{Call: _e.mock.On("UpdateUserDetails",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockUserServiceClient_UpdateUserDetails_Call) Run(run func(ctx context.Context, in *user.UpdateUserRequest, opts ...grpc.CallOption)) *MockUserServiceClient_UpdateUserDetails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*user.UpdateUserRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockUserServiceClient_UpdateUserDetails_Call) Return(_a0 *user.UpdateUserResponse, _a1 error) *MockUserServiceClient_UpdateUserDetails_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserServiceClient_UpdateUserDetails_Call) RunAndReturn(run func(context.Context, *user.UpdateUserRequest, ...grpc.CallOption) (*user.UpdateUserResponse, error)) *MockUserServiceClient_UpdateUserDetails_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserServiceClient creates a new instance of MockUserServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserServiceClient {
	mock := &MockUserServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
