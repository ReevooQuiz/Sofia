package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/users-service/service"
	"net/http"
	"strconv"
	"sync"
)

type UsersController struct {
	usersService service.UsersService
}

func (u *UsersController) SetUsersService(usersService service.UsersService) {
	u.usersService = usersService
}

func (u *UsersController) Init(group *sync.WaitGroup, usersService service.UsersService) (server *http.Server) {
	u.usersService = usersService
	err := u.usersService.Init()
	if err != nil {
		log.Info(err)
		return server
	}
	server = &http.Server{Addr: ":9092"}
	http.HandleFunc("/ban", u.Ban)
	http.HandleFunc("/banned", u.Banned)
	http.HandleFunc("/checkSession", u.CheckSession)
	http.HandleFunc("/checkToken", u.CheckToken)
	http.HandleFunc("/follow", u.Follow)
	http.HandleFunc("/followed", u.Followed)
	http.HandleFunc("/followers", u.Followers)
	http.HandleFunc("/infoList", u.InfoList)
	http.HandleFunc("/like", u.Like)
	http.HandleFunc("/login", u.Login)
	http.HandleFunc("/notifications", u.Notifications)
	http.HandleFunc("/oauth/github", u.OAuthGithub)
	http.HandleFunc("/passwd", u.Passwd)
	http.HandleFunc("/publicInfo", u.PublicInfo)
	http.HandleFunc("/refreshToken", u.RefreshToken)
	http.HandleFunc("/register", u.Register)
	http.HandleFunc("/userAnswers", u.UserAnswers)
	http.HandleFunc("/userQuestions", u.UserQuestions)
	http.HandleFunc("/verificationCode", u.VerificationCode)
	http.HandleFunc("/verify", u.Verify)
	http.HandleFunc("/wordBan", u.WordBan)
	http.HandleFunc("/wordsBanned", u.WordsBanned)
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}

func (u *UsersController) Destruct() {
	u.usersService.Destruct()
}

func (u *UsersController) Ban(w http.ResponseWriter, r *http.Request) {
	var req service.ReqBan
	var res service.ResBan
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Ban(r.Header.Get("Authorization"), req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Banned(w http.ResponseWriter, r *http.Request) {
	var res service.ResBanned
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var page int64
	page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)
	if err != nil || page < 0 {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Banned(r.Header.Get("Authorization"), page)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) CheckSession(w http.ResponseWriter, r *http.Request) {
	res, err := u.usersService.CheckSession(r.Header.Get("Authorization"))
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) CheckToken(w http.ResponseWriter, r *http.Request) {
	var res service.ResCheckToken
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Successful = false
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.CheckToken(r.FormValue("token"))
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Follow(w http.ResponseWriter, r *http.Request) {
	var res service.ResFollow
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var uid int64
	uid, err = strconv.ParseInt(r.FormValue("uid"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Follow(r.Header.Get("Authorization"), uid, r.FormValue("follow") == "true")
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Followed(w http.ResponseWriter, r *http.Request) {
	var res service.ResFollowed
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var uid int64
	uid, err = strconv.ParseInt(r.FormValue("uid"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Followed(r.Header.Get("Authorization"), uid)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Followers(w http.ResponseWriter, r *http.Request) {
	var res service.ResFollowers
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var uid int64
	uid, err = strconv.ParseInt(r.FormValue("uid"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Followers(r.Header.Get("Authorization"), uid)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) InfoList(w http.ResponseWriter, r *http.Request) {
	var req service.ReqInfoList
	var res service.ResInfoList
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.InfoList(req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Like(w http.ResponseWriter, r *http.Request) {
	var req service.ReqLike
	var res service.ResLike
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Like(r.Header.Get("Authorization"), req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Login(w http.ResponseWriter, r *http.Request) {
	var req service.ReqLogin
	var res service.ResLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Login(req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Notifications(w http.ResponseWriter, r *http.Request) {
	var res service.ResNotifications
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var page int64
	page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)
	if err != nil || page < 0 {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Notifications(r.Header.Get("Authorization"), page)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) OAuthGithub(w http.ResponseWriter, r *http.Request) {
	var res service.ResOAuthGithub
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.OAuthGithub(r.FormValue("code"), r.FormValue("error"))
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Passwd(w http.ResponseWriter, r *http.Request) {
	var req service.ReqPasswd
	var res service.ResPasswd
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Passwd(r.Header.Get("Authorization"), req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) PublicInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var res service.ResPublicInfoGet
		err := r.ParseForm()
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		var uid int64
		uid, err = strconv.ParseInt(r.FormValue("uid"), 10, 64)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		res, err = u.usersService.PublicInfoGet(r.Header.Get("Authorization"), uid)
		if err != nil {
			log.Info(err)
		}
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
	}
	if r.Method == "PUT" {
		var req service.ReqPublicInfoPut
		var res service.ResPublicInfoPut
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		res, err = u.usersService.PublicInfoPut(r.Header.Get("Authorization"), req)
		if err != nil {
			log.Info(err)
		}
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
	}
}

func (u *UsersController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req service.ReqRefreshToken
	var res service.ResRefreshToken
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.RefreshToken(req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Register(w http.ResponseWriter, r *http.Request) {
	var req service.ReqRegister
	var res service.ResRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Register(req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) UserAnswers(w http.ResponseWriter, r *http.Request) {
	var res service.ResUserAnswers
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var uid int64
	uid, err = strconv.ParseInt(r.FormValue("uid"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var page int64
	page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)
	if err != nil || page < 0 {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.UserAnswers(r.Header.Get("Authorization"), uid, page)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) UserQuestions(w http.ResponseWriter, r *http.Request) {
	var res service.ResUserQuestions
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var uid int64
	uid, err = strconv.ParseInt(r.FormValue("uid"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var page int64
	page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)
	if err != nil || page < 0 {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.UserQuestions(r.Header.Get("Authorization"), uid, page)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) VerificationCode(w http.ResponseWriter, r *http.Request) {
	var res service.ResVerificationCode
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.VerificationCode(r.FormValue("register") == "true", r.FormValue("email"))
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Verify(w http.ResponseWriter, r *http.Request) {
	var res service.ResVerify
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var code int64
	code, err = strconv.ParseInt(r.FormValue("code"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.Verify(r.FormValue("email"), code)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) WordBan(w http.ResponseWriter, r *http.Request) {
	var req service.ReqWordBan
	var res service.ResWordBan
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.WordBan(r.Header.Get("Authorization"), req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) WordsBanned(w http.ResponseWriter, r *http.Request) {
	var res service.ResWordsBanned
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var page int64
	page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)
	if err != nil || page < 0 {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.WordsBanned(r.Header.Get("Authorization"), page)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}