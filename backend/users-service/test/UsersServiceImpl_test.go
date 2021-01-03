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

func TestServiceBan(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Role: entity.ADMIN},
		{Uid: 2, Role: entity.USER},
		{Uid: 2, Role: entity.DISABLE},
	}
	token0, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	token1, _ := util.SignToken(users[1].Uid, users[1].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(users[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[2]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[2].Uid).Return(users[2], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[1]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		req   service.ReqBan
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResBan
	}{
		{"BanNormal", args{token: token0, req: service.ReqBan{Uid: strconv.FormatInt(users[1].Uid, 10), Ban: true}}, service.ResBan{Code: 0}},
		{"LiftNormal", args{token: token0, req: service.ReqBan{Uid: strconv.FormatInt(users[2].Uid, 10), Ban: false}}, service.ResBan{Code: 0}},
		{"NotAdmin", args{token: token1, req: service.ReqBan{Uid: strconv.FormatInt(users[0].Uid, 10), Ban: true}}, service.ResBan{Code: 1}},
		{"UserNotFound", args{token: token0, req: service.ReqBan{Uid: strconv.FormatInt(users[1].Uid, 10), Ban: true}}, service.ResBan{Code: 1}},
		{"WrongToken", args{req: service.ReqBan{Uid: strconv.FormatInt(users[1].Uid, 10), Ban: true}}, service.ResBan{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Ban(tt.args.token, tt.args.req); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceCheckToken(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResCheckToken
	}{
		{"Normal", args{}, service.ResCheckToken{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.CheckToken(tt.args.token); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceFollow(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Name: "test", Role: entity.USER, FollowerCount: 0, FollowingCount: 0},
		{Uid: 1, Name: "test", Role: entity.USER, FollowerCount: 0, FollowingCount: 1},
		{Uid: 2, FollowerCount: 0, FollowingCount: 0},
		{Uid: 2, FollowerCount: 1, FollowingCount: 0},
	}
	follows := []entity.Follows{
		{Uid: 2, Follower: 1},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().InsertFollow(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[1]).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[2].Uid).Return(users[2], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[3]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowByUidAndFollower(gomock.Any(), follows[0].Uid, follows[0].Follower).Return(follows[0], nil),
		mockUsersDao.EXPECT().RemoveFollowByUidAndFollower(gomock.Any(), follows[0].Uid, follows[0].Follower).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(users[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[0]).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[3].Uid).Return(users[3], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[2]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().InsertFollow(gomock.Any(), gomock.Any()).Return(errors.New("error 1062: Duplicate entry '2-1' for key 'PRIMARY'")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowByUidAndFollower(gomock.Any(), follows[0].Uid, follows[0].Follower).Return(entity.Follows{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().InsertFollow(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().InsertFollow(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[1]).Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[2].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token  string
		uid    int64
		follow bool
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResFollow
	}{
		{"FollowNormal", args{token: token, uid: users[2].Uid, follow: true}, service.ResFollow{Code: 0}},
		{"UnfollowNormal", args{token: token, uid: users[2].Uid, follow: false}, service.ResFollow{Code: 0}},
		{"FollowMyself", args{token: token, uid: users[0].Uid, follow: true}, service.ResFollow{Code: 1}},
		{"FollowFound", args{token: token, uid: users[2].Uid, follow: true}, service.ResFollow{Code: 1}},
		{"FollowNotFound", args{token: token, uid: users[2].Uid, follow: false}, service.ResFollow{Code: 1}},
		{"FollowerNotFound", args{token: token, uid: users[2].Uid, follow: true}, service.ResFollow{Code: 1}},
		{"UidNotFound", args{token: token, uid: users[2].Uid, follow: true}, service.ResFollow{Code: 1}},
		{"WrongToken", args{uid: users[2].Uid, follow: true}, service.ResFollow{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Follow(tt.args.token, tt.args.uid, tt.args.follow); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceFollowed(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	follows := []entity.Follows{
		{Uid: 2, Follower: 1},
	}
	users := []entity.Users{
		{Uid: 1, Role: entity.USER},
		{Uid: 2, Name: "test", Nickname: "test", Profile: "test"},
	}
	userDetails := []entity.UserDetails{
		{Uid: 2, Icon: "test"},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowsByFollower(gomock.Any(), users[0].Uid).Return(follows, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(users[1], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[1].Uid).Return(userDetails[0], nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowsByFollower(gomock.Any(), users[0].Uid).Return(follows, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowsByFollower(gomock.Any(), users[0].Uid).Return(follows, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(users[1], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[1].Uid).Return(entity.UserDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		uid   int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResFollowed
	}{
		{"Normal", args{token: token, uid: users[0].Uid}, service.ResFollowed{Code: 0}},
		{"UserNotFound", args{token: token, uid: users[0].Uid}, service.ResFollowed{Code: 1}},
		{"UserDetailNotFound", args{token: token, uid: users[0].Uid}, service.ResFollowed{Code: 1}},
		{"WrongToken", args{uid: users[0].Uid}, service.ResFollowed{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Followed(tt.args.token, tt.args.uid); res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceFollowers(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	follows := []entity.Follows{
		{Uid: 1, Follower: 2},
	}
	users := []entity.Users{
		{Uid: 1, Role: entity.USER},
		{Uid: 2, Name: "test", Nickname: "test", Profile: "test"},
	}
	userDetails := []entity.UserDetails{
		{Uid: 2, Icon: "test"},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowsByUid(gomock.Any(), users[0].Uid).Return(follows, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(users[1], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[1].Uid).Return(userDetails[0], nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowsByUid(gomock.Any(), users[0].Uid).Return(follows, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindFollowsByUid(gomock.Any(), users[0].Uid).Return(follows, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[1].Uid).Return(users[1], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[1].Uid).Return(entity.UserDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		uid   int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResFollowers
	}{
		{"Normal", args{token: token, uid: users[0].Uid}, service.ResFollowers{Code: 0}},
		{"UserNotFound", args{token: token, uid: users[0].Uid}, service.ResFollowers{Code: 1}},
		{"UserDetailNotFound", args{token: token, uid: users[0].Uid}, service.ResFollowers{Code: 1}},
		{"WrongToken", args{uid: users[0].Uid}, service.ResFollowers{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Followers(tt.args.token, tt.args.uid); res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceInfoList(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Name: "test", Nickname: "test", Role: entity.USER},
	}
	userDetails := []entity.UserDetails{
		{Uid: 1, Icon: "test"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(userDetails[0], nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(entity.UserDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		req service.ReqInfoList
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResInfoList
	}{
		{"Normal", args{req: service.ReqInfoList{Uids: []int64{users[0].Uid}}}, service.ResInfoList{Code: 0}},
		{"UserNotFound", args{req: service.ReqInfoList{Uids: []int64{users[0].Uid}}}, service.ResInfoList{Code: 1}},
		{"UserDetailNotFound", args{req: service.ReqInfoList{Uids: []int64{users[0].Uid}}}, service.ResInfoList{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.InfoList(tt.args.req); res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
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
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(users[0], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[1].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[1].Name).Return(users[1], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[2].Name).Return(users[2], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
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

func TestServiceNotifications(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Role: entity.USER, NotificationTime: 1e9},
	}
	notifications := [][]dao.Notifications{
		{{Type: 0}, {Type: 0}, {Type: 1}, {Type: 1}, {Type: 2}, {Type: 2}, {Type: 3}, {Type: 3}, {Type: 4}, {Type: 4}, {Type: 5}, {Type: 5}},
		{{Type: 0}},
		{{Type: 1}},
		{{Type: 2}},
		{{Type: 3}},
		{{Type: 4}},
		{{Type: 5}},
		{},
		{{Type: 6}},
	}
	answers := []entity.Answers{
		{Time: 1e9 + 1},
		{Time: 1e9 - 1},
	}
	likeAnswers := []entity.LikeAnswers{
		{Time: 1e9 + 1},
		{Time: 1e9 - 1},
	}
	approveAnswers := []entity.ApproveAnswers{
		{Time: 1e9 + 1},
		{Time: 1e9 - 1},
	}
	comments := []entity.Comments{
		{Time: 1e9 + 1},
		{Time: 1e9 - 1},
	}
	criticisms := []entity.Criticisms{
		{Time: 1e9 + 1},
		{Time: 1e9 - 1},
	}
	follows := []entity.Follows{
		{Time: 1e9 + 1},
		{Time: 1e9 - 1},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[0], nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), notifications[0][0].Id0).Return(answers[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(entity.QuestionDetails{}, nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), notifications[0][1].Id0).Return(answers[0], nil),
		mockUsersDao.EXPECT().FindLikeAnswerByUidAndAid(gomock.Any(), notifications[0][2].Id0, notifications[0][2].Id1).Return(likeAnswers[0], nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), likeAnswers[0].Aid).Return(answers[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(entity.QuestionDetails{}, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), likeAnswers[0].Aid).Return(entity.AnswerDetails{}, nil),
		mockUsersDao.EXPECT().FindLikeAnswerByUidAndAid(gomock.Any(), notifications[0][3].Id0, notifications[0][3].Id1).Return(likeAnswers[0], nil),
		mockUsersDao.EXPECT().FindApproveAnswerByUidAndAid(gomock.Any(), notifications[0][4].Id0, notifications[0][4].Id1).Return(approveAnswers[0], nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), approveAnswers[0].Aid).Return(answers[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(entity.QuestionDetails{}, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), approveAnswers[0].Aid).Return(entity.AnswerDetails{}, nil),
		mockUsersDao.EXPECT().FindApproveAnswerByUidAndAid(gomock.Any(), notifications[0][5].Id0, notifications[0][5].Id1).Return(approveAnswers[0], nil),
		mockUsersDao.EXPECT().FindCommentByCmid(gomock.Any(), notifications[0][6].Id0).Return(comments[0], nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), comments[0].Aid).Return(answers[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(entity.QuestionDetails{}, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), comments[0].Aid).Return(entity.AnswerDetails{}, nil),
		mockUsersDao.EXPECT().FindCommentByCmid(gomock.Any(), notifications[0][7].Id0).Return(comments[0], nil),
		mockUsersDao.EXPECT().FindCriticismByCtid(gomock.Any(), notifications[0][8].Id0).Return(criticisms[0], nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), criticisms[0].Aid).Return(answers[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(entity.QuestionDetails{}, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), criticisms[0].Aid).Return(entity.AnswerDetails{}, nil),
		mockUsersDao.EXPECT().FindCriticismByCtid(gomock.Any(), notifications[0][9].Id0).Return(criticisms[0], nil),
		mockUsersDao.EXPECT().FindFollowByUidAndFollower(gomock.Any(), users[0].Uid, notifications[0][10].Id0).Return(follows[0], nil),
		mockUsersDao.EXPECT().FindFollowByUidAndFollower(gomock.Any(), users[0].Uid, notifications[0][11].Id0).Return(follows[0], nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[1], nil),
		mockUsersDao.EXPECT().FindAnswerByAid(gomock.Any(), notifications[1][0].Id0).Return(answers[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[2], nil),
		mockUsersDao.EXPECT().FindLikeAnswerByUidAndAid(gomock.Any(), notifications[2][0].Id0, notifications[2][0].Id1).Return(likeAnswers[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[3], nil),
		mockUsersDao.EXPECT().FindApproveAnswerByUidAndAid(gomock.Any(), notifications[3][0].Id0, notifications[3][0].Id1).Return(approveAnswers[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[4], nil),
		mockUsersDao.EXPECT().FindCommentByCmid(gomock.Any(), notifications[4][0].Id0).Return(comments[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[5], nil),
		mockUsersDao.EXPECT().FindCriticismByCtid(gomock.Any(), notifications[5][0].Id0).Return(criticisms[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[6], nil),
		mockUsersDao.EXPECT().FindFollowByUidAndFollower(gomock.Any(), users[0].Uid, notifications[6][0].Id0).Return(follows[1], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[7], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindNotificationsByUidPageable(gomock.Any(), users[0].Uid, dao.Pageable{Size: 10}).Return(notifications[8], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		page  int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResNotifications
	}{
		{"Normal", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalFinishIn0", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalFinishIn1", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalFinishIn2", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalFinishIn3", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalFinishIn4", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalFinishIn5", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"NormalAndLenLessThan10", args{token: token, page: 0}, service.ResNotifications{Code: 0}},
		{"UserNotFound", args{token: token, page: 0}, service.ResNotifications{Code: 1}},
		{"WrongType", args{token: token, page: 0}, service.ResNotifications{Code: 1}},
		{"WrongToken", args{page: 0}, service.ResNotifications{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.Notifications(tt.args.token, tt.args.page); res.Code != tt.wantRes.Code {
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
	userDetails := []entity.UserDetails{
		{Uid: 1, Icon: ""},
	}
	favorites := []entity.Favorites{
		{Uid: users[0].Uid, Title: "Default"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByOidAndAccountType(gomock.Any(), users[0].Oid, users[0].AccountType).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(users[0].Uid, nil),
		mockUsersDao.EXPECT().InsertUserDetail(gomock.Any(), userDetails[0]).Return(nil),
		mockUsersDao.EXPECT().InsertFavorite(gomock.Any(), favorites[0]).Return(favorites[0].Fid, nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByOidAndAccountType(gomock.Any(), users[1].Oid, users[1].AccountType).Return(users[1], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
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
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[0]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
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

func TestServicePublicInfoGet(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Name: "test", Role: entity.USER},
	}
	userDetails := []entity.UserDetails{
		{Uid: 1, Icon: "test"},
	}
	labels := []entity.Labels{
		{Lid: 1, Title: "test"},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(userDetails[0], nil),
		mockUsersDao.EXPECT().FindLabelsByUid(gomock.Any(), users[0].Uid).Return(labels, nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(entity.UserDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		uid   int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResPublicInfoGet
	}{
		{"Normal", args{token, users[0].Uid}, service.ResPublicInfoGet{Code: 0}},
		{"UserNotFound", args{token, users[0].Uid}, service.ResPublicInfoGet{Code: 1}},
		{"UserDetailNotFound", args{token, users[0].Uid}, service.ResPublicInfoGet{Code: 1}},
		{"WrongToken", args{uid: users[0].Uid}, service.ResPublicInfoGet{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.PublicInfoGet(tt.args.token, tt.args.uid); res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServicePublicInfoPut(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Name: "test", Role: entity.USER},
		{Uid: 2},
	}
	userDetails := []entity.UserDetails{
		{Uid: 1, Icon: "test"},
	}
	labels := []entity.Labels{
		{Lid: 1, Title: "test"},
	}
	userLabels := []entity.UserLabels{
		{Uid: 1, Lid: 1},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[0]).Return(nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(userDetails[0], nil),
		mockUsersDao.EXPECT().UpdateUserDetailByUid(gomock.Any(), userDetails[0]).Return(nil),
		mockUsersDao.EXPECT().RemoveUserLabelsByUid(gomock.Any(), users[0].Uid).Return(nil),
		mockUsersDao.EXPECT().FindLabelByTitle(gomock.Any(), labels[0].Title).Return(labels[0], nil),
		mockUsersDao.EXPECT().InsertUserLabel(gomock.Any(), userLabels[0]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[0]).Return(nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(userDetails[0], nil),
		mockUsersDao.EXPECT().UpdateUserDetailByUid(gomock.Any(), userDetails[0]).Return(nil),
		mockUsersDao.EXPECT().RemoveUserLabelsByUid(gomock.Any(), users[0].Uid).Return(nil),
		mockUsersDao.EXPECT().FindLabelByTitle(gomock.Any(), labels[0].Title).Return(entity.Labels{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().InsertLabel(gomock.Any(), gomock.Any()).Return(labels[0].Lid, nil),
		mockUsersDao.EXPECT().InsertUserLabel(gomock.Any(), userLabels[0]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(users[1], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[0]).Return(nil),
		mockUsersDao.EXPECT().FindUserDetailByUid(gomock.Any(), users[0].Uid).Return(entity.UserDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		req   service.ReqPublicInfoPut
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResPublicInfoPut
	}{
		{"NormalAndLabelFound", args{token, service.ReqPublicInfoPut{Name: users[0].Name, Nickname: users[0].Nickname, Profile: users[0].Profile, Icon: userDetails[0].Icon, Gender: users[0].Gender, Email: users[0].Email, Labels: []string{labels[0].Title}}}, service.ResPublicInfoPut{Code: 0}},
		{"NormalAndLabelNotFound", args{token, service.ReqPublicInfoPut{Name: users[0].Name, Nickname: users[0].Nickname, Profile: users[0].Profile, Icon: userDetails[0].Icon, Gender: users[0].Gender, Email: users[0].Email, Labels: []string{labels[0].Title}}}, service.ResPublicInfoPut{Code: 0}},
		{"NameFound", args{token, service.ReqPublicInfoPut{Name: users[0].Name, Nickname: users[0].Nickname, Profile: users[0].Profile, Icon: userDetails[0].Icon, Gender: users[0].Gender, Email: users[0].Email}}, service.ResPublicInfoPut{Code: 1, Result: service.ResultPublicInfoPut{Type: 0}}},
		{"UserNotFound", args{token, service.ReqPublicInfoPut{Name: users[0].Name, Nickname: users[0].Nickname, Profile: users[0].Profile, Icon: userDetails[0].Icon, Gender: users[0].Gender, Email: users[0].Email}}, service.ResPublicInfoPut{Code: 1, Result: service.ResultPublicInfoPut{Type: 1}}},
		{"UserDetailNotFound", args{token, service.ReqPublicInfoPut{Name: users[0].Name, Nickname: users[0].Nickname, Profile: users[0].Profile, Icon: userDetails[0].Icon, Gender: users[0].Gender, Email: users[0].Email}}, service.ResPublicInfoPut{Code: 1, Result: service.ResultPublicInfoPut{Type: 1}}},
		{"WrongToken", args{req: service.ReqPublicInfoPut{Name: users[0].Name, Nickname: users[0].Nickname, Profile: users[0].Profile, Icon: userDetails[0].Icon, Gender: users[0].Gender, Email: users[0].Email}}, service.ResPublicInfoPut{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.PublicInfoPut(tt.args.token, tt.args.req); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceRefreshToken(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Role: entity.DISABLE},
	}
	refreshToken, _ := util.SignToken(users[0].Uid, users[0].Role, true)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		req service.ReqRefreshToken
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResRefreshToken
	}{
		{"Disable", args{req: service.ReqRefreshToken{Refresh: refreshToken}}, service.ResRefreshToken{Code: 1, Result: service.ResultRefreshToken{Type: 0}}},
		{"UserNotFound", args{req: service.ReqRefreshToken{Refresh: refreshToken}}, service.ResRefreshToken{Code: 1, Result: service.ResultRefreshToken{Type: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.RefreshToken(tt.args.req); res != tt.wantRes {
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
		{Name: "test", Email: "test@sjtu.edu.cn", Role: entity.USER, ActiveCode: 0},
		{Name: "test", Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, ActiveCode: 1e5},
	}
	userDetails := []entity.UserDetails{
		{Uid: 1},
	}
	favorites := []entity.Favorites{
		{Uid: users[0].Uid, Title: "Default"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().InsertUserDetail(gomock.Any(), userDetails[0]).Return(nil),
		mockUsersDao.EXPECT().InsertFavorite(gomock.Any(), favorites[0]).Return(favorites[0].Fid, nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[0].Name).Return(users[0], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[1].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[1].Email).Return(users[1], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByName(gomock.Any(), users[2].Name).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[2].Email).Return(users[2], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
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

func TestServiceUserAnswers(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Role: entity.USER},
		{Uid: 2},
	}
	answers := []entity.Answers{
		{Aid: 1},
	}
	answerDetails := []entity.AnswerDetails{
		{Aid: 1},
	}
	questions := []entity.Questions{
		{Qid: 1},
	}
	questionDetails := []entity.QuestionDetails{
		{Qid: 1},
	}
	labels := []entity.Labels{
		{Lid: 1},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindAnswersByAnswererOrderByTimeDescPageable(gomock.Any(), users[1].Uid, dao.Pageable{Size: 10}).Return(answers, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), answers[0].Aid).Return(answerDetails[0], nil),
		mockUsersDao.EXPECT().FindLikeAnswerByUidAndAid(gomock.Any(), users[0].Uid, answers[0].Aid).Return(entity.LikeAnswers{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindApproveAnswerByUidAndAid(gomock.Any(), users[0].Uid, answers[0].Aid).Return(entity.ApproveAnswers{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindQuestionByQid(gomock.Any(), answers[0].Qid).Return(questions[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(questionDetails[0], nil),
		mockUsersDao.EXPECT().FindLabelsByUid(gomock.Any(), users[0].Uid).Return(labels, nil),
		mockUsersDao.EXPECT().FindLabelsByQid(gomock.Any(), answers[0].Qid).Return(labels, nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindAnswersByAnswererOrderByTimeDescPageable(gomock.Any(), users[1].Uid, dao.Pageable{Size: 10}).Return(answers, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), answers[0].Aid).Return(entity.AnswerDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindAnswersByAnswererOrderByTimeDescPageable(gomock.Any(), users[1].Uid, dao.Pageable{Size: 10}).Return(answers, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), answers[0].Aid).Return(answerDetails[0], nil),
		mockUsersDao.EXPECT().FindLikeAnswerByUidAndAid(gomock.Any(), users[0].Uid, answers[0].Aid).Return(entity.LikeAnswers{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindApproveAnswerByUidAndAid(gomock.Any(), users[0].Uid, answers[0].Aid).Return(entity.ApproveAnswers{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindQuestionByQid(gomock.Any(), answers[0].Qid).Return(entity.Questions{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByUid(gomock.Any(), users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().FindAnswersByAnswererOrderByTimeDescPageable(gomock.Any(), users[1].Uid, dao.Pageable{Size: 10}).Return(answers, nil),
		mockUsersDao.EXPECT().FindAnswerDetailByAid(gomock.Any(), answers[0].Aid).Return(answerDetails[0], nil),
		mockUsersDao.EXPECT().FindLikeAnswerByUidAndAid(gomock.Any(), users[0].Uid, answers[0].Aid).Return(entity.LikeAnswers{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindApproveAnswerByUidAndAid(gomock.Any(), users[0].Uid, answers[0].Aid).Return(entity.ApproveAnswers{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().FindQuestionByQid(gomock.Any(), answers[0].Qid).Return(questions[0], nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), answers[0].Qid).Return(entity.QuestionDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		uid   int64
		page  int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResUserAnswers
	}{
		{"Normal", args{token: token, uid: users[1].Uid, page: 0}, service.ResUserAnswers{Code: 0}},
		{"UserNotFound", args{token: token, uid: users[1].Uid, page: 0}, service.ResUserAnswers{Code: 1}},
		{"AnswerDetailNotFound", args{token: token, uid: users[1].Uid, page: 0}, service.ResUserAnswers{Code: 1}},
		{"QuestionNotFound", args{token: token, uid: users[1].Uid, page: 0}, service.ResUserAnswers{Code: 1}},
		{"QuestionDetailNotFound", args{token: token, uid: users[1].Uid, page: 0}, service.ResUserAnswers{Code: 1}},
		{"WrongToken", args{uid: users[1].Uid, page: 0}, service.ResUserAnswers{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.UserAnswers(tt.args.token, tt.args.uid, tt.args.page); res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestServiceUserQuestions(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Role: entity.USER},
	}
	questions := []entity.Questions{
		{Qid: 1},
	}
	questionDetails := []entity.QuestionDetails{
		{Qid: 1},
	}
	labels := []entity.Labels{
		{Lid: 1},
	}
	token, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindQuestionsByRaiserOrderByTimeDescPageable(gomock.Any(), users[0].Uid, dao.Pageable{Number: 0, Size: 10}).Return(questions, nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), questions[0].Qid).Return(questionDetails[0], nil),
		mockUsersDao.EXPECT().FindLabelsByQid(gomock.Any(), questions[0].Qid).Return(labels, nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindQuestionsByRaiserOrderByTimeDescPageable(gomock.Any(), users[0].Uid, dao.Pageable{Number: 0, Size: 10}).Return(questions, nil),
		mockUsersDao.EXPECT().FindQuestionDetailByQid(gomock.Any(), questions[0].Qid).Return(entity.QuestionDetails{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(true).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		uid   int64
		page  int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResUserQuestions
	}{
		{"Normal", args{token: token, uid: users[0].Uid, page: 0}, service.ResUserQuestions{Code: 0}},
		{"QuestionDetailNotFound", args{token: token, uid: users[0].Uid, page: 0}, service.ResUserQuestions{Code: 1}},
		{"WrongToken", args{uid: users[0].Uid, page: 0}, service.ResUserQuestions{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.UserQuestions(tt.args.token, tt.args.uid, tt.args.page); res.Code != tt.wantRes.Code {
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
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(users[0].Uid, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
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
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[1]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[2].Email).Return(users[2], nil),
		mockUsersDao.EXPECT().UpdateUserByUid(gomock.Any(), users[3]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().FindUserByEmail(gomock.Any(), users[2].Email).Return(users[2], nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
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

func TestServiceWordBan(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: 1, Role: entity.ADMIN},
		{Uid: 2, Role: entity.USER},
	}
	banWords := []entity.BanWords{
		{Word: "test"},
	}
	token0, _ := util.SignToken(users[0].Uid, users[0].Role, false)
	token1, _ := util.SignToken(users[1].Uid, users[1].Role, false)
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().InsertBanWord(gomock.Any(), banWords[0]).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().RemoveBanWordByWord(gomock.Any(), banWords[0].Word).Return(nil),
		mockUsersDao.EXPECT().Commit(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Begin(false).Return(dao.TransactionContext{}, nil),
		mockUsersDao.EXPECT().Rollback(gomock.Any()).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	var u service.UsersServiceImpl
	_ = u.Init(mockUsersDao)
	defer u.Destruct()
	type args struct {
		token string
		req   service.ReqWordBan
	}
	tests := []struct {
		name    string
		args    args
		wantRes service.ResWordBan
	}{
		{"BanNormal", args{token: token0, req: service.ReqWordBan{Word: banWords[0].Word, Ban: true}}, service.ResWordBan{Code: 0}},
		{"LiftNormal", args{token: token0, req: service.ReqWordBan{Word: banWords[0].Word, Ban: false}}, service.ResWordBan{Code: 0}},
		{"NotAdmin", args{token: token1, req: service.ReqWordBan{Word: banWords[0].Word, Ban: true}}, service.ResWordBan{Code: 1}},
		{"WrongToken", args{req: service.ReqWordBan{Word: banWords[0].Word, Ban: true}}, service.ResWordBan{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, _ := u.WordBan(tt.args.token, tt.args.req); res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}
