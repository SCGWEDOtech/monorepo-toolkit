// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/whatthefar/monorepo-toolkit/pkg/interactor (interfaces: ListChangesInteractor)

// Package mock_interactor is a generated GoMock package.
package mock_interactor

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockListChangesInteractor is a mock of ListChangesInteractor interface
type MockListChangesInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockListChangesInteractorMockRecorder
}

// MockListChangesInteractorMockRecorder is the mock recorder for MockListChangesInteractor
type MockListChangesInteractorMockRecorder struct {
	mock *MockListChangesInteractor
}

// NewMockListChangesInteractor creates a new mock instance
func NewMockListChangesInteractor(ctrl *gomock.Controller) *MockListChangesInteractor {
	mock := &MockListChangesInteractor{ctrl: ctrl}
	mock.recorder = &MockListChangesInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListChangesInteractor) EXPECT() *MockListChangesInteractorMockRecorder {
	return m.recorder
}

// ListChanges mocks base method
func (m *MockListChangesInteractor) ListChanges(arg0 context.Context, arg1 []string, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChanges", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChanges indicates an expected call of ListChanges
func (mr *MockListChangesInteractorMockRecorder) ListChanges(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChanges", reflect.TypeOf((*MockListChangesInteractor)(nil).ListChanges), arg0, arg1, arg2)
}