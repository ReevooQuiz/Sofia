package controller

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//func TestInit(t *testing.T) {
//	type fields struct {
//		usersService service.UsersService
//	}
//	type args struct {
//		group        *sync.WaitGroup
//		usersService service.UsersService
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *http.Server
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := &UsersController{
//				usersService: tt.fields.usersService,
//			}
//			if got := u.Init(tt.args.group, tt.args.usersService); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Init() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestActivate(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	reader := strings.NewReader(`{"token":"1"}`)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersService := mock.NewMockUsersService(mockCtrl)
	users := [][]entity.Users{
		{{1, "root", "root", "root@sjtu.edu.cn", 3}, {1, "root", "root", "root@sjtu.edu.cn", 1}},
		{{}, {}},
	}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindById(users[0][0].Id).Return(users[0][0], nil),
		mockUsersService.EXPECT().Update(users[0][1]).Return(nil),
		mockUsersService.EXPECT().Destruct(),
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindById(users[1][0].Id).Return(users[1][0], errors.New("sql: no rows in result set")),
		mockUsersService.EXPECT().Destruct(),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/activate", u.activate)
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
		{"NotFound", args{"0"}, http.StatusOK, res{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, "/activate?token="+tt.args.token, reader)
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

//func TestRegister(t *testing.T) {
//	type fields struct {
//		usersService service.UsersService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := &UsersController{
//				usersService: tt.fields.usersService,
//			}
//		})
//	}
//}
