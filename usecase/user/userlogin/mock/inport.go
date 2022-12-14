// Code generated by MockGen. DO NOT EDIT.
// Source: goplay-backend-engineer-test/usecase/user/userlogin (interfaces: Inport)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	userlogin "goplay-backend-engineer-test/usecase/user/userlogin"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInport is a mock of Inport interface.
type MockInport struct {
	ctrl     *gomock.Controller
	recorder *MockInportMockRecorder
}

// MockInportMockRecorder is the mock recorder for MockInport.
type MockInportMockRecorder struct {
	mock *MockInport
}

// NewMockInport creates a new mock instance.
func NewMockInport(ctrl *gomock.Controller) *MockInport {
	mock := &MockInport{ctrl: ctrl}
	mock.recorder = &MockInportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInport) EXPECT() *MockInportMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockInport) Execute(arg0 context.Context, arg1 userlogin.InportRequest) (userlogin.InportResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(userlogin.InportResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockInportMockRecorder) Execute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockInport)(nil).Execute), arg0, arg1)
}
