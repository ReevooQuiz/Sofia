package controller

import (
	"encoding/json"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/service"
	"gopkg.in/gomail.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type UsersController struct {
	usersService service.UsersService
}

var emailPassword string
var emailUsername string
var oAuthGithubId string
var oAuthGithubSecret string

func init() {
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	emailPassword = os.Getenv("EMAIL_PASSWORD")
	emailUsername = os.Getenv("EMAIL_USERNAME")
	oAuthGithubId = os.Getenv("OAUTH_GITHUB_ID")
	oAuthGithubSecret = os.Getenv("OAUTH_GITHUB_SECRET")
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
	token := bson.ObjectIdHex(r.FormValue("token"))
	var user entity.Users
	user, err = u.usersService.FindUserByUid(token)
	if err != nil {
		log.Info(err)
		res.Code = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	user.Role = entity.USER
	err = u.usersService.UpdateUser(user)
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
	request, err = http.NewRequest("POST", "https://github.com/login/oauth/access_token?client_id="+oAuthGithubId+"&client_secret="+oAuthGithubSecret+"&code="+r.FormValue("code"), nil)
	if err != nil {
		log.Info(err)
		return
	}
	request.Header.Set("Accept", "application/json")
	client := http.Client{}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		log.Info(err)
		return
	}
	var responseBodyJson []byte
	responseBodyJson, err = ioutil.ReadAll(response.Body)
	var responseBodyToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	err = json.Unmarshal(responseBodyJson, &responseBodyToken)
	if err != nil {
		log.Info(err)
		return
	}
	request, err = http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Info(err)
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "token "+responseBodyToken.AccessToken)
	response, err = client.Do(request)
	if err != nil {
		log.Info(err)
		return
	}
	responseBodyJson, err = ioutil.ReadAll(response.Body)
	var responseBodyInfo struct {
		Login             string `json:"login"`
		Id                int64  `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
		Name              string `json:"name"`
		Company           string `json:"company"`
		Blog              string `json:"blog"`
		Location          string `json:"location"`
		Email             string `json:"email"`
		Hireable          string `json:"hireable"`
		Bio               string `json:"bio"`
		TwitterUsername   string `json:"twitter_username"`
		PublicRepos       int64  `json:"public_repos"`
		PublicGists       int64  `json:"public_gists"`
		Followers         int64  `json:"followers"`
		Following         int64  `json:"following"`
		CreatedAt         string `json:"created_at"`
		UpdatedAt         string `json:"updated_at"`
	}
	err = json.Unmarshal(responseBodyJson, &responseBodyInfo)
	if err != nil {
		log.Info(err)
		return
	}
	log.Info(responseBodyInfo)
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
	user.Nickname = req.Username
	user.Password = req.Password
	user.Email = req.Email
	user.Role = entity.NOTACTIVE
	user.Type = entity.SOFIA
	_, err = u.usersService.FindUserByNickname(user.Nickname)
	if err == nil {
		res.Code = 1
		res.Result.Type = 0
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	_, err = u.usersService.FindUserByEmail(user.Email)
	if err == nil {
		res.Code = 1
		res.Result.Type = 1
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	user.Uid, err = u.usersService.InsertUser(user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
		return
	}
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(emailUsername, "Sofia"))
	message.SetHeader("To", user.Email)
	message.SetHeader("Subject", "Sofia")
	message.SetBody("text/html", `<a href="http://localhost:9092/activate?token=`+user.Uid.Hex()+`">activate</a>`)
	err = gomail.NewDialer("smtp.qq.com", 587, emailUsername, emailPassword).DialAndSend(message)
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
