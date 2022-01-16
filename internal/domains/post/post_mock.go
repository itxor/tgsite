package post

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/itxor/tgsite/pkg/telegram"
)

// MockPostUseCaseForUpdateTelegramLoopInterface is a mock of PostUseCaseForUpdateTelegramLoopInterface interface.
type MockPostUseCaseForUpdateTelegramLoopInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder
}

// MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder is the mock recorder for MockPostUseCaseForUpdateTelegramLoopInterface.
type MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder struct {
	mock *MockPostUseCaseForUpdateTelegramLoopInterface
}

// NewMockPostUseCaseForUpdateTelegramLoopInterface creates a new mock instance.
func NewMockPostUseCaseForUpdateTelegramLoopInterface(ctrl *gomock.Controller) *MockPostUseCaseForUpdateTelegramLoopInterface {
	mock := &MockPostUseCaseForUpdateTelegramLoopInterface{ctrl: ctrl}
	mock.recorder = &MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostUseCaseForUpdateTelegramLoopInterface) EXPECT() *MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder {
	return m.recorder
}

// BuildNewPostFromMessage mocks base method.
func (m *MockPostUseCaseForUpdateTelegramLoopInterface) BuildNewPostFromMessage(dto telegram.MessageDTO) (*Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildNewPostFromMessage", dto)
	ret0, _ := ret[0].(*Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildNewPostFromMessage indicates an expected call of BuildNewPostFromMessage.
func (mr *MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder) BuildNewPostFromMessage(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildNewPostFromMessage", reflect.TypeOf((*MockPostUseCaseForUpdateTelegramLoopInterface)(nil).BuildNewPostFromMessage), dto)
}

// DispatchAddPost mocks base method.
func (m *MockPostUseCaseForUpdateTelegramLoopInterface) DispatchAddPost(post Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchAddPost", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchAddPost indicates an expected call of DispatchAddPost.
func (mr *MockPostUseCaseForUpdateTelegramLoopInterfaceMockRecorder) DispatchAddPost(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchAddPost", reflect.TypeOf((*MockPostUseCaseForUpdateTelegramLoopInterface)(nil).DispatchAddPost), post)
}

// MockPostUseCaseForSubscribeNewPostsInterface is a mock of PostUseCaseForSubscribeNewPostsInterface interface.
type MockPostUseCaseForSubscribeNewPostsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPostUseCaseForSubscribeNewPostsInterfaceMockRecorder
}

// MockPostUseCaseForSubscribeNewPostsInterfaceMockRecorder is the mock recorder for MockPostUseCaseForSubscribeNewPostsInterface.
type MockPostUseCaseForSubscribeNewPostsInterfaceMockRecorder struct {
	mock *MockPostUseCaseForSubscribeNewPostsInterface
}

// NewMockPostUseCaseForSubscribeNewPostsInterface creates a new mock instance.
func NewMockPostUseCaseForSubscribeNewPostsInterface(ctrl *gomock.Controller) *MockPostUseCaseForSubscribeNewPostsInterface {
	mock := &MockPostUseCaseForSubscribeNewPostsInterface{ctrl: ctrl}
	mock.recorder = &MockPostUseCaseForSubscribeNewPostsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostUseCaseForSubscribeNewPostsInterface) EXPECT() *MockPostUseCaseForSubscribeNewPostsInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPostUseCaseForSubscribeNewPostsInterface) Add(post Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPostUseCaseForSubscribeNewPostsInterfaceMockRecorder) Add(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPostUseCaseForSubscribeNewPostsInterface)(nil).Add), post)
}

// MockPostRepositoryInterface is a mock of PostRepositoryInterface interface.
type MockPostRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryInterfaceMockRecorder
}

// MockPostRepositoryInterfaceMockRecorder is the mock recorder for MockPostRepositoryInterface.
type MockPostRepositoryInterfaceMockRecorder struct {
	mock *MockPostRepositoryInterface
}

// NewMockPostRepositoryInterface creates a new mock instance.
func NewMockPostRepositoryInterface(ctrl *gomock.Controller) *MockPostRepositoryInterface {
	mock := &MockPostRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepositoryInterface) EXPECT() *MockPostRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPostRepositoryInterface) Add(post Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPostRepositoryInterfaceMockRecorder) Add(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPostRepositoryInterface)(nil).Add), post)
}
