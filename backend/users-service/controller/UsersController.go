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

func (u *UsersController) Init(group *sync.WaitGroup, usersService service.UsersService) (server *http.Server) {
	u.usersService = usersService
	server = &http.Server{Addr: ":9092"}
	http.HandleFunc("/activate", u.Activate)
	http.HandleFunc("/oauth/github", u.OAuthGithub)
	http.HandleFunc("/register", u.Register)
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}

func (u *UsersController) Activate(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Code int8 `json:"code"`
	}
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
	var token int64
	token, err = strconv.ParseInt(r.FormValue("token"), 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var user entity.Users
	user, err = u.usersService.FindById(token)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	user.Role = entity.USER
	err = u.usersService.Update(user)
	if err != nil {
		log.Info(err)
		res.Code = 1
	} else {
		res.Code = 0
	}
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}

func (u *UsersController) OAuthGithub(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Info(err)
		return
	}
	var request *http.Request
	request, err = http.NewRequest("POST", "https://github.com/login/oauth/access_token?client_id=51f0dde36e2f4fcee97c&client_secret=04aee9d3c62d4ea10577113dedbf62b842f8a855&code="+r.FormValue("code"), nil)
	if err != nil {
		log.Info(err)
		return
	}
	request.Header.Set("accept", "application/json")
	client := http.Client{}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		log.Info(err)
		return
	}
	responseBody := make([]byte, 92)
	_, err = response.Body.Read(responseBody)
	_, _ = w.Write(responseBody)
}

func (u *UsersController) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	var res struct {
		Code   int8 `json:"code"`
		Result struct {
			Type int8 `json:"type"`
		} `json:"result"`
	}
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
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	var user entity.Users
	user.Username = req.Username
	user.Password = req.Password
	user.Email = req.Email
	user.Role = entity.NOTACTIVE
	_, err = u.usersService.FindByUsername(user.Username)
	if err == nil {
		res.Code = 1
		res.Result.Type = 0
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	_, err = u.usersService.FindByEmail(user.Email)
	if err == nil {
		res.Code = 1
		res.Result.Type = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	user.Id, err = u.usersService.Insert(user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
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
	object, _ := json.Marshal(res)
	_, _ = w.Write(object)
}
