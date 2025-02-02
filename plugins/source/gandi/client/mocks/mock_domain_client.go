// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/gandi/client (interfaces: DomainClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/go-gandi/go-gandi/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockDomainClient is a mock of DomainClient interface.
type MockDomainClient struct {
	ctrl     *gomock.Controller
	recorder *MockDomainClientMockRecorder
}

// MockDomainClientMockRecorder is the mock recorder for MockDomainClient.
type MockDomainClientMockRecorder struct {
	mock *MockDomainClient
}

// NewMockDomainClient creates a new mock instance.
func NewMockDomainClient(ctrl *gomock.Controller) *MockDomainClient {
	mock := &MockDomainClient{ctrl: ctrl}
	mock.recorder = &MockDomainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainClient) EXPECT() *MockDomainClientMockRecorder {
	return m.recorder
}

// GetDomain mocks base method.
func (m *MockDomainClient) GetDomain(arg0 string) (domain.Details, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomain", arg0)
	ret0, _ := ret[0].(domain.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomain indicates an expected call of GetDomain.
func (mr *MockDomainClientMockRecorder) GetDomain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomain", reflect.TypeOf((*MockDomainClient)(nil).GetDomain), arg0)
}

// GetLiveDNS mocks base method.
func (m *MockDomainClient) GetLiveDNS(arg0 string) (domain.LiveDNS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveDNS", arg0)
	ret0, _ := ret[0].(domain.LiveDNS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLiveDNS indicates an expected call of GetLiveDNS.
func (mr *MockDomainClientMockRecorder) GetLiveDNS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveDNS", reflect.TypeOf((*MockDomainClient)(nil).GetLiveDNS), arg0)
}

// ListDNSSECKeys mocks base method.
func (m *MockDomainClient) ListDNSSECKeys(arg0 string) ([]domain.DNSSECKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDNSSECKeys", arg0)
	ret0, _ := ret[0].([]domain.DNSSECKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDNSSECKeys indicates an expected call of ListDNSSECKeys.
func (mr *MockDomainClientMockRecorder) ListDNSSECKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDNSSECKeys", reflect.TypeOf((*MockDomainClient)(nil).ListDNSSECKeys), arg0)
}

// ListDomains mocks base method.
func (m *MockDomainClient) ListDomains() ([]domain.ListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDomains")
	ret0, _ := ret[0].([]domain.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomains indicates an expected call of ListDomains.
func (mr *MockDomainClientMockRecorder) ListDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomains", reflect.TypeOf((*MockDomainClient)(nil).ListDomains))
}

// ListGlueRecords mocks base method.
func (m *MockDomainClient) ListGlueRecords(arg0 string) ([]domain.GlueRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGlueRecords", arg0)
	ret0, _ := ret[0].([]domain.GlueRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGlueRecords indicates an expected call of ListGlueRecords.
func (mr *MockDomainClientMockRecorder) ListGlueRecords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGlueRecords", reflect.TypeOf((*MockDomainClient)(nil).ListGlueRecords), arg0)
}

// ListWebRedirections mocks base method.
func (m *MockDomainClient) ListWebRedirections(arg0 string) ([]domain.WebRedirection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWebRedirections", arg0)
	ret0, _ := ret[0].([]domain.WebRedirection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebRedirections indicates an expected call of ListWebRedirections.
func (mr *MockDomainClientMockRecorder) ListWebRedirections(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebRedirections", reflect.TypeOf((*MockDomainClient)(nil).ListWebRedirections), arg0)
}
