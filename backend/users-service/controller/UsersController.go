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

func (u *UsersController) Init(group *sync.WaitGroup, usersService service.UsersService) (server *http.Server) {
	u.usersService = usersService
	server = &http.Server{Addr: ":9092"}
	http.HandleFunc("/login", u.Login)
	http.HandleFunc("/oauth/github", u.OAuthGithub)
	http.HandleFunc("/passwd", u.Passwd)
	http.HandleFunc("/register", u.Register)
	http.HandleFunc("/verificationCode", u.VerificationCode)
	http.HandleFunc("/verify", u.Verify)
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
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

func (u *UsersController) Register(w http.ResponseWriter, r *http.Request) {
	var req service.ReqRegister
	var res service.ResRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
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
