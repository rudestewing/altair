// Code generated by MockGen. DO NOT EDIT.
// Source: core/formatter.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entity "github.com/codefluence-x/altair/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOauthApplicationFormater is a mock of OauthApplicationFormater interface
type MockOauthApplicationFormater struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationFormaterMockRecorder
}

// MockOauthApplicationFormaterMockRecorder is the mock recorder for MockOauthApplicationFormater
type MockOauthApplicationFormaterMockRecorder struct {
	mock *MockOauthApplicationFormater
}

// NewMockOauthApplicationFormater creates a new mock instance
func NewMockOauthApplicationFormater(ctrl *gomock.Controller) *MockOauthApplicationFormater {
	mock := &MockOauthApplicationFormater{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationFormaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthApplicationFormater) EXPECT() *MockOauthApplicationFormaterMockRecorder {
	return m.recorder
}

// ApplicationList mocks base method
func (m *MockOauthApplicationFormater) ApplicationList(ctx context.Context, applications []entity.OauthApplication) []entity.OauthApplicationJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationList", ctx, applications)
	ret0, _ := ret[0].([]entity.OauthApplicationJSON)
	return ret0
}

// ApplicationList indicates an expected call of ApplicationList
func (mr *MockOauthApplicationFormaterMockRecorder) ApplicationList(ctx, applications interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationList", reflect.TypeOf((*MockOauthApplicationFormater)(nil).ApplicationList), ctx, applications)
}

// Application mocks base method
func (m *MockOauthApplicationFormater) Application(ctx context.Context, application entity.OauthApplication) entity.OauthApplicationJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", ctx, application)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	return ret0
}

// Application indicates an expected call of Application
func (mr *MockOauthApplicationFormaterMockRecorder) Application(ctx, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockOauthApplicationFormater)(nil).Application), ctx, application)
}

// MockOauthFormatter is a mock of OauthFormatter interface
type MockOauthFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockOauthFormatterMockRecorder
}

// MockOauthFormatterMockRecorder is the mock recorder for MockOauthFormatter
type MockOauthFormatterMockRecorder struct {
	mock *MockOauthFormatter
}

// NewMockOauthFormatter creates a new mock instance
func NewMockOauthFormatter(ctrl *gomock.Controller) *MockOauthFormatter {
	mock := &MockOauthFormatter{ctrl: ctrl}
	mock.recorder = &MockOauthFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthFormatter) EXPECT() *MockOauthFormatterMockRecorder {
	return m.recorder
}

// AccessGrant mocks base method
func (m *MockOauthFormatter) AccessGrant(r entity.AuthorizationRequestJSON, e entity.OauthAccessGrant) entity.OauthAccessGrantJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGrant", r, e)
	ret0, _ := ret[0].(entity.OauthAccessGrantJSON)
	return ret0
}

// AccessGrant indicates an expected call of AccessGrant
func (mr *MockOauthFormatterMockRecorder) AccessGrant(r, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGrant", reflect.TypeOf((*MockOauthFormatter)(nil).AccessGrant), r, e)
}

// AccessToken mocks base method
func (m *MockOauthFormatter) AccessToken(r entity.AuthorizationRequestJSON, e entity.OauthAccessToken) entity.OauthAccessTokenJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken", r, e)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	return ret0
}

// AccessToken indicates an expected call of AccessToken
func (mr *MockOauthFormatterMockRecorder) AccessToken(r, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockOauthFormatter)(nil).AccessToken), r, e)
}
