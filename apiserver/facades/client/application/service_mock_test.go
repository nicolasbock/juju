// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/application (interfaces: MachineService,ApplicationService,StoragePoolGetter,NetworkService,ExternalControllerService,ModelConfigService,ModelAgentService)
//
// Generated by this command:
//
//	mockgen -typed -package application -destination service_mock_test.go github.com/juju/juju/apiserver/facades/client/application MachineService,ApplicationService,StoragePoolGetter,NetworkService,ExternalControllerService,ModelConfigService,ModelAgentService
//

// Package application is a generated GoMock package.
package application

import (
	context "context"
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	assumes "github.com/juju/juju/core/assumes"
	charm "github.com/juju/juju/core/charm"
	crossmodel "github.com/juju/juju/core/crossmodel"
	life "github.com/juju/juju/core/life"
	machine "github.com/juju/juju/core/machine"
	network "github.com/juju/juju/core/network"
	service "github.com/juju/juju/domain/application/service"
	config "github.com/juju/juju/environs/config"
	charm0 "github.com/juju/juju/internal/charm"
	storage "github.com/juju/juju/internal/storage"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// CreateMachine mocks base method.
func (m *MockMachineService) CreateMachine(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockMachineServiceMockRecorder) CreateMachine(arg0, arg1 any) *MockMachineServiceCreateMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockMachineService)(nil).CreateMachine), arg0, arg1)
	return &MockMachineServiceCreateMachineCall{Call: call}
}

// MockMachineServiceCreateMachineCall wrap *gomock.Call
type MockMachineServiceCreateMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceCreateMachineCall) Return(arg0 string, arg1 error) *MockMachineServiceCreateMachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceCreateMachineCall) Do(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceCreateMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceCreateMachineCall) DoAndReturn(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceCreateMachineCall {
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

// AddUnits mocks base method.
func (m *MockApplicationService) AddUnits(arg0 context.Context, arg1 string, arg2 ...service.AddUnitArg) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUnits", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUnits indicates an expected call of AddUnits.
func (mr *MockApplicationServiceMockRecorder) AddUnits(arg0, arg1 any, arg2 ...any) *MockApplicationServiceAddUnitsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnits", reflect.TypeOf((*MockApplicationService)(nil).AddUnits), varargs...)
	return &MockApplicationServiceAddUnitsCall{Call: call}
}

// MockApplicationServiceAddUnitsCall wrap *gomock.Call
type MockApplicationServiceAddUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceAddUnitsCall) Return(arg0 error) *MockApplicationServiceAddUnitsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceAddUnitsCall) Do(f func(context.Context, string, ...service.AddUnitArg) error) *MockApplicationServiceAddUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceAddUnitsCall) DoAndReturn(f func(context.Context, string, ...service.AddUnitArg) error) *MockApplicationServiceAddUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ChangeApplicationScale mocks base method.
func (m *MockApplicationService) ChangeApplicationScale(arg0 context.Context, arg1 string, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeApplicationScale", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeApplicationScale indicates an expected call of ChangeApplicationScale.
func (mr *MockApplicationServiceMockRecorder) ChangeApplicationScale(arg0, arg1, arg2 any) *MockApplicationServiceChangeApplicationScaleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeApplicationScale", reflect.TypeOf((*MockApplicationService)(nil).ChangeApplicationScale), arg0, arg1, arg2)
	return &MockApplicationServiceChangeApplicationScaleCall{Call: call}
}

// MockApplicationServiceChangeApplicationScaleCall wrap *gomock.Call
type MockApplicationServiceChangeApplicationScaleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceChangeApplicationScaleCall) Return(arg0 int, arg1 error) *MockApplicationServiceChangeApplicationScaleCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceChangeApplicationScaleCall) Do(f func(context.Context, string, int) (int, error)) *MockApplicationServiceChangeApplicationScaleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceChangeApplicationScaleCall) DoAndReturn(f func(context.Context, string, int) (int, error)) *MockApplicationServiceChangeApplicationScaleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateApplication mocks base method.
func (m *MockApplicationService) CreateApplication(arg0 context.Context, arg1 string, arg2 charm0.Charm, arg3 charm.Origin, arg4 service.AddApplicationArgs, arg5 ...service.AddUnitArg) (application.ID, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateApplication", varargs...)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationServiceMockRecorder) CreateApplication(arg0, arg1, arg2, arg3, arg4 any, arg5 ...any) *MockApplicationServiceCreateApplicationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3, arg4}, arg5...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationService)(nil).CreateApplication), varargs...)
	return &MockApplicationServiceCreateApplicationCall{Call: call}
}

// MockApplicationServiceCreateApplicationCall wrap *gomock.Call
type MockApplicationServiceCreateApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceCreateApplicationCall) Return(arg0 application.ID, arg1 error) *MockApplicationServiceCreateApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceCreateApplicationCall) Do(f func(context.Context, string, charm0.Charm, charm.Origin, service.AddApplicationArgs, ...service.AddUnitArg) (application.ID, error)) *MockApplicationServiceCreateApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceCreateApplicationCall) DoAndReturn(f func(context.Context, string, charm0.Charm, charm.Origin, service.AddApplicationArgs, ...service.AddUnitArg) (application.ID, error)) *MockApplicationServiceCreateApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteApplication mocks base method.
func (m *MockApplicationService) DeleteApplication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockApplicationServiceMockRecorder) DeleteApplication(arg0, arg1 any) *MockApplicationServiceDeleteApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockApplicationService)(nil).DeleteApplication), arg0, arg1)
	return &MockApplicationServiceDeleteApplicationCall{Call: call}
}

// MockApplicationServiceDeleteApplicationCall wrap *gomock.Call
type MockApplicationServiceDeleteApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceDeleteApplicationCall) Return(arg0 error) *MockApplicationServiceDeleteApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceDeleteApplicationCall) Do(f func(context.Context, string) error) *MockApplicationServiceDeleteApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceDeleteApplicationCall) DoAndReturn(f func(context.Context, string) error) *MockApplicationServiceDeleteApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DestroyApplication mocks base method.
func (m *MockApplicationService) DestroyApplication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyApplication indicates an expected call of DestroyApplication.
func (mr *MockApplicationServiceMockRecorder) DestroyApplication(arg0, arg1 any) *MockApplicationServiceDestroyApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyApplication", reflect.TypeOf((*MockApplicationService)(nil).DestroyApplication), arg0, arg1)
	return &MockApplicationServiceDestroyApplicationCall{Call: call}
}

// MockApplicationServiceDestroyApplicationCall wrap *gomock.Call
type MockApplicationServiceDestroyApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceDestroyApplicationCall) Return(arg0 error) *MockApplicationServiceDestroyApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceDestroyApplicationCall) Do(f func(context.Context, string) error) *MockApplicationServiceDestroyApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceDestroyApplicationCall) DoAndReturn(f func(context.Context, string) error) *MockApplicationServiceDestroyApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DestroyUnit mocks base method.
func (m *MockApplicationService) DestroyUnit(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyUnit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyUnit indicates an expected call of DestroyUnit.
func (mr *MockApplicationServiceMockRecorder) DestroyUnit(arg0, arg1 any) *MockApplicationServiceDestroyUnitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyUnit", reflect.TypeOf((*MockApplicationService)(nil).DestroyUnit), arg0, arg1)
	return &MockApplicationServiceDestroyUnitCall{Call: call}
}

// MockApplicationServiceDestroyUnitCall wrap *gomock.Call
type MockApplicationServiceDestroyUnitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceDestroyUnitCall) Return(arg0 error) *MockApplicationServiceDestroyUnitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceDestroyUnitCall) Do(f func(context.Context, string) error) *MockApplicationServiceDestroyUnitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceDestroyUnitCall) DoAndReturn(f func(context.Context, string) error) *MockApplicationServiceDestroyUnitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationLife mocks base method.
func (m *MockApplicationService) GetApplicationLife(arg0 context.Context, arg1 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationLife", arg0, arg1)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationLife indicates an expected call of GetApplicationLife.
func (mr *MockApplicationServiceMockRecorder) GetApplicationLife(arg0, arg1 any) *MockApplicationServiceGetApplicationLifeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationLife", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationLife), arg0, arg1)
	return &MockApplicationServiceGetApplicationLifeCall{Call: call}
}

// MockApplicationServiceGetApplicationLifeCall wrap *gomock.Call
type MockApplicationServiceGetApplicationLifeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationLifeCall) Return(arg0 life.Value, arg1 error) *MockApplicationServiceGetApplicationLifeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationLifeCall) Do(f func(context.Context, string) (life.Value, error)) *MockApplicationServiceGetApplicationLifeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationLifeCall) DoAndReturn(f func(context.Context, string) (life.Value, error)) *MockApplicationServiceGetApplicationLifeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSupportedFeatures mocks base method.
func (m *MockApplicationService) GetSupportedFeatures(arg0 context.Context) (assumes.FeatureSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedFeatures", arg0)
	ret0, _ := ret[0].(assumes.FeatureSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedFeatures indicates an expected call of GetSupportedFeatures.
func (mr *MockApplicationServiceMockRecorder) GetSupportedFeatures(arg0 any) *MockApplicationServiceGetSupportedFeaturesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedFeatures", reflect.TypeOf((*MockApplicationService)(nil).GetSupportedFeatures), arg0)
	return &MockApplicationServiceGetSupportedFeaturesCall{Call: call}
}

// MockApplicationServiceGetSupportedFeaturesCall wrap *gomock.Call
type MockApplicationServiceGetSupportedFeaturesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetSupportedFeaturesCall) Return(arg0 assumes.FeatureSet, arg1 error) *MockApplicationServiceGetSupportedFeaturesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetSupportedFeaturesCall) Do(f func(context.Context) (assumes.FeatureSet, error)) *MockApplicationServiceGetSupportedFeaturesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetSupportedFeaturesCall) DoAndReturn(f func(context.Context) (assumes.FeatureSet, error)) *MockApplicationServiceGetSupportedFeaturesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitLife mocks base method.
func (m *MockApplicationService) GetUnitLife(arg0 context.Context, arg1 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitLife", arg0, arg1)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitLife indicates an expected call of GetUnitLife.
func (mr *MockApplicationServiceMockRecorder) GetUnitLife(arg0, arg1 any) *MockApplicationServiceGetUnitLifeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitLife", reflect.TypeOf((*MockApplicationService)(nil).GetUnitLife), arg0, arg1)
	return &MockApplicationServiceGetUnitLifeCall{Call: call}
}

// MockApplicationServiceGetUnitLifeCall wrap *gomock.Call
type MockApplicationServiceGetUnitLifeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetUnitLifeCall) Return(arg0 life.Value, arg1 error) *MockApplicationServiceGetUnitLifeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetUnitLifeCall) Do(f func(context.Context, string) (life.Value, error)) *MockApplicationServiceGetUnitLifeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetUnitLifeCall) DoAndReturn(f func(context.Context, string) (life.Value, error)) *MockApplicationServiceGetUnitLifeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetApplicationScale mocks base method.
func (m *MockApplicationService) SetApplicationScale(arg0 context.Context, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetApplicationScale", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetApplicationScale indicates an expected call of SetApplicationScale.
func (mr *MockApplicationServiceMockRecorder) SetApplicationScale(arg0, arg1, arg2 any) *MockApplicationServiceSetApplicationScaleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApplicationScale", reflect.TypeOf((*MockApplicationService)(nil).SetApplicationScale), arg0, arg1, arg2)
	return &MockApplicationServiceSetApplicationScaleCall{Call: call}
}

// MockApplicationServiceSetApplicationScaleCall wrap *gomock.Call
type MockApplicationServiceSetApplicationScaleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceSetApplicationScaleCall) Return(arg0 error) *MockApplicationServiceSetApplicationScaleCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceSetApplicationScaleCall) Do(f func(context.Context, string, int) error) *MockApplicationServiceSetApplicationScaleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceSetApplicationScaleCall) DoAndReturn(f func(context.Context, string, int) error) *MockApplicationServiceSetApplicationScaleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateApplicationCharm mocks base method.
func (m *MockApplicationService) UpdateApplicationCharm(arg0 context.Context, arg1 string, arg2 service.UpdateCharmParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationCharm indicates an expected call of UpdateApplicationCharm.
func (mr *MockApplicationServiceMockRecorder) UpdateApplicationCharm(arg0, arg1, arg2 any) *MockApplicationServiceUpdateApplicationCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationCharm", reflect.TypeOf((*MockApplicationService)(nil).UpdateApplicationCharm), arg0, arg1, arg2)
	return &MockApplicationServiceUpdateApplicationCharmCall{Call: call}
}

// MockApplicationServiceUpdateApplicationCharmCall wrap *gomock.Call
type MockApplicationServiceUpdateApplicationCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceUpdateApplicationCharmCall) Return(arg0 error) *MockApplicationServiceUpdateApplicationCharmCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceUpdateApplicationCharmCall) Do(f func(context.Context, string, service.UpdateCharmParams) error) *MockApplicationServiceUpdateApplicationCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceUpdateApplicationCharmCall) DoAndReturn(f func(context.Context, string, service.UpdateCharmParams) error) *MockApplicationServiceUpdateApplicationCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockStoragePoolGetter is a mock of StoragePoolGetter interface.
type MockStoragePoolGetter struct {
	ctrl     *gomock.Controller
	recorder *MockStoragePoolGetterMockRecorder
}

// MockStoragePoolGetterMockRecorder is the mock recorder for MockStoragePoolGetter.
type MockStoragePoolGetterMockRecorder struct {
	mock *MockStoragePoolGetter
}

// NewMockStoragePoolGetter creates a new mock instance.
func NewMockStoragePoolGetter(ctrl *gomock.Controller) *MockStoragePoolGetter {
	mock := &MockStoragePoolGetter{ctrl: ctrl}
	mock.recorder = &MockStoragePoolGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoragePoolGetter) EXPECT() *MockStoragePoolGetterMockRecorder {
	return m.recorder
}

// GetStoragePoolByName mocks base method.
func (m *MockStoragePoolGetter) GetStoragePoolByName(arg0 context.Context, arg1 string) (*storage.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePoolByName", arg0, arg1)
	ret0, _ := ret[0].(*storage.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePoolByName indicates an expected call of GetStoragePoolByName.
func (mr *MockStoragePoolGetterMockRecorder) GetStoragePoolByName(arg0, arg1 any) *MockStoragePoolGetterGetStoragePoolByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolByName", reflect.TypeOf((*MockStoragePoolGetter)(nil).GetStoragePoolByName), arg0, arg1)
	return &MockStoragePoolGetterGetStoragePoolByNameCall{Call: call}
}

// MockStoragePoolGetterGetStoragePoolByNameCall wrap *gomock.Call
type MockStoragePoolGetterGetStoragePoolByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStoragePoolGetterGetStoragePoolByNameCall) Return(arg0 *storage.Config, arg1 error) *MockStoragePoolGetterGetStoragePoolByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStoragePoolGetterGetStoragePoolByNameCall) Do(f func(context.Context, string) (*storage.Config, error)) *MockStoragePoolGetterGetStoragePoolByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStoragePoolGetterGetStoragePoolByNameCall) DoAndReturn(f func(context.Context, string) (*storage.Config, error)) *MockStoragePoolGetterGetStoragePoolByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSpaces mocks base method.
func (m *MockNetworkService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockNetworkServiceMockRecorder) GetAllSpaces(arg0 any) *MockNetworkServiceGetAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockNetworkService)(nil).GetAllSpaces), arg0)
	return &MockNetworkServiceGetAllSpacesCall{Call: call}
}

// MockNetworkServiceGetAllSpacesCall wrap *gomock.Call
type MockNetworkServiceGetAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSpacesCall) Return(arg0 network.SpaceInfos, arg1 error) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSpacesCall) Do(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSpacesCall) DoAndReturn(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Space mocks base method.
func (m *MockNetworkService) Space(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Space", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Space indicates an expected call of Space.
func (mr *MockNetworkServiceMockRecorder) Space(arg0, arg1 any) *MockNetworkServiceSpaceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Space", reflect.TypeOf((*MockNetworkService)(nil).Space), arg0, arg1)
	return &MockNetworkServiceSpaceCall{Call: call}
}

// MockNetworkServiceSpaceCall wrap *gomock.Call
type MockNetworkServiceSpaceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceSpaceCall) Return(arg0 *network.SpaceInfo, arg1 error) *MockNetworkServiceSpaceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceSpaceCall) Do(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceSpaceCall) DoAndReturn(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SpaceByName mocks base method.
func (m *MockNetworkService) SpaceByName(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceByName", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceByName indicates an expected call of SpaceByName.
func (mr *MockNetworkServiceMockRecorder) SpaceByName(arg0, arg1 any) *MockNetworkServiceSpaceByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceByName", reflect.TypeOf((*MockNetworkService)(nil).SpaceByName), arg0, arg1)
	return &MockNetworkServiceSpaceByNameCall{Call: call}
}

// MockNetworkServiceSpaceByNameCall wrap *gomock.Call
type MockNetworkServiceSpaceByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceSpaceByNameCall) Return(arg0 *network.SpaceInfo, arg1 error) *MockNetworkServiceSpaceByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceSpaceByNameCall) Do(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceSpaceByNameCall) DoAndReturn(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockExternalControllerService is a mock of ExternalControllerService interface.
type MockExternalControllerService struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllerServiceMockRecorder
}

// MockExternalControllerServiceMockRecorder is the mock recorder for MockExternalControllerService.
type MockExternalControllerServiceMockRecorder struct {
	mock *MockExternalControllerService
}

// NewMockExternalControllerService creates a new mock instance.
func NewMockExternalControllerService(ctrl *gomock.Controller) *MockExternalControllerService {
	mock := &MockExternalControllerService{ctrl: ctrl}
	mock.recorder = &MockExternalControllerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalControllerService) EXPECT() *MockExternalControllerServiceMockRecorder {
	return m.recorder
}

// UpdateExternalController mocks base method.
func (m *MockExternalControllerService) UpdateExternalController(arg0 context.Context, arg1 crossmodel.ControllerInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalController indicates an expected call of UpdateExternalController.
func (mr *MockExternalControllerServiceMockRecorder) UpdateExternalController(arg0, arg1 any) *MockExternalControllerServiceUpdateExternalControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalController", reflect.TypeOf((*MockExternalControllerService)(nil).UpdateExternalController), arg0, arg1)
	return &MockExternalControllerServiceUpdateExternalControllerCall{Call: call}
}

// MockExternalControllerServiceUpdateExternalControllerCall wrap *gomock.Call
type MockExternalControllerServiceUpdateExternalControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExternalControllerServiceUpdateExternalControllerCall) Return(arg0 error) *MockExternalControllerServiceUpdateExternalControllerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExternalControllerServiceUpdateExternalControllerCall) Do(f func(context.Context, crossmodel.ControllerInfo) error) *MockExternalControllerServiceUpdateExternalControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExternalControllerServiceUpdateExternalControllerCall) DoAndReturn(f func(context.Context, crossmodel.ControllerInfo) error) *MockExternalControllerServiceUpdateExternalControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *MockModelConfigServiceModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
	return &MockModelConfigServiceModelConfigCall{Call: call}
}

// MockModelConfigServiceModelConfigCall wrap *gomock.Call
type MockModelConfigServiceModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelAgentService is a mock of ModelAgentService interface.
type MockModelAgentService struct {
	ctrl     *gomock.Controller
	recorder *MockModelAgentServiceMockRecorder
}

// MockModelAgentServiceMockRecorder is the mock recorder for MockModelAgentService.
type MockModelAgentServiceMockRecorder struct {
	mock *MockModelAgentService
}

// NewMockModelAgentService creates a new mock instance.
func NewMockModelAgentService(ctrl *gomock.Controller) *MockModelAgentService {
	mock := &MockModelAgentService{ctrl: ctrl}
	mock.recorder = &MockModelAgentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelAgentService) EXPECT() *MockModelAgentServiceMockRecorder {
	return m.recorder
}

// GetModelAgentVersion mocks base method.
func (m *MockModelAgentService) GetModelAgentVersion(arg0 context.Context) (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelAgentVersion", arg0)
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelAgentVersion indicates an expected call of GetModelAgentVersion.
func (mr *MockModelAgentServiceMockRecorder) GetModelAgentVersion(arg0 any) *MockModelAgentServiceGetModelAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelAgentVersion", reflect.TypeOf((*MockModelAgentService)(nil).GetModelAgentVersion), arg0)
	return &MockModelAgentServiceGetModelAgentVersionCall{Call: call}
}

// MockModelAgentServiceGetModelAgentVersionCall wrap *gomock.Call
type MockModelAgentServiceGetModelAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelAgentServiceGetModelAgentVersionCall) Return(arg0 version.Number, arg1 error) *MockModelAgentServiceGetModelAgentVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelAgentServiceGetModelAgentVersionCall) Do(f func(context.Context) (version.Number, error)) *MockModelAgentServiceGetModelAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelAgentServiceGetModelAgentVersionCall) DoAndReturn(f func(context.Context) (version.Number, error)) *MockModelAgentServiceGetModelAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
