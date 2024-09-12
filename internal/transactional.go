// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/storage (interfaces: Transactional)
//
// Generated by this command:
//
//	mockgen -package internal -destination internal/transactional.go github.com/ory/fosite/storage Transactional
//
// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTransactional is a mock of Transactional interface.
type MockTransactional struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionalMockRecorder
}

// MockTransactionalMockRecorder is the mock recorder for MockTransactional.
type MockTransactionalMockRecorder struct {
	mock *MockTransactional
}

// NewMockTransactional creates a new mock instance.
func NewMockTransactional(ctrl *gomock.Controller) *MockTransactional {
	mock := &MockTransactional{ctrl: ctrl}
	mock.recorder = &MockTransactionalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactional) EXPECT() *MockTransactionalMockRecorder {
	return m.recorder
}

// BeginTX mocks base method.
func (m *MockTransactional) BeginTX(arg0 context.Context) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTX", arg0)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTX indicates an expected call of BeginTX.
func (mr *MockTransactionalMockRecorder) BeginTX(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTX", reflect.TypeOf((*MockTransactional)(nil).BeginTX), arg0)
}

// Commit mocks base method.
func (m *MockTransactional) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionalMockRecorder) Commit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransactional)(nil).Commit), arg0)
}

// Rollback mocks base method.
func (m *MockTransactional) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionalMockRecorder) Rollback(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransactional)(nil).Rollback), arg0)
}
