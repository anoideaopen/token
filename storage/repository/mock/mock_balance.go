// Code generated by MockGen. DO NOT EDIT.
// Source: repository/balance.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	big "math/big"
	reflect "reflect"

	model "github.com/anoideaopen/token/model"
	gomock "go.uber.org/mock/gomock"
)

// MockBalance is a mock of Balance interface.
type MockBalance struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceMockRecorder
}

// MockBalanceMockRecorder is the mock recorder for MockBalance.
type MockBalanceMockRecorder struct {
	mock *MockBalance
}

// NewMockBalance creates a new mock instance.
func NewMockBalance(ctrl *gomock.Controller) *MockBalance {
	mock := &MockBalance{ctrl: ctrl}
	mock.recorder = &MockBalanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBalance) EXPECT() *MockBalanceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockBalance) List(ctx context.Context, addr model.Address, acc model.Account) (map[model.Currency]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, addr, acc)
	ret0, _ := ret[0].(map[model.Currency]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockBalanceMockRecorder) List(ctx, addr, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBalance)(nil).List), ctx, addr, acc)
}

// Load mocks base method.
func (m *MockBalance) Load(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", ctx, addr, acc, curr)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockBalanceMockRecorder) Load(ctx, addr, acc, curr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockBalance)(nil).Load), ctx, addr, acc, curr)
}

// Save mocks base method.
func (m *MockBalance) Save(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency, val *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, addr, acc, curr, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockBalanceMockRecorder) Save(ctx, addr, acc, curr, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockBalance)(nil).Save), ctx, addr, acc, curr, val)
}