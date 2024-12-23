// Code generated by MockGen. DO NOT EDIT.
// Source: ./port.go

// Package mock_core_port is a generated GoMock package.
package mock_core_port

import (
	reflect "reflect"

	core_port "github.com/eduardobcolombo/learning-grpc/cmd/client/core/port"
	portpb "github.com/eduardobcolombo/learning-grpc/portpb"
	gomock "github.com/golang/mock/gomock"
)

// MockCoreService is a mock of CoreService interface.
type MockCoreService struct {
	ctrl     *gomock.Controller
	recorder *MockCoreServiceMockRecorder
}

// MockCoreServiceMockRecorder is the mock recorder for MockCoreService.
type MockCoreServiceMockRecorder struct {
	mock *MockCoreService
}

// NewMockCoreService creates a new mock instance.
func NewMockCoreService(ctrl *gomock.Controller) *MockCoreService {
	mock := &MockCoreService{ctrl: ctrl}
	mock.recorder = &MockCoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoreService) EXPECT() *MockCoreServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCoreService) Create(port core_port.PortCore) (*core_port.PortCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", port)
	ret0, _ := ret[0].(*core_port.PortCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCoreServiceMockRecorder) Create(port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCoreService)(nil).Create), port)
}

// GetByUnloc mocks base method.
func (m *MockCoreService) GetByUnloc(unloc string) (*portpb.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUnloc", unloc)
	ret0, _ := ret[0].(*portpb.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUnloc indicates an expected call of GetByUnloc.
func (mr *MockCoreServiceMockRecorder) GetByUnloc(unloc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUnloc", reflect.TypeOf((*MockCoreService)(nil).GetByUnloc), unloc)
}

// Retrieve mocks base method.
func (m *MockCoreService) Retrieve() ([]*portpb.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve")
	ret0, _ := ret[0].([]*portpb.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve.
func (mr *MockCoreServiceMockRecorder) Retrieve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockCoreService)(nil).Retrieve))
}

// Update mocks base method.
func (m *MockCoreService) Update(fileName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", fileName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCoreServiceMockRecorder) Update(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCoreService)(nil).Update), fileName)
}