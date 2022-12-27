// Code generated by MockGen. DO NOT EDIT.
// Source: modules/product/handler.go

// Package mock_product is a generated GoMock package.
package mock_product

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "phabricator.sirclo.com/source/Hercules.git/models"
	mp_adapter "phabricator.sirclo.com/source/Hercules.git/modules/mp_adapter"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// UpsertProduct mocks base method.
func (m *MockHandler) UpsertProduct(credential models.Credentials, productRequests []mp_adapter.UpsertProductRequest) (*[]mp_adapter.UpsertProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertProduct", credential, productRequests)
	ret0, _ := ret[0].(*[]mp_adapter.UpsertProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertProduct indicates an expected call of UpsertProduct.
func (mr *MockHandlerMockRecorder) UpsertProduct(credential, productRequests interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertProduct", reflect.TypeOf((*MockHandler)(nil).UpsertProduct), credential, productRequests)
}

// VariantFetcherChannel mocks base method.
func (m *MockHandler) VariantFetcherChannel(credential models.Credentials) ([]mp_adapter.UpsertProductRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariantFetcherChannel", credential)
	ret0, _ := ret[0].([]mp_adapter.UpsertProductRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariantFetcherChannel indicates an expected call of VariantFetcherChannel.
func (mr *MockHandlerMockRecorder) VariantFetcherChannel(credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariantFetcherChannel", reflect.TypeOf((*MockHandler)(nil).VariantFetcherChannel), credential)
}
