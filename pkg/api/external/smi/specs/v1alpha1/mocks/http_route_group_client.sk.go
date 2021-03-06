// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/smi/specs/v1alpha1/http_route_group_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1alpha1 "github.com/solo-io/supergloo/pkg/api/external/smi/specs/v1alpha1"
)

// MockHTTPRouteGroupWatcher is a mock of HTTPRouteGroupWatcher interface
type MockHTTPRouteGroupWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupWatcherMockRecorder
}

// MockHTTPRouteGroupWatcherMockRecorder is the mock recorder for MockHTTPRouteGroupWatcher
type MockHTTPRouteGroupWatcherMockRecorder struct {
	mock *MockHTTPRouteGroupWatcher
}

// NewMockHTTPRouteGroupWatcher creates a new mock instance
func NewMockHTTPRouteGroupWatcher(ctrl *gomock.Controller) *MockHTTPRouteGroupWatcher {
	mock := &MockHTTPRouteGroupWatcher{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupWatcher) EXPECT() *MockHTTPRouteGroupWatcherMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockHTTPRouteGroupWatcher) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha1.HTTPRouteGroupList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha1.HTTPRouteGroupList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockHTTPRouteGroupWatcherMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockHTTPRouteGroupWatcher)(nil).Watch), namespace, opts)
}

// MockHTTPRouteGroupClient is a mock of HTTPRouteGroupClient interface
type MockHTTPRouteGroupClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupClientMockRecorder
}

// MockHTTPRouteGroupClientMockRecorder is the mock recorder for MockHTTPRouteGroupClient
type MockHTTPRouteGroupClientMockRecorder struct {
	mock *MockHTTPRouteGroupClient
}

// NewMockHTTPRouteGroupClient creates a new mock instance
func NewMockHTTPRouteGroupClient(ctrl *gomock.Controller) *MockHTTPRouteGroupClient {
	mock := &MockHTTPRouteGroupClient{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupClient) EXPECT() *MockHTTPRouteGroupClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockHTTPRouteGroupClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockHTTPRouteGroupClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockHTTPRouteGroupClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockHTTPRouteGroupClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).Register))
}

// Read mocks base method
func (m *MockHTTPRouteGroupClient) Read(namespace, name string, opts clients.ReadOpts) (*v1alpha1.HTTPRouteGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", namespace, name, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPRouteGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockHTTPRouteGroupClientMockRecorder) Read(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).Read), namespace, name, opts)
}

// Write mocks base method
func (m *MockHTTPRouteGroupClient) Write(resource *v1alpha1.HTTPRouteGroup, opts clients.WriteOpts) (*v1alpha1.HTTPRouteGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPRouteGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockHTTPRouteGroupClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockHTTPRouteGroupClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockHTTPRouteGroupClientMockRecorder) Delete(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).Delete), namespace, name, opts)
}

// List mocks base method
func (m *MockHTTPRouteGroupClient) List(namespace string, opts clients.ListOpts) (v1alpha1.HTTPRouteGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace, opts)
	ret0, _ := ret[0].(v1alpha1.HTTPRouteGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockHTTPRouteGroupClientMockRecorder) List(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).List), namespace, opts)
}

// Watch mocks base method
func (m *MockHTTPRouteGroupClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha1.HTTPRouteGroupList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha1.HTTPRouteGroupList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockHTTPRouteGroupClientMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).Watch), namespace, opts)
}
