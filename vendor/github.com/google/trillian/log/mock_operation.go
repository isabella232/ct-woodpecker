// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian/log (interfaces: Operation)

package log

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOperation is a mock of Operation interface.
type MockOperation struct {
	ctrl     *gomock.Controller
	recorder *MockOperationMockRecorder
}

// MockOperationMockRecorder is the mock recorder for MockOperation.
type MockOperationMockRecorder struct {
	mock *MockOperation
}

// NewMockOperation creates a new mock instance.
func NewMockOperation(ctrl *gomock.Controller) *MockOperation {
	mock := &MockOperation{ctrl: ctrl}
	mock.recorder = &MockOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperation) EXPECT() *MockOperationMockRecorder {
	return m.recorder
}

// ExecutePass mocks base method.
func (m *MockOperation) ExecutePass(arg0 context.Context, arg1 int64, arg2 *OperationInfo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecutePass", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecutePass indicates an expected call of ExecutePass.
func (mr *MockOperationMockRecorder) ExecutePass(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecutePass", reflect.TypeOf((*MockOperation)(nil).ExecutePass), arg0, arg1, arg2)
}
