package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/service"
	"gopkg.in/gomail.v2"
	"net/http"
	"strconv"
	"sync"
)

type UsersController struct {
	usersService service.UsersService
}

func (u *UsersController) activate(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Code int8 `json:"code"`
	}
	var object []byte
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	var token int64
	token, err = strconv.ParseInt(r.Form["token"][0], 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	var user entity.Users
	user, err = u.usersService.FindById(token)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	user.Role = 1
	err = u.usersService.Update(user)
	if err != nil {
		log.Info(err)
		res.Code = 1
	} else {
		res.Code = 0
	}
	object, err = json.Marshal(res)
	if err != nil {
		log.Info(err)
		return
	}
	_, err = w.Write(object)
	if err != nil {
		log.Info(err)
	}
}

func (u *UsersController) register(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Code   int8 `json:"code"`
		Result struct {
			Type int8 `json:"type"`
		} `json:"result"`
	}
	var object []byte
	err := u.usersService.Init()
	defer u.usersService.Destruct()
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	var user entity.Users
	user.Username = r.PostFormValue("username")
	user.Password = r.PostFormValue("password")
	user.Email = r.PostFormValue("email")
	user.Role = 3
	_, err = u.usersService.FindByUsername(user.Username)
	if err == nil {
		res.Code = 1
		res.Result.Type = 0
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	log.Info(err)
	_, err = u.usersService.FindByEmail(user.Email)
	if err == nil {
		res.Code = 1
		res.Result.Type = 1
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	log.Info(err)
	user.Id, err = u.usersService.Insert(user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, err = json.Marshal(res)
		if err != nil {
			log.Info(err)
			return
		}
		_, err = w.Write(object)
		if err != nil {
			log.Info(err)
		}
		return
	}
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress("308011618@qq.com", "Sofia"))
	message.SetHeader("To", user.Email)
	message.SetHeader("Subject", "Sofia")
	message.SetBody("text/html", `<a href="http://localhost:9090/activate?token=`+strconv.FormatInt(user.Id, 10)+`">activate</a>`)
	err = gomail.NewDialer("smtp.qq.com", 587, "308011618@qq.com", "czanokubhfrubhjj").DialAndSend(message)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
	} else {
		res.Code = 0
	}
	object, err = json.Marshal(res)
	if err != nil {
		log.Info(err)
		return
	}
	_, err = w.Write(object)
	if err != nil {
		log.Info(err)
	}
}

func (u *UsersController) Init(group *sync.WaitGroup, usersService service.UsersService) *http.Server {
	u.usersService = usersService
	server := &http.Server{Addr: ":9090"}
	http.HandleFunc("/activate", u.activate)
	http.HandleFunc("/register", u.register)
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}
