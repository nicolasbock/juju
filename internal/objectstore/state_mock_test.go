// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/objectstore (interfaces: MongoSession)
//
// Generated by this command:
//
//	mockgen -package objectstore -destination state_mock_test.go github.com/juju/juju/internal/objectstore MongoSession
//

// Package objectstore is a generated GoMock package.
package objectstore

import (
	reflect "reflect"

	mgo "github.com/juju/mgo/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockMongoSession is a mock of MongoSession interface.
type MockMongoSession struct {
	ctrl     *gomock.Controller
	recorder *MockMongoSessionMockRecorder
}

// MockMongoSessionMockRecorder is the mock recorder for MockMongoSession.
type MockMongoSessionMockRecorder struct {
	mock *MockMongoSession
}

// NewMockMongoSession creates a new mock instance.
func NewMockMongoSession(ctrl *gomock.Controller) *MockMongoSession {
	mock := &MockMongoSession{ctrl: ctrl}
	mock.recorder = &MockMongoSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMongoSession) EXPECT() *MockMongoSessionMockRecorder {
	return m.recorder
}

// MongoSession mocks base method.
func (m *MockMongoSession) MongoSession() *mgo.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoSession")
	ret0, _ := ret[0].(*mgo.Session)
	return ret0
}

// MongoSession indicates an expected call of MongoSession.
func (mr *MockMongoSessionMockRecorder) MongoSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoSession", reflect.TypeOf((*MockMongoSession)(nil).MongoSession))
}
