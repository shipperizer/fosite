// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: DeviceUserResponder)

// Package internal is a generated GoMock package.
package internal

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeviceUserResponder is a mock of DeviceUserResponder interface.
type MockDeviceUserResponder struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceUserResponderMockRecorder
}

// MockDeviceUserResponderMockRecorder is the mock recorder for MockDeviceUserResponder.
type MockDeviceUserResponderMockRecorder struct {
	mock *MockDeviceUserResponder
}

// NewMockDeviceUserResponder creates a new mock instance.
func NewMockDeviceUserResponder(ctrl *gomock.Controller) *MockDeviceUserResponder {
	mock := &MockDeviceUserResponder{ctrl: ctrl}
	mock.recorder = &MockDeviceUserResponderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceUserResponder) EXPECT() *MockDeviceUserResponderMockRecorder {
	return m.recorder
}

// AddHeader mocks base method.
func (m *MockDeviceUserResponder) AddHeader(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHeader", arg0, arg1)
}

// AddHeader indicates an expected call of AddHeader.
func (mr *MockDeviceUserResponderMockRecorder) AddHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHeader", reflect.TypeOf((*MockDeviceUserResponder)(nil).AddHeader), arg0, arg1)
}

// GetHeader mocks base method.
func (m *MockDeviceUserResponder) GetHeader() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockDeviceUserResponderMockRecorder) GetHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockDeviceUserResponder)(nil).GetHeader))
}
