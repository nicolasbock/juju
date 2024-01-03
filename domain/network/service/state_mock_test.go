// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/network/service (interfaces: State,Logger)
//
// Generated by this command:
//
//	mockgen -package service -destination state_mock_test.go github.com/juju/juju/domain/network/service State,Logger
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	network "github.com/juju/juju/core/network"
	utils "github.com/juju/utils/v3"
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

// AddSpace mocks base method.
func (m *MockState) AddSpace(arg0 context.Context, arg1, arg2 string, arg3 network.Id, arg4 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSpace", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSpace indicates an expected call of AddSpace.
func (mr *MockStateMockRecorder) AddSpace(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSpace", reflect.TypeOf((*MockState)(nil).AddSpace), arg0, arg1, arg2, arg3, arg4)
}

// AddSubnet mocks base method.
func (m *MockState) AddSubnet(arg0 context.Context, arg1 utils.UUID, arg2 string, arg3, arg4 network.Id, arg5 int, arg6 []string, arg7 string, arg8 *network.FanCIDRs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubnet", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSubnet indicates an expected call of AddSubnet.
func (mr *MockStateMockRecorder) AddSubnet(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockState)(nil).AddSubnet), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// DeleteSpace mocks base method.
func (m *MockState) DeleteSpace(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSpace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSpace indicates an expected call of DeleteSpace.
func (mr *MockStateMockRecorder) DeleteSpace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSpace", reflect.TypeOf((*MockState)(nil).DeleteSpace), arg0, arg1)
}

// GetAllSpaces mocks base method.
func (m *MockState) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockStateMockRecorder) GetAllSpaces(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockState)(nil).GetAllSpaces), arg0)
}

// GetSpace mocks base method.
func (m *MockState) GetSpace(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpace", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpace indicates an expected call of GetSpace.
func (mr *MockStateMockRecorder) GetSpace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpace", reflect.TypeOf((*MockState)(nil).GetSpace), arg0, arg1)
}

// GetSpaceByName mocks base method.
func (m *MockState) GetSpaceByName(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpaceByName", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpaceByName indicates an expected call of GetSpaceByName.
func (mr *MockStateMockRecorder) GetSpaceByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpaceByName", reflect.TypeOf((*MockState)(nil).GetSpaceByName), arg0, arg1)
}

// UpdateSpace mocks base method.
func (m *MockState) UpdateSpace(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSpace", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSpace indicates an expected call of UpdateSpace.
func (mr *MockStateMockRecorder) UpdateSpace(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSpace", reflect.TypeOf((*MockState)(nil).UpdateSpace), arg0, arg1, arg2)
}

// UpdateSubnet mocks base method.
func (m *MockState) UpdateSubnet(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubnet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubnet indicates an expected call of UpdateSubnet.
func (mr *MockStateMockRecorder) UpdateSubnet(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubnet", reflect.TypeOf((*MockState)(nil).UpdateSubnet), arg0, arg1, arg2)
}

// UpsertSubnets mocks base method.
func (m *MockState) UpsertSubnets(arg0 context.Context, arg1 []network.SubnetInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSubnets", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSubnets indicates an expected call of UpsertSubnets.
func (mr *MockStateMockRecorder) UpsertSubnets(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSubnets", reflect.TypeOf((*MockState)(nil).UpsertSubnets), arg0, arg1)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}
