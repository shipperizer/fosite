// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: ClientCredentialsGrantStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockClientCredentialsGrantStorage is a mock of ClientCredentialsGrantStorage interface.
type MockClientCredentialsGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *MockClientCredentialsGrantStorageMockRecorder
}

// MockClientCredentialsGrantStorageMockRecorder is the mock recorder for MockClientCredentialsGrantStorage.
type MockClientCredentialsGrantStorageMockRecorder struct {
	mock *MockClientCredentialsGrantStorage
}

// NewMockClientCredentialsGrantStorage creates a new mock instance.
func NewMockClientCredentialsGrantStorage(ctrl *gomock.Controller) *MockClientCredentialsGrantStorage {
	mock := &MockClientCredentialsGrantStorage{ctrl: ctrl}
	mock.recorder = &MockClientCredentialsGrantStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientCredentialsGrantStorage) EXPECT() *MockClientCredentialsGrantStorageMockRecorder {
	return m.recorder
}

// CreateAccessTokenSession mocks base method.
func (m *MockClientCredentialsGrantStorage) CreateAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessTokenSession indicates an expected call of CreateAccessTokenSession.
func (mr *MockClientCredentialsGrantStorageMockRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessTokenSession", reflect.TypeOf((*MockClientCredentialsGrantStorage)(nil).CreateAccessTokenSession), arg0, arg1, arg2)
}

// DeleteAccessTokenSession mocks base method.
func (m *MockClientCredentialsGrantStorage) DeleteAccessTokenSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessTokenSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessTokenSession indicates an expected call of DeleteAccessTokenSession.
func (mr *MockClientCredentialsGrantStorageMockRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessTokenSession", reflect.TypeOf((*MockClientCredentialsGrantStorage)(nil).DeleteAccessTokenSession), arg0, arg1)
}

// GetAccessTokenSession mocks base method.
func (m *MockClientCredentialsGrantStorage) GetAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessTokenSession indicates an expected call of GetAccessTokenSession.
func (mr *MockClientCredentialsGrantStorageMockRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenSession", reflect.TypeOf((*MockClientCredentialsGrantStorage)(nil).GetAccessTokenSession), arg0, arg1, arg2)
}
