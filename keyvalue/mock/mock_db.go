// Code generated by MockGen. DO NOT EDIT.
// Source: db.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	keyvalue "github.com/anoideaopen/token/keyvalue"
	gomock "go.uber.org/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockDB) Del(arg0 context.Context, arg1 keyvalue.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockDBMockRecorder) Del(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockDB)(nil).Del), arg0, arg1)
}

// Get mocks base method.
func (m *MockDB) Get(arg0 context.Context, arg1 keyvalue.Key) (keyvalue.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(keyvalue.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDBMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDB)(nil).Get), arg0, arg1)
}

// Iter mocks base method.
func (m *MockDB) Iter(arg0 context.Context, arg1 keyvalue.Prefix) (keyvalue.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter", arg0, arg1)
	ret0, _ := ret[0].(keyvalue.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iter indicates an expected call of Iter.
func (mr *MockDBMockRecorder) Iter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockDB)(nil).Iter), arg0, arg1)
}

// Set mocks base method.
func (m *MockDB) Set(arg0 context.Context, arg1 keyvalue.Key, arg2 keyvalue.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockDBMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDB)(nil).Set), arg0, arg1, arg2)
}

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// HasNext mocks base method.
func (m *MockIterator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIteratorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIterator)(nil).HasNext))
}

// Next mocks base method.
func (m *MockIterator) Next() (keyvalue.Key, keyvalue.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(keyvalue.Key)
	ret1, _ := ret[1].(keyvalue.Value)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}
