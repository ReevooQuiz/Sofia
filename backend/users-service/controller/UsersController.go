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
	server = &http.Server{Addr: ":9092"}
	http.HandleFunc("/checkToken", u.CheckToken)
	http.HandleFunc("/infoList", u.InfoList)
	http.HandleFunc("/login", u.Login)
	http.HandleFunc("/oauth/github", u.OAuthGithub)
	http.HandleFunc("/passwd", u.Passwd)
	http.HandleFunc("/publicInfo", u.PublicInfo)
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

func (u *UsersController) CheckToken(w http.ResponseWriter, r *http.Request) {
	var res service.ResCheckToken
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
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

func (u *UsersController) InfoList(w http.ResponseWriter, r *http.Request) {
	var req service.ReqInfoList
	var res service.ResInfoList
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	res, err = u.usersService.InfoList(r.Header.Get("Authorization"), req)
	if err != nil {
		log.Info(err)
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) Login(w http.ResponseWriter, r *http.Request) {
	var req service.ReqLogin
	var res service.ResLogin
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
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
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = r.ParseForm()
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
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
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
		err := u.usersService.Init()
		defer u.usersService.Destruct()
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		err = r.ParseForm()
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
		err := u.usersService.Init()
		defer u.usersService.Destruct()
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&req)
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

func (u *UsersController) Register(w http.ResponseWriter, r *http.Request) {
	var req service.ReqRegister
	var res service.ResRegister
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
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
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = r.ParseForm()
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
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	err = r.ParseForm()
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
