// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/generate/schemagen/gen (interfaces: APIServer,Registry,PackageRegistry,Linker)
//
// Generated by this command:
//
//	mockgen -package gen -destination describeapi_mock.go -write_package_comment=false github.com/juju/juju/generate/schemagen/gen APIServer,Registry,PackageRegistry,Linker
package gen

import (
	reflect "reflect"

	facade "github.com/juju/juju/apiserver/facade"
	gomock "go.uber.org/mock/gomock"
	packages "golang.org/x/tools/go/packages"
)

// MockAPIServer is a mock of APIServer interface.
type MockAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockAPIServerMockRecorder
}

// MockAPIServerMockRecorder is the mock recorder for MockAPIServer.
type MockAPIServerMockRecorder struct {
	mock *MockAPIServer
}

// NewMockAPIServer creates a new mock instance.
func NewMockAPIServer(ctrl *gomock.Controller) *MockAPIServer {
	mock := &MockAPIServer{ctrl: ctrl}
	mock.recorder = &MockAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIServer) EXPECT() *MockAPIServerMockRecorder {
	return m.recorder
}

// AdminFacadeDetails mocks base method.
func (m *MockAPIServer) AdminFacadeDetails() []facade.Details {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminFacadeDetails")
	ret0, _ := ret[0].([]facade.Details)
	return ret0
}

// AdminFacadeDetails indicates an expected call of AdminFacadeDetails.
func (mr *MockAPIServerMockRecorder) AdminFacadeDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminFacadeDetails", reflect.TypeOf((*MockAPIServer)(nil).AdminFacadeDetails))
}

// AllFacades mocks base method.
func (m *MockAPIServer) AllFacades() Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllFacades")
	ret0, _ := ret[0].(Registry)
	return ret0
}

// AllFacades indicates an expected call of AllFacades.
func (mr *MockAPIServerMockRecorder) AllFacades() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllFacades", reflect.TypeOf((*MockAPIServer)(nil).AllFacades))
}

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// GetType mocks base method.
func (m *MockRegistry) GetType(arg0 string, arg1 int) (reflect.Type, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType", arg0, arg1)
	ret0, _ := ret[0].(reflect.Type)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetType indicates an expected call of GetType.
func (mr *MockRegistryMockRecorder) GetType(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockRegistry)(nil).GetType), arg0, arg1)
}

// List mocks base method.
func (m *MockRegistry) List() []facade.Description {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]facade.Description)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRegistryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRegistry)(nil).List))
}

// ListDetails mocks base method.
func (m *MockRegistry) ListDetails() []facade.Details {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDetails")
	ret0, _ := ret[0].([]facade.Details)
	return ret0
}

// ListDetails indicates an expected call of ListDetails.
func (mr *MockRegistryMockRecorder) ListDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetails", reflect.TypeOf((*MockRegistry)(nil).ListDetails))
}

// MockPackageRegistry is a mock of PackageRegistry interface.
type MockPackageRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockPackageRegistryMockRecorder
}

// MockPackageRegistryMockRecorder is the mock recorder for MockPackageRegistry.
type MockPackageRegistryMockRecorder struct {
	mock *MockPackageRegistry
}

// NewMockPackageRegistry creates a new mock instance.
func NewMockPackageRegistry(ctrl *gomock.Controller) *MockPackageRegistry {
	mock := &MockPackageRegistry{ctrl: ctrl}
	mock.recorder = &MockPackageRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageRegistry) EXPECT() *MockPackageRegistryMockRecorder {
	return m.recorder
}

// LoadPackage mocks base method.
func (m *MockPackageRegistry) LoadPackage() (*packages.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPackage")
	ret0, _ := ret[0].(*packages.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPackage indicates an expected call of LoadPackage.
func (mr *MockPackageRegistryMockRecorder) LoadPackage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPackage", reflect.TypeOf((*MockPackageRegistry)(nil).LoadPackage))
}

// MockLinker is a mock of Linker interface.
type MockLinker struct {
	ctrl     *gomock.Controller
	recorder *MockLinkerMockRecorder
}

// MockLinkerMockRecorder is the mock recorder for MockLinker.
type MockLinkerMockRecorder struct {
	mock *MockLinker
}

// NewMockLinker creates a new mock instance.
func NewMockLinker(ctrl *gomock.Controller) *MockLinker {
	mock := &MockLinker{ctrl: ctrl}
	mock.recorder = &MockLinkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinker) EXPECT() *MockLinkerMockRecorder {
	return m.recorder
}

// Links mocks base method.
func (m *MockLinker) Links(arg0 string, arg1 facade.MultiModelFactory) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Links", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Links indicates an expected call of Links.
func (mr *MockLinkerMockRecorder) Links(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Links", reflect.TypeOf((*MockLinker)(nil).Links), arg0, arg1)
}
