package nats

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNatsServiceForTelegramUpdateLoopInterface is a mock of NatsServiceForTelegramUpdateLoopInterface interface.
type MockNatsServiceForTelegramUpdateLoopInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNatsServiceForTelegramUpdateLoopInterfaceMockRecorder
}

// MockNatsServiceForTelegramUpdateLoopInterfaceMockRecorder is the mock recorder for MockNatsServiceForTelegramUpdateLoopInterface.
type MockNatsServiceForTelegramUpdateLoopInterfaceMockRecorder struct {
	mock *MockNatsServiceForTelegramUpdateLoopInterface
}

// NewMockNatsServiceForTelegramUpdateLoopInterface creates a new mock instance.
func NewMockNatsServiceForTelegramUpdateLoopInterface(ctrl *gomock.Controller) *MockNatsServiceForTelegramUpdateLoopInterface {
	mock := &MockNatsServiceForTelegramUpdateLoopInterface{ctrl: ctrl}
	mock.recorder = &MockNatsServiceForTelegramUpdateLoopInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatsServiceForTelegramUpdateLoopInterface) EXPECT() *MockNatsServiceForTelegramUpdateLoopInterfaceMockRecorder {
	return m.recorder
}

// Dispatch mocks base method.
func (m *MockNatsServiceForTelegramUpdateLoopInterface) Dispatch(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockNatsServiceForTelegramUpdateLoopInterfaceMockRecorder) Dispatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockNatsServiceForTelegramUpdateLoopInterface)(nil).Dispatch), arg0, arg1)
}

// MockNatsServiceForNewPostsSubscribersInterface is a mock of NatsServiceForNewPostsSubscribersInterface interface.
type MockNatsServiceForNewPostsSubscribersInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNatsServiceForNewPostsSubscribersInterfaceMockRecorder
}

// MockNatsServiceForNewPostsSubscribersInterfaceMockRecorder is the mock recorder for MockNatsServiceForNewPostsSubscribersInterface.
type MockNatsServiceForNewPostsSubscribersInterfaceMockRecorder struct {
	mock *MockNatsServiceForNewPostsSubscribersInterface
}

// NewMockNatsServiceForNewPostsSubscribersInterface creates a new mock instance.
func NewMockNatsServiceForNewPostsSubscribersInterface(ctrl *gomock.Controller) *MockNatsServiceForNewPostsSubscribersInterface {
	mock := &MockNatsServiceForNewPostsSubscribersInterface{ctrl: ctrl}
	mock.recorder = &MockNatsServiceForNewPostsSubscribersInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatsServiceForNewPostsSubscribersInterface) EXPECT() *MockNatsServiceForNewPostsSubscribersInterfaceMockRecorder {
	return m.recorder
}

// SubscribeToNewPostQueue mocks base method.
func (m *MockNatsServiceForNewPostsSubscribersInterface) SubscribeToNewPostQueue() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewPostQueue")
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeToNewPostQueue indicates an expected call of SubscribeToNewPostQueue.
func (mr *MockNatsServiceForNewPostsSubscribersInterfaceMockRecorder) SubscribeToNewPostQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewPostQueue", reflect.TypeOf((*MockNatsServiceForNewPostsSubscribersInterface)(nil).SubscribeToNewPostQueue))
}
