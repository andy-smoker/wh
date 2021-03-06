// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	structs "github.com/andy-smoker/wh-server/pkg/structs"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthorization is a mock of Authorization interface
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockAuthorization) CreateUser(user structs.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method
func (m *MockAuthorization) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken
func (mr *MockAuthorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), accessToken)
}

// MockWarehouse is a mock of Warehouse interface
type MockWarehouse struct {
	ctrl     *gomock.Controller
	recorder *MockWarehouseMockRecorder
}

// MockWarehouseMockRecorder is the mock recorder for MockWarehouse
type MockWarehouseMockRecorder struct {
	mock *MockWarehouse
}

// NewMockWarehouse creates a new mock instance
func NewMockWarehouse(ctrl *gomock.Controller) *MockWarehouse {
	mock := &MockWarehouse{ctrl: ctrl}
	mock.recorder = &MockWarehouseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWarehouse) EXPECT() *MockWarehouseMockRecorder {
	return m.recorder
}

// CreateItem mocks base method
func (m *MockWarehouse) CreateItem(item structs.WHitem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem
func (mr *MockWarehouseMockRecorder) CreateItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockWarehouse)(nil).CreateItem), item)
}

// GetItem mocks base method
func (m *MockWarehouse) GetItem(itemID int) (structs.WHitem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", itemID)
	ret0, _ := ret[0].(structs.WHitem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem
func (mr *MockWarehouseMockRecorder) GetItem(itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockWarehouse)(nil).GetItem), itemID)
}

// UpdateItem mocks base method
func (m *MockWarehouse) UpdateItem(item structs.WHitem) (structs.WHitem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", item)
	ret0, _ := ret[0].(structs.WHitem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem
func (mr *MockWarehouseMockRecorder) UpdateItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockWarehouse)(nil).UpdateItem), item)
}

// DeleteItem mocks base method
func (m *MockWarehouse) DeleteItem(id int, itemType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", id, itemType)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem
func (mr *MockWarehouseMockRecorder) DeleteItem(id, itemType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockWarehouse)(nil).DeleteItem), id, itemType)
}

// GetItemsList mocks base method
func (m *MockWarehouse) GetItemsList(itemsType string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemsList", itemsType)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemsList indicates an expected call of GetItemsList
func (mr *MockWarehouseMockRecorder) GetItemsList(itemsType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemsList", reflect.TypeOf((*MockWarehouse)(nil).GetItemsList), itemsType)
}
