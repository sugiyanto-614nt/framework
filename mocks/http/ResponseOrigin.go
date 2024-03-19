// Code generated by mockery. DO NOT EDIT.

package http

import (
	bytes "bytes"

	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"
)

// ResponseOrigin is an autogenerated mock type for the ResponseOrigin type
type ResponseOrigin struct {
	mock.Mock
}

type ResponseOrigin_Expecter struct {
	mock *mock.Mock
}

func (_m *ResponseOrigin) EXPECT() *ResponseOrigin_Expecter {
	return &ResponseOrigin_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields:
func (_m *ResponseOrigin) Body() *bytes.Buffer {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 *bytes.Buffer
	if rf, ok := ret.Get(0).(func() *bytes.Buffer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bytes.Buffer)
		}
	}

	return r0
}

// ResponseOrigin_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type ResponseOrigin_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
func (_e *ResponseOrigin_Expecter) Body() *ResponseOrigin_Body_Call {
	return &ResponseOrigin_Body_Call{Call: _e.mock.On("Body")}
}

func (_c *ResponseOrigin_Body_Call) Run(run func()) *ResponseOrigin_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseOrigin_Body_Call) Return(_a0 *bytes.Buffer) *ResponseOrigin_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseOrigin_Body_Call) RunAndReturn(run func() *bytes.Buffer) *ResponseOrigin_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *ResponseOrigin) Header() nethttp.Header {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 nethttp.Header
	if rf, ok := ret.Get(0).(func() nethttp.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nethttp.Header)
		}
	}

	return r0
}

// ResponseOrigin_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type ResponseOrigin_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *ResponseOrigin_Expecter) Header() *ResponseOrigin_Header_Call {
	return &ResponseOrigin_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *ResponseOrigin_Header_Call) Run(run func()) *ResponseOrigin_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseOrigin_Header_Call) Return(_a0 nethttp.Header) *ResponseOrigin_Header_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseOrigin_Header_Call) RunAndReturn(run func() nethttp.Header) *ResponseOrigin_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Size provides a mock function with given fields:
func (_m *ResponseOrigin) Size() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Size")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ResponseOrigin_Size_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Size'
type ResponseOrigin_Size_Call struct {
	*mock.Call
}

// Size is a helper method to define mock.On call
func (_e *ResponseOrigin_Expecter) Size() *ResponseOrigin_Size_Call {
	return &ResponseOrigin_Size_Call{Call: _e.mock.On("Size")}
}

func (_c *ResponseOrigin_Size_Call) Run(run func()) *ResponseOrigin_Size_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseOrigin_Size_Call) Return(_a0 int) *ResponseOrigin_Size_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseOrigin_Size_Call) RunAndReturn(run func() int) *ResponseOrigin_Size_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields:
func (_m *ResponseOrigin) Status() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ResponseOrigin_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type ResponseOrigin_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
func (_e *ResponseOrigin_Expecter) Status() *ResponseOrigin_Status_Call {
	return &ResponseOrigin_Status_Call{Call: _e.mock.On("Status")}
}

func (_c *ResponseOrigin_Status_Call) Run(run func()) *ResponseOrigin_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseOrigin_Status_Call) Return(_a0 int) *ResponseOrigin_Status_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseOrigin_Status_Call) RunAndReturn(run func() int) *ResponseOrigin_Status_Call {
	_c.Call.Return(run)
	return _c
}

// NewResponseOrigin creates a new instance of ResponseOrigin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResponseOrigin(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResponseOrigin {
	mock := &ResponseOrigin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
