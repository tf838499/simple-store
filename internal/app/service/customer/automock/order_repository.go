// Code generated by MockGen. DO NOT EDIT.
// Source: simple-store/internal/app/service/customer (interfaces: OrderRepository)

// Package automock is a generated GoMock package.
package automock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	PostgresDB "simple-store/internal/adapter/repository/PostgresDB"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// GetGetOrderByOwner mocks base method.
func (m *MockOrderRepository) GetGetOrderByOwner(arg0 context.Context, arg1 sql.NullString) ([]PostgresDB.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGetOrderByOwner", arg0, arg1)
	ret0, _ := ret[0].([]PostgresDB.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGetOrderByOwner indicates an expected call of GetGetOrderByOwner.
func (mr *MockOrderRepositoryMockRecorder) GetGetOrderByOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGetOrderByOwner", reflect.TypeOf((*MockOrderRepository)(nil).GetGetOrderByOwner), arg0, arg1)
}

// GetGoodByName mocks base method.
func (m *MockOrderRepository) GetGoodByName(arg0 context.Context, arg1 sql.NullString) (PostgresDB.Good, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoodByName", arg0, arg1)
	ret0, _ := ret[0].(PostgresDB.Good)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoodByName indicates an expected call of GetGoodByName.
func (mr *MockOrderRepositoryMockRecorder) GetGoodByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoodByName", reflect.TypeOf((*MockOrderRepository)(nil).GetGoodByName), arg0, arg1)
}

// InsertOrder mocks base method.
func (m *MockOrderRepository) InsertOrder(arg0 context.Context, arg1 PostgresDB.InsertOrderParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrder indicates an expected call of InsertOrder.
func (mr *MockOrderRepositoryMockRecorder) InsertOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrder", reflect.TypeOf((*MockOrderRepository)(nil).InsertOrder), arg0, arg1)
}