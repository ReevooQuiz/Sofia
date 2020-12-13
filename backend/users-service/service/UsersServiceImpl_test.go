package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/mock"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActivate(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: bson.ObjectIdHex("000000000000000000000000"), Role: entity.USER},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(users[0]).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Destruct(),
	)
	u := UsersServiceImpl{mockUsersDao}
	mux.HandleFunc("/activate", u.Verify)
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

func TestLogin(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Name: "test", HashPassword: "test", Role: entity.DISABLE},
		{Name: "test", HashPassword: "test", Role: entity.USER},
		{Name: "test", HashPassword: "test", Role: entity.NOT_ACTIVE},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[0].Name).Return(users[0], nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[0].Name).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[1].Name).Return(users[1], nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[2].Name).Return(users[2], nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	u := UsersServiceImpl{mockUsersDao}
	mux.HandleFunc("/login", u.Login)
	type args struct {
		name     string
		password string
	}
	type result struct {
		Type         int8          `json:"type"`
		Role         int8          `json:"role"`
		Uid          bson.ObjectId `json:"uid"`
		Icon         string        `json:"icon"`
		Name         string        `json:"name"`
		Nickname     string        `json:"nickname"`
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
		{"Disable", args{users[0].Name, users[0].HashPassword}, http.StatusOK, res{1, result{Type: 0}}},
		{"NameNotFound", args{users[0].Name, users[0].HashPassword}, http.StatusOK, res{1, result{Type: 1}}},
		{"WrongPassword", args{users[1].Name, ""}, http.StatusOK, res{1, result{Type: 1}}},
		{"NotActive", args{users[2].Name, users[2].HashPassword}, http.StatusOK, res{1, result{Type: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req struct {
				Name     string `json:"name"`
				Password string `json:"password"`
			}
			req.Name = tt.args.name
			req.Password = tt.args.password
			requestBody, _ := json.Marshal(req)
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
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: bson.ObjectIdHex("000000000000000000000000"), Oid: "0", Role: entity.USER, AccountType: entity.GITHUB},
		{Oid: "0", Role: entity.DISABLE, AccountType: entity.GITHUB},
	}
	favorites := []entity.Favorites{
		{Uid: users[0].Uid.Hex(), Title: "Default"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByOidAndAccountType(users[0].Oid, users[0].AccountType).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().InsertUser(gomock.Any()).Return(users[0].Uid, nil),
		mockUsersDao.EXPECT().InsertFavorite(favorites[0]).Return(favorites[0].Fid, nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByOidAndAccountType(users[1].Oid, users[1].AccountType).Return(users[1], nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	u := UsersServiceImpl{mockUsersDao}
	mux.HandleFunc("/oauth/github", u.OAuthGithub)
	type args struct {
		code  string
		error string
	}
	type result struct {
		Type         int8          `json:"type"`
		First        bool          `json:"first"`
		Role         int8          `json:"role"`
		Uid          bson.ObjectId `json:"uid"`
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
		{"NormalAndFirst", args{}, http.StatusOK, res{0, result{First: true, Role: users[0].Role, Uid: users[0].Uid}}},
		{"Disable", args{}, http.StatusOK, res{1, result{Type: 0}}},
		{"AccessDenied", args{error: "access_denied"}, http.StatusOK, res{1, result{Type: 1}}},
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
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: bson.ObjectIdHex("000000000000000000000000"), HashPassword: "test", Role: entity.USER},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().UpdateUser(users[0]).Return(nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByUid(users[0].Uid).Return(users[0], nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	u := UsersServiceImpl{mockUsersDao}
	mux.HandleFunc("/passwd", u.Passwd)
	type args struct {
		old string
		new string
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
		{"Normal", args{"test", "test"}, http.StatusOK, res{Code: 0}},
		{"WrongName", args{"test", "test"}, http.StatusOK, res{1, result{1}}},
		{"WrongPassword", args{}, http.StatusOK, res{1, result{0}}},
		{"WrongToken", args{"test", "test"}, http.StatusOK, res{Code: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req struct {
				Old string `json:"old"`
				New string `json:"new"`
			}
			req.Old = tt.args.old
			req.New = tt.args.new
			requestBody, _ := json.Marshal(req)
			var token string
			if tt.name != "WrongToken" {
				token, _ = util.SignToken(users[0].Uid, users[0].Role, false)
			}
			r, _ := http.NewRequest("PUT", "/passwd", bytes.NewReader(requestBody))
			r.Header.Set("Authorization", token)
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
	mockUsersDao := mock.NewMockUsersDao(mockCtrl)
	users := []entity.Users{
		{Uid: bson.ObjectIdHex("000000000000000000000000"), Name: "test", Email: "test@sjtu.edu.cn", Role: entity.NOT_ACTIVE, AccountType: entity.SOFIA},
	}
	favorites := []entity.Favorites{
		{Uid: users[0].Uid.Hex(), Title: "Default"},
	}
	gomock.InOrder(
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[0].Name).Return(users[0], nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[0].Name).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(users[0], nil),
		mockUsersDao.EXPECT().Destruct(),
		mockUsersDao.EXPECT().Init().Return(nil),
		mockUsersDao.EXPECT().FindUserByName(users[0].Name).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().FindUserByEmail(users[0].Email).Return(entity.Users{}, errors.New("mongo: no rows in result set")),
		mockUsersDao.EXPECT().InsertUser(gomock.Any()).Return(users[0].Uid, nil),
		mockUsersDao.EXPECT().InsertFavorite(favorites[0]).Return(favorites[0].Fid, nil),
		mockUsersDao.EXPECT().Destruct(),
	)
	u := UsersServiceImpl{mockUsersDao}
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
		{"NameFound", args{users[0].Name, users[0].Nickname, users[0].HashPassword, users[0].Email, users[0].Icon, users[0].Gender}, http.StatusOK, res{1, result{0}}},
		{"EmailFound", args{users[0].Name, users[0].Nickname, users[0].HashPassword, users[0].Email, users[0].Icon, users[0].Gender}, http.StatusOK, res{1, result{1}}},
		{"MailFail", args{users[0].Name, users[0].Nickname, users[0].HashPassword, users[0].Email, users[0].Icon, users[0].Gender}, http.StatusOK, res{1, result{2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req struct {
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Password string `json:"password"`
				Email    string `json:"email"`
				Icon     string `json:"icon"`
				Gender   int8   `json:"gender"`
			}
			req.Name = tt.args.name
			req.Nickname = tt.args.nickname
			req.Password = tt.args.password
			req.Email = tt.args.email
			req.Icon = tt.args.icon
			req.Gender = tt.args.gender
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
