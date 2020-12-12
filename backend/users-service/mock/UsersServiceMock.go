// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zhanghanchong/users-service/service (interfaces: UsersService)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	dao "github.com/zhanghanchong/users-service/dao"
	entity "github.com/zhanghanchong/users-service/entity"
	bson "gopkg.in/mgo.v2/bson"
	reflect "reflect"
)

// MockUsersService is a mock of UsersService interface
type MockUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceMockRecorder
}

// MockUsersServiceMockRecorder is the mock recorder for MockUsersService
type MockUsersServiceMockRecorder struct {
	mock *MockUsersService
}

// NewMockUsersService creates a new mock instance
func NewMockUsersService(ctrl *gomock.Controller) *MockUsersService {
	mock := &MockUsersService{ctrl: ctrl}
	mock.recorder = &MockUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersService) EXPECT() *MockUsersServiceMockRecorder {
	return m.recorder
}

// Destruct mocks base method
func (m *MockUsersService) Destruct() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destruct")
}

// Destruct indicates an expected call of Destruct
func (mr *MockUsersServiceMockRecorder) Destruct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destruct", reflect.TypeOf((*MockUsersService)(nil).Destruct))
}

// FindUserByEmail mocks base method
func (m *MockUsersService) FindUserByEmail(arg0 string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail
func (mr *MockUsersServiceMockRecorder) FindUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUsersService)(nil).FindUserByEmail), arg0)
}

// FindUserByName mocks base method
func (m *MockUsersService) FindUserByName(arg0 string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByName", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName
func (mr *MockUsersServiceMockRecorder) FindUserByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockUsersService)(nil).FindUserByName), arg0)
}

// FindUserByOidAndAccountType mocks base method
func (m *MockUsersService) FindUserByOidAndAccountType(arg0 string, arg1 int8) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByOidAndAccountType", arg0, arg1)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByOidAndAccountType indicates an expected call of FindUserByOidAndAccountType
func (mr *MockUsersServiceMockRecorder) FindUserByOidAndAccountType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByOidAndAccountType", reflect.TypeOf((*MockUsersService)(nil).FindUserByOidAndAccountType), arg0, arg1)
}

// FindUserByUid mocks base method
func (m *MockUsersService) FindUserByUid(arg0 bson.ObjectId) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUid", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUid indicates an expected call of FindUserByUid
func (mr *MockUsersServiceMockRecorder) FindUserByUid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUid", reflect.TypeOf((*MockUsersService)(nil).FindUserByUid), arg0)
}

// Init mocks base method
func (m *MockUsersService) Init(arg0 ...dao.UsersDao) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Init", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockUsersServiceMockRecorder) Init(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockUsersService)(nil).Init), arg0...)
}

// InsertUser mocks base method
func (m *MockUsersService) InsertUser(arg0 entity.Users) (bson.ObjectId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0)
	ret0, _ := ret[0].(bson.ObjectId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockUsersServiceMockRecorder) InsertUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUsersService)(nil).InsertUser), arg0)
}

// UpdateUser mocks base method
func (m *MockUsersService) UpdateUser(arg0 entity.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUsersServiceMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsersService)(nil).UpdateUser), arg0)
}
