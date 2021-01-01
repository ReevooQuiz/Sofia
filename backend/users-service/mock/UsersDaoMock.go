// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zhanghanchong/users-service/dao (interfaces: UsersDao)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	dao "github.com/zhanghanchong/users-service/dao"
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

// Begin mocks base method
func (m *MockUsersDao) Begin(arg0 bool) (dao.TransactionContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(dao.TransactionContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin
func (mr *MockUsersDaoMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockUsersDao)(nil).Begin), arg0)
}

// Commit mocks base method
func (m *MockUsersDao) Commit(arg0 *dao.TransactionContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockUsersDaoMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockUsersDao)(nil).Commit), arg0)
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
func (m *MockUsersDao) FindLabelByTitle(arg0 dao.TransactionContext, arg1 string) (entity.Labels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLabelByTitle", arg0, arg1)
	ret0, _ := ret[0].(entity.Labels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLabelByTitle indicates an expected call of FindLabelByTitle
func (mr *MockUsersDaoMockRecorder) FindLabelByTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLabelByTitle", reflect.TypeOf((*MockUsersDao)(nil).FindLabelByTitle), arg0, arg1)
}

// FindLabelsByUid mocks base method
func (m *MockUsersDao) FindLabelsByUid(arg0 dao.TransactionContext, arg1 int64) ([]entity.Labels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLabelsByUid", arg0, arg1)
	ret0, _ := ret[0].([]entity.Labels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLabelsByUid indicates an expected call of FindLabelsByUid
func (mr *MockUsersDaoMockRecorder) FindLabelsByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLabelsByUid", reflect.TypeOf((*MockUsersDao)(nil).FindLabelsByUid), arg0, arg1)
}

// FindUserByEmail mocks base method
func (m *MockUsersDao) FindUserByEmail(arg0 dao.TransactionContext, arg1 string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail
func (mr *MockUsersDaoMockRecorder) FindUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUsersDao)(nil).FindUserByEmail), arg0, arg1)
}

// FindUserByName mocks base method
func (m *MockUsersDao) FindUserByName(arg0 dao.TransactionContext, arg1 string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByName", arg0, arg1)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName
func (mr *MockUsersDaoMockRecorder) FindUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockUsersDao)(nil).FindUserByName), arg0, arg1)
}

// FindUserByOidAndAccountType mocks base method
func (m *MockUsersDao) FindUserByOidAndAccountType(arg0 dao.TransactionContext, arg1 string, arg2 int8) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByOidAndAccountType", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByOidAndAccountType indicates an expected call of FindUserByOidAndAccountType
func (mr *MockUsersDaoMockRecorder) FindUserByOidAndAccountType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByOidAndAccountType", reflect.TypeOf((*MockUsersDao)(nil).FindUserByOidAndAccountType), arg0, arg1, arg2)
}

// FindUserByUid mocks base method
func (m *MockUsersDao) FindUserByUid(arg0 dao.TransactionContext, arg1 int64) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUid", arg0, arg1)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUid indicates an expected call of FindUserByUid
func (mr *MockUsersDaoMockRecorder) FindUserByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUid", reflect.TypeOf((*MockUsersDao)(nil).FindUserByUid), arg0, arg1)
}

// FindUserDetailByUid mocks base method
func (m *MockUsersDao) FindUserDetailByUid(arg0 dao.TransactionContext, arg1 int64) (entity.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserDetailByUid", arg0, arg1)
	ret0, _ := ret[0].(entity.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserDetailByUid indicates an expected call of FindUserDetailByUid
func (mr *MockUsersDaoMockRecorder) FindUserDetailByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserDetailByUid", reflect.TypeOf((*MockUsersDao)(nil).FindUserDetailByUid), arg0, arg1)
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
func (m *MockUsersDao) InsertFavorite(arg0 dao.TransactionContext, arg1 entity.Favorites) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFavorite", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertFavorite indicates an expected call of InsertFavorite
func (mr *MockUsersDaoMockRecorder) InsertFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFavorite", reflect.TypeOf((*MockUsersDao)(nil).InsertFavorite), arg0, arg1)
}

// InsertLabel mocks base method
func (m *MockUsersDao) InsertLabel(arg0 dao.TransactionContext, arg1 entity.Labels) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLabel", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertLabel indicates an expected call of InsertLabel
func (mr *MockUsersDaoMockRecorder) InsertLabel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLabel", reflect.TypeOf((*MockUsersDao)(nil).InsertLabel), arg0, arg1)
}

// InsertUser mocks base method
func (m *MockUsersDao) InsertUser(arg0 dao.TransactionContext, arg1 entity.Users) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockUsersDaoMockRecorder) InsertUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUsersDao)(nil).InsertUser), arg0, arg1)
}

// InsertUserDetail mocks base method
func (m *MockUsersDao) InsertUserDetail(arg0 dao.TransactionContext, arg1 entity.UserDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserDetail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserDetail indicates an expected call of InsertUserDetail
func (mr *MockUsersDaoMockRecorder) InsertUserDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserDetail", reflect.TypeOf((*MockUsersDao)(nil).InsertUserDetail), arg0, arg1)
}

// InsertUserLabel mocks base method
func (m *MockUsersDao) InsertUserLabel(arg0 dao.TransactionContext, arg1 entity.UserLabels) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserLabel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserLabel indicates an expected call of InsertUserLabel
func (mr *MockUsersDaoMockRecorder) InsertUserLabel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserLabel", reflect.TypeOf((*MockUsersDao)(nil).InsertUserLabel), arg0, arg1)
}

// RemoveUserLabelsByUid mocks base method
func (m *MockUsersDao) RemoveUserLabelsByUid(arg0 dao.TransactionContext, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserLabelsByUid", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserLabelsByUid indicates an expected call of RemoveUserLabelsByUid
func (mr *MockUsersDaoMockRecorder) RemoveUserLabelsByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserLabelsByUid", reflect.TypeOf((*MockUsersDao)(nil).RemoveUserLabelsByUid), arg0, arg1)
}

// Rollback mocks base method
func (m *MockUsersDao) Rollback(arg0 *dao.TransactionContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockUsersDaoMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockUsersDao)(nil).Rollback), arg0)
}

// UpdateUserByUid mocks base method
func (m *MockUsersDao) UpdateUserByUid(arg0 dao.TransactionContext, arg1 entity.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByUid", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByUid indicates an expected call of UpdateUserByUid
func (mr *MockUsersDaoMockRecorder) UpdateUserByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByUid", reflect.TypeOf((*MockUsersDao)(nil).UpdateUserByUid), arg0, arg1)
}

// UpdateUserDetailByUid mocks base method
func (m *MockUsersDao) UpdateUserDetailByUid(arg0 dao.TransactionContext, arg1 entity.UserDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDetailByUid", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDetailByUid indicates an expected call of UpdateUserDetailByUid
func (mr *MockUsersDaoMockRecorder) UpdateUserDetailByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDetailByUid", reflect.TypeOf((*MockUsersDao)(nil).UpdateUserDetailByUid), arg0, arg1)
}
