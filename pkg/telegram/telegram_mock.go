package telegram

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTelegramClientInterface is a mock of TelegramClientInterface interface.
type MockTelegramClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramClientInterfaceMockRecorder
}

// MockTelegramClientInterfaceMockRecorder is the mock recorder for MockTelegramClientInterface.
type MockTelegramClientInterfaceMockRecorder struct {
	mock *MockTelegramClientInterface
}

// NewMockTelegramClientInterface creates a new mock instance.
func NewMockTelegramClientInterface(ctrl *gomock.Controller) *MockTelegramClientInterface {
	mock := &MockTelegramClientInterface{ctrl: ctrl}
	mock.recorder = &MockTelegramClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramClientInterface) EXPECT() *MockTelegramClientInterfaceMockRecorder {
	return m.recorder
}

// GetStickerURL mocks base method.
func (m *MockTelegramClientInterface) GetStickerURL(message MessageDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStickerURL", message)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStickerURL indicates an expected call of GetStickerURL.
func (mr *MockTelegramClientInterfaceMockRecorder) GetStickerURL(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStickerURL", reflect.TypeOf((*MockTelegramClientInterface)(nil).GetStickerURL), message)
}

// GetUpdateChan mocks base method.
func (m *MockTelegramClientInterface) GetUpdateChan() chan MessageDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateChan")
	ret0, _ := ret[0].(chan MessageDTO)
	return ret0
}

// GetUpdateChan indicates an expected call of GetUpdateChan.
func (mr *MockTelegramClientInterfaceMockRecorder) GetUpdateChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateChan", reflect.TypeOf((*MockTelegramClientInterface)(nil).GetUpdateChan))
}

// GetVoiceURL mocks base method.
func (m *MockTelegramClientInterface) GetVoiceURL(message MessageDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVoiceURL", message)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoiceURL indicates an expected call of GetVoiceURL.
func (mr *MockTelegramClientInterfaceMockRecorder) GetVoiceURL(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoiceURL", reflect.TypeOf((*MockTelegramClientInterface)(nil).GetVoiceURL), message)
}
