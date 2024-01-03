// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/uniter/relation (interfaces: SubordinateDestroyer)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mock_subordinate_destroyer.go github.com/juju/juju/internal/worker/uniter/relation SubordinateDestroyer
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSubordinateDestroyer is a mock of SubordinateDestroyer interface.
type MockSubordinateDestroyer struct {
	ctrl     *gomock.Controller
	recorder *MockSubordinateDestroyerMockRecorder
}

// MockSubordinateDestroyerMockRecorder is the mock recorder for MockSubordinateDestroyer.
type MockSubordinateDestroyerMockRecorder struct {
	mock *MockSubordinateDestroyer
}

// NewMockSubordinateDestroyer creates a new mock instance.
func NewMockSubordinateDestroyer(ctrl *gomock.Controller) *MockSubordinateDestroyer {
	mock := &MockSubordinateDestroyer{ctrl: ctrl}
	mock.recorder = &MockSubordinateDestroyerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubordinateDestroyer) EXPECT() *MockSubordinateDestroyerMockRecorder {
	return m.recorder
}

// DestroyAllSubordinates mocks base method.
func (m *MockSubordinateDestroyer) DestroyAllSubordinates() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyAllSubordinates")
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyAllSubordinates indicates an expected call of DestroyAllSubordinates.
func (mr *MockSubordinateDestroyerMockRecorder) DestroyAllSubordinates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyAllSubordinates", reflect.TypeOf((*MockSubordinateDestroyer)(nil).DestroyAllSubordinates))
}
