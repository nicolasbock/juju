// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/upgradestepsmachine (interfaces: StatusSetter)

// Package upgradestepsmachine is a generated GoMock package.
package upgradestepsmachine

import (
	reflect "reflect"

	status "github.com/juju/juju/core/status"
	gomock "go.uber.org/mock/gomock"
)

// MockStatusSetter is a mock of StatusSetter interface.
type MockStatusSetter struct {
	ctrl     *gomock.Controller
	recorder *MockStatusSetterMockRecorder
}

// MockStatusSetterMockRecorder is the mock recorder for MockStatusSetter.
type MockStatusSetterMockRecorder struct {
	mock *MockStatusSetter
}

// NewMockStatusSetter creates a new mock instance.
func NewMockStatusSetter(ctrl *gomock.Controller) *MockStatusSetter {
	mock := &MockStatusSetter{ctrl: ctrl}
	mock.recorder = &MockStatusSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusSetter) EXPECT() *MockStatusSetterMockRecorder {
	return m.recorder
}

// SetStatus mocks base method.
func (m *MockStatusSetter) SetStatus(arg0 status.Status, arg1 string, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockStatusSetterMockRecorder) SetStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockStatusSetter)(nil).SetStatus), arg0, arg1, arg2)
}
