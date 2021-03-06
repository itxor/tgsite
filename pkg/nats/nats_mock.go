// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/nats/interface.go

// Package mock_nats is a generated GoMock package.
package nats

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
)

// MockNatsClientInterface is a mock of NatsClientInterface interface.
type MockNatsClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNatsClientInterfaceMockRecorder
}

// MockNatsClientInterfaceMockRecorder is the mock recorder for MockNatsClientInterface.
type MockNatsClientInterfaceMockRecorder struct {
	mock *MockNatsClientInterface
}

// NewMockNatsClientInterface creates a new mock instance.
func NewMockNatsClientInterface(ctrl *gomock.Controller) *MockNatsClientInterface {
	mock := &MockNatsClientInterface{ctrl: ctrl}
	mock.recorder = &MockNatsClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatsClientInterface) EXPECT() *MockNatsClientInterfaceMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockNatsClientInterface) Connect() (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockNatsClientInterfaceMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNatsClientInterface)(nil).Connect))
}

// GetConnect mocks base method.
func (m *MockNatsClientInterface) GetConnect() *nats.EncodedConn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnect")
	ret0, _ := ret[0].(*nats.EncodedConn)
	return ret0
}

// GetConnect indicates an expected call of GetConnect.
func (mr *MockNatsClientInterfaceMockRecorder) GetConnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnect", reflect.TypeOf((*MockNatsClientInterface)(nil).GetConnect))
}

// Publish mocks base method.
func (m *MockNatsClientInterface) Publish(subject string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", subject, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockNatsClientInterfaceMockRecorder) Publish(subject, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockNatsClientInterface)(nil).Publish), subject, value)
}
