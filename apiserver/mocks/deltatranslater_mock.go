// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver (interfaces: DeltaTranslater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	params "github.com/juju/juju/apiserver/params"
	multiwatcher "github.com/juju/juju/core/multiwatcher"
)

// MockDeltaTranslater is a mock of DeltaTranslater interface.
type MockDeltaTranslater struct {
	ctrl     *gomock.Controller
	recorder *MockDeltaTranslaterMockRecorder
}

// MockDeltaTranslaterMockRecorder is the mock recorder for MockDeltaTranslater.
type MockDeltaTranslaterMockRecorder struct {
	mock *MockDeltaTranslater
}

// NewMockDeltaTranslater creates a new mock instance.
func NewMockDeltaTranslater(ctrl *gomock.Controller) *MockDeltaTranslater {
	mock := &MockDeltaTranslater{ctrl: ctrl}
	mock.recorder = &MockDeltaTranslaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeltaTranslater) EXPECT() *MockDeltaTranslaterMockRecorder {
	return m.recorder
}

// TranslateAction mocks base method.
func (m *MockDeltaTranslater) TranslateAction(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateAction", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateAction indicates an expected call of TranslateAction.
func (mr *MockDeltaTranslaterMockRecorder) TranslateAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateAction", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateAction), arg0)
}

// TranslateAnnotation mocks base method.
func (m *MockDeltaTranslater) TranslateAnnotation(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateAnnotation", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateAnnotation indicates an expected call of TranslateAnnotation.
func (mr *MockDeltaTranslaterMockRecorder) TranslateAnnotation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateAnnotation", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateAnnotation), arg0)
}

// TranslateApplication mocks base method.
func (m *MockDeltaTranslater) TranslateApplication(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateApplication", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateApplication indicates an expected call of TranslateApplication.
func (mr *MockDeltaTranslaterMockRecorder) TranslateApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateApplication", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateApplication), arg0)
}

// TranslateApplicationOffer mocks base method.
func (m *MockDeltaTranslater) TranslateApplicationOffer(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateApplicationOffer", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateApplicationOffer indicates an expected call of TranslateApplicationOffer.
func (mr *MockDeltaTranslaterMockRecorder) TranslateApplicationOffer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateApplicationOffer", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateApplicationOffer), arg0)
}

// TranslateBlock mocks base method.
func (m *MockDeltaTranslater) TranslateBlock(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateBlock", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateBlock indicates an expected call of TranslateBlock.
func (mr *MockDeltaTranslaterMockRecorder) TranslateBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateBlock", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateBlock), arg0)
}

// TranslateBranch mocks base method.
func (m *MockDeltaTranslater) TranslateBranch(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateBranch", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateBranch indicates an expected call of TranslateBranch.
func (mr *MockDeltaTranslaterMockRecorder) TranslateBranch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateBranch", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateBranch), arg0)
}

// TranslateCharm mocks base method.
func (m *MockDeltaTranslater) TranslateCharm(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateCharm", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateCharm indicates an expected call of TranslateCharm.
func (mr *MockDeltaTranslaterMockRecorder) TranslateCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateCharm", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateCharm), arg0)
}

// TranslateMachine mocks base method.
func (m *MockDeltaTranslater) TranslateMachine(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateMachine", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateMachine indicates an expected call of TranslateMachine.
func (mr *MockDeltaTranslaterMockRecorder) TranslateMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateMachine", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateMachine), arg0)
}

// TranslateModel mocks base method.
func (m *MockDeltaTranslater) TranslateModel(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateModel", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateModel indicates an expected call of TranslateModel.
func (mr *MockDeltaTranslaterMockRecorder) TranslateModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateModel", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateModel), arg0)
}

// TranslateRelation mocks base method.
func (m *MockDeltaTranslater) TranslateRelation(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateRelation", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateRelation indicates an expected call of TranslateRelation.
func (mr *MockDeltaTranslaterMockRecorder) TranslateRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateRelation", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateRelation), arg0)
}

// TranslateRemoteApplication mocks base method.
func (m *MockDeltaTranslater) TranslateRemoteApplication(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateRemoteApplication", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateRemoteApplication indicates an expected call of TranslateRemoteApplication.
func (mr *MockDeltaTranslaterMockRecorder) TranslateRemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateRemoteApplication", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateRemoteApplication), arg0)
}

// TranslateUnit mocks base method.
func (m *MockDeltaTranslater) TranslateUnit(arg0 multiwatcher.EntityInfo) params.EntityInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateUnit", arg0)
	ret0, _ := ret[0].(params.EntityInfo)
	return ret0
}

// TranslateUnit indicates an expected call of TranslateUnit.
func (mr *MockDeltaTranslaterMockRecorder) TranslateUnit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateUnit", reflect.TypeOf((*MockDeltaTranslater)(nil).TranslateUnit), arg0)
}
