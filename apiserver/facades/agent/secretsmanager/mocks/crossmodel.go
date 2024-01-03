// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/secretsmanager (interfaces: CrossModelState,CrossModelSecretsClient)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/crossmodel.go github.com/juju/juju/apiserver/facades/agent/secretsmanager CrossModelState,CrossModelSecretsClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	secrets0 "github.com/juju/juju/internal/secrets"
	provider "github.com/juju/juju/internal/secrets/provider"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
	macaroon "gopkg.in/macaroon.v2"
)

// MockCrossModelState is a mock of CrossModelState interface.
type MockCrossModelState struct {
	ctrl     *gomock.Controller
	recorder *MockCrossModelStateMockRecorder
}

// MockCrossModelStateMockRecorder is the mock recorder for MockCrossModelState.
type MockCrossModelStateMockRecorder struct {
	mock *MockCrossModelState
}

// NewMockCrossModelState creates a new mock instance.
func NewMockCrossModelState(ctrl *gomock.Controller) *MockCrossModelState {
	mock := &MockCrossModelState{ctrl: ctrl}
	mock.recorder = &MockCrossModelStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossModelState) EXPECT() *MockCrossModelStateMockRecorder {
	return m.recorder
}

// GetMacaroon mocks base method.
func (m *MockCrossModelState) GetMacaroon(arg0 names.Tag) (*macaroon.Macaroon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMacaroon", arg0)
	ret0, _ := ret[0].(*macaroon.Macaroon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMacaroon indicates an expected call of GetMacaroon.
func (mr *MockCrossModelStateMockRecorder) GetMacaroon(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMacaroon", reflect.TypeOf((*MockCrossModelState)(nil).GetMacaroon), arg0)
}

// GetRemoteEntity mocks base method.
func (m *MockCrossModelState) GetRemoteEntity(arg0 string) (names.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteEntity", arg0)
	ret0, _ := ret[0].(names.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteEntity indicates an expected call of GetRemoteEntity.
func (mr *MockCrossModelStateMockRecorder) GetRemoteEntity(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteEntity", reflect.TypeOf((*MockCrossModelState)(nil).GetRemoteEntity), arg0)
}

// GetToken mocks base method.
func (m *MockCrossModelState) GetToken(arg0 names.Tag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockCrossModelStateMockRecorder) GetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockCrossModelState)(nil).GetToken), arg0)
}

// MockCrossModelSecretsClient is a mock of CrossModelSecretsClient interface.
type MockCrossModelSecretsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCrossModelSecretsClientMockRecorder
}

// MockCrossModelSecretsClientMockRecorder is the mock recorder for MockCrossModelSecretsClient.
type MockCrossModelSecretsClientMockRecorder struct {
	mock *MockCrossModelSecretsClient
}

// NewMockCrossModelSecretsClient creates a new mock instance.
func NewMockCrossModelSecretsClient(ctrl *gomock.Controller) *MockCrossModelSecretsClient {
	mock := &MockCrossModelSecretsClient{ctrl: ctrl}
	mock.recorder = &MockCrossModelSecretsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossModelSecretsClient) EXPECT() *MockCrossModelSecretsClientMockRecorder {
	return m.recorder
}

// GetRemoteSecretContentInfo mocks base method.
func (m *MockCrossModelSecretsClient) GetRemoteSecretContentInfo(arg0 *secrets.URI, arg1 int, arg2, arg3 bool, arg4 string, arg5 int, arg6 macaroon.Slice) (*secrets0.ContentParams, *provider.ModelBackendConfig, int, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteSecretContentInfo", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*secrets0.ContentParams)
	ret1, _ := ret[1].(*provider.ModelBackendConfig)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(bool)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetRemoteSecretContentInfo indicates an expected call of GetRemoteSecretContentInfo.
func (mr *MockCrossModelSecretsClientMockRecorder) GetRemoteSecretContentInfo(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteSecretContentInfo", reflect.TypeOf((*MockCrossModelSecretsClient)(nil).GetRemoteSecretContentInfo), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetSecretAccessScope mocks base method.
func (m *MockCrossModelSecretsClient) GetSecretAccessScope(arg0 *secrets.URI, arg1 string, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretAccessScope", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretAccessScope indicates an expected call of GetSecretAccessScope.
func (mr *MockCrossModelSecretsClientMockRecorder) GetSecretAccessScope(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretAccessScope", reflect.TypeOf((*MockCrossModelSecretsClient)(nil).GetSecretAccessScope), arg0, arg1, arg2)
}
