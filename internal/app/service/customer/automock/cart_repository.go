// Code generated by MockGen. DO NOT EDIT.
// Source: simple-store/internal/app/service/customer (interfaces: CartRepository)

// Package automock is a generated GoMock package.
package automock

import (
	context "context"
	reflect "reflect"
	redisclient "simple-store/internal/adapter/redisclient"

	gomock "github.com/golang/mock/gomock"
)

// MockCartRepository is a mock of CartRepository interface.
type MockCartRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepositoryMockRecorder
}

// MockCartRepositoryMockRecorder is the mock recorder for MockCartRepository.
type MockCartRepositoryMockRecorder struct {
	mock *MockCartRepository
}

// NewMockCartRepository creates a new mock instance.
func NewMockCartRepository(ctrl *gomock.Controller) *MockCartRepository {
	mock := &MockCartRepository{ctrl: ctrl}
	mock.recorder = &MockCartRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepository) EXPECT() *MockCartRepositoryMockRecorder {
	return m.recorder
}

// DeleteGood mocks base method.
func (m *MockCartRepository) DeleteGood(arg0 context.Context, arg1 []redisclient.GoodInCartParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGood", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGood indicates an expected call of DeleteGood.
func (mr *MockCartRepositoryMockRecorder) DeleteGood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGood", reflect.TypeOf((*MockCartRepository)(nil).DeleteGood), arg0, arg1)
}

// GetCartListCache mocks base method.
func (m *MockCartRepository) GetCartListCache(arg0 context.Context, arg1 string) ([]redisclient.GoodInRedisParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartListCache", arg0, arg1)
	ret0, _ := ret[0].([]redisclient.GoodInRedisParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartListCache indicates an expected call of GetCartListCache.
func (mr *MockCartRepositoryMockRecorder) GetCartListCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartListCache", reflect.TypeOf((*MockCartRepository)(nil).GetCartListCache), arg0, arg1)
}

// GetGoodPrice mocks base method.
func (m *MockCartRepository) GetGoodPrice(arg0 context.Context, arg1 []string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoodPrice", arg0, arg1)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoodPrice indicates an expected call of GetGoodPrice.
func (mr *MockCartRepositoryMockRecorder) GetGoodPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoodPrice", reflect.TypeOf((*MockCartRepository)(nil).GetGoodPrice), arg0, arg1)
}

// SetGood mocks base method.
func (m *MockCartRepository) SetGood(arg0 context.Context, arg1 []redisclient.GoodInCartParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGood", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGood indicates an expected call of SetGood.
func (mr *MockCartRepositoryMockRecorder) SetGood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGood", reflect.TypeOf((*MockCartRepository)(nil).SetGood), arg0, arg1)
}

// SetGoodPrice mocks base method.
func (m *MockCartRepository) SetGoodPrice(arg0 context.Context, arg1 redisclient.GoodPriceInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGoodPrice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGoodPrice indicates an expected call of SetGoodPrice.
func (mr *MockCartRepositoryMockRecorder) SetGoodPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGoodPrice", reflect.TypeOf((*MockCartRepository)(nil).SetGoodPrice), arg0, arg1)
}
