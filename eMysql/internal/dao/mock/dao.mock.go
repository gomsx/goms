// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fuwensun/goms/eMysql/internal/dao (interfaces: Dao)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "github.com/fuwensun/goms/eMysql/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDao is a mock of Dao interface
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDao) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDaoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDao)(nil).Close))
}

// Ping mocks base method
func (m *MockDao) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockDaoMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDao)(nil).Ping), arg0)
}

// ReadPingCount mocks base method
func (m *MockDao) ReadPingCount(arg0 context.Context, arg1 model.PingType) (model.PingCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPingCount", arg0, arg1)
	ret0, _ := ret[0].(model.PingCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPingCount indicates an expected call of ReadPingCount
func (mr *MockDaoMockRecorder) ReadPingCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPingCount", reflect.TypeOf((*MockDao)(nil).ReadPingCount), arg0, arg1)
}

// UpdatePingCount mocks base method
func (m *MockDao) UpdatePingCount(arg0 context.Context, arg1 model.PingType, arg2 model.PingCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePingCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePingCount indicates an expected call of UpdatePingCount
func (mr *MockDaoMockRecorder) UpdatePingCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePingCount", reflect.TypeOf((*MockDao)(nil).UpdatePingCount), arg0, arg1, arg2)
}