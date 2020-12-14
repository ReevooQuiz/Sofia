package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/mock"
	"github.com/zhanghanchong/users-service/service"
	"github.com/zhanghanchong/users-service/util"
	"strconv"
	"testing"
)

func TestServiceInit(t *testing.T) {
	u := service.UsersServiceImpl{}
	type args struct {
		usersDAO []dao.UsersDao
	}
	tests := []struct {
		name string
		args args
	}{
		{"Initialize", args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = u.Init(tt.args.usersDAO...)
		})
	}
}

func TestServiceLogin(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Name: "test", HashPassword: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Role: entity.DISABLE},
		{Name: "test", HashPassword: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Role: entity.USER},
		{Name: "test", HashPassword: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Role: entity.NOT_ACTIVE},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[0].Name).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserByName(users[1].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByName(users[1].Name).Return(users[1], nil),
		mockUsersDao.EXPECT().FindUserByName(users[2].Name).Return(users[2], nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		req service.ReqLogin
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResLogin
	}{
		{"Disable", args{req: service.ReqLogin{Name: users[0].Name, Password: "test"}}, service.ResLogin{Code: 1, Result: service.ResultLogin{Type: 0}}},
		{"NameNotFound", args{req: service.ReqLogin{Name: users[1].Name, Password: "test"}}, service.ResLogin{Code: 1, Result: service.ResultLogin{Type: 1}}},
		{"WrongPassword", args{req: service.ReqLogin{Name: users[1].Name}}, service.ResLogin{Code: 1, Result: service.ResultLogin{Type: 1}}},
		{"NotActive", args{req: service.ReqLogin{Name: users[2].Name, Password: "test"}}, service.ResLogin{Code: 1, Result: service.ResultLogin{Type: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Login(tt.args.req); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceOAuthGithub(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Oid: "0", Role: entity.USER, AccountType: entity.GITHUB},
		{Oid: "0", Role: entity.DISABLE, AccountType: entity.GITHUB},
	}
	favorites := []entity.Favorites{
		{Uid: users[0].Uid, Title: "Default"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByOidAndAccountType(users[0].Oid, users[0].AccountType).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().InsertUser(gomock.Any()).Return(users[0].Uid, nil),
		mockUsersDao.EXPECT().InsertFavorite(favorites[0]).Return(favorites[0].Fid, nil),
		mockUsersDao.EXPECT().FindUserByOidAndAccountType(users[1].Oid, users[1].AccountType).Return(users[1], nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		code  string
		error string
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResOAuthGithub
	}{
		{"NormalAndFirst", args{}, service.ResOAuthGithub{Code: 0, Result: service.ResultOAuthGithub{First: true, Role: users[0].Role, Uid: strconv.FormatInt(users[0].Uid, 10)}}},
		{"Disable", args{}, service.ResOAuthGithub{Code: 1, Result: service.ResultOAuthGithub{Type: 0}}},
		{"AccessDenied", args{error: "access_denied"}, service.ResOAuthGithub{Code: 1, Result: service.ResultOAuthGithub{Type: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.OAuthGithub(tt.args.code, tt.args.error); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServicePasswd(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, HashPassword: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Role: entity.USER},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(users[0]).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		req   service.ReqPasswd
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResPasswd
	}{
		{"Normal", args{token: token, req: service.ReqPasswd{Old: "test", New: "test"}}, service.ResPasswd{Code: 0}},
		{"WrongName", args{token: token, req: service.ReqPasswd{Old: "test", New: "test"}}, service.ResPasswd{Code: 1, Result: service.ResultPasswd{Type: 1}}},
		{"WrongPassword", args{token: token}, service.ResPasswd{Code: 1, Result: service.ResultPasswd{}}},
		{"WrongToken", args{req: service.ReqPasswd{Old: "test", New: "test"}}, service.ResPasswd{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Passwd(tt.args.token, tt.args.req); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceRegister(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Name: "test", Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, ActiveCode: 0},
		{Email: "test@sjtu.edu.cn", Role: entity.USER, ActiveCode: 0},
		{Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, ActiveCode: 1e5},
	}
	userDetails := []entity.UserDetails{
		{Uid: 1},
	}
	favorites := []entity.Favorites{
		{Uid: users[0].Uid, Title: "Default"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().InsertUserDetail(userDetails[0]).Return(nil),
		mockUsersDao.EXPECT().InsertFavorite(favorites[0]).Return(favorites[0].Fid, nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("error 1062: Duplicate entry '"+users[0].Name+"' for key 'name'")),
		mockUsersDao.EXPECT().FindUserByEmail(users[1].Email).Return(users[1], nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[2].Email).Return(users[2], nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		req service.ReqRegister
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResRegister
	}{
		{"Normal", args{service.ReqRegister{Name: users[0].Name, Nickname: users[0].Nickname, Password: "test", Email: users[0].Email, Icon: userDetails[0].Icon, Gender: users[0].Gender}}, service.ResRegister{Code: 0}},
		{"NameFound", args{service.ReqRegister{Name: users[0].Name, Nickname: users[0].Nickname, Password: users[0].HashPassword, Email: users[0].Email, Icon: userDetails[0].Icon, Gender: users[0].Gender}}, service.ResRegister{Code: 1, Result: service.ResultRegister{Type: 0}}},
		{"EmailFound", args{service.ReqRegister{Name: users[1].Name, Nickname: users[1].Nickname, Password: users[1].HashPassword, Email: users[1].Email, Icon: userDetails[0].Icon, Gender: users[1].Gender}}, service.ResRegister{Code: 1, Result: service.ResultRegister{Type: 1}}},
		{"NotActive", args{service.ReqRegister{Name: users[2].Name, Nickname: users[2].Nickname, Password: users[2].HashPassword, Email: users[2].Email, Icon: userDetails[0].Icon, Gender: users[2].Gender}}, service.ResRegister{Code: 1, Result: service.ResultRegister{Type: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Register(tt.args.req); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceVerificationCode(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Name: "test@sjtu.edu.cn", Email: "test@sjtu.edu.cn"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().InsertUser(gomock.Any()).Return(int64(0), errors.New("error 1062: Duplicate entry '"+users[0].Name+"' for key 'name'")),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().InsertUser(gomock.Any()).Return(users[0].Uid, nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		register bool
		email    string
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResVerificationCode
	}{
		{"EmailFound", args{true, users[0].Email}, service.ResVerificationCode{Code: 1, Result: service.ResultVerificationCode{Type: 0}}},
		{"EmailNotFound", args{false, users[0].Email}, service.ResVerificationCode{Code: 1, Result: service.ResultVerificationCode{Type: 0}}},
		{"RegisterMailFail", args{true, users[0].Email}, service.ResVerificationCode{Code: 1, Result: service.ResultVerificationCode{Type: 1}}},
		{"NotRegisterMailFail", args{false, users[0].Email}, service.ResVerificationCode{Code: 1, Result: service.ResultVerificationCode{Type: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.VerificationCode(tt.args.register, tt.args.email); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceVerify(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, ActiveCode: 1e5, PasswdCode: 0},
		{Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, ActiveCode: 0, PasswdCode: 0},
		{Email: "test@sjtu.edu.cn", Role: entity.USER, ActiveCode: 0, PasswdCode: 1e5},
		{Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, ActiveCode: 0, PasswdCode: 0},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(users[1]).Return(nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[2].Email).Return(users[2], nil),
		mockUsersDao.EXPECT().UpdateUser(users[3]).Return(nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserByEmail(users[2].Email).Return(users[2], nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		email string
		code  int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResVerify
	}{
		{"RegisterNormal", args{users[0].Email, users[0].ActiveCode}, service.ResVerify{Code: 0}},
		{"NotRegisterNormal", args{users[2].Email, users[2].PasswdCode}, service.ResVerify{Code: 0}},
		{"MailNotFound", args{users[0].Email, users[0].ActiveCode}, service.ResVerify{Code: 1}},
		{"RegisterFail", args{users[0].Email, users[0].ActiveCode + 1}, service.ResVerify{Code: 1}},
		{"NotRegisterFail", args{users[2].Email, users[2].PasswdCode + 1}, service.ResVerify{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Verify(tt.args.email, tt.args.code); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}
