// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_streamingpb

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	streamingpb "github.com/milvus-io/milvus/pkg/streaming/proto/streamingpb"
)

// MockStreamingNodeHandlerService_ProduceServer is an autogenerated mock type for the StreamingNodeHandlerService_ProduceServer type
type MockStreamingNodeHandlerService_ProduceServer struct {
	mock.Mock
}

type MockStreamingNodeHandlerService_ProduceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStreamingNodeHandlerService_ProduceServer) EXPECT() *MockStreamingNodeHandlerService_ProduceServer_Expecter {
	return &MockStreamingNodeHandlerService_ProduceServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockStreamingNodeHandlerService_ProduceServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockStreamingNodeHandlerService_ProduceServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockStreamingNodeHandlerService_ProduceServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) Context() *MockStreamingNodeHandlerService_ProduceServer_Context_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Context_Call) Run(run func()) *MockStreamingNodeHandlerService_ProduceServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Context_Call) Return(_a0 context.Context) *MockStreamingNodeHandlerService_ProduceServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Context_Call) RunAndReturn(run func() context.Context) *MockStreamingNodeHandlerService_ProduceServer_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *MockStreamingNodeHandlerService_ProduceServer) Recv() (*streamingpb.ProduceRequest, error) {
	ret := _m.Called()

	var r0 *streamingpb.ProduceRequest
	var r1 error
	if rf, ok := ret.Get(0).(func() (*streamingpb.ProduceRequest, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *streamingpb.ProduceRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*streamingpb.ProduceRequest)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStreamingNodeHandlerService_ProduceServer_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockStreamingNodeHandlerService_ProduceServer_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) Recv() *MockStreamingNodeHandlerService_ProduceServer_Recv_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Recv_Call) Run(run func()) *MockStreamingNodeHandlerService_ProduceServer_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Recv_Call) Return(_a0 *streamingpb.ProduceRequest, _a1 error) *MockStreamingNodeHandlerService_ProduceServer_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Recv_Call) RunAndReturn(run func() (*streamingpb.ProduceRequest, error)) *MockStreamingNodeHandlerService_ProduceServer_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockStreamingNodeHandlerService_ProduceServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) RecvMsg(m interface{}) *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call) Run(run func(m interface{})) *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call) Return(_a0 error) *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockStreamingNodeHandlerService_ProduceServer_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockStreamingNodeHandlerService_ProduceServer) Send(_a0 *streamingpb.ProduceResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*streamingpb.ProduceResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingNodeHandlerService_ProduceServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockStreamingNodeHandlerService_ProduceServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *streamingpb.ProduceResponse
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) Send(_a0 interface{}) *MockStreamingNodeHandlerService_ProduceServer_Send_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Send_Call) Run(run func(_a0 *streamingpb.ProduceResponse)) *MockStreamingNodeHandlerService_ProduceServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*streamingpb.ProduceResponse))
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Send_Call) Return(_a0 error) *MockStreamingNodeHandlerService_ProduceServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_Send_Call) RunAndReturn(run func(*streamingpb.ProduceResponse) error) *MockStreamingNodeHandlerService_ProduceServer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockStreamingNodeHandlerService_ProduceServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) SendHeader(_a0 interface{}) *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call) Return(_a0 error) *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockStreamingNodeHandlerService_ProduceServer_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockStreamingNodeHandlerService_ProduceServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) SendMsg(m interface{}) *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call) Run(run func(m interface{})) *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call) Return(_a0 error) *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockStreamingNodeHandlerService_ProduceServer_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockStreamingNodeHandlerService_ProduceServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) SetHeader(_a0 interface{}) *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call) Return(_a0 error) *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockStreamingNodeHandlerService_ProduceServer_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockStreamingNodeHandlerService_ProduceServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockStreamingNodeHandlerService_ProduceServer_Expecter) SetTrailer(_a0 interface{}) *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call {
	return &MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call) Return() *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *MockStreamingNodeHandlerService_ProduceServer_SetTrailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStreamingNodeHandlerService_ProduceServer creates a new instance of MockStreamingNodeHandlerService_ProduceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStreamingNodeHandlerService_ProduceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStreamingNodeHandlerService_ProduceServer {
	mock := &MockStreamingNodeHandlerService_ProduceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
