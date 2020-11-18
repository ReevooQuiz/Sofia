package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/mock"
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
		{1, "root", "root", "root@sjtu.edu.cn", entity.NOTACTIVE},
		{1, "root", "root", "root@sjtu.edu.cn", entity.USER},
	}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindById(users[0].Uid).Return(users[0], nil),
		mockUsersService.EXPECT().Update(users[1]).Return(nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindById(int64(0)).Return(entity.Users{}, errors.New("sql: no rows in result set")),
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
		{"Normal", args{"1"}, http.StatusOK, res{0}},
		{"TokenNotFound", args{"nil"}, http.StatusOK, res{1}},
		{"IdNotFound", args{"0"}, http.StatusOK, res{1}},
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
	u := UsersController{mockUsersService}
	mux.HandleFunc("/oauth/github", u.OAuthGithub)
	type args struct {
		code string
	}
	type res struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"Normal", args{""}, http.StatusOK, res{"", "", ""}},
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
		{0, "root", "root", "root@sjtu.edu.cn", entity.NOTACTIVE},
		{0, "root", "root", "root", entity.NOTACTIVE},
	}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindByUsername(users[0].Username).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersService.EXPECT().FindByEmail(users[0].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersService.EXPECT().Insert(users[0]).Return(int64(1), nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindByUsername(users[0].Username).Return(users[0], nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindByUsername(users[0].Username).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersService.EXPECT().FindByEmail(users[0].Email).Return(users[0], nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindByUsername(users[1].Username).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersService.EXPECT().FindByEmail(users[1].Email).Return(entity.Users{}, errors.New("sql: no rows in result set")),
		mockUsersService.EXPECT().Insert(users[1]).Return(int64(1), nil),
		mockUsersService.EXPECT().Destruct(),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/register", u.Register)
	type args struct {
		username string
		password string
		email    string
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
		{"Normal", args{users[0].Username, users[0].Password, users[0].Email}, http.StatusOK, res{0, result{0}}},
		{"UsernameFound", args{users[0].Username, users[0].Password, users[0].Email}, http.StatusOK, res{1, result{0}}},
		{"EmailFound", args{users[0].Username, users[0].Password, users[0].Email}, http.StatusOK, res{1, result{1}}},
		{"MailFail", args{users[1].Username, users[1].Password, users[1].Email}, http.StatusOK, res{1, result{2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req struct {
				Username string `json:"username"`
				Password string `json:"password"`
				Email    string `json:"email"`
			}
			req.Username = tt.args.username
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
