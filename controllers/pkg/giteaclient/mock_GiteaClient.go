// Code generated by mockery v2.37.1. DO NOT EDIT.

package giteaclient

import (
	context "context"

	gitea "code.gitea.io/sdk/gitea"

	mock "github.com/stretchr/testify/mock"
)

// MockGiteaClient is an autogenerated mock type for the GiteaClient type
type MockGiteaClient struct {
	mock.Mock
}

type MockGiteaClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGiteaClient) EXPECT() *MockGiteaClient_Expecter {
	return &MockGiteaClient_Expecter{mock: &_m.Mock}
}

// CreateRepo provides a mock function with given fields: createRepoOption
func (_m *MockGiteaClient) CreateRepo(createRepoOption gitea.CreateRepoOption) (*gitea.Repository, *gitea.Response, error) {
	ret := _m.Called(createRepoOption)

	var r0 *gitea.Repository
	var r1 *gitea.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(gitea.CreateRepoOption) (*gitea.Repository, *gitea.Response, error)); ok {
		return rf(createRepoOption)
	}
	if rf, ok := ret.Get(0).(func(gitea.CreateRepoOption) *gitea.Repository); ok {
		r0 = rf(createRepoOption)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(gitea.CreateRepoOption) *gitea.Response); ok {
		r1 = rf(createRepoOption)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitea.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(gitea.CreateRepoOption) error); ok {
		r2 = rf(createRepoOption)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockGiteaClient_CreateRepo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRepo'
type MockGiteaClient_CreateRepo_Call struct {
	*mock.Call
}

// CreateRepo is a helper method to define mock.On call
//   - createRepoOption gitea.CreateRepoOption
func (_e *MockGiteaClient_Expecter) CreateRepo(createRepoOption interface{}) *MockGiteaClient_CreateRepo_Call {
	return &MockGiteaClient_CreateRepo_Call{Call: _e.mock.On("CreateRepo", createRepoOption)}
}

func (_c *MockGiteaClient_CreateRepo_Call) Run(run func(createRepoOption gitea.CreateRepoOption)) *MockGiteaClient_CreateRepo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gitea.CreateRepoOption))
	})
	return _c
}

func (_c *MockGiteaClient_CreateRepo_Call) Return(_a0 *gitea.Repository, _a1 *gitea.Response, _a2 error) *MockGiteaClient_CreateRepo_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockGiteaClient_CreateRepo_Call) RunAndReturn(run func(gitea.CreateRepoOption) (*gitea.Repository, *gitea.Response, error)) *MockGiteaClient_CreateRepo_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAccessToken provides a mock function with given fields: value
func (_m *MockGiteaClient) DeleteAccessToken(value interface{}) (*gitea.Response, error) {
	ret := _m.Called(value)

	var r0 *gitea.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (*gitea.Response, error)); ok {
		return rf(value)
	}
	if rf, ok := ret.Get(0).(func(interface{}) *gitea.Response); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGiteaClient_DeleteAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccessToken'
type MockGiteaClient_DeleteAccessToken_Call struct {
	*mock.Call
}

// DeleteAccessToken is a helper method to define mock.On call
//   - value interface{}
func (_e *MockGiteaClient_Expecter) DeleteAccessToken(value interface{}) *MockGiteaClient_DeleteAccessToken_Call {
	return &MockGiteaClient_DeleteAccessToken_Call{Call: _e.mock.On("DeleteAccessToken", value)}
}

func (_c *MockGiteaClient_DeleteAccessToken_Call) Run(run func(value interface{})) *MockGiteaClient_DeleteAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockGiteaClient_DeleteAccessToken_Call) Return(_a0 *gitea.Response, _a1 error) *MockGiteaClient_DeleteAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGiteaClient_DeleteAccessToken_Call) RunAndReturn(run func(interface{}) (*gitea.Response, error)) *MockGiteaClient_DeleteAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRepo provides a mock function with given fields: owner, repo
func (_m *MockGiteaClient) DeleteRepo(owner string, repo string) (*gitea.Response, error) {
	ret := _m.Called(owner, repo)

	var r0 *gitea.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*gitea.Response, error)); ok {
		return rf(owner, repo)
	}
	if rf, ok := ret.Get(0).(func(string, string) *gitea.Response); ok {
		r0 = rf(owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGiteaClient_DeleteRepo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRepo'
type MockGiteaClient_DeleteRepo_Call struct {
	*mock.Call
}

// DeleteRepo is a helper method to define mock.On call
//   - owner string
//   - repo string
func (_e *MockGiteaClient_Expecter) DeleteRepo(owner interface{}, repo interface{}) *MockGiteaClient_DeleteRepo_Call {
	return &MockGiteaClient_DeleteRepo_Call{Call: _e.mock.On("DeleteRepo", owner, repo)}
}

func (_c *MockGiteaClient_DeleteRepo_Call) Run(run func(owner string, repo string)) *MockGiteaClient_DeleteRepo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockGiteaClient_DeleteRepo_Call) Return(_a0 *gitea.Response, _a1 error) *MockGiteaClient_DeleteRepo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGiteaClient_DeleteRepo_Call) RunAndReturn(run func(string, string) (*gitea.Response, error)) *MockGiteaClient_DeleteRepo_Call {
	_c.Call.Return(run)
	return _c
}

// EditRepo provides a mock function with given fields: userName, repoCRName, editRepoOption
func (_m *MockGiteaClient) EditRepo(userName string, repoCRName string, editRepoOption gitea.EditRepoOption) (*gitea.Repository, *gitea.Response, error) {
	ret := _m.Called(userName, repoCRName, editRepoOption)

	var r0 *gitea.Repository
	var r1 *gitea.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, gitea.EditRepoOption) (*gitea.Repository, *gitea.Response, error)); ok {
		return rf(userName, repoCRName, editRepoOption)
	}
	if rf, ok := ret.Get(0).(func(string, string, gitea.EditRepoOption) *gitea.Repository); ok {
		r0 = rf(userName, repoCRName, editRepoOption)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, gitea.EditRepoOption) *gitea.Response); ok {
		r1 = rf(userName, repoCRName, editRepoOption)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitea.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, gitea.EditRepoOption) error); ok {
		r2 = rf(userName, repoCRName, editRepoOption)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockGiteaClient_EditRepo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditRepo'
type MockGiteaClient_EditRepo_Call struct {
	*mock.Call
}

// EditRepo is a helper method to define mock.On call
//   - userName string
//   - repoCRName string
//   - editRepoOption gitea.EditRepoOption
func (_e *MockGiteaClient_Expecter) EditRepo(userName interface{}, repoCRName interface{}, editRepoOption interface{}) *MockGiteaClient_EditRepo_Call {
	return &MockGiteaClient_EditRepo_Call{Call: _e.mock.On("EditRepo", userName, repoCRName, editRepoOption)}
}

func (_c *MockGiteaClient_EditRepo_Call) Run(run func(userName string, repoCRName string, editRepoOption gitea.EditRepoOption)) *MockGiteaClient_EditRepo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(gitea.EditRepoOption))
	})
	return _c
}

func (_c *MockGiteaClient_EditRepo_Call) Return(_a0 *gitea.Repository, _a1 *gitea.Response, _a2 error) *MockGiteaClient_EditRepo_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockGiteaClient_EditRepo_Call) RunAndReturn(run func(string, string, gitea.EditRepoOption) (*gitea.Repository, *gitea.Response, error)) *MockGiteaClient_EditRepo_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields:
func (_m *MockGiteaClient) Get() *gitea.Client {
	ret := _m.Called()

	var r0 *gitea.Client
	if rf, ok := ret.Get(0).(func() *gitea.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.Client)
		}
	}

	return r0
}

// MockGiteaClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockGiteaClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockGiteaClient_Expecter) Get() *MockGiteaClient_Get_Call {
	return &MockGiteaClient_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockGiteaClient_Get_Call) Run(run func()) *MockGiteaClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGiteaClient_Get_Call) Return(_a0 *gitea.Client) *MockGiteaClient_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGiteaClient_Get_Call) RunAndReturn(run func() *gitea.Client) *MockGiteaClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetMyUserInfo provides a mock function with given fields:
func (_m *MockGiteaClient) GetMyUserInfo() (*gitea.User, *gitea.Response, error) {
	ret := _m.Called()

	var r0 *gitea.User
	var r1 *gitea.Response
	var r2 error
	if rf, ok := ret.Get(0).(func() (*gitea.User, *gitea.Response, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gitea.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.User)
		}
	}

	if rf, ok := ret.Get(1).(func() *gitea.Response); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitea.Response)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockGiteaClient_GetMyUserInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMyUserInfo'
type MockGiteaClient_GetMyUserInfo_Call struct {
	*mock.Call
}

// GetMyUserInfo is a helper method to define mock.On call
func (_e *MockGiteaClient_Expecter) GetMyUserInfo() *MockGiteaClient_GetMyUserInfo_Call {
	return &MockGiteaClient_GetMyUserInfo_Call{Call: _e.mock.On("GetMyUserInfo")}
}

func (_c *MockGiteaClient_GetMyUserInfo_Call) Run(run func()) *MockGiteaClient_GetMyUserInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGiteaClient_GetMyUserInfo_Call) Return(_a0 *gitea.User, _a1 *gitea.Response, _a2 error) *MockGiteaClient_GetMyUserInfo_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockGiteaClient_GetMyUserInfo_Call) RunAndReturn(run func() (*gitea.User, *gitea.Response, error)) *MockGiteaClient_GetMyUserInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetRepo provides a mock function with given fields: userName, repoCRName
func (_m *MockGiteaClient) GetRepo(userName string, repoCRName string) (*gitea.Repository, *gitea.Response, error) {
	ret := _m.Called(userName, repoCRName)

	var r0 *gitea.Repository
	var r1 *gitea.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*gitea.Repository, *gitea.Response, error)); ok {
		return rf(userName, repoCRName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *gitea.Repository); ok {
		r0 = rf(userName, repoCRName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitea.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *gitea.Response); ok {
		r1 = rf(userName, repoCRName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitea.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(userName, repoCRName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockGiteaClient_GetRepo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRepo'
type MockGiteaClient_GetRepo_Call struct {
	*mock.Call
}

// GetRepo is a helper method to define mock.On call
//   - userName string
//   - repoCRName string
func (_e *MockGiteaClient_Expecter) GetRepo(userName interface{}, repoCRName interface{}) *MockGiteaClient_GetRepo_Call {
	return &MockGiteaClient_GetRepo_Call{Call: _e.mock.On("GetRepo", userName, repoCRName)}
}

func (_c *MockGiteaClient_GetRepo_Call) Run(run func(userName string, repoCRName string)) *MockGiteaClient_GetRepo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockGiteaClient_GetRepo_Call) Return(_a0 *gitea.Repository, _a1 *gitea.Response, _a2 error) *MockGiteaClient_GetRepo_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockGiteaClient_GetRepo_Call) RunAndReturn(run func(string, string) (*gitea.Repository, *gitea.Response, error)) *MockGiteaClient_GetRepo_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockGiteaClient) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockGiteaClient_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockGiteaClient_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGiteaClient_Expecter) Start(ctx interface{}) *MockGiteaClient_Start_Call {
	return &MockGiteaClient_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockGiteaClient_Start_Call) Run(run func(ctx context.Context)) *MockGiteaClient_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGiteaClient_Start_Call) Return() *MockGiteaClient_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockGiteaClient_Start_Call) RunAndReturn(run func(context.Context)) *MockGiteaClient_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGiteaClient creates a new instance of MockGiteaClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGiteaClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGiteaClient {
	mock := &MockGiteaClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
