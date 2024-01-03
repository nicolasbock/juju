// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1 (interfaces: ApiextensionsV1beta1Interface,CustomResourceDefinitionInterface)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/apiextensionsv1beta1_mock.go -mock_names=CustomResourceDefinitionInterface=MockCustomResourceDefinitionV1Beta1Interface k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1 ApiextensionsV1beta1Interface,CustomResourceDefinitionInterface
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	v1beta10 "k8s.io/apiextensions-apiserver/pkg/client/applyconfiguration/apiextensions/v1beta1"
	v1beta11 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MockApiextensionsV1beta1Interface is a mock of ApiextensionsV1beta1Interface interface.
type MockApiextensionsV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockApiextensionsV1beta1InterfaceMockRecorder
}

// MockApiextensionsV1beta1InterfaceMockRecorder is the mock recorder for MockApiextensionsV1beta1Interface.
type MockApiextensionsV1beta1InterfaceMockRecorder struct {
	mock *MockApiextensionsV1beta1Interface
}

// NewMockApiextensionsV1beta1Interface creates a new mock instance.
func NewMockApiextensionsV1beta1Interface(ctrl *gomock.Controller) *MockApiextensionsV1beta1Interface {
	mock := &MockApiextensionsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockApiextensionsV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiextensionsV1beta1Interface) EXPECT() *MockApiextensionsV1beta1InterfaceMockRecorder {
	return m.recorder
}

// CustomResourceDefinitions mocks base method.
func (m *MockApiextensionsV1beta1Interface) CustomResourceDefinitions() v1beta11.CustomResourceDefinitionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomResourceDefinitions")
	ret0, _ := ret[0].(v1beta11.CustomResourceDefinitionInterface)
	return ret0
}

// CustomResourceDefinitions indicates an expected call of CustomResourceDefinitions.
func (mr *MockApiextensionsV1beta1InterfaceMockRecorder) CustomResourceDefinitions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomResourceDefinitions", reflect.TypeOf((*MockApiextensionsV1beta1Interface)(nil).CustomResourceDefinitions))
}

// RESTClient mocks base method.
func (m *MockApiextensionsV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockApiextensionsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockApiextensionsV1beta1Interface)(nil).RESTClient))
}

// MockCustomResourceDefinitionV1Beta1Interface is a mock of CustomResourceDefinitionInterface interface.
type MockCustomResourceDefinitionV1Beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder
}

// MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder is the mock recorder for MockCustomResourceDefinitionV1Beta1Interface.
type MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder struct {
	mock *MockCustomResourceDefinitionV1Beta1Interface
}

// NewMockCustomResourceDefinitionV1Beta1Interface creates a new mock instance.
func NewMockCustomResourceDefinitionV1Beta1Interface(ctrl *gomock.Controller) *MockCustomResourceDefinitionV1Beta1Interface {
	mock := &MockCustomResourceDefinitionV1Beta1Interface{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionV1Beta1Interface) EXPECT() *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Apply(arg0 context.Context, arg1 *v1beta10.CustomResourceDefinitionApplyConfiguration, arg2 v1.ApplyOptions) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Apply(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) ApplyStatus(arg0 context.Context, arg1 *v1beta10.CustomResourceDefinitionApplyConfiguration, arg2 v1.ApplyOptions) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Create(arg0 context.Context, arg1 *v1beta1.CustomResourceDefinition, arg2 v1.CreateOptions) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Create(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Delete(arg0 context.Context, arg1 string, arg2 v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) DeleteCollection(arg0 context.Context, arg1 v1.DeleteOptions, arg2 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Get(arg0 context.Context, arg1 string, arg2 v1.GetOptions) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) List(arg0 context.Context, arg1 v1.ListOptions) (*v1beta1.CustomResourceDefinitionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinitionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v1.PatchOptions, arg5 ...string) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 any, arg5 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Update(arg0 context.Context, arg1 *v1beta1.CustomResourceDefinition, arg2 v1.UpdateOptions) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) UpdateStatus(arg0 context.Context, arg1 *v1beta1.CustomResourceDefinition, arg2 v1.UpdateOptions) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockCustomResourceDefinitionV1Beta1Interface) Watch(arg0 context.Context, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockCustomResourceDefinitionV1Beta1InterfaceMockRecorder) Watch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCustomResourceDefinitionV1Beta1Interface)(nil).Watch), arg0, arg1)
}
