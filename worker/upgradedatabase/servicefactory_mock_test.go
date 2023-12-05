// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/servicefactory (interfaces: ControllerServiceFactory)

// Package upgradedatabase is a generated GoMock package.
package upgradedatabase

import (
	reflect "reflect"

	service "github.com/juju/juju/domain/autocert/service"
	service0 "github.com/juju/juju/domain/cloud/service"
	service1 "github.com/juju/juju/domain/controllerconfig/service"
	service2 "github.com/juju/juju/domain/controllernode/service"
	service3 "github.com/juju/juju/domain/credential/service"
	service4 "github.com/juju/juju/domain/externalcontroller/service"
	service5 "github.com/juju/juju/domain/flag/service"
	service6 "github.com/juju/juju/domain/model/service"
	service7 "github.com/juju/juju/domain/modeldefaults/service"
	service8 "github.com/juju/juju/domain/modelmanager/service"
	service9 "github.com/juju/juju/domain/objectstore/service"
	service10 "github.com/juju/juju/domain/upgrade/service"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerServiceFactory is a mock of ControllerServiceFactory interface.
type MockControllerServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockControllerServiceFactoryMockRecorder
}

// MockControllerServiceFactoryMockRecorder is the mock recorder for MockControllerServiceFactory.
type MockControllerServiceFactoryMockRecorder struct {
	mock *MockControllerServiceFactory
}

// NewMockControllerServiceFactory creates a new mock instance.
func NewMockControllerServiceFactory(ctrl *gomock.Controller) *MockControllerServiceFactory {
	mock := &MockControllerServiceFactory{ctrl: ctrl}
	mock.recorder = &MockControllerServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerServiceFactory) EXPECT() *MockControllerServiceFactoryMockRecorder {
	return m.recorder
}

// AgentObjectStore mocks base method.
func (m *MockControllerServiceFactory) AgentObjectStore() *service9.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentObjectStore")
	ret0, _ := ret[0].(*service9.Service)
	return ret0
}

// AgentObjectStore indicates an expected call of AgentObjectStore.
func (mr *MockControllerServiceFactoryMockRecorder) AgentObjectStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentObjectStore", reflect.TypeOf((*MockControllerServiceFactory)(nil).AgentObjectStore))
}

// AutocertCache mocks base method.
func (m *MockControllerServiceFactory) AutocertCache() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockControllerServiceFactoryMockRecorder) AutocertCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockControllerServiceFactory)(nil).AutocertCache))
}

// Cloud mocks base method.
func (m *MockControllerServiceFactory) Cloud() *service0.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(*service0.Service)
	return ret0
}

// Cloud indicates an expected call of Cloud.
func (mr *MockControllerServiceFactoryMockRecorder) Cloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockControllerServiceFactory)(nil).Cloud))
}

// ControllerConfig mocks base method.
func (m *MockControllerServiceFactory) ControllerConfig() *service1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service1.Service)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerConfig))
}

// ControllerNode mocks base method.
func (m *MockControllerServiceFactory) ControllerNode() *service2.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service2.Service)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerNode))
}

// Credential mocks base method.
func (m *MockControllerServiceFactory) Credential() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockControllerServiceFactoryMockRecorder) Credential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockControllerServiceFactory)(nil).Credential))
}

// ExternalController mocks base method.
func (m *MockControllerServiceFactory) ExternalController() *service4.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service4.Service)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockControllerServiceFactoryMockRecorder) ExternalController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockControllerServiceFactory)(nil).ExternalController))
}

// Flag mocks base method.
func (m *MockControllerServiceFactory) Flag() *service5.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flag")
	ret0, _ := ret[0].(*service5.Service)
	return ret0
}

// Flag indicates an expected call of Flag.
func (mr *MockControllerServiceFactoryMockRecorder) Flag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flag", reflect.TypeOf((*MockControllerServiceFactory)(nil).Flag))
}

// Model mocks base method.
func (m *MockControllerServiceFactory) Model() *service6.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*service6.Service)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockControllerServiceFactoryMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockControllerServiceFactory)(nil).Model))
}

// ModelDefaults mocks base method.
func (m *MockControllerServiceFactory) ModelDefaults() *service7.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelDefaults")
	ret0, _ := ret[0].(*service7.Service)
	return ret0
}

// ModelDefaults indicates an expected call of ModelDefaults.
func (mr *MockControllerServiceFactoryMockRecorder) ModelDefaults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelDefaults", reflect.TypeOf((*MockControllerServiceFactory)(nil).ModelDefaults))
}

// ModelManager mocks base method.
func (m *MockControllerServiceFactory) ModelManager() *service8.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelManager")
	ret0, _ := ret[0].(*service8.Service)
	return ret0
}

// ModelManager indicates an expected call of ModelManager.
func (mr *MockControllerServiceFactoryMockRecorder) ModelManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelManager", reflect.TypeOf((*MockControllerServiceFactory)(nil).ModelManager))
}

// Upgrade mocks base method.
func (m *MockControllerServiceFactory) Upgrade() *service10.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade")
	ret0, _ := ret[0].(*service10.Service)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockControllerServiceFactoryMockRecorder) Upgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockControllerServiceFactory)(nil).Upgrade))
}
