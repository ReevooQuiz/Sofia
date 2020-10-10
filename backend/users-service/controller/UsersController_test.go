package controller

import (
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
	user := entity.Users{Id: 1, Username: "root", Password: "root", Email: "root@sjtu.edu.cn", Role: 3}
	gomock.InOrder(
		mockUsersService.EXPECT().Init().Return(nil),
		mockUsersService.EXPECT().FindById(1).Return(user, nil),
		mockUsersService.EXPECT().Update(user).Return(nil),
		mockUsersService.EXPECT().Destruct(),
	)
	u := UsersController{mockUsersService}
	mux.HandleFunc("/activate", u.activate)
	type args struct {
		token string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{"Normal", args{"1"}, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, "/activate?token="+tt.args.token, reader)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			resp := w.Result()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("Actual: %v, expect: %v.", resp.StatusCode, tt.wantStatus)
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
