// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/uniter/relation (interfaces: StateTrackerClient)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mock_state_tracker.go github.com/juju/juju/internal/worker/uniter/relation StateTrackerClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	api "github.com/juju/juju/internal/worker/uniter/api"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockStateTrackerClient is a mock of StateTrackerClient interface.
type MockStateTrackerClient struct {
	ctrl     *gomock.Controller
	recorder *MockStateTrackerClientMockRecorder
}

// MockStateTrackerClientMockRecorder is the mock recorder for MockStateTrackerClient.
type MockStateTrackerClientMockRecorder struct {
	mock *MockStateTrackerClient
}

// NewMockStateTrackerClient creates a new mock instance.
func NewMockStateTrackerClient(ctrl *gomock.Controller) *MockStateTrackerClient {
	mock := &MockStateTrackerClient{ctrl: ctrl}
	mock.recorder = &MockStateTrackerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateTrackerClient) EXPECT() *MockStateTrackerClientMockRecorder {
	return m.recorder
}

// Relation mocks base method.
func (m *MockStateTrackerClient) Relation(arg0 names.RelationTag) (api.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relation", arg0)
	ret0, _ := ret[0].(api.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation indicates an expected call of Relation.
func (mr *MockStateTrackerClientMockRecorder) Relation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation", reflect.TypeOf((*MockStateTrackerClient)(nil).Relation), arg0)
}

// RelationById mocks base method.
func (m *MockStateTrackerClient) RelationById(arg0 int) (api.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationById", arg0)
	ret0, _ := ret[0].(api.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelationById indicates an expected call of RelationById.
func (mr *MockStateTrackerClientMockRecorder) RelationById(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationById", reflect.TypeOf((*MockStateTrackerClient)(nil).RelationById), arg0)
}

// Unit mocks base method.
func (m *MockStateTrackerClient) Unit(arg0 names.UnitTag) (api.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(api.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockStateTrackerClientMockRecorder) Unit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockStateTrackerClient)(nil).Unit), arg0)
}
