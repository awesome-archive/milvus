// Code generated by mockery v2.32.4. DO NOT EDIT.

package session

import (
	types "github.com/milvus-io/milvus/internal/types"
	mock "github.com/stretchr/testify/mock"
)

// MockWorkerManager is an autogenerated mock type for the WorkerManager type
type MockWorkerManager struct {
	mock.Mock
}

type MockWorkerManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkerManager) EXPECT() *MockWorkerManager_Expecter {
	return &MockWorkerManager_Expecter{mock: &_m.Mock}
}

// AddNode provides a mock function with given fields: nodeID, address
func (_m *MockWorkerManager) AddNode(nodeID int64, address string) error {
	ret := _m.Called(nodeID, address)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(nodeID, address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkerManager_AddNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNode'
type MockWorkerManager_AddNode_Call struct {
	*mock.Call
}

// AddNode is a helper method to define mock.On call
//   - nodeID int64
//   - address string
func (_e *MockWorkerManager_Expecter) AddNode(nodeID interface{}, address interface{}) *MockWorkerManager_AddNode_Call {
	return &MockWorkerManager_AddNode_Call{Call: _e.mock.On("AddNode", nodeID, address)}
}

func (_c *MockWorkerManager_AddNode_Call) Run(run func(nodeID int64, address string)) *MockWorkerManager_AddNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(string))
	})
	return _c
}

func (_c *MockWorkerManager_AddNode_Call) Return(_a0 error) *MockWorkerManager_AddNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerManager_AddNode_Call) RunAndReturn(run func(int64, string) error) *MockWorkerManager_AddNode_Call {
	_c.Call.Return(run)
	return _c
}

// ClientSupportDisk provides a mock function with given fields:
func (_m *MockWorkerManager) ClientSupportDisk() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockWorkerManager_ClientSupportDisk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientSupportDisk'
type MockWorkerManager_ClientSupportDisk_Call struct {
	*mock.Call
}

// ClientSupportDisk is a helper method to define mock.On call
func (_e *MockWorkerManager_Expecter) ClientSupportDisk() *MockWorkerManager_ClientSupportDisk_Call {
	return &MockWorkerManager_ClientSupportDisk_Call{Call: _e.mock.On("ClientSupportDisk")}
}

func (_c *MockWorkerManager_ClientSupportDisk_Call) Run(run func()) *MockWorkerManager_ClientSupportDisk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorkerManager_ClientSupportDisk_Call) Return(_a0 bool) *MockWorkerManager_ClientSupportDisk_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerManager_ClientSupportDisk_Call) RunAndReturn(run func() bool) *MockWorkerManager_ClientSupportDisk_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllClients provides a mock function with given fields:
func (_m *MockWorkerManager) GetAllClients() map[int64]types.IndexNodeClient {
	ret := _m.Called()

	var r0 map[int64]types.IndexNodeClient
	if rf, ok := ret.Get(0).(func() map[int64]types.IndexNodeClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64]types.IndexNodeClient)
		}
	}

	return r0
}

// MockWorkerManager_GetAllClients_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllClients'
type MockWorkerManager_GetAllClients_Call struct {
	*mock.Call
}

// GetAllClients is a helper method to define mock.On call
func (_e *MockWorkerManager_Expecter) GetAllClients() *MockWorkerManager_GetAllClients_Call {
	return &MockWorkerManager_GetAllClients_Call{Call: _e.mock.On("GetAllClients")}
}

func (_c *MockWorkerManager_GetAllClients_Call) Run(run func()) *MockWorkerManager_GetAllClients_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorkerManager_GetAllClients_Call) Return(_a0 map[int64]types.IndexNodeClient) *MockWorkerManager_GetAllClients_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerManager_GetAllClients_Call) RunAndReturn(run func() map[int64]types.IndexNodeClient) *MockWorkerManager_GetAllClients_Call {
	_c.Call.Return(run)
	return _c
}

// GetClientByID provides a mock function with given fields: nodeID
func (_m *MockWorkerManager) GetClientByID(nodeID int64) (types.IndexNodeClient, bool) {
	ret := _m.Called(nodeID)

	var r0 types.IndexNodeClient
	var r1 bool
	if rf, ok := ret.Get(0).(func(int64) (types.IndexNodeClient, bool)); ok {
		return rf(nodeID)
	}
	if rf, ok := ret.Get(0).(func(int64) types.IndexNodeClient); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.IndexNodeClient)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) bool); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockWorkerManager_GetClientByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClientByID'
type MockWorkerManager_GetClientByID_Call struct {
	*mock.Call
}

// GetClientByID is a helper method to define mock.On call
//   - nodeID int64
func (_e *MockWorkerManager_Expecter) GetClientByID(nodeID interface{}) *MockWorkerManager_GetClientByID_Call {
	return &MockWorkerManager_GetClientByID_Call{Call: _e.mock.On("GetClientByID", nodeID)}
}

func (_c *MockWorkerManager_GetClientByID_Call) Run(run func(nodeID int64)) *MockWorkerManager_GetClientByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockWorkerManager_GetClientByID_Call) Return(_a0 types.IndexNodeClient, _a1 bool) *MockWorkerManager_GetClientByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerManager_GetClientByID_Call) RunAndReturn(run func(int64) (types.IndexNodeClient, bool)) *MockWorkerManager_GetClientByID_Call {
	_c.Call.Return(run)
	return _c
}

// PickClient provides a mock function with given fields:
func (_m *MockWorkerManager) PickClient() (int64, types.IndexNodeClient) {
	ret := _m.Called()

	var r0 int64
	var r1 types.IndexNodeClient
	if rf, ok := ret.Get(0).(func() (int64, types.IndexNodeClient)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() types.IndexNodeClient); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(types.IndexNodeClient)
		}
	}

	return r0, r1
}

// MockWorkerManager_PickClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PickClient'
type MockWorkerManager_PickClient_Call struct {
	*mock.Call
}

// PickClient is a helper method to define mock.On call
func (_e *MockWorkerManager_Expecter) PickClient() *MockWorkerManager_PickClient_Call {
	return &MockWorkerManager_PickClient_Call{Call: _e.mock.On("PickClient")}
}

func (_c *MockWorkerManager_PickClient_Call) Run(run func()) *MockWorkerManager_PickClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorkerManager_PickClient_Call) Return(_a0 int64, _a1 types.IndexNodeClient) *MockWorkerManager_PickClient_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerManager_PickClient_Call) RunAndReturn(run func() (int64, types.IndexNodeClient)) *MockWorkerManager_PickClient_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveNode provides a mock function with given fields: nodeID
func (_m *MockWorkerManager) RemoveNode(nodeID int64) {
	_m.Called(nodeID)
}

// MockWorkerManager_RemoveNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveNode'
type MockWorkerManager_RemoveNode_Call struct {
	*mock.Call
}

// RemoveNode is a helper method to define mock.On call
//   - nodeID int64
func (_e *MockWorkerManager_Expecter) RemoveNode(nodeID interface{}) *MockWorkerManager_RemoveNode_Call {
	return &MockWorkerManager_RemoveNode_Call{Call: _e.mock.On("RemoveNode", nodeID)}
}

func (_c *MockWorkerManager_RemoveNode_Call) Run(run func(nodeID int64)) *MockWorkerManager_RemoveNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockWorkerManager_RemoveNode_Call) Return() *MockWorkerManager_RemoveNode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWorkerManager_RemoveNode_Call) RunAndReturn(run func(int64)) *MockWorkerManager_RemoveNode_Call {
	_c.Call.Return(run)
	return _c
}

// StoppingNode provides a mock function with given fields: nodeID
func (_m *MockWorkerManager) StoppingNode(nodeID int64) {
	_m.Called(nodeID)
}

// MockWorkerManager_StoppingNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoppingNode'
type MockWorkerManager_StoppingNode_Call struct {
	*mock.Call
}

// StoppingNode is a helper method to define mock.On call
//   - nodeID int64
func (_e *MockWorkerManager_Expecter) StoppingNode(nodeID interface{}) *MockWorkerManager_StoppingNode_Call {
	return &MockWorkerManager_StoppingNode_Call{Call: _e.mock.On("StoppingNode", nodeID)}
}

func (_c *MockWorkerManager_StoppingNode_Call) Run(run func(nodeID int64)) *MockWorkerManager_StoppingNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockWorkerManager_StoppingNode_Call) Return() *MockWorkerManager_StoppingNode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWorkerManager_StoppingNode_Call) RunAndReturn(run func(int64)) *MockWorkerManager_StoppingNode_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorkerManager creates a new instance of MockWorkerManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkerManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkerManager {
	mock := &MockWorkerManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
