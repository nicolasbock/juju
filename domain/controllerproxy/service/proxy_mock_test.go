// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/proxy (interfaces: Proxier)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination proxy_mock_test.go github.com/juju/juju/internal/proxy Proxier
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProxier is a mock of Proxier interface.
type MockProxier struct {
	ctrl     *gomock.Controller
	recorder *MockProxierMockRecorder
}

// MockProxierMockRecorder is the mock recorder for MockProxier.
type MockProxierMockRecorder struct {
	mock *MockProxier
}

// NewMockProxier creates a new mock instance.
func NewMockProxier(ctrl *gomock.Controller) *MockProxier {
	mock := &MockProxier{ctrl: ctrl}
	mock.recorder = &MockProxierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProxier) EXPECT() *MockProxierMockRecorder {
	return m.recorder
}

// Insecure mocks base method.
func (m *MockProxier) Insecure() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Insecure")
}

// Insecure indicates an expected call of Insecure.
func (mr *MockProxierMockRecorder) Insecure() *MockProxierInsecureCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insecure", reflect.TypeOf((*MockProxier)(nil).Insecure))
	return &MockProxierInsecureCall{Call: call}
}

// MockProxierInsecureCall wrap *gomock.Call
type MockProxierInsecureCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProxierInsecureCall) Return() *MockProxierInsecureCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProxierInsecureCall) Do(f func()) *MockProxierInsecureCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProxierInsecureCall) DoAndReturn(f func()) *MockProxierInsecureCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MarshalYAML mocks base method.
func (m *MockProxier) MarshalYAML() (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalYAML")
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalYAML indicates an expected call of MarshalYAML.
func (mr *MockProxierMockRecorder) MarshalYAML() *MockProxierMarshalYAMLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalYAML", reflect.TypeOf((*MockProxier)(nil).MarshalYAML))
	return &MockProxierMarshalYAMLCall{Call: call}
}

// MockProxierMarshalYAMLCall wrap *gomock.Call
type MockProxierMarshalYAMLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProxierMarshalYAMLCall) Return(arg0 any, arg1 error) *MockProxierMarshalYAMLCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProxierMarshalYAMLCall) Do(f func() (any, error)) *MockProxierMarshalYAMLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProxierMarshalYAMLCall) DoAndReturn(f func() (any, error)) *MockProxierMarshalYAMLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RawConfig mocks base method.
func (m *MockProxier) RawConfig() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawConfig")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawConfig indicates an expected call of RawConfig.
func (mr *MockProxierMockRecorder) RawConfig() *MockProxierRawConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawConfig", reflect.TypeOf((*MockProxier)(nil).RawConfig))
	return &MockProxierRawConfigCall{Call: call}
}

// MockProxierRawConfigCall wrap *gomock.Call
type MockProxierRawConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProxierRawConfigCall) Return(arg0 map[string]any, arg1 error) *MockProxierRawConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProxierRawConfigCall) Do(f func() (map[string]any, error)) *MockProxierRawConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProxierRawConfigCall) DoAndReturn(f func() (map[string]any, error)) *MockProxierRawConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Start mocks base method.
func (m *MockProxier) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockProxierMockRecorder) Start(arg0 any) *MockProxierStartCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProxier)(nil).Start), arg0)
	return &MockProxierStartCall{Call: call}
}

// MockProxierStartCall wrap *gomock.Call
type MockProxierStartCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProxierStartCall) Return(arg0 error) *MockProxierStartCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProxierStartCall) Do(f func(context.Context) error) *MockProxierStartCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProxierStartCall) DoAndReturn(f func(context.Context) error) *MockProxierStartCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *MockProxier) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockProxierMockRecorder) Stop() *MockProxierStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProxier)(nil).Stop))
	return &MockProxierStopCall{Call: call}
}

// MockProxierStopCall wrap *gomock.Call
type MockProxierStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProxierStopCall) Return() *MockProxierStopCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProxierStopCall) Do(f func()) *MockProxierStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProxierStopCall) DoAndReturn(f func()) *MockProxierStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Type mocks base method.
func (m *MockProxier) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockProxierMockRecorder) Type() *MockProxierTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockProxier)(nil).Type))
	return &MockProxierTypeCall{Call: call}
}

// MockProxierTypeCall wrap *gomock.Call
type MockProxierTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProxierTypeCall) Return(arg0 string) *MockProxierTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProxierTypeCall) Do(f func() string) *MockProxierTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProxierTypeCall) DoAndReturn(f func() string) *MockProxierTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
