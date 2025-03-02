// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: AccessRequester)

// Package internal is a generated GoMock package.
package internal

import (
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockAccessRequester is a mock of AccessRequester interface.
type MockAccessRequester struct {
	ctrl     *gomock.Controller
	recorder *MockAccessRequesterMockRecorder
}

// MockAccessRequesterMockRecorder is the mock recorder for MockAccessRequester.
type MockAccessRequesterMockRecorder struct {
	mock *MockAccessRequester
}

// NewMockAccessRequester creates a new mock instance.
func NewMockAccessRequester(ctrl *gomock.Controller) *MockAccessRequester {
	mock := &MockAccessRequester{ctrl: ctrl}
	mock.recorder = &MockAccessRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessRequester) EXPECT() *MockAccessRequesterMockRecorder {
	return m.recorder
}

// AppendRequestedScope mocks base method.
func (m *MockAccessRequester) AppendRequestedScope(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendRequestedScope", arg0)
}

// AppendRequestedScope indicates an expected call of AppendRequestedScope.
func (mr *MockAccessRequesterMockRecorder) AppendRequestedScope(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendRequestedScope", reflect.TypeOf((*MockAccessRequester)(nil).AppendRequestedScope), arg0)
}

// GetClient mocks base method.
func (m *MockAccessRequester) GetClient() fosite.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(fosite.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockAccessRequesterMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockAccessRequester)(nil).GetClient))
}

// GetGrantTypes mocks base method.
func (m *MockAccessRequester) GetGrantTypes() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrantTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantTypes indicates an expected call of GetGrantTypes.
func (mr *MockAccessRequesterMockRecorder) GetGrantTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantTypes", reflect.TypeOf((*MockAccessRequester)(nil).GetGrantTypes))
}

// GetGrantedAudience mocks base method.
func (m *MockAccessRequester) GetGrantedAudience() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrantedAudience")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantedAudience indicates an expected call of GetGrantedAudience.
func (mr *MockAccessRequesterMockRecorder) GetGrantedAudience() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantedAudience", reflect.TypeOf((*MockAccessRequester)(nil).GetGrantedAudience))
}

// GetGrantedScopes mocks base method.
func (m *MockAccessRequester) GetGrantedScopes() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrantedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantedScopes indicates an expected call of GetGrantedScopes.
func (mr *MockAccessRequesterMockRecorder) GetGrantedScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantedScopes", reflect.TypeOf((*MockAccessRequester)(nil).GetGrantedScopes))
}

// GetID mocks base method.
func (m *MockAccessRequester) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockAccessRequesterMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockAccessRequester)(nil).GetID))
}

// GetRequestForm mocks base method.
func (m *MockAccessRequester) GetRequestForm() url.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestForm")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

// GetRequestForm indicates an expected call of GetRequestForm.
func (mr *MockAccessRequesterMockRecorder) GetRequestForm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestForm", reflect.TypeOf((*MockAccessRequester)(nil).GetRequestForm))
}

// GetRequestedAt mocks base method.
func (m *MockAccessRequester) GetRequestedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetRequestedAt indicates an expected call of GetRequestedAt.
func (mr *MockAccessRequesterMockRecorder) GetRequestedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedAt", reflect.TypeOf((*MockAccessRequester)(nil).GetRequestedAt))
}

// GetRequestedAudience mocks base method.
func (m *MockAccessRequester) GetRequestedAudience() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedAudience")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetRequestedAudience indicates an expected call of GetRequestedAudience.
func (mr *MockAccessRequesterMockRecorder) GetRequestedAudience() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedAudience", reflect.TypeOf((*MockAccessRequester)(nil).GetRequestedAudience))
}

// GetRequestedScopes mocks base method.
func (m *MockAccessRequester) GetRequestedScopes() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetRequestedScopes indicates an expected call of GetRequestedScopes.
func (mr *MockAccessRequesterMockRecorder) GetRequestedScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedScopes", reflect.TypeOf((*MockAccessRequester)(nil).GetRequestedScopes))
}

// GetSession mocks base method.
func (m *MockAccessRequester) GetSession() fosite.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession")
	ret0, _ := ret[0].(fosite.Session)
	return ret0
}

// GetSession indicates an expected call of GetSession.
func (mr *MockAccessRequesterMockRecorder) GetSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockAccessRequester)(nil).GetSession))
}

// GrantAudience mocks base method.
func (m *MockAccessRequester) GrantAudience(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GrantAudience", arg0)
}

// GrantAudience indicates an expected call of GrantAudience.
func (mr *MockAccessRequesterMockRecorder) GrantAudience(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantAudience", reflect.TypeOf((*MockAccessRequester)(nil).GrantAudience), arg0)
}

// GrantScope mocks base method.
func (m *MockAccessRequester) GrantScope(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GrantScope", arg0)
}

// GrantScope indicates an expected call of GrantScope.
func (mr *MockAccessRequesterMockRecorder) GrantScope(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantScope", reflect.TypeOf((*MockAccessRequester)(nil).GrantScope), arg0)
}

// Merge mocks base method.
func (m *MockAccessRequester) Merge(arg0 fosite.Requester) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", arg0)
}

// Merge indicates an expected call of Merge.
func (mr *MockAccessRequesterMockRecorder) Merge(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockAccessRequester)(nil).Merge), arg0)
}

// Sanitize mocks base method.
func (m *MockAccessRequester) Sanitize(arg0 []string) fosite.Requester {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sanitize", arg0)
	ret0, _ := ret[0].(fosite.Requester)
	return ret0
}

// Sanitize indicates an expected call of Sanitize.
func (mr *MockAccessRequesterMockRecorder) Sanitize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sanitize", reflect.TypeOf((*MockAccessRequester)(nil).Sanitize), arg0)
}

// SetID mocks base method.
func (m *MockAccessRequester) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID.
func (mr *MockAccessRequesterMockRecorder) SetID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockAccessRequester)(nil).SetID), arg0)
}

// SetRequestedAudience mocks base method.
func (m *MockAccessRequester) SetRequestedAudience(arg0 fosite.Arguments) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequestedAudience", arg0)
}

// SetRequestedAudience indicates an expected call of SetRequestedAudience.
func (mr *MockAccessRequesterMockRecorder) SetRequestedAudience(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequestedAudience", reflect.TypeOf((*MockAccessRequester)(nil).SetRequestedAudience), arg0)
}

// SetRequestedScopes mocks base method.
func (m *MockAccessRequester) SetRequestedScopes(arg0 fosite.Arguments) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequestedScopes", arg0)
}

// SetRequestedScopes indicates an expected call of SetRequestedScopes.
func (mr *MockAccessRequesterMockRecorder) SetRequestedScopes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequestedScopes", reflect.TypeOf((*MockAccessRequester)(nil).SetRequestedScopes), arg0)
}

// SetSession mocks base method.
func (m *MockAccessRequester) SetSession(arg0 fosite.Session) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSession", arg0)
}

// SetSession indicates an expected call of SetSession.
func (mr *MockAccessRequesterMockRecorder) SetSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockAccessRequester)(nil).SetSession), arg0)
}
