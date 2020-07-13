// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/whatthefar/monorepo-toolkit/pkg/interface/controller (interfaces: CI)

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCI is a mock of CI interface
type MockCI struct {
	ctrl     *gomock.Controller
	recorder *MockCIMockRecorder
}

// MockCIMockRecorder is the mock recorder for MockCI
type MockCIMockRecorder struct {
	mock *MockCI
}

// NewMockCI creates a new mock instance
func NewMockCI(ctrl *gomock.Controller) *MockCI {
	mock := &MockCI{ctrl: ctrl}
	mock.recorder = &MockCIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCI) EXPECT() *MockCIMockRecorder {
	return m.recorder
}

// ListProjects mocks base method
func (m *MockCI) ListProjects(arg0 context.Context, arg1 []string, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects
func (mr *MockCIMockRecorder) ListProjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockCI)(nil).ListProjects), arg0, arg1, arg2)
}

// ListProjectsJoined mocks base method
func (m *MockCI) ListProjectsJoined(arg0 context.Context, arg1 []string, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsJoined", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsJoined indicates an expected call of ListProjectsJoined
func (mr *MockCIMockRecorder) ListProjectsJoined(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsJoined", reflect.TypeOf((*MockCI)(nil).ListProjectsJoined), arg0, arg1, arg2)
}