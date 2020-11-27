package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/mock"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	tests := []struct {
		name string
	}{
		{"Initialize"},
	}
	u := UsersController{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpServerExitDone := &sync.WaitGroup{}
			httpServerExitDone.Add(1)
			server := u.Init(httpServerExitDone, mockUsersService)
			time.Sleep(500 * time.Microsecond)
			if err := server.Shutdown(context.Background()); err != nil {
				t.Error(err)
			}
			httpServerExitDone.Wait()
		})
	}
}

func TestActivate(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	users := []entity.Users{
		{Uid: bson.ObjectIdHex("000000000000000000000000"), Role: entity.USER},
	}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByUid(users[0].Uid).Return(users[0], nil),
		mockUsersService.EXPECT().UpdateUser(users[0]).Return(nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByUid(users[0].Uid).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersService.EXPECT().Destruct(),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/activate", u.Activate)
	type args struct {
		token string
	}
	type res struct {
		Code int8 `json:"code"`
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"Normal", args{"000000000000000000000000"}, http.StatusOK, res{0}},
		{"UidNotFound", args{"000000000000000000000000"}, http.StatusOK, res{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/activate?token="+tt.args.token, nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			body := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(body)
			var res res
			_ = json.Unmarshal(body, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestOAuthGithub(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	users := []entity.Users{
		{Oid: "0", Role: entity.NOTACTIVE, AccountType: entity.GITHUB},
	}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByOidAndAccountType(users[0].Oid, users[0].AccountType).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersService.EXPECT().InsertUser(users[0]).Return(users[0].Uid, nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByOidAndAccountType(users[0].Oid, users[0].AccountType).Return(users[0], nil),
		mockUsersService.EXPECT().Destruct(),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/oauth/github", u.OAuthGithub)
	type args struct {
		code string
	}
	type result struct {
		First        bool          `json:"first"`
		Role         int8          `json:"role"`
		Id           bson.ObjectId `json:"id"`
		Token        string        `json:"token"`
		RefreshToken string        `json:"refresh_token"`
	}
	type res struct {
		Code   int8   `json:"code"`
		Result result `json:"result"`
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"NormalAndFirst", args{""}, http.StatusOK, res{0, result{First: true, Role: users[0].Role, Id: users[0].Uid}}},
		{"NormalAndNotFirst", args{""}, http.StatusOK, res{0, result{First: false, Role: users[0].Role, Id: users[0].Uid}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/oauth/github?code="+tt.args.code, nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res res
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	users := []entity.Users{
		{Nickname: "test", Email: "test@sjtu.edu.cn", Role: entity.NOTACTIVE, AccountType: entity.SOFIA},
	}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByNickname(users[0].Nickname).Return(users[0], nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByNickname(users[0].Nickname).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersService.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindUserByNickname(users[0].Nickname).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersService.EXPECT().FindUserByEmail(users[0].Email).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersService.EXPECT().InsertUser(users[0]).Return(users[0].Uid, nil),
		mockUsersService.EXPECT().Destruct(),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/register", u.Register)
	type args struct {
		name     string
		nickname string
		password string
		email    string
		icon     string
		gender   int8
	}
	type result struct {
		Type int8 `json:"type"`
	}
	type res struct {
		Code   int8   `json:"code"`
		Result result `json:"result"`
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"NicknameFound", args{users[0].Name, users[0].Nickname, users[0].Password, users[0].Email, users[0].Icon, users[0].Gender}, http.StatusOK, res{1, result{0}}},
		{"EmailFound", args{users[0].Name, users[0].Nickname, users[0].Password, users[0].Email, users[0].Icon, users[0].Gender}, http.StatusOK, res{1, result{1}}},
		{"MailFail", args{users[0].Name, users[0].Nickname, users[0].Password, users[0].Email, users[0].Icon, users[0].Gender}, http.StatusOK, res{1, result{2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req struct {
				Nickname string `json:"nickname"`
				Password string `json:"password"`
				Email    string `json:"email"`
			}
			req.Nickname = tt.args.nickname
			req.Password = tt.args.password
			req.Email = tt.args.email
			requestBody, _ := json.Marshal(req)
			r, _ := http.NewRequest("POST", "/register", bytes.NewReader(requestBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res res
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}
