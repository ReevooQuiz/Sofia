// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SKFE396/qa-service/service (interfaces: QaService)

// Package mock is a generated GoMock package.
package mock

import (
	dao "github.com/SKFE396/qa-service/dao"
	rpc "github.com/SKFE396/qa-service/rpc"
	service "github.com/SKFE396/qa-service/service"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockQaService is a mock of QaService interface
type MockQaService struct {
	ctrl     *gomock.Controller
	recorder *MockQaServiceMockRecorder
}

// MockQaServiceMockRecorder is the mock recorder for MockQaService
type MockQaServiceMockRecorder struct {
	mock *MockQaService
}

// NewMockQaService creates a new mock instance
func NewMockQaService(ctrl *gomock.Controller) *MockQaService {
	mock := &MockQaService{ctrl: ctrl}
	mock.recorder = &MockQaServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQaService) EXPECT() *MockQaServiceMockRecorder {
	return m.recorder
}

// AddQuestion mocks base method
func (m *MockQaService) AddQuestion(arg0 string, arg1 service.ReqQuestionsPost) (int8, interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddQuestion", arg0, arg1)
	ret0, _ := ret[0].(int8)
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

// AddQuestion indicates an expected call of AddQuestion
func (mr *MockQaServiceMockRecorder) AddQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddQuestion", reflect.TypeOf((*MockQaService)(nil).AddQuestion), arg0, arg1)
}

// Destruct mocks base method
func (m *MockQaService) Destruct() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destruct")
}

// Destruct indicates an expected call of Destruct
func (mr *MockQaServiceMockRecorder) Destruct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destruct", reflect.TypeOf((*MockQaService)(nil).Destruct))
}

// Init mocks base method
func (m *MockQaService) Init(arg0 dao.QaDao, arg1 rpc.UsersRPC) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockQaServiceMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockQaService)(nil).Init), arg0, arg1)
}

// MainPage mocks base method
func (m *MockQaService) MainPage(arg0 string, arg1 int64) (int8, interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainPage", arg0, arg1)
	ret0, _ := ret[0].(int8)
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

// MainPage indicates an expected call of MainPage
func (mr *MockQaServiceMockRecorder) MainPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainPage", reflect.TypeOf((*MockQaService)(nil).MainPage), arg0, arg1)
}

// ModifyQuestion mocks base method
func (m *MockQaService) ModifyQuestion(arg0 string, arg1 service.ReqQuestionsPut) (int8, interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyQuestion", arg0, arg1)
	ret0, _ := ret[0].(int8)
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

// ModifyQuestion indicates an expected call of ModifyQuestion
func (mr *MockQaServiceMockRecorder) ModifyQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyQuestion", reflect.TypeOf((*MockQaService)(nil).ModifyQuestion), arg0, arg1)
}
