// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver (interfaces: ApplicationServiceGetter,ApplicationService)
//
// Generated by this command:
//
//	mockgen -typed -package apiserver -destination application_service_mock_test.go github.com/juju/juju/apiserver ApplicationServiceGetter,ApplicationService
//

// Package apiserver is a generated GoMock package.
package apiserver

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockApplicationServiceGetter is a mock of ApplicationServiceGetter interface.
type MockApplicationServiceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceGetterMockRecorder
}

// MockApplicationServiceGetterMockRecorder is the mock recorder for MockApplicationServiceGetter.
type MockApplicationServiceGetterMockRecorder struct {
	mock *MockApplicationServiceGetter
}

// NewMockApplicationServiceGetter creates a new mock instance.
func NewMockApplicationServiceGetter(ctrl *gomock.Controller) *MockApplicationServiceGetter {
	mock := &MockApplicationServiceGetter{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationServiceGetter) EXPECT() *MockApplicationServiceGetterMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockApplicationServiceGetter) Application(arg0 context.Context) (ApplicationService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(ApplicationService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockApplicationServiceGetterMockRecorder) Application(arg0 any) *MockApplicationServiceGetterApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockApplicationServiceGetter)(nil).Application), arg0)
	return &MockApplicationServiceGetterApplicationCall{Call: call}
}

// MockApplicationServiceGetterApplicationCall wrap *gomock.Call
type MockApplicationServiceGetterApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetterApplicationCall) Return(arg0 ApplicationService, arg1 error) *MockApplicationServiceGetterApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetterApplicationCall) Do(f func(context.Context) (ApplicationService, error)) *MockApplicationServiceGetterApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetterApplicationCall) DoAndReturn(f func(context.Context) (ApplicationService, error)) *MockApplicationServiceGetterApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetCharmArchiveBySHA256Prefix mocks base method.
func (m *MockApplicationService) GetCharmArchiveBySHA256Prefix(arg0 context.Context, arg1 string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmArchiveBySHA256Prefix", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmArchiveBySHA256Prefix indicates an expected call of GetCharmArchiveBySHA256Prefix.
func (mr *MockApplicationServiceMockRecorder) GetCharmArchiveBySHA256Prefix(arg0, arg1 any) *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmArchiveBySHA256Prefix", reflect.TypeOf((*MockApplicationService)(nil).GetCharmArchiveBySHA256Prefix), arg0, arg1)
	return &MockApplicationServiceGetCharmArchiveBySHA256PrefixCall{Call: call}
}

// MockApplicationServiceGetCharmArchiveBySHA256PrefixCall wrap *gomock.Call
type MockApplicationServiceGetCharmArchiveBySHA256PrefixCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall) Return(arg0 io.ReadCloser, arg1 error) *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall) Do(f func(context.Context, string) (io.ReadCloser, error)) *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall) DoAndReturn(f func(context.Context, string) (io.ReadCloser, error)) *MockApplicationServiceGetCharmArchiveBySHA256PrefixCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
