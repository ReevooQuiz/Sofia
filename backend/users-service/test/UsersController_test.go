package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/controller"
	"github.com/zhanghanchong/users-service/mock"
	"github.com/zhanghanchong/users-service/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestControllerInit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
	)
	tests := []struct {
		name string
	}{
		{"Initialize"},
	}
	u := controller.UsersController{}
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

func TestControllerBan(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Ban(gomock.Any(), gomock.Any()).Return(service.ResBan{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/ban", u.Ban)
	type args struct {
		token string
		req   service.ReqBan
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResBan
	}{
		{"Normal", args{}, http.StatusOK, service.ResBan{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("PUT", "/ban", bytes.NewReader(requestBody))
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResBan
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerBanned(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Banned(gomock.Any(), gomock.Any()).Return(service.ResBanned{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/banned", u.Banned)
	type args struct {
		token string
		page  int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResBanned
	}{
		{"Normal", args{}, http.StatusOK, service.ResBanned{}},
		{"WrongPage", args{page: -1}, http.StatusOK, service.ResBanned{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/banned?page="+strconv.FormatInt(tt.args.page, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResBanned
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerCheckToken(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().CheckToken(gomock.Any()).Return(service.ResCheckToken{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/checkToken", u.CheckToken)
	type args struct {
		token string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResCheckToken
	}{
		{"Normal", args{}, http.StatusOK, service.ResCheckToken{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/checkToken?token="+tt.args.token, nil)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResCheckToken
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerFollow(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Follow(gomock.Any(), gomock.Any(), gomock.Any()).Return(service.ResFollow{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/follow", u.Follow)
	type args struct {
		token  string
		uid    int64
		follow bool
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResFollow
	}{
		{"Normal", args{}, http.StatusOK, service.ResFollow{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var follow string
			if tt.args.follow {
				follow = "true"
			} else {
				follow = "false"
			}
			r, _ := http.NewRequest("PUT", "/follow?uid="+strconv.FormatInt(tt.args.uid, 10)+"&follow="+follow, bytes.NewReader([]byte{}))
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResFollow
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerFollowed(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Followed(gomock.Any(), gomock.Any()).Return(service.ResFollowed{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/followed", u.Followed)
	type args struct {
		token string
		uid   int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResFollowed
	}{
		{"Normal", args{}, http.StatusOK, service.ResFollowed{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/followed?uid="+strconv.FormatInt(tt.args.uid, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResFollowed
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerFollowers(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Followers(gomock.Any(), gomock.Any()).Return(service.ResFollowers{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/followers", u.Followers)
	type args struct {
		token string
		uid   int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResFollowers
	}{
		{"Normal", args{}, http.StatusOK, service.ResFollowers{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/followers?uid="+strconv.FormatInt(tt.args.uid, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResFollowers
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerInfoList(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().InfoList(gomock.Any()).Return(service.ResInfoList{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/infoList", u.InfoList)
	type args struct {
		req service.ReqInfoList
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResInfoList
	}{
		{"Normal", args{}, http.StatusOK, service.ResInfoList{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("POST", "/infoList", bytes.NewReader(requestBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResInfoList
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerLogin(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Login(gomock.Any()).Return(service.ResLogin{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/login", u.Login)
	type args struct {
		req service.ReqLogin
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResLogin
	}{
		{"Normal", args{}, http.StatusOK, service.ResLogin{}},
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
			var res service.ResLogin
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerNotifications(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Notifications(gomock.Any(), gomock.Any()).Return(service.ResNotifications{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/notifications", u.Notifications)
	type args struct {
		token string
		page  int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResNotifications
	}{
		{"Normal", args{}, http.StatusOK, service.ResNotifications{}},
		{"WrongPage", args{page: -1}, http.StatusOK, service.ResNotifications{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/notifications?page="+strconv.FormatInt(tt.args.page, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResNotifications
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerOAuthGithub(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().OAuthGithub(gomock.Any(), gomock.Any()).Return(service.ResOAuthGithub{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/oauth/github", u.OAuthGithub)
	type args struct {
		code  string
		error string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResOAuthGithub
	}{
		{"Normal", args{}, http.StatusOK, service.ResOAuthGithub{}},
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
			var res service.ResOAuthGithub
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerPasswd(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Passwd(gomock.Any(), gomock.Any()).Return(service.ResPasswd{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/passwd", u.Passwd)
	type args struct {
		token string
		req   service.ReqPasswd
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResPasswd
	}{
		{"Normal", args{}, http.StatusOK, service.ResPasswd{}},
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
			var res service.ResPasswd
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerPublicInfoGet(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().PublicInfoGet(gomock.Any(), gomock.Any()).Return(service.ResPublicInfoGet{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/publicInfo", u.PublicInfo)
	type args struct {
		token string
		uid   int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResPublicInfoGet
	}{
		{"Normal", args{}, http.StatusOK, service.ResPublicInfoGet{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/publicInfo?uid="+strconv.FormatInt(tt.args.uid, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResPublicInfoGet
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerPublicInfoPut(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().PublicInfoPut(gomock.Any(), gomock.Any()).Return(service.ResPublicInfoPut{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/publicInfo", u.PublicInfo)
	type args struct {
		token string
		req   service.ReqPublicInfoPut
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResPublicInfoPut
	}{
		{"Normal", args{}, http.StatusOK, service.ResPublicInfoPut{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("PUT", "/publicInfo", bytes.NewReader(requestBody))
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResPublicInfoPut
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerRefreshToken(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().RefreshToken(gomock.Any()).Return(service.ResRefreshToken{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/refreshToken", u.RefreshToken)
	type args struct {
		req service.ReqRefreshToken
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResRefreshToken
	}{
		{"Normal", args{}, http.StatusOK, service.ResRefreshToken{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("POST", "/refreshToken", bytes.NewReader(requestBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResRefreshToken
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerRegister(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Register(gomock.Any()).Return(service.ResRegister{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/register", u.Register)
	type args struct {
		req service.ReqRegister
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResRegister
	}{
		{"Normal", args{}, http.StatusOK, service.ResRegister{}},
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
			var res service.ResRegister
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerUserAnswers(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().UserAnswers(gomock.Any(), gomock.Any(), gomock.Any()).Return(service.ResUserAnswers{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/userAnswers", u.UserAnswers)
	type args struct {
		token string
		uid   int64
		page  int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResUserAnswers
	}{
		{"Normal", args{}, http.StatusOK, service.ResUserAnswers{}},
		{"WrongPage", args{page: -1}, http.StatusOK, service.ResUserAnswers{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/userAnswers?uid="+strconv.FormatInt(tt.args.uid, 10)+"&page="+strconv.FormatInt(tt.args.page, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResUserAnswers
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerUserQuestions(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().UserQuestions(gomock.Any(), gomock.Any(), gomock.Any()).Return(service.ResUserQuestions{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/userQuestions", u.UserQuestions)
	type args struct {
		token string
		uid   int64
		page  int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResUserQuestions
	}{
		{"Normal", args{}, http.StatusOK, service.ResUserQuestions{}},
		{"WrongPage", args{page: -1}, http.StatusOK, service.ResUserQuestions{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/userQuestions?uid="+strconv.FormatInt(tt.args.uid, 10)+"&page="+strconv.FormatInt(tt.args.page, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResUserQuestions
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerVerificationCode(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().VerificationCode(gomock.Any(), gomock.Any()).Return(service.ResVerificationCode{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/verificationCode", u.VerificationCode)
	type args struct {
		register bool
		email    string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResVerificationCode
	}{
		{"Normal", args{}, http.StatusOK, service.ResVerificationCode{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var register string
			if tt.args.register {
				register = "true"
			} else {
				register = "false"
			}
			r, _ := http.NewRequest("GET", "/verificationCode?register="+register+"&email="+tt.args.email, nil)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResVerificationCode
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerVerify(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().Verify(gomock.Any(), gomock.Any()).Return(service.ResVerify{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/verify", u.Verify)
	type args struct {
		email string
		code  int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResVerify
	}{
		{"Normal", args{}, http.StatusOK, service.ResVerify{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/verify?email="+tt.args.email+"&code="+strconv.FormatInt(tt.args.code, 10), nil)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResVerify
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerWordBan(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().WordBan(gomock.Any(), gomock.Any()).Return(service.ResWordBan{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/wordBan", u.WordBan)
	type args struct {
		token string
		req   service.ReqWordBan
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResWordBan
	}{
		{"Normal", args{}, http.StatusOK, service.ResWordBan{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.req)
			r, _ := http.NewRequest("PUT", "/wordBan", bytes.NewReader(requestBody))
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResWordBan
			_ = json.Unmarshal(responseBody, &res)
			if res != tt.wantRes {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}

func TestControllerWordsBanned(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	gomock.InOrder(
		mockUsersService.EXPECT().WordsBanned(gomock.Any(), gomock.Any()).Return(service.ResWordsBanned{}, nil),
	)
	var u controller.UsersController
	u.SetUsersService(mockUsersService)
	mux.HandleFunc("/wordsBanned", u.WordsBanned)
	type args struct {
		token string
		page  int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRes    service.ResWordsBanned
	}{
		{"Normal", args{}, http.StatusOK, service.ResWordsBanned{}},
		{"WrongPage", args{page: -1}, http.StatusOK, service.ResWordsBanned{Code: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/wordsBanned?page="+strconv.FormatInt(tt.args.page, 10), nil)
			r.Header.Set("Authorization", tt.args.token)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, tt.wantStatus)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res service.ResWordsBanned
			_ = json.Unmarshal(responseBody, &res)
			if res.Code != tt.wantRes.Code {
				t.Errorf("Actual: %v, expect: %v.", res, tt.wantRes)
			}
		})
	}
}
