// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zhanghanchong/users-service/dao (interfaces: UsersDao)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/zhanghanchong/users-service/entity"
	reflect "reflect"
)

// MockUsersDao is a mock of UsersDao interface
type MockUsersDao struct {
	ctrl     *gomock.Controller
	recorder *MockUsersDaoMockRecorder
}

// MockUsersDaoMockRecorder is the mock recorder for MockUsersDao
type MockUsersDaoMockRecorder struct {
	mock *MockUsersDao
}

// NewMockUsersDao creates a new mock instance
func NewMockUsersDao(ctrl *gomock.Controller) *MockUsersDao {
	mock := &MockUsersDao{ctrl: ctrl}
	mock.recorder = &MockUsersDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersDao) EXPECT() *MockUsersDaoMockRecorder {
	return m.recorder
}

// Destruct mocks base method
func (m *MockUsersDao) Destruct() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destruct")
}

// Destruct indicates an expected call of Destruct
func (mr *MockUsersDaoMockRecorder) Destruct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destruct", reflect.TypeOf((*MockUsersDao)(nil).Destruct))
}

// FindLabelByTitle mocks base method
func (m *MockUsersDao) FindLabelByTitle(arg0 string) (entity.Labels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLabelByTitle", arg0)
	ret0, _ := ret[0].(entity.Labels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLabelByTitle indicates an expected call of FindLabelByTitle
func (mr *MockUsersDaoMockRecorder) FindLabelByTitle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLabelByTitle", reflect.TypeOf((*MockUsersDao)(nil).FindLabelByTitle), arg0)
}

// FindUserByEmail mocks base method
func (m *MockUsersDao) FindUserByEmail(arg0 string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail
func (mr *MockUsersDaoMockRecorder) FindUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUsersDao)(nil).FindUserByEmail), arg0)
}

// FindUserByName mocks base method
func (m *MockUsersDao) FindUserByName(arg0 string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByName", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName
func (mr *MockUsersDaoMockRecorder) FindUserByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockUsersDao)(nil).FindUserByName), arg0)
}

// FindUserByOidAndAccountType mocks base method
func (m *MockUsersDao) FindUserByOidAndAccountType(arg0 string, arg1 int8) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByOidAndAccountType", arg0, arg1)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByOidAndAccountType indicates an expected call of FindUserByOidAndAccountType
func (mr *MockUsersDaoMockRecorder) FindUserByOidAndAccountType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByOidAndAccountType", reflect.TypeOf((*MockUsersDao)(nil).FindUserByOidAndAccountType), arg0, arg1)
}

// FindUserByUid mocks base method
func (m *MockUsersDao) FindUserByUid(arg0 int64) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUid", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUid indicates an expected call of FindUserByUid
func (mr *MockUsersDaoMockRecorder) FindUserByUid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUid", reflect.TypeOf((*MockUsersDao)(nil).FindUserByUid), arg0)
}

// FindUserDetailByUid mocks base method
func (m *MockUsersDao) FindUserDetailByUid(arg0 int64) (entity.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserDetailByUid", arg0)
	ret0, _ := ret[0].(entity.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserDetailByUid indicates an expected call of FindUserDetailByUid
func (mr *MockUsersDaoMockRecorder) FindUserDetailByUid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserDetailByUid", reflect.TypeOf((*MockUsersDao)(nil).FindUserDetailByUid), arg0)
}

// Init mocks base method
func (m *MockUsersDao) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockUsersDaoMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockUsersDao)(nil).Init))
}

// InsertFavorite mocks base method
func (m *MockUsersDao) InsertFavorite(arg0 entity.Favorites) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFavorite", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertFavorite indicates an expected call of InsertFavorite
func (mr *MockUsersDaoMockRecorder) InsertFavorite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFavorite", reflect.TypeOf((*MockUsersDao)(nil).InsertFavorite), arg0)
}

// InsertLabel mocks base method
func (m *MockUsersDao) InsertLabel(arg0 entity.Labels) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLabel", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertLabel indicates an expected call of InsertLabel
func (mr *MockUsersDaoMockRecorder) InsertLabel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLabel", reflect.TypeOf((*MockUsersDao)(nil).InsertLabel), arg0)
}

// InsertUser mocks base method
func (m *MockUsersDao) InsertUser(arg0 entity.Users) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockUsersDaoMockRecorder) InsertUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUsersDao)(nil).InsertUser), arg0)
}

// InsertUserDetail mocks base method
func (m *MockUsersDao) InsertUserDetail(arg0 entity.UserDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserDetail", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserDetail indicates an expected call of InsertUserDetail
func (mr *MockUsersDaoMockRecorder) InsertUserDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserDetail", reflect.TypeOf((*MockUsersDao)(nil).InsertUserDetail), arg0)
}

// InsertUserLabel mocks base method
func (m *MockUsersDao) InsertUserLabel(arg0 entity.UserLabels) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserLabel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserLabel indicates an expected call of InsertUserLabel
func (mr *MockUsersDaoMockRecorder) InsertUserLabel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserLabel", reflect.TypeOf((*MockUsersDao)(nil).InsertUserLabel), arg0)
}

// RemoveUserLabelsByUid mocks base method
func (m *MockUsersDao) RemoveUserLabelsByUid(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserLabelsByUid", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserLabelsByUid indicates an expected call of RemoveUserLabelsByUid
func (mr *MockUsersDaoMockRecorder) RemoveUserLabelsByUid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserLabelsByUid", reflect.TypeOf((*MockUsersDao)(nil).RemoveUserLabelsByUid), arg0)
}

// UpdateUser mocks base method
func (m *MockUsersDao) UpdateUser(arg0 entity.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUsersDaoMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsersDao)(nil).UpdateUser), arg0)
}

// UpdateUserDetail mocks base method
func (m *MockUsersDao) UpdateUserDetail(arg0 entity.UserDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDetail", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDetail indicates an expected call of UpdateUserDetail
func (mr *MockUsersDaoMockRecorder) UpdateUserDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDetail", reflect.TypeOf((*MockUsersDao)(nil).UpdateUserDetail), arg0)
}
