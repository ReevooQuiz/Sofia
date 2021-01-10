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
	"strings"
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

type ReqApprove struct {
	Aid     string `json:"aid"`
	Approve bool   `json:"approve"`
}

type ReqBan struct {
	Uid string `json:"uid"`
	Ban bool   `json:"ban"`
}

type ReqFavorite struct {
	Qid      string `json:"qid"`
	Favorite bool   `json:"favorite"`
}

type ReqFollow struct {
	Uid    string `json:"uid"`
	Follow bool   `json:"follow"`
}

type ReqInfoList struct {
	Uids []int64 `json:"uids"`
}

type ReqLike struct {
	Aid  string `json:"aid"`
	Like bool   `json:"like"`
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
	Gender   string   `json:"gender"`
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
	Gender   string `json:"gender"`
}

type ReqWordBan struct {
	Word string `json:"word"`
	Ban  bool   `json:"ban"`
}

type ResApprove struct {
	Code int8 `json:"code"`
}

type ResBan struct {
	Code int8 `json:"code"`
}

type ResBanned struct {
	Code   int8           `json:"code"`
	Result []ResultBanned `json:"result"`
}

type ResCheckSession struct {
	Code int8 `json:"code"`
}

type ResCheckToken struct {
	Successful bool  `json:"successful"`
	Uid        int64 `json:"uid"`
	Role       int8  `json:"role"`
}

type ResCollection struct {
	Code   int8               `json:"code"`
	Result []ResultCollection `json:"result"`
}

type ResFavorite struct {
	Code int8 `json:"code"`
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

type ResLike struct {
	Code int8 `json:"code"`
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

type ResUserAnswers struct {
	Code   int8                `json:"code"`
	Result []ResultUserAnswers `json:"result"`
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

type ResWordBan struct {
	Code int8 `json:"code"`
}

type ResWordsBanned struct {
	Code   int8     `json:"code"`
	Result []string `json:"result"`
}

type ResultBanned struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type ResultCollection struct {
	Qid           string                 `json:"qid"`
	QuestionTitle string                 `json:"question_title"`
	Raiser        ResultCollectionRaiser `json:"raiser"`
	QuestionHead  string                 `json:"question_head"`
}

type ResultCollectionRaiser struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type ResultFollowed struct {
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}

type ResultFollowers struct {
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Uid      string `json:"uid"`
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
	NewApprovalCount  int64     `json:"new_approval_count"`
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
	Follow         bool     `json:"follow"`
}

type ResultPublicInfoPut struct {
	Type int8 `json:"type"`
}

type ResultRefreshToken struct {
	Type         int8   `json:"type"`
	Role         int8   `json:"role"`
	Uid          string `json:"uid"`
	Icon         string `json:"icon"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResultRegister struct {
	Type int8 `json:"type"`
}

type ResultUserAnswers struct {
	Question ResultUserAnswersQuestion `json:"question"`
	Answer   ResultUserAnswersAnswer   `json:"answer"`
}

type ResultUserAnswersAnswer struct {
	Aid            string    `json:"aid"`
	LikeCount      int64     `json:"like_count"`
	CriticismCount int64     `json:"criticism_count"`
	ApprovalCount  int64     `json:"approval_count"`
	CommentCount   int64     `json:"comment_count"`
	Head           string    `json:"head"`
	Time           time.Time `json:"time"`
	PictureUrls    []string  `json:"picture_urls"`
	Liked          bool      `json:"liked"`
	Approved       bool      `json:"approved"`
	Approvable     bool      `json:"approvable"`
}

type ResultUserAnswersQuestion struct {
	Qid      string   `json:"qid"`
	Title    string   `json:"title"`
	Category string   `json:"category"`
	Labels   []string `json:"labels"`
	Head     string   `json:"head"`
}

type ResultUserQuestions struct {
	Qid           string    `json:"qid"`
	HasKeywords   bool      `json:"has_keywords"`
	Closed        bool      `json:"closed"`
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

func (u *UsersServiceImpl) Approve(token string, req ReqApprove) (res ResApprove, err error) {
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
	var aid int64
	aid, err = strconv.ParseInt(req.Aid, 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var approveAnswer entity.ApproveAnswers
	if req.Approve {
		approveAnswer.Uid = user.Uid
		approveAnswer.Aid = aid
		approveAnswer.Time = time.Now().Unix()
		err = u.usersDao.InsertApproveAnswer(ctx, approveAnswer)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	} else {
		approveAnswer, err = u.usersDao.FindApproveAnswerByUidAndAid(ctx, user.Uid, aid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		err = u.usersDao.RemoveApproveAnswerByUidAndAid(ctx, user.Uid, aid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	}
	var answer entity.Answers
	answer, err = u.usersDao.FindAnswerByAid(ctx, aid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if req.Approve {
		answer.ApprovalCount++
	} else {
		answer.ApprovalCount--
	}
	err = u.usersDao.UpdateAnswerByAid(ctx, answer)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Ban(token string, req ReqBan) (res ResBan, err error) {
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
	if user.Role != entity.ADMIN {
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var uid int64
	uid, err = strconv.ParseInt(req.Uid, 10, 64)
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
	if req.Ban {
		user.Role = entity.DISABLE
	} else {
		user.Role = entity.USER
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

func (u *UsersServiceImpl) Banned(token string, page int64) (res ResBanned, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
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
	if user.Role != entity.ADMIN {
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var users []entity.Users
	users, err = u.usersDao.FindUsersByRolePageable(ctx, entity.DISABLE, dao.Pageable{Number: page, Size: 10})
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultBanned{}
	for _, user = range users {
		var userDetail entity.UserDetails
		userDetail, err = u.usersDao.FindUserDetailByUid(ctx, user.Uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		res.Result = append(res.Result, ResultBanned{strconv.FormatInt(user.Uid, 10), user.Name, user.Nickname, userDetail.Icon})
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) CheckSession(token string) (res ResCheckSession, err error) {
	var successful bool
	successful, _, _, err = util.ParseToken(token)
	if err != nil || !successful {
		res.Code = 2
	} else {
		res.Code = 0
	}
	return res, err
}

func (u *UsersServiceImpl) CheckToken(token string) (res ResCheckToken, err error) {
	res.Successful, res.Uid, res.Role, err = util.ParseToken(token)
	return res, err
}

func (u *UsersServiceImpl) Collection(token string, page int64) (res ResCollection, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var uid int64
	var successful bool
	successful, uid, _, err = util.ParseToken(token)
	if err != nil || !successful {
		if err != nil {
			log.Info(err)
		}
		res.Code = 2
		return res, u.usersDao.Rollback(&ctx)
	}
	var favorite entity.Favorites
	favorite, err = u.usersDao.FindFavoriteByUidAndTitle(ctx, uid, "Default")
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var favoriteItems []entity.FavoriteItems
	favoriteItems, err = u.usersDao.FindFavoriteItemsByFidPageable(ctx, favorite.Fid, dao.Pageable{Number: page, Size: 10})
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultCollection{}
	for _, favoriteItem := range favoriteItems {
		var question entity.Questions
		question, err = u.usersDao.FindQuestionByQid(ctx, favoriteItem.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var questionDetail entity.QuestionDetails
		questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, favoriteItem.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var user entity.Users
		user, err = u.usersDao.FindUserByUid(ctx, question.Raiser)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var userDetail entity.UserDetails
		userDetail, err = u.usersDao.FindUserDetailByUid(ctx, user.Uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		res.Result = append(res.Result, ResultCollection{strconv.FormatInt(favoriteItem.Qid, 10), questionDetail.Title, ResultCollectionRaiser{strconv.FormatInt(user.Uid, 10), user.Name, user.Nickname, userDetail.Icon}, questionDetail.Head})
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Favorite(token string, req ReqFavorite) (res ResFavorite, err error) {
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
	var qid int64
	qid, err = strconv.ParseInt(req.Qid, 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var favorite entity.Favorites
	favorite, err = u.usersDao.FindFavoriteByUidAndTitle(ctx, user.Uid, "Default")
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var favoriteItem entity.FavoriteItems
	if req.Favorite {
		favoriteItem.Fid = favorite.Fid
		favoriteItem.Qid = qid
		err = u.usersDao.InsertFavoriteItem(ctx, favoriteItem)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	} else {
		favoriteItem, err = u.usersDao.FindFavoriteItemByFidAndQid(ctx, favorite.Fid, qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		err = u.usersDao.RemoveFavoriteItemByFidAndQid(ctx, favorite.Fid, qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	}
	var question entity.Questions
	question, err = u.usersDao.FindQuestionByQid(ctx, qid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if req.Favorite {
		question.FavoriteCount++
	} else {
		question.FavoriteCount--
	}
	err = u.usersDao.UpdateQuestionByQid(ctx, question)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) Follow(token string, req ReqFollow) (res ResFollow, err error) {
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
	var uid int64
	uid, err = strconv.ParseInt(req.Uid, 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if uid == user.Uid {
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if req.Follow {
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
	if req.Follow {
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
	if req.Follow {
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
		res.Result = append(res.Result, ResultFollowed{userDetail.Icon, user.Name, strconv.FormatInt(user.Uid, 10), user.Nickname, user.Profile})
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
		res.Result = append(res.Result, ResultFollowers{userDetail.Icon, user.Name, strconv.FormatInt(user.Uid, 10), user.Nickname, user.Profile})
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

func (u *UsersServiceImpl) Like(token string, req ReqLike) (res ResLike, err error) {
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
	var aid int64
	aid, err = strconv.ParseInt(req.Aid, 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var likeAnswer entity.LikeAnswers
	if req.Like {
		likeAnswer.Uid = user.Uid
		likeAnswer.Aid = aid
		likeAnswer.Time = time.Now().Unix()
		err = u.usersDao.InsertLikeAnswer(ctx, likeAnswer)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	} else {
		likeAnswer, err = u.usersDao.FindLikeAnswerByUidAndAid(ctx, user.Uid, aid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		err = u.usersDao.RemoveLikeAnswerByUidAndAid(ctx, user.Uid, aid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	}
	var answer entity.Answers
	answer, err = u.usersDao.FindAnswerByAid(ctx, aid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if req.Like {
		answer.LikeCount++
	} else {
		answer.LikeCount--
	}
	err = u.usersDao.UpdateAnswerByAid(ctx, answer)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
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
	user, err = u.usersDao.FindUserByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var notifications []dao.Notifications
	notifications, err = u.usersDao.FindNotificationsByUidPageable(ctx, user.Uid, dao.Pageable{Number: page, Size: 10})
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultNotifications{}
	for _, notification := range notifications {
		switch notification.Type {
		case 0:
			var answer entity.Answers
			answer, err = u.usersDao.FindAnswerByAid(ctx, notification.Id0)
			if err != nil {
				log.Info(err)
				res.Code = 1
				return res, u.usersDao.Rollback(&ctx)
			}
			if answer.Time < user.NotificationTime {
				user.NotificationTime = time.Now().Unix()
				err = u.usersDao.UpdateUserByUid(ctx, user)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Code = 0
				return res, u.usersDao.Commit(&ctx)
			}
			flag := true
			for i, result := range res.Result {
				if result.Type == 0 {
					var qid int64
					qid, err = strconv.ParseInt(result.Qid, 10, 64)
					if err != nil {
						log.Info(err)
						res.Code = 1
						return res, u.usersDao.Rollback(&ctx)
					}
					if qid == answer.Qid {
						res.Result[i].NewAnswerCount++
						flag = false
						break
					}
				}
			}
			if flag {
				var questionDetail entity.QuestionDetails
				questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, answer.Qid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Result = append(res.Result, ResultNotifications{Type: 0, Time: time.Unix(answer.Time, 0), Qid: strconv.FormatInt(answer.Qid, 10), QuestionTitle: questionDetail.Title, NewAnswerCount: 1})
			}
		case 1:
			var likeAnswer entity.LikeAnswers
			likeAnswer, err = u.usersDao.FindLikeAnswerByUidAndAid(ctx, notification.Id0, notification.Id1)
			if err != nil {
				log.Info(err)
				res.Code = 1
				return res, u.usersDao.Rollback(&ctx)
			}
			if likeAnswer.Time < user.NotificationTime {
				user.NotificationTime = time.Now().Unix()
				err = u.usersDao.UpdateUserByUid(ctx, user)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Code = 0
				return res, u.usersDao.Commit(&ctx)
			}
			flag := true
			for i, result := range res.Result {
				if result.Type == 1 {
					var aid int64
					aid, err = strconv.ParseInt(result.Aid, 10, 64)
					if err != nil {
						log.Info(err)
						res.Code = 1
						return res, u.usersDao.Rollback(&ctx)
					}
					if aid == likeAnswer.Aid {
						res.Result[i].NewLikeCount++
						flag = false
						break
					}
				}
			}
			if flag {
				var answer entity.Answers
				answer, err = u.usersDao.FindAnswerByAid(ctx, likeAnswer.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var questionDetail entity.QuestionDetails
				questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, answer.Qid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var answerDetail entity.AnswerDetails
				answerDetail, err = u.usersDao.FindAnswerDetailByAid(ctx, likeAnswer.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Result = append(res.Result, ResultNotifications{Type: 1, Time: time.Unix(likeAnswer.Time, 0), Qid: strconv.FormatInt(answer.Qid, 10), QuestionTitle: questionDetail.Title, Aid: strconv.FormatInt(likeAnswer.Aid, 10), AnswerHead: answerDetail.Head, NewLikeCount: 1})
			}
		case 2:
			var approveAnswer entity.ApproveAnswers
			approveAnswer, err = u.usersDao.FindApproveAnswerByUidAndAid(ctx, notification.Id0, notification.Id1)
			if err != nil {
				log.Info(err)
				res.Code = 1
				return res, u.usersDao.Rollback(&ctx)
			}
			if approveAnswer.Time < user.NotificationTime {
				user.NotificationTime = time.Now().Unix()
				err = u.usersDao.UpdateUserByUid(ctx, user)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Code = 0
				return res, u.usersDao.Commit(&ctx)
			}
			flag := true
			for i, result := range res.Result {
				if result.Type == 2 {
					var aid int64
					aid, err = strconv.ParseInt(result.Aid, 10, 64)
					if err != nil {
						log.Info(err)
						res.Code = 1
						return res, u.usersDao.Rollback(&ctx)
					}
					if aid == approveAnswer.Aid {
						res.Result[i].NewApprovalCount++
						flag = false
						break
					}
				}
			}
			if flag {
				var answer entity.Answers
				answer, err = u.usersDao.FindAnswerByAid(ctx, approveAnswer.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var questionDetail entity.QuestionDetails
				questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, answer.Qid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var answerDetail entity.AnswerDetails
				answerDetail, err = u.usersDao.FindAnswerDetailByAid(ctx, approveAnswer.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Result = append(res.Result, ResultNotifications{Type: 2, Time: time.Unix(approveAnswer.Time, 0), Qid: strconv.FormatInt(answer.Qid, 10), QuestionTitle: questionDetail.Title, Aid: strconv.FormatInt(approveAnswer.Aid, 10), AnswerHead: answerDetail.Head, NewApprovalCount: 1})
			}
		case 3:
			var comment entity.Comments
			comment, err = u.usersDao.FindCommentByCmid(ctx, notification.Id0)
			if err != nil {
				log.Info(err)
				res.Code = 1
				return res, u.usersDao.Rollback(&ctx)
			}
			if comment.Time < user.NotificationTime {
				user.NotificationTime = time.Now().Unix()
				err = u.usersDao.UpdateUserByUid(ctx, user)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Code = 0
				return res, u.usersDao.Commit(&ctx)
			}
			flag := true
			for i, result := range res.Result {
				if result.Type == 3 {
					var aid int64
					aid, err = strconv.ParseInt(result.Aid, 10, 64)
					if err != nil {
						log.Info(err)
						res.Code = 1
						return res, u.usersDao.Rollback(&ctx)
					}
					if aid == comment.Aid {
						res.Result[i].NewCommentCount++
						flag = false
						break
					}
				}
			}
			if flag {
				var answer entity.Answers
				answer, err = u.usersDao.FindAnswerByAid(ctx, comment.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var questionDetail entity.QuestionDetails
				questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, answer.Qid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var answerDetail entity.AnswerDetails
				answerDetail, err = u.usersDao.FindAnswerDetailByAid(ctx, comment.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Result = append(res.Result, ResultNotifications{Type: 3, Time: time.Unix(comment.Time, 0), Qid: strconv.FormatInt(answer.Qid, 10), QuestionTitle: questionDetail.Title, Aid: strconv.FormatInt(comment.Aid, 10), AnswerHead: answerDetail.Head, NewCommentCount: 1})
			}
		case 4:
			var criticism entity.Criticisms
			criticism, err = u.usersDao.FindCriticismByCtid(ctx, notification.Id0)
			if err != nil {
				log.Info(err)
				res.Code = 1
				return res, u.usersDao.Rollback(&ctx)
			}
			if criticism.Time < user.NotificationTime {
				user.NotificationTime = time.Now().Unix()
				err = u.usersDao.UpdateUserByUid(ctx, user)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Code = 0
				return res, u.usersDao.Commit(&ctx)
			}
			flag := true
			for i, result := range res.Result {
				if result.Type == 4 {
					var aid int64
					aid, err = strconv.ParseInt(result.Aid, 10, 64)
					if err != nil {
						log.Info(err)
						res.Code = 1
						return res, u.usersDao.Rollback(&ctx)
					}
					if aid == criticism.Aid {
						res.Result[i].NewCriticismCount++
						flag = false
						break
					}
				}
			}
			if flag {
				var answer entity.Answers
				answer, err = u.usersDao.FindAnswerByAid(ctx, criticism.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var questionDetail entity.QuestionDetails
				questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, answer.Qid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				var answerDetail entity.AnswerDetails
				answerDetail, err = u.usersDao.FindAnswerDetailByAid(ctx, criticism.Aid)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Result = append(res.Result, ResultNotifications{Type: 4, Time: time.Unix(criticism.Time, 0), Qid: strconv.FormatInt(answer.Qid, 10), QuestionTitle: questionDetail.Title, Aid: strconv.FormatInt(criticism.Aid, 10), AnswerHead: answerDetail.Head, NewCriticismCount: 1})
			}
		case 5:
			var follow entity.Follows
			follow, err = u.usersDao.FindFollowByUidAndFollower(ctx, user.Uid, notification.Id0)
			if err != nil {
				log.Info(err)
				res.Code = 1
				return res, u.usersDao.Rollback(&ctx)
			}
			if follow.Time < user.NotificationTime {
				user.NotificationTime = time.Now().Unix()
				err = u.usersDao.UpdateUserByUid(ctx, user)
				if err != nil {
					log.Info(err)
					res.Code = 1
					return res, u.usersDao.Rollback(&ctx)
				}
				res.Code = 0
				return res, u.usersDao.Commit(&ctx)
			}
			flag := true
			for i, result := range res.Result {
				if result.Type == 5 {
					res.Result[i].NewFollowerCount++
					flag = false
					break
				}
			}
			if flag {
				res.Result = append(res.Result, ResultNotifications{Type: 5, Time: time.Unix(follow.Time, 0), NewFollowerCount: 1})
			}
		default:
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
	}
	if len(notifications) < 10 {
		user.NotificationTime = time.Now().Unix()
		err = u.usersDao.UpdateUserByUid(ctx, user)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
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
	var myUid int64
	var successful bool
	successful, myUid, _, err = util.ParseToken(token)
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
	_, err = u.usersDao.FindFollowByUidAndFollower(ctx, user.Uid, myUid)
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
	res.Result.Follow = err == nil
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
	var gender int64
	gender, err = strconv.ParseInt(req.Gender, 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	user.Name = req.Name
	user.Nickname = req.Nickname
	user.Profile = req.Profile
	user.Gender = int8(gender)
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
	var userDetail entity.UserDetails
	userDetail, err = u.usersDao.FindUserDetailByUid(ctx, uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 1
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
	res.Result.Icon = userDetail.Icon
	res.Result.Name = user.Name
	res.Result.Nickname = user.Nickname
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
	var gender int64
	gender, err = strconv.ParseInt(req.Gender, 10, 64)
	if err != nil {
		log.Info(err)
		res.Code = 1
		res.Result.Type = 3
		return res, u.usersDao.Rollback(&ctx)
	}
	user.Name = req.Name
	user.Nickname = req.Nickname
	user.Salt = u.generateSalt()
	user.HashPassword = u.encryptPassword(req.Password, user.Salt)
	user.Gender = int8(gender)
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

func (u *UsersServiceImpl) UserAnswers(token string, uid int64, page int64) (res ResUserAnswers, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
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
	user, err = u.usersDao.FindUserByUid(ctx, user.Uid)
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var answers []entity.Answers
	answers, err = u.usersDao.FindAnswersByAnswererOrderByTimeDescPageable(ctx, uid, dao.Pageable{Number: page, Size: 10})
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Result = []ResultUserAnswers{}
	for _, answer := range answers {
		var answerDetail entity.AnswerDetails
		answerDetail, err = u.usersDao.FindAnswerDetailByAid(ctx, answer.Aid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		_, err = u.usersDao.FindLikeAnswerByUidAndAid(ctx, user.Uid, answer.Aid)
		liked := err == nil
		_, err = u.usersDao.FindApproveAnswerByUidAndAid(ctx, user.Uid, answer.Aid)
		approved := err == nil
		var question entity.Questions
		question, err = u.usersDao.FindQuestionByQid(ctx, answer.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var questionDetail entity.QuestionDetails
		questionDetail, err = u.usersDao.FindQuestionDetailByQid(ctx, answer.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var userLabels []entity.Labels
		userLabels, err = u.usersDao.FindLabelsByUid(ctx, user.Uid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var questionLabels []entity.Labels
		questionLabels, err = u.usersDao.FindLabelsByQid(ctx, answer.Qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			return res, u.usersDao.Rollback(&ctx)
		}
		var questionLabelTitles []string
		approvable := true
		for _, questionLabel := range questionLabels {
			questionLabelTitles = append(questionLabelTitles, questionLabel.Title)
			flag := false
			for _, userLabel := range userLabels {
				if questionLabel.Lid == userLabel.Lid {
					flag = true
					break
				}
			}
			approvable = approvable && flag
		}
		res.Result = append(res.Result, ResultUserAnswers{ResultUserAnswersQuestion{strconv.FormatInt(answer.Qid, 10), questionDetail.Title, question.Category, questionLabelTitles, fmt.Sprintf("%.20s", questionDetail.Content)}, ResultUserAnswersAnswer{strconv.FormatInt(answer.Aid, 10), answer.LikeCount, answer.CriticismCount, answer.ApprovalCount, answer.CommentCount, fmt.Sprintf("%.20s", answerDetail.Content), time.Unix(answer.Time, 0), []string{answerDetail.PictureUrl}, liked, approved, approvable}})
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
	var banWords []entity.BanWords
	banWords, err = u.usersDao.FindBanWords(ctx)
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
		hasKeywords := false
		for _, banWord := range banWords {
			if strings.Contains(questionDetail.Content, banWord.Word) {
				hasKeywords = true
				break
			}
		}
		res.Result = append(res.Result, ResultUserQuestions{strconv.FormatInt(question.Qid, 10), hasKeywords, question.Closed == 1, questionDetail.Title, time.Unix(question.Time, 0), question.AnswerCount, question.ViewCount, question.FavoriteCount, question.Category, labelTitles, fmt.Sprintf("%.20s", questionDetail.Content), []string{questionDetail.PictureUrl}})
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
	message.SetBody("text/html", `<!doctype html>
<html>
<head>
<meta charset='UTF-8'><meta name='viewport' content='width=device-width initial-scale=1'>
<title>欢迎来到Reevoo σοφία</title><style type='text/css'>html {overflow-x: initial !important;}:root { --bg-color:#ffffff; --text-color:#333333; --select-text-bg-color:#B5D6FC; --select-text-font-color:auto; --monospace:"Lucida Console",Consolas,"Courier",monospace; --title-bar-height:20px; }
.mac-os-11 { --title-bar-height:28px; }
html { font-size: 14px; background-color: var(--bg-color); color: var(--text-color); font-family: "Helvetica Neue", Helvetica, Arial, sans-serif; -webkit-font-smoothing: antialiased; }
body { margin: 0px; padding: 0px; height: auto; bottom: 0px; top: 0px; left: 0px; right: 0px; font-size: 1rem; line-height: 1.42857; overflow-x: hidden; background: inherit; tab-size: 4; }
iframe { margin: auto; }
a.url { word-break: break-all; }
a:active, a:hover { outline: 0px; }
.in-text-selection, ::selection { text-shadow: none; background: var(--select-text-bg-color); color: var(--select-text-font-color); }
#write { margin: 0px auto; height: auto; width: inherit; word-break: normal; overflow-wrap: break-word; position: relative; white-space: normal; overflow-x: visible; padding-top: 36px; }
#write.first-line-indent p { text-indent: 2em; }
#write.first-line-indent li p, #write.first-line-indent p * { text-indent: 0px; }
#write.first-line-indent li { margin-left: 2em; }
.for-image #write { padding-left: 8px; padding-right: 8px; }
body.typora-export { padding-left: 30px; padding-right: 30px; }
.typora-export .footnote-line, .typora-export li, .typora-export p { white-space: pre-wrap; }
.typora-export .task-list-item input { pointer-events: none; }
@media screen and (max-width: 500px) {
  body.typora-export { padding-left: 0px; padding-right: 0px; }
  #write { padding-left: 20px; padding-right: 20px; }
  .CodeMirror-sizer { margin-left: 0px !important; }
  .CodeMirror-gutters { display: none !important; }
}
#write li > figure:last-child { margin-bottom: 0.5rem; }
#write ol, #write ul { position: relative; }
img { max-width: 100%; vertical-align: middle; image-orientation: from-image; }
button, input, select, textarea { color: inherit; font: inherit; }
input[type="checkbox"], input[type="radio"] { line-height: normal; padding: 0px; }
*, ::after, ::before { box-sizing: border-box; }
#write h1, #write h2, #write h3, #write h4, #write h5, #write h6, #write p, #write pre { width: inherit; }
#write h1, #write h2, #write h3, #write h4, #write h5, #write h6, #write p { position: relative; }
p { line-height: inherit; }
h1, h2, h3, h4, h5, h6 { break-after: avoid-page; break-inside: avoid; orphans: 4; }
p { orphans: 4; }
h1 { font-size: 2rem; }
h2 { font-size: 1.8rem; }
h3 { font-size: 1.6rem; }
h4 { font-size: 1.4rem; }
h5 { font-size: 1.2rem; }
h6 { font-size: 1rem; }
.md-math-block, .md-rawblock, h1, h2, h3, h4, h5, h6, p { margin-top: 1rem; margin-bottom: 1rem; }
.hidden { display: none; }
.md-blockmeta { color: rgb(204, 204, 204); font-weight: 700; font-style: italic; }
a { cursor: pointer; }
sup.md-footnote { padding: 2px 4px; background-color: rgba(238, 238, 238, 0.7); color: rgb(85, 85, 85); border-radius: 4px; cursor: pointer; }
sup.md-footnote a, sup.md-footnote a:hover { color: inherit; text-transform: inherit; text-decoration: inherit; }
#write input[type="checkbox"] { cursor: pointer; width: inherit; height: inherit; }
figure { overflow-x: auto; margin: 1.2em 0px; max-width: calc(100% + 16px); padding: 0px; }
figure > table { margin: 0px; }
tr { break-inside: avoid; break-after: auto; }
thead { display: table-header-group; }
table { border-collapse: collapse; border-spacing: 0px; width: 100%; overflow: auto; break-inside: auto; text-align: left; }
table.md-table td { min-width: 32px; }
.CodeMirror-gutters { border-right: 0px; background-color: inherit; }
.CodeMirror-linenumber { user-select: none; }
.CodeMirror { text-align: left; }
.CodeMirror-placeholder { opacity: 0.3; }
.CodeMirror pre { padding: 0px 4px; }
.CodeMirror-lines { padding: 0px; }
div.hr:focus { cursor: none; }
#write pre { white-space: pre-wrap; }
#write.fences-no-line-wrapping pre { white-space: pre; }
#write pre.ty-contain-cm { white-space: normal; }
.CodeMirror-gutters { margin-right: 4px; }
.md-fences { font-size: 0.9rem; display: block; break-inside: avoid; text-align: left; overflow: visible; white-space: pre; background: inherit; position: relative !important; }
.md-diagram-panel { width: 100%; margin-top: 10px; text-align: center; padding-top: 0px; padding-bottom: 8px; overflow-x: auto; }
#write .md-fences.mock-cm { white-space: pre-wrap; }
.md-fences.md-fences-with-lineno { padding-left: 0px; }
#write.fences-no-line-wrapping .md-fences.mock-cm { white-space: pre; overflow-x: auto; }
.md-fences.mock-cm.md-fences-with-lineno { padding-left: 8px; }
.CodeMirror-line, twitterwidget { break-inside: avoid; }
.footnotes { opacity: 0.8; font-size: 0.9rem; margin-top: 1em; margin-bottom: 1em; }
.footnotes + .footnotes { margin-top: 0px; }
.md-reset { margin: 0px; padding: 0px; border: 0px; outline: 0px; vertical-align: top; background: 0px 0px; text-decoration: none; text-shadow: none; float: none; position: static; width: auto; height: auto; white-space: nowrap; cursor: inherit; -webkit-tap-highlight-color: transparent; line-height: normal; font-weight: 400; text-align: left; box-sizing: content-box; direction: ltr; }
li div { padding-top: 0px; }
blockquote { margin: 1rem 0px; }
li .mathjax-block, li p { margin: 0.5rem 0px; }
li blockquote { margin: 1rem 0px; }
li { margin: 0px; position: relative; }
blockquote > :last-child { margin-bottom: 0px; }
blockquote > :first-child, li > :first-child { margin-top: 0px; }
.footnotes-area { color: rgb(136, 136, 136); margin-top: 0.714rem; padding-bottom: 0.143rem; white-space: normal; }
#write .footnote-line { white-space: pre-wrap; }
@media print {
  body, html { border: 1px solid transparent; height: 99%; break-after: avoid; break-before: avoid; font-variant-ligatures: no-common-ligatures; }
  #write { margin-top: 0px; padding-top: 0px; border-color: transparent !important; }
  .typora-export * { -webkit-print-color-adjust: exact; }
  .typora-export #write { break-after: avoid; }
  .typora-export #write::after { height: 0px; }
  .is-mac table { break-inside: avoid; }
}
.footnote-line { margin-top: 0.714em; font-size: 0.7em; }
a img, img a { cursor: pointer; }
pre.md-meta-block { font-size: 0.8rem; min-height: 0.8rem; white-space: pre-wrap; background: rgb(204, 204, 204); display: block; overflow-x: hidden; }
p > .md-image:only-child:not(.md-img-error) img, p > img:only-child { display: block; margin: auto; }
#write.first-line-indent p > .md-image:only-child:not(.md-img-error) img { left: -2em; position: relative; }
p > .md-image:only-child { display: inline-block; width: 100%; }
#write .MathJax_Display { margin: 0.8em 0px 0px; }
.md-math-block { width: 100%; }
.md-math-block:not(:empty)::after { display: none; }
.MathJax_ref { fill: currentcolor; }
[contenteditable="true"]:active, [contenteditable="true"]:focus, [contenteditable="false"]:active, [contenteditable="false"]:focus { outline: 0px; box-shadow: none; }
.md-task-list-item { position: relative; list-style-type: none; }
.task-list-item.md-task-list-item { padding-left: 0px; }
.md-task-list-item > input { position: absolute; top: 0px; left: 0px; margin-left: -1.2em; margin-top: calc(1em - 10px); border: none; }
.math { font-size: 1rem; }
.md-toc { min-height: 3.58rem; position: relative; font-size: 0.9rem; border-radius: 10px; }
.md-toc-content { position: relative; margin-left: 0px; }
.md-toc-content::after, .md-toc::after { display: none; }
.md-toc-item { display: block; color: rgb(65, 131, 196); }
.md-toc-item a { text-decoration: none; }
.md-toc-inner:hover { text-decoration: underline; }
.md-toc-inner { display: inline-block; cursor: pointer; }
.md-toc-h1 .md-toc-inner { margin-left: 0px; font-weight: 700; }
.md-toc-h2 .md-toc-inner { margin-left: 2em; }
.md-toc-h3 .md-toc-inner { margin-left: 4em; }
.md-toc-h4 .md-toc-inner { margin-left: 6em; }
.md-toc-h5 .md-toc-inner { margin-left: 8em; }
.md-toc-h6 .md-toc-inner { margin-left: 10em; }
@media screen and (max-width: 48em) {
  .md-toc-h3 .md-toc-inner { margin-left: 3.5em; }
  .md-toc-h4 .md-toc-inner { margin-left: 5em; }
  .md-toc-h5 .md-toc-inner { margin-left: 6.5em; }
  .md-toc-h6 .md-toc-inner { margin-left: 8em; }
}
a.md-toc-inner { font-size: inherit; font-style: inherit; font-weight: inherit; line-height: inherit; }
.footnote-line a:not(.reversefootnote) { color: inherit; }
.md-attr { display: none; }
.md-fn-count::after { content: "."; }
code, pre, samp, tt { font-family: var(--monospace); }
kbd { margin: 0px 0.1em; padding: 0.1em 0.6em; font-size: 0.8em; color: rgb(36, 39, 41); background: rgb(255, 255, 255); border: 1px solid rgb(173, 179, 185); border-radius: 3px; box-shadow: rgba(12, 13, 14, 0.2) 0px 1px 0px, rgb(255, 255, 255) 0px 0px 0px 2px inset; white-space: nowrap; vertical-align: middle; }
.md-comment { color: rgb(162, 127, 3); opacity: 0.8; font-family: var(--monospace); }
code { text-align: left; vertical-align: initial; }
a.md-print-anchor { white-space: pre !important; border-width: initial !important; border-style: none !important; border-color: initial !important; display: inline-block !important; position: absolute !important; width: 1px !important; right: 0px !important; outline: 0px !important; background: 0px 0px !important; text-decoration: initial !important; text-shadow: initial !important; }
.md-inline-math .MathJax_SVG .noError { display: none !important; }
.html-for-mac .inline-math-svg .MathJax_SVG { vertical-align: 0.2px; }
.md-math-block .MathJax_SVG_Display { text-align: center; margin: 0px; position: relative; text-indent: 0px; max-width: none; max-height: none; min-height: 0px; min-width: 100%; width: auto; overflow-y: hidden; display: block !important; }
.MathJax_SVG_Display, .md-inline-math .MathJax_SVG_Display { width: auto; margin: inherit; display: inline-block !important; }
.MathJax_SVG .MJX-monospace { font-family: var(--monospace); }
.MathJax_SVG .MJX-sans-serif { font-family: sans-serif; }
.MathJax_SVG { display: inline; font-style: normal; font-weight: 400; line-height: normal; zoom: 90%; text-indent: 0px; text-align: left; text-transform: none; letter-spacing: normal; word-spacing: normal; overflow-wrap: normal; white-space: nowrap; float: none; direction: ltr; max-width: none; max-height: none; min-width: 0px; min-height: 0px; border: 0px; padding: 0px; margin: 0px; }
.MathJax_SVG * { transition: none 0s ease 0s; }
.MathJax_SVG_Display svg { vertical-align: middle !important; margin-bottom: 0px !important; margin-top: 0px !important; }
.os-windows.monocolor-emoji .md-emoji { font-family: "Segoe UI Symbol", sans-serif; }
.md-diagram-panel > svg { max-width: 100%; }
[lang="flow"] svg, [lang="mermaid"] svg { max-width: 100%; height: auto; }
[lang="mermaid"] .node text { font-size: 1rem; }
table tr th { border-bottom: 0px; }
video { max-width: 100%; display: block; margin: 0px auto; }
iframe { max-width: 100%; width: 100%; border: none; }
.highlight td, .highlight tr { border: 0px; }
mark { background: rgb(255, 255, 0); color: rgb(0, 0, 0); }
.md-html-inline .md-plain, .md-html-inline strong, mark .md-inline-math, mark strong { color: inherit; }
mark .md-meta { color: rgb(0, 0, 0); opacity: 0.3 !important; }
@media print {
  .typora-export h1, .typora-export h2, .typora-export h3, .typora-export h4, .typora-export h5, .typora-export h6 { break-inside: avoid; }
}
.md-diagram-panel .messageText { stroke: none !important; }
.md-diagram-panel .start-state { fill: var(--node-fill); }
.md-diagram-panel .edgeLabel rect { opacity: 1 !important; }
.md-require-zoom-fix foreignobject { font-size: var(--mermaid-font-zoom); }


html {
	font-size: 19px;
}

html, body {
	margin: auto;
	background: #fefefe;
}
body {
	font-family: "Vollkorn", Palatino, Times;
	color: #333;
	line-height: 1.4;
	text-align: justify;
}

#write {
	max-width: 960px;
	margin: 0 auto;
	margin-bottom: 2em;
	line-height: 1.53;
	padding-top: 40px;
}

@media only screen and (min-width: 1400px) {
	#write {
		max-width: 1100px;
	}
}

@media print {
	html {
		font-size: 13px;
	}
}

/* Typography
-------------------------------------------------------- */

#write>h1:first-child,
h1 {
	margin-top: 1.6em;
	font-weight: normal;
}

h1 {
	font-size:3em;
}

h2 {
	margin-top:2em;
	font-weight: normal;
}

h3 {
	font-weight: normal;
	font-style: italic;
	margin-top: 3em;
}

h1, 
h2, 
h3{
	text-align: center;
}

h2:after{
	border-bottom: 1px solid #2f2f2f;
    content: '';
    width: 100px;
    display: block;
    margin: 0 auto;
    height: 1px;
}

h1+h2, h2+h3 {
	margin-top: 0.83em;
}

p,
.mathjax-block {
	margin-top: 0;
	-webkit-hypens: auto;
	-moz-hypens: auto;
	hyphens: auto;
}
ul {
	list-style: square;
	padding-left: 1.2em;
}
ol {
	padding-left: 1.2em;
}
blockquote {
	margin-left: 1em;
	padding-left: 1em;
	border-left: 1px solid #ddd;
}
code,
pre {
	font-family: "Consolas", "Menlo", "Monaco", monospace, serif;
	font-size: .9em;
	background: white;
}
.md-fences{
	margin-left: 1em;
	padding-left: 1em;
	border: 1px solid #ddd;
	padding-bottom: 8px;
	padding-top: 6px;
	margin-bottom: 1.5em;
}

a {
	color: #2484c1;
	text-decoration: none;
}
a:hover {
	text-decoration: underline;
}
a img {
	border: none;
}
h1 a,
h1 a:hover {
	color: #333;
	text-decoration: none;
}
hr {
	color: #ddd;
	height: 1px;
	margin: 2em 0;
	border-top: solid 1px #ddd;
	border-bottom: none;
	border-left: 0;
	border-right: 0;
}
.ty-table-edit {
	background: #ededed;
    padding-top: 4px;
}
table {
	margin-bottom: 1.333333rem
}
table th,
table td {
	padding: 8px;
	line-height: 1.333333rem;
	vertical-align: top;
	border-top: 1px solid #ddd
}
table th {
	font-weight: bold
}
table thead th {
	vertical-align: bottom
}
table caption+thead tr:first-child th,
table caption+thead tr:first-child td,
table colgroup+thead tr:first-child th,
table colgroup+thead tr:first-child td,
table thead:first-child tr:first-child th,
table thead:first-child tr:first-child td {
	border-top: 0
}
table tbody+tbody {
	border-top: 2px solid #ddd
}

.task-list{
	padding:0;
}

.md-task-list-item {
	padding-left: 1.6rem;
}

.md-task-list-item > input:before {
	content: '\221A';
	display: inline-block;
	width: 1.33333333rem;
  	height: 1.6rem;
	vertical-align: middle;
	text-align: center;
	color: #ddd;
	background-color: #fefefe;
}

.md-task-list-item > input:checked:before,
.md-task-list-item > input[checked]:before{
	color: inherit;
}
.md-tag {
	color: inherit;
	font: inherit;
}
#write pre.md-meta-block {
	min-height: 35px;
	padding: 0.5em 1em;
}
#write pre.md-meta-block {
	white-space: pre;
	background: #f8f8f8;
	border: 0px;
	color: #999;
	
	width: 100vw;
	max-width: calc(100% + 60px);
	margin-left: -30px;
	border-left: 30px #f8f8f8 solid;
	border-right: 30px #f8f8f8 solid;

	margin-bottom: 2em;
	margin-top: -1.3333333333333rem;
	padding-top: 26px;
	padding-bottom: 10px;
	line-height: 1.8em;
	font-size: 0.9em;
	font-size: 0.76em;
	padding-left: 0;
}
.md-img-error.md-image>.md-meta{
	vertical-align: bottom;
}
#write>h5.md-focus:before {
	top: 2px;
}

.md-toc {
	margin-top: 40px;
}

.md-toc-content {
	padding-bottom: 20px;
}

.outline-expander:before {
	color: inherit;
	font-size: 14px;
	top: auto;
	content: "\f0da";
	font-family: FontAwesome;
}

.outline-expander:hover:before,
.outline-item-open>.outline-item>.outline-expander:before {
  	content: "\f0d7";
}

/** source code mode */
#typora-source {
	font-family: Courier, monospace;
    color: #6A6A6A;
}

.html-for-mac #typora-sidebar {
    -webkit-box-shadow: 0 6px 12px rgba(0, 0, 0, .175);
    box-shadow: 0 6px 12px rgba(0, 0, 0, .175);
}

.cm-s-typora-default .cm-header, 
.cm-s-typora-default .cm-property,
.CodeMirror.cm-s-typora-default div.CodeMirror-cursor {
	color: #428bca;
}

.cm-s-typora-default .cm-atom, .cm-s-typora-default .cm-number {
	color: #777777;
}

.typora-node .file-list-item-parent-loc, 
.typora-node .file-list-item-time, 
.typora-node .file-list-item-summary {
	font-family: arial, sans-serif;
}

.md-task-list-item>input {
    margin-left: -1.3em;
    margin-top: calc(1rem - 12px);
}

.md-mathjax-midline {
	background: #fafafa;
}

.md-fences .code-tooltip {
	bottom: -2em !important;
}

.dropdown-menu .divider {
	border-color: #e5e5e5;
}

 :root {--mermaid-font-zoom:1em ;} 
</style>
</head>
<body class='typora-export os-windows'>
<div id='write'  class=''><h1><a name="欢迎来到reevoo-σοφία" class="md-header-anchor"></a><span>欢迎来到Reevoo σοφία</span></h1><h2><a name="你的验证码是" class="md-header-anchor"></a><span>你的验证码是：</span></h2><div align="center">`+strconv.FormatInt(code, 10)+`</div><p><img src="C:\Users\Radiu\OneDrive - sjtu.edu.cn\设计\素材\reevoo-logo.jpg" alt="reevoo-logo" style="zoom: 2%;" /></p><footer>
        <p style="text-align:center">
            <a href="http://reevoo-frontend.s3-website-us-east-1.amazonaws.com">Reevoo σοφία</a>© 2021 All rights
            reserved.
        </p>
    </footer></div>
</body>
</html>`)
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

func (u *UsersServiceImpl) WordBan(token string, req ReqWordBan) (res ResWordBan, err error) {
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
	if user.Role != entity.ADMIN {
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	if req.Ban {
		err = u.usersDao.InsertBanWord(ctx, entity.BanWords{Word: req.Word})
	} else {
		err = u.usersDao.RemoveBanWordByWord(ctx, req.Word)
	}
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	return res, u.usersDao.Commit(&ctx)
}

func (u *UsersServiceImpl) WordsBanned(token string, page int64) (res ResWordsBanned, err error) {
	var ctx dao.TransactionContext
	ctx, err = u.usersDao.Begin(true)
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
	if user.Role != entity.ADMIN {
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	var banWords []entity.BanWords
	banWords, err = u.usersDao.FindBanWordsPageable(ctx, dao.Pageable{Number: page, Size: 10})
	if err != nil {
		log.Info(err)
		res.Code = 1
		return res, u.usersDao.Rollback(&ctx)
	}
	res.Code = 0
	res.Result = []string{}
	for _, banWord := range banWords {
		res.Result = append(res.Result, banWord.Word)
	}
	return res, u.usersDao.Commit(&ctx)
}
