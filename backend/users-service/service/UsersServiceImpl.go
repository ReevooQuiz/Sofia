package service

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
	"github.com/zhanghanchong/users-service/util"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	emailPassword     string
	emailUsername     string
	oAuthGithubId     string
	oAuthGithubSecret string
)

type UsersServiceImpl struct {
	usersDao dao.UsersDao
}

type ReqInfoList struct {
	Uids []int64 `json:"uids"`
}

type ReqLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ReqPasswd struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type ReqPublicInfoPut struct {
	Name     string   `json:"name"`
	Nickname string   `json:"nickname"`
	Profile  string   `json:"profile"`
	Icon     string   `json:"icon"`
	Gender   int8     `json:"gender"`
	Email    string   `json:"email"`
	Labels   []string `json:"labels"`
}

type ReqRefreshToken struct {
	Refresh string `json:"refresh"`
}

type ReqRegister struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Icon     string `json:"icon"`
	Gender   int8   `json:"gender"`
}

type ResCheckToken struct {
	Successful bool  `json:"successful"`
	Uid        int64 `json:"uid"`
	Role       int8  `json:"role"`
}

type ResFollow struct {
	Code int8 `json:"code"`
}

type ResFollowed struct {
	Code   int8             `json:"code"`
	Result []ResultFollowed `json:"result"`
}

type ResFollowers struct {
	Code   int8              `json:"code"`
	Result []ResultFollowers `json:"result"`
}

type ResInfoList struct {
	Code   int8             `json:"code"`
	Result []ResultInfoList `json:"result"`
}

type ResLogin struct {
	Code   int8        `json:"code"`
	Result ResultLogin `json:"result"`
}

type ResNotifications struct {
	Code   int8                  `json:"code"`
	Result []ResultNotifications `json:"result"`
}

type ResOAuthGithub struct {
	Code   int8              `json:"code"`
	Result ResultOAuthGithub `json:"result"`
}

type ResPasswd struct {
	Code   int8         `json:"code"`
	Result ResultPasswd `json:"result"`
}

type ResPublicInfoGet struct {
	Code   int8                `json:"code"`
	Result ResultPublicInfoGet `json:"result"`
}

type ResPublicInfoPut struct {
	Code   int8                `json:"code"`
	Result ResultPublicInfoPut `json:"result"`
}

type ResRefreshToken struct {
	Code   int8               `json:"code"`
	Result ResultRefreshToken `json:"result"`
}

type ResRegister struct {
	Code   int8           `json:"code"`
	Result ResultRegister `json:"result"`
}

type ResUserQuestions struct {
	Code   int8                  `json:"code"`
	Result []ResultUserQuestions `json:"result"`
}

type ResVerificationCode struct {
	Code   int8                   `json:"code"`
	Result ResultVerificationCode `json:"result"`
}

type ResVerify struct {
	Code int8 `json:"code"`
}

type ResultFollowed struct {
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}

type ResultFollowers struct {
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}

type ResultInfoList struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type ResultLogin struct {
	Type         int8   `json:"type"`
	Role         int8   `json:"role"`
	Uid          string `json:"uid"`
	Icon         string `json:"icon"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResultNotifications struct {
	Type              int8      `json:"type"`
	Time              time.Time `json:"time"`
	Qid               string    `json:"qid"`
	QuestionTitle     string    `json:"question_title"`
	Aid               string    `json:"aid"`
	AnswerHead        string    `json:"answer_head"`
	NewAnswerCount    int64     `json:"new_answer_count"`
	NewLikeCount      int64     `json:"new_like_count"`
	NewCommentCount   int64     `json:"new_comment_count"`
	NewCriticismCount int64     `json:"new_criticism_count"`
	NewFollowerCount  int64     `json:"new_follower_count"`
}

type ResultOAuthGithub struct {
	Type         int8   `json:"type"`
	First        bool   `json:"first"`
	Role         int8   `json:"role"`
	Uid          string `json:"uid"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResultPasswd struct {
	Type int8 `json:"type"`
}

type ResultPublicInfoGet struct {
	Name           string   `json:"name"`
	Nickname       string   `json:"nickname"`
	Profile        string   `json:"profile"`
	Icon           string   `json:"icon"`
	Level          int64    `json:"level"`
	Gender         int8     `json:"gender"`
	Email          string   `json:"email"`
	AccountType    int8     `json:"account_type"`
	Labels         []string `json:"labels"`
	QuestionCount  int64    `json:"question_count"`
	AnswerCount    int64    `json:"answer_count"`
	FollowerCount  int64    `json:"follower_count"`
	FollowingCount int64    `json:"following_count"`
	LikeCount      int64    `json:"like_count"`
	ApprovalCount  int64    `json:"approval_count"`
}

type ResultPublicInfoPut struct {
	Type int8 `json:"type"`
}

type ResultRefreshToken struct {
	Type         int8   `json:"type"`
	Role         int8   `json:"role"`
	Uid          string `json:"uid"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResultRegister struct {
	Type int8 `json:"type"`
}

type ResultUserQuestions struct {
	Qid           string    `json:"qid"`
	Title         string    `json:"title"`
	Time          time.Time `json:"time"`
	AnswerCount   int64     `json:"answer_count"`
	ViewCount     int64     `json:"view_count"`
	FavoriteCount int64     `json:"favorite_count"`
	Category      string    `json:"category"`
	Labels        []string  `json:"labels"`
	Head          string    `json:"head"`
	PictureUrls   []string  `json:"picture_urls"`
}

type ResultVerificationCode struct {
	Type int8 `json:"type"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	emailPassword = os.Getenv("EMAIL_PASSWORD")
	emailUsername = os.Getenv("EMAIL_USERNAME")
	oAuthGithubId = os.Getenv("OAUTH_GITHUB_ID")
	oAuthGithubSecret = os.Getenv("OAUTH_GITHUB_SECRET")
}

func (u *UsersServiceImpl) encryptPassword(password string, salt string) (hashPassword string) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password+salt)))
}

func (u *UsersServiceImpl) generateCode() (code int64) {
	return rand.Int63n(9e5) + 1e5
}

func (u *UsersServiceImpl) generateSalt() (salt string) {
	b := make([]byte, 16)
	for i := range b {
		b[i] = byte(rand.Uint32() & 0xFF)
	}
	return fmt.Sprintf("%x", b)
}

func (u *UsersServiceImpl) Init(usersDao ...dao.UsersDao) (err error) {
	if len(usersDao) == 0 {
		usersDao = append(usersDao, &dao.UsersDaoImpl{})
	}
	u.usersDao = usersDao[0]
	return u.usersDao.Init()
}

func (u *UsersServiceImpl) Destruct() {
	u.usersDao.Destruct()
}

func (u *UsersServiceImpl) CheckToken(token string) (res ResCheckToken, err error) {
	res.Successful, res.Uid, res.Role, err = util.ParseToken(token)
	return res, err
}

func (u *UsersServiceImpl) Follow(token string, uid int64, follow bool) (res ResFollow, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	var successful bool
	successful, user.Uid, user.Role, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	if uid == user.Uid {
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if follow {
		var follow entity.Follows
		follow.Uid = uid
		follow.Follower = user.Uid
		follow.Time = time.Now().Unix()
		err = u.usersDao.InsertFollow(ctx, follow)
	} else {
		_, err = u.usersDao.FindFollowByUidAndFollower(ctx, uid, user.Uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		err = u.usersDao.RemoveFollowByUidAndFollower(ctx, uid, user.Uid)
	}
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	user, err = u.usersDao.FindUserByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if follow {
		user.FollowingCount++
	} else {
		user.FollowingCount--
	}
	err = u.usersDao.UpdateUserByUid(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	user, err = u.usersDao.FindUserByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if follow {
		user.FollowerCount++
	} else {
		user.FollowerCount--
	}
	err = u.usersDao.UpdateUserByUid(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Followed(token string, uid int64) (res ResFollowed, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var successful bool
	successful, _, _, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var follows []entity.Follows
	follows, err = u.usersDao.FindFollowsByFollower(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultFollowed{}
	for _, follow := range follows {
		var user entity.Users
		user, err = u.usersDao.FindUserByUid(ctx, follow.Uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var userDetail entity.UserDetails
		userDetail, err = u.usersDao.FindUserDetailByUid(ctx, follow.Uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		res.Result = append(res.Result, ResultFollowed{userDetail.Icon, user.Name, user.Nickname, user.Profile})
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Followers(token string, uid int64) (res ResFollowers, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var successful bool
	successful, _, _, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var follows []entity.Follows
	follows, err = u.usersDao.FindFollowsByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultFollowers{}
	for _, follow := range follows {
		var user entity.Users
		user, err = u.usersDao.FindUserByUid(ctx, follow.Follower)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var userDetail entity.UserDetails
		userDetail, err = u.usersDao.FindUserDetailByUid(ctx, follow.Follower)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		res.Result = append(res.Result, ResultFollowers{userDetail.Icon, user.Name, user.Nickname, user.Profile})
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) InfoList(req ReqInfoList) (res ResInfoList, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultInfoList{}
	for _, uid := range req.Uids {
		var user entity.Users
		user, err = u.usersDao.FindUserByUid(ctx, uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var userDetail entity.UserDetails
		userDetail, err = u.usersDao.FindUserDetailByUid(ctx, uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		res.Result = append(res.Result, ResultInfoList{user.Name, user.Nickname, userDetail.Icon})
	}
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Login(req ReqLogin) (res ResLogin, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	user, err = u.usersDao.FindUserByName(ctx, req.Name)
	if err != nil || u.encryptPassword(req.Password, user.Salt) != user.HashPassword {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if user.Role == entity.DISABLE {
		res.Code = 1
		res.Result.Type = 0
		return res, u.usersDao.Rollback(&ctx)
	}
	if user.Role == entity.NOT_ACTIVE {
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var userDetail entity.UserDetails
	userDetail, err = u.usersDao.FindUserDetailByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	var token string
	token, err = util.SignToken(user.Uid, user.Role, false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	var refreshToken string
	refreshToken, err = util.SignToken(user.Uid, user.Role, true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	res.Result.Role = user.Role
	res.Result.Uid = strconv.FormatInt(user.Uid, 10)
	res.Result.Icon = userDetail.Icon
	res.Result.Name = user.Name
	res.Result.Nickname = user.Nickname
	res.Result.Token = token
	res.Result.RefreshToken = refreshToken
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Notifications(token string, page int64) (res ResNotifications, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	var successful bool
	successful, user.Uid, user.Role, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) OAuthGithub(code string, error string) (res ResOAuthGithub, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	if error == "access_denied" {
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var request *http.Request
	request, err = http.NewRequest("POST", "https://github.com/login/oauth/access_token?client_id="+oAuthGithubId+"&client_secret="+oAuthGithubSecret+"&code="+code, nil)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	request.Header.Set("Accept", "application/json")
	client := http.Client{}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
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
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	request, err = http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "token "+responseBodyToken.AccessToken)
	response, err = client.Do(request)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
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
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	user, err = u.usersDao.FindUserByOidAndAccountType(ctx, strconv.FormatInt(responseBodyInfo.Id, 10), entity.GITHUB)
	if err == nil {
		if user.Role == entity.DISABLE {
			res.Code = 1
			res.Result.Type = 0
			return res, u.usersDao.Rollback(&ctx)
		}
		var token string
		token, err = util.SignToken(user.Uid, user.Role, false)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 2
			return res, u.usersDao.Rollback(&ctx)
		}
		var refreshToken string
		refreshToken, err = util.SignToken(user.Uid, user.Role, true)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 2
			return res, u.usersDao.Rollback(&ctx)
		}
		res.Code = 0
		res.Result.First = false
		res.Result.Role = user.Role
		res.Result.Uid = strconv.FormatInt(user.Uid, 10)
		res.Result.Token = token
		res.Result.RefreshToken = refreshToken
		return res, u.usersDao.Commit(&ctx)
	}
	user = entity.Users{Oid: strconv.FormatInt(responseBodyInfo.Id, 10), Profile: "", Role: entity.USER, ActiveCode: 0, PasswdCode: 0, AccountType: entity.GITHUB, Exp: 0, FollowerCount: 0, FollowingCount: 0, QuestionCount: 0, AnswerCount: 0, LikeCount: 0, ApprovalCount: 0, NotificationTime: time.Now().Unix()}
	user.Uid, err = u.usersDao.InsertUser(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var userDetail entity.UserDetails
	userDetail.Uid = user.Uid
	userDetail.Icon = ""
	err = u.usersDao.InsertUserDetail(ctx, userDetail)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var favorite entity.Favorites
	favorite.Uid = user.Uid
	favorite.Title = "Default"
	favorite.Fid, err = u.usersDao.InsertFavorite(ctx, favorite)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	res.Result.First = true
	res.Result.Role = user.Role
	res.Result.Uid = strconv.FormatInt(user.Uid, 10)
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Passwd(token string, req ReqPasswd) (res ResPasswd, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	var successful bool
	successful, user.Uid, user.Role, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	user, err = u.usersDao.FindUserByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if u.encryptPassword(req.Old, user.Salt) != user.HashPassword {
		res.Code = 1
		res.Result.Type = 0
		return res, u.usersDao.Rollback(&ctx)
	}
	user.HashPassword = u.encryptPassword(req.New, user.Salt)
	err = u.usersDao.UpdateUserByUid(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) PublicInfoGet(token string, uid int64) (res ResPublicInfoGet, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var successful bool
	successful, _, _, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	user, err = u.usersDao.FindUserByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var userDetail entity.UserDetails
	userDetail, err = u.usersDao.FindUserDetailByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var labels []entity.Labels
	labels, err = u.usersDao.FindLabelsByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	res.Result.Name = user.Name
	res.Result.Nickname = user.Nickname
	res.Result.Profile = user.Profile
	res.Result.Icon = userDetail.Icon
	res.Result.Level = user.Exp
	res.Result.Gender = user.Gender
	res.Result.Email = user.Email
	res.Result.AccountType = user.AccountType
	res.Result.Labels = []string{}
	for _, label := range labels {
		res.Result.Labels = append(res.Result.Labels, label.Title)
	}
	res.Result.QuestionCount = user.QuestionCount
	res.Result.AnswerCount = user.AnswerCount
	res.Result.FollowerCount = user.FollowerCount
	res.Result.FollowingCount = user.FollowingCount
	res.Result.LikeCount = user.LikeCount
	res.Result.ApprovalCount = user.ApprovalCount
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) PublicInfoPut(token string, req ReqPublicInfoPut) (res ResPublicInfoPut, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	var successful bool
	successful, user.Uid, user.Role, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	user, err = u.usersDao.FindUserByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var userByName entity.Users
	userByName, err = u.usersDao.FindUserByName(ctx, req.Name)
	if err == nil && user.Uid != userByName.Uid {
		res.Code = 1
		res.Result.Type = 0
		return res, u.usersDao.Rollback(&ctx)
	}
	user.Name = req.Name
	user.Nickname = req.Nickname
	user.Profile = req.Profile
	user.Gender = req.Gender
	user.Email = req.Email
	err = u.usersDao.UpdateUserByUid(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var userDetail entity.UserDetails
	userDetail, err = u.usersDao.FindUserDetailByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	userDetail.Icon = req.Icon
	err = u.usersDao.UpdateUserDetailByUid(ctx, userDetail)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	err = u.usersDao.RemoveUserLabelsByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	for _, labelTitle := range req.Labels {
		var label entity.Labels
		label, err = u.usersDao.FindLabelByTitle(ctx, labelTitle)
		if err != nil {
			label = entity.Labels{Title: labelTitle}
			label.Lid, err = u.usersDao.InsertLabel(ctx, label)
			if err != nil {
				log.Info(err)
				res.Code = 1
				res.Result.Type = 1
				return res, u.usersDao.Rollback(&ctx)
			}
		}
		err = u.usersDao.InsertUserLabel(ctx, entity.UserLabels{Uid: user.Uid, Lid: label.Lid})
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) RefreshToken(req ReqRefreshToken) (res ResRefreshToken, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var uid int64
	var successful bool
	successful, uid, _, err = util.ParseToken(req.Refresh)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	user, err = u.usersDao.FindUserByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if user.Role == entity.DISABLE {
		res.Code = 1
		res.Result.Type = 0
		return res, u.usersDao.Rollback(&ctx)
	}
	var token string
	token, err = util.SignToken(user.Uid, user.Role, false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var refreshToken string
	refreshToken, err = util.SignToken(user.Uid, user.Role, true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	res.Result.Role = user.Role
	res.Result.Uid = strconv.FormatInt(user.Uid, 10)
	res.Result.Token = token
	res.Result.RefreshToken = refreshToken
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Register(req ReqRegister) (res ResRegister, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	_, err = u.usersDao.FindUserByName(ctx, req.Name)
	if err == nil {
		res.Code = 1
		res.Result.Type = 0
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	user, err = u.usersDao.FindUserByEmail(ctx, req.Email)
	if err != nil || user.ActiveCode > 0 {
		if err != nil {
			log.Info(err)
		}
		res.Code = 1
		res.Result.Type = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	if user.Role != entity.NOT_ACTIVE {
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	user.Name = req.Name
	user.Nickname = req.Nickname
	user.Salt = u.generateSalt()
	user.HashPassword = u.encryptPassword(req.Password, user.Salt)
	user.Gender = req.Gender
	user.Role = entity.USER
	err = u.usersDao.UpdateUserByUid(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	var userDetail entity.UserDetails
	userDetail.Uid = user.Uid
	userDetail.Icon = req.Icon
	err = u.usersDao.InsertUserDetail(ctx, userDetail)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	var favorite entity.Favorites
	favorite.Uid = user.Uid
	favorite.Title = "Default"
	favorite.Fid, err = u.usersDao.InsertFavorite(ctx, favorite)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) UserQuestions(token string, uid int64, page int64) (res ResUserQuestions, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var successful bool
	successful, _, _, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var questions []entity.Questions
	questions, err = u.usersDao.FindQuestionsByRaiserOrderByTimeDescPageable(ctx, uid, dao.Pageable{Number: page, Size: 10})
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultUserQuestions{}
	for _, question := range questions {
		var questionDetail entity.QuestionDetails
		questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, question.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var labels []entity.Labels
		labels, err = u.usersDao.FindLabelsByQid(ctx, question.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var labelTitles []string
		for _, label := range labels {
			labelTitles = append(labelTitles, label.Title)
		}
		res.Result = append(res.Result, ResultUserQuestions{strconv.FormatInt(question.Qid, 10), questionDetail.Title, time.Unix(question.Time, 0), question.AnswerCount, question.ViewCount, question.FavoriteCount, question.Category, labelTitles, questionDetail.Content, nil})
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) VerificationCode(register bool, email string) (res ResVerificationCode, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var code int64
	if register {
		_, err = u.usersDao.FindUserByEmail(ctx, email)
		if err == nil {
			res.Code = 1
			res.Result.Type = 0
			return res, u.usersDao.Rollback(&ctx)
		}
		var user entity.Users
		user.Email = email
		user.Profile = ""
		user.Role = entity.NOT_ACTIVE
		user.AccountType = entity.SOFIA
		user.ActiveCode = u.generateCode()
		user.PasswdCode = 0
		user.Exp = 0
		user.FollowerCount = 0
		user.FollowingCount = 0
		user.QuestionCount = 0
		user.AnswerCount = 0
		user.LikeCount = 0
		user.ApprovalCount = 0
		user.NotificationTime = time.Now().Unix()
		_, err = u.usersDao.InsertUser(ctx, user)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		code = user.ActiveCode
	} else {
		var user entity.Users
		user, err = u.usersDao.FindUserByEmail(ctx, email)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 0
			return res, u.usersDao.Rollback(&ctx)
		}
		user.PasswdCode = u.generateCode()
		err = u.usersDao.UpdateUserByUid(ctx, user)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Type = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		code = user.PasswdCode
	}
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(emailUsername, "Sofia"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Sofia")
	message.SetBody("text/html", strconv.FormatInt(code, 10))
	err = gomail.NewDialer("smtp.qq.com", 587, emailUsername, emailPassword).DialAndSend(message)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Verify(email string, code int64) (res ResVerify, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(false)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var user entity.Users
	user, err = u.usersDao.FindUserByEmail(ctx, email)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if user.ActiveCode > 0 {
		if code != user.ActiveCode {
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		user.ActiveCode = 0
	}
	if user.PasswdCode > 0 {
		if code != user.PasswdCode {
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		user.Role = entity.NOT_ACTIVE
		user.PasswdCode = 0
	}
	err = u.usersDao.UpdateUserByUid(ctx, user)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}
