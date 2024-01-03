// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/modelcmd (interfaces: ModelConfigAPI)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/modelconfig_mock.go github.com/juju/juju/cmd/modelcmd ModelConfigAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockModelConfigAPI is a mock of ModelConfigAPI interface.
type MockModelConfigAPI struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigAPIMockRecorder
}

// MockModelConfigAPIMockRecorder is the mock recorder for MockModelConfigAPI.
type MockModelConfigAPIMockRecorder struct {
	mock *MockModelConfigAPI
}

// NewMockModelConfigAPI creates a new mock instance.
func NewMockModelConfigAPI(ctrl *gomock.Controller) *MockModelConfigAPI {
	mock := &MockModelConfigAPI{ctrl: ctrl}
	mock.recorder = &MockModelConfigAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigAPI) EXPECT() *MockModelConfigAPIMockRecorder {
	return m.recorder
}

// ModelGet mocks base method.
func (m *MockModelConfigAPI) ModelGet() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet.
func (mr *MockModelConfigAPIMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockModelConfigAPI)(nil).ModelGet))
}
