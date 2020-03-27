// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/addons/addons.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockworkspaceService is a mock of workspaceService interface
type MockworkspaceService struct {
	ctrl     *gomock.Controller
	recorder *MockworkspaceServiceMockRecorder
}

// MockworkspaceServiceMockRecorder is the mock recorder for MockworkspaceService
type MockworkspaceServiceMockRecorder struct {
	mock *MockworkspaceService
}

// NewMockworkspaceService creates a new mock instance
func NewMockworkspaceService(ctrl *gomock.Controller) *MockworkspaceService {
	mock := &MockworkspaceService{ctrl: ctrl}
	mock.recorder = &MockworkspaceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockworkspaceService) EXPECT() *MockworkspaceServiceMockRecorder {
	return m.recorder
}

// ReadAddonsDir mocks base method
func (m *MockworkspaceService) ReadAddonsDir(appName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAddonsDir", appName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAddonsDir indicates an expected call of ReadAddonsDir
func (mr *MockworkspaceServiceMockRecorder) ReadAddonsDir(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAddonsDir", reflect.TypeOf((*MockworkspaceService)(nil).ReadAddonsDir), appName)
}

// ReadAddonsFile mocks base method
func (m *MockworkspaceService) ReadAddonsFile(appName, fileName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAddonsFile", appName, fileName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAddonsFile indicates an expected call of ReadAddonsFile
func (mr *MockworkspaceServiceMockRecorder) ReadAddonsFile(appName, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAddonsFile", reflect.TypeOf((*MockworkspaceService)(nil).ReadAddonsFile), appName, fileName)
}