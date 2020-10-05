// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/makasim/amqpextra (interfaces: AMQPConnection)

// Package mock_amqpextra is a generated GoMock package.
package mock_amqpextra

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	amqp "github.com/streadway/amqp"
)

// MockAMQPConnection is a mock of AMQPConnection interface.
type MockAMQPConnection struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPConnectionMockRecorder
}

// MockAMQPConnectionMockRecorder is the mock recorder for MockAMQPConnection.
type MockAMQPConnectionMockRecorder struct {
	mock *MockAMQPConnection
}

// NewMockAMQPConnection creates a new mock instance.
func NewMockAMQPConnection(ctrl *gomock.Controller) *MockAMQPConnection {
	mock := &MockAMQPConnection{ctrl: ctrl}
	mock.recorder = &MockAMQPConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPConnection) EXPECT() *MockAMQPConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAMQPConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPConnection)(nil).Close))
}

// NotifyClose mocks base method.
func (m *MockAMQPConnection) NotifyClose(arg0 chan *amqp.Error) chan *amqp.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyClose", arg0)
	ret0, _ := ret[0].(chan *amqp.Error)
	return ret0
}

// NotifyClose indicates an expected call of NotifyClose.
func (mr *MockAMQPConnectionMockRecorder) NotifyClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClose", reflect.TypeOf((*MockAMQPConnection)(nil).NotifyClose), arg0)
}