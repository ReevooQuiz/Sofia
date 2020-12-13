package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/mock"
	"github.com/zhanghanchong/users-service/service"
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

func TestLogin(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Login(gomock.Any()).Return(service.ResLogin{}, nil),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/login", u.Login)
	type args struct {
		req service.ReqLogin
	}
	type res struct {
		res service.ResLogin
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"Normal", args{}, http.StatusOK, res{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("POST", "/login", bytes.NewReader(requestBody))
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

func TestOAuthGithub(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().OAuthGithub(gomock.Any(), gomock.Any()).Return(service.ResOAuthGithub{}, nil),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/oauth/github", u.OAuthGithub)
	type args struct {
		code  string
		error string
	}
	type res struct {
		res service.ResOAuthGithub
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"Normal", args{}, http.StatusOK, res{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/oauth/github?code="+tt.args.code+"&error="+tt.args.error, nil)
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

func TestPasswd(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Passwd(gomock.Any(), gomock.Any()).Return(service.ResPasswd{}, nil),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/passwd", u.Passwd)
	type args struct {
		token string
		req   service.ReqPasswd
	}
	type res struct {
		res service.ResPasswd
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"Normal", args{}, http.StatusOK, res{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("PUT", "/passwd", bytes.NewReader(requestBody))
			r.Header.Set("Authorization", tt.args.token)
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

func TestRegister(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Register(gomock.Any()).Return(service.ResRegister{}, nil),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/register", u.Register)
	type args struct {
		req service.ReqRegister
	}
	type res struct {
		res service.ResRegister
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    res
	}{
		{"Normal", args{}, http.StatusOK, res{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
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
