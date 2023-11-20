// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/gateway/cart_repository.go

// Package mock_gateway is a generated GoMock package.
package mock_gateway

import (
	reflect "reflect"
	entity "star_store/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockCartRepositoryInterface is a mock of CartRepositoryInterface interface.
type MockCartRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepositoryInterfaceMockRecorder
}

// MockCartRepositoryInterfaceMockRecorder is the mock recorder for MockCartRepositoryInterface.
type MockCartRepositoryInterfaceMockRecorder struct {
	mock *MockCartRepositoryInterface
}

// NewMockCartRepositoryInterface creates a new mock instance.
func NewMockCartRepositoryInterface(ctrl *gomock.Controller) *MockCartRepositoryInterface {
	mock := &MockCartRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCartRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepositoryInterface) EXPECT() *MockCartRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCartRepositoryInterface) Create(cart *entity.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCartRepositoryInterfaceMockRecorder) Create(cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCartRepositoryInterface)(nil).Create), cart)
}

// GetByID mocks base method.
func (m *MockCartRepositoryInterface) GetByID(id string) (*entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCartRepositoryInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCartRepositoryInterface)(nil).GetByID), id)
}

// GetByUser mocks base method.
func (m *MockCartRepositoryInterface) GetByUser(userID string) (*entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUser", userID)
	ret0, _ := ret[0].(*entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUser indicates an expected call of GetByUser.
func (mr *MockCartRepositoryInterfaceMockRecorder) GetByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUser", reflect.TypeOf((*MockCartRepositoryInterface)(nil).GetByUser), userID)
}

// Update mocks base method.
func (m *MockCartRepositoryInterface) Update(cart *entity.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCartRepositoryInterfaceMockRecorder) Update(cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCartRepositoryInterface)(nil).Update), cart)
}