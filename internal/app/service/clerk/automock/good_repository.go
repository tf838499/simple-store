// Code generated by MockGen. DO NOT EDIT.
// Source: simple-store/internal/app/service/clerk (interfaces: GoodRepository)

// Package automock is a generated GoMock package.
package automock

import (
	context "context"
	reflect "reflect"
	PostgresDB "simple-store/internal/adapter/repository/PostgresDB"

	gomock "github.com/golang/mock/gomock"
)

// MockGoodRepository is a mock of GoodRepository interface.
type MockGoodRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGoodRepositoryMockRecorder
}

// MockGoodRepositoryMockRecorder is the mock recorder for MockGoodRepository.
type MockGoodRepositoryMockRecorder struct {
	mock *MockGoodRepository
}

// NewMockGoodRepository creates a new mock instance.
func NewMockGoodRepository(ctrl *gomock.Controller) *MockGoodRepository {
	mock := &MockGoodRepository{ctrl: ctrl}
	mock.recorder = &MockGoodRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoodRepository) EXPECT() *MockGoodRepositoryMockRecorder {
	return m.recorder
}

// DeleteGood mocks base method.
func (m *MockGoodRepository) DeleteGood(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGood", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGood indicates an expected call of DeleteGood.
func (mr *MockGoodRepositoryMockRecorder) DeleteGood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGood", reflect.TypeOf((*MockGoodRepository)(nil).DeleteGood), arg0, arg1)
}

// GetGoodListByPage mocks base method.
func (m *MockGoodRepository) GetGoodListByPage(arg0 context.Context, arg1 PostgresDB.GetGoodListByPageParams) ([]PostgresDB.Good, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoodListByPage", arg0, arg1)
	ret0, _ := ret[0].([]PostgresDB.Good)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoodListByPage indicates an expected call of GetGoodListByPage.
func (mr *MockGoodRepositoryMockRecorder) GetGoodListByPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoodListByPage", reflect.TypeOf((*MockGoodRepository)(nil).GetGoodListByPage), arg0, arg1)
}

// InsertGoods mocks base method.
func (m *MockGoodRepository) InsertGoods(arg0 context.Context, arg1 PostgresDB.InsertGoodsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertGoods", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertGoods indicates an expected call of InsertGoods.
func (mr *MockGoodRepositoryMockRecorder) InsertGoods(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertGoods", reflect.TypeOf((*MockGoodRepository)(nil).InsertGoods), arg0, arg1)
}

// InsertGoodsWithTx mocks base method.
func (m *MockGoodRepository) InsertGoodsWithTx(arg0 context.Context, arg1 []PostgresDB.InsertGoodsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertGoodsWithTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertGoodsWithTx indicates an expected call of InsertGoodsWithTx.
func (mr *MockGoodRepositoryMockRecorder) InsertGoodsWithTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertGoodsWithTx", reflect.TypeOf((*MockGoodRepository)(nil).InsertGoodsWithTx), arg0, arg1)
}

// UpdateGood mocks base method.
func (m *MockGoodRepository) UpdateGood(arg0 context.Context, arg1 PostgresDB.UpdateGoodParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGood", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGood indicates an expected call of UpdateGood.
func (mr *MockGoodRepositoryMockRecorder) UpdateGood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGood", reflect.TypeOf((*MockGoodRepository)(nil).UpdateGood), arg0, arg1)
}
