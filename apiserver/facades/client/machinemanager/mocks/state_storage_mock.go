// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state/binarystorage (interfaces: StorageCloser)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/state_storage_mock.go github.com/juju/juju/state/binarystorage StorageCloser
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	binarystorage "github.com/juju/juju/state/binarystorage"
	gomock "go.uber.org/mock/gomock"
)

// MockStorageCloser is a mock of StorageCloser interface.
type MockStorageCloser struct {
	ctrl     *gomock.Controller
	recorder *MockStorageCloserMockRecorder
}

// MockStorageCloserMockRecorder is the mock recorder for MockStorageCloser.
type MockStorageCloserMockRecorder struct {
	mock *MockStorageCloser
}

// NewMockStorageCloser creates a new mock instance.
func NewMockStorageCloser(ctrl *gomock.Controller) *MockStorageCloser {
	mock := &MockStorageCloser{ctrl: ctrl}
	mock.recorder = &MockStorageCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageCloser) EXPECT() *MockStorageCloserMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStorageCloser) Add(arg0 context.Context, arg1 io.Reader, arg2 binarystorage.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStorageCloserMockRecorder) Add(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStorageCloser)(nil).Add), arg0, arg1, arg2)
}

// AllMetadata mocks base method.
func (m *MockStorageCloser) AllMetadata() ([]binarystorage.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMetadata")
	ret0, _ := ret[0].([]binarystorage.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMetadata indicates an expected call of AllMetadata.
func (mr *MockStorageCloserMockRecorder) AllMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMetadata", reflect.TypeOf((*MockStorageCloser)(nil).AllMetadata))
}

// Close mocks base method.
func (m *MockStorageCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageCloser)(nil).Close))
}

// Metadata mocks base method.
func (m *MockStorageCloser) Metadata(arg0 string) (binarystorage.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", arg0)
	ret0, _ := ret[0].(binarystorage.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata.
func (mr *MockStorageCloserMockRecorder) Metadata(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockStorageCloser)(nil).Metadata), arg0)
}

// Open mocks base method.
func (m *MockStorageCloser) Open(arg0 context.Context, arg1 string) (binarystorage.Metadata, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(binarystorage.Metadata)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Open indicates an expected call of Open.
func (mr *MockStorageCloserMockRecorder) Open(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockStorageCloser)(nil).Open), arg0, arg1)
}
