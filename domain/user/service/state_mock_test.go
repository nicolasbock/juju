// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/user/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -package service -destination state_mock_test.go github.com/juju/juju/domain/user/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	user "github.com/juju/juju/core/user"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockState) AddUser(arg0 context.Context, arg1 user.UUID, arg2 user.User, arg3 user.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockStateMockRecorder) AddUser(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockState)(nil).AddUser), arg0, arg1, arg2, arg3)
}

// AddUserWithActivationKey mocks base method.
func (m *MockState) AddUserWithActivationKey(arg0 context.Context, arg1 user.UUID, arg2 user.User, arg3 user.UUID, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserWithActivationKey", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserWithActivationKey indicates an expected call of AddUserWithActivationKey.
func (mr *MockStateMockRecorder) AddUserWithActivationKey(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserWithActivationKey", reflect.TypeOf((*MockState)(nil).AddUserWithActivationKey), arg0, arg1, arg2, arg3, arg4)
}

// AddUserWithPasswordHash mocks base method.
func (m *MockState) AddUserWithPasswordHash(arg0 context.Context, arg1 user.UUID, arg2 user.User, arg3 user.UUID, arg4 string, arg5 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserWithPasswordHash", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserWithPasswordHash indicates an expected call of AddUserWithPasswordHash.
func (mr *MockStateMockRecorder) AddUserWithPasswordHash(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserWithPasswordHash", reflect.TypeOf((*MockState)(nil).AddUserWithPasswordHash), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DisableUserAuthentication mocks base method.
func (m *MockState) DisableUserAuthentication(arg0 context.Context, arg1 user.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUserAuthentication indicates an expected call of DisableUserAuthentication.
func (mr *MockStateMockRecorder) DisableUserAuthentication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUserAuthentication", reflect.TypeOf((*MockState)(nil).DisableUserAuthentication), arg0, arg1)
}

// EnableUserAuthentication mocks base method.
func (m *MockState) EnableUserAuthentication(arg0 context.Context, arg1 user.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUserAuthentication indicates an expected call of EnableUserAuthentication.
func (mr *MockStateMockRecorder) EnableUserAuthentication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUserAuthentication", reflect.TypeOf((*MockState)(nil).EnableUserAuthentication), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockState) GetUser(arg0 context.Context, arg1 user.UUID) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStateMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockState)(nil).GetUser), arg0, arg1)
}

// GetUserByName mocks base method.
func (m *MockState) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockStateMockRecorder) GetUserByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockState)(nil).GetUserByName), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockState) RemoveUser(arg0 context.Context, arg1 user.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockStateMockRecorder) RemoveUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockState)(nil).RemoveUser), arg0, arg1)
}

// SetActivationKey mocks base method.
func (m *MockState) SetActivationKey(arg0 context.Context, arg1 user.UUID, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetActivationKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetActivationKey indicates an expected call of SetActivationKey.
func (mr *MockStateMockRecorder) SetActivationKey(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetActivationKey", reflect.TypeOf((*MockState)(nil).SetActivationKey), arg0, arg1, arg2)
}

// SetPasswordHash mocks base method.
func (m *MockState) SetPasswordHash(arg0 context.Context, arg1 user.UUID, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPasswordHash", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPasswordHash indicates an expected call of SetPasswordHash.
func (mr *MockStateMockRecorder) SetPasswordHash(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPasswordHash", reflect.TypeOf((*MockState)(nil).SetPasswordHash), arg0, arg1, arg2, arg3)
}
