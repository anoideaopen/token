// Code generated by MockGen. DO NOT EDIT.
// Source: repository/notification.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/anoideaopen/token/model"
	gomock "go.uber.org/mock/gomock"
)

// MockNotification is a mock of Notification interface.
type MockNotification struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationMockRecorder
}

// MockNotificationMockRecorder is the mock recorder for MockNotification.
type MockNotificationMockRecorder struct {
	mock *MockNotification
}

// NewMockNotification creates a new mock instance.
func NewMockNotification(ctrl *gomock.Controller) *MockNotification {
	mock := &MockNotification{ctrl: ctrl}
	mock.recorder = &MockNotificationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotification) EXPECT() *MockNotificationMockRecorder {
	return m.recorder
}

// SaveBalancesUpdate mocks base method.
func (m *MockNotification) SaveBalancesUpdate(ctx context.Context, bu model.Notification[model.BalancesUpdate]) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBalancesUpdate", ctx, bu)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBalancesUpdate indicates an expected call of SaveBalancesUpdate.
func (mr *MockNotificationMockRecorder) SaveBalancesUpdate(ctx, bu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBalancesUpdate", reflect.TypeOf((*MockNotification)(nil).SaveBalancesUpdate), ctx, bu)
}
