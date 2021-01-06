package service

import (
	"github.com/zhanghanchong/users-service/dao"
)

type UsersService interface {
	Init(usersDAO ...dao.UsersDao) (err error)
	Destruct()
	Ban(token string, req ReqBan) (res ResBan, err error)
	Banned(token string, page int64) (res ResBanned, err error)
	CheckSession(token string) (res ResCheckSession, err error)
	CheckToken(token string) (res ResCheckToken, err error)
	Follow(token string, uid int64, follow bool) (res ResFollow, err error)
	Followed(token string, uid int64) (res ResFollowed, err error)
	Followers(token string, uid int64) (res ResFollowers, err error)
	InfoList(req ReqInfoList) (res ResInfoList, err error)
	Login(req ReqLogin) (res ResLogin, err error)
	Notifications(token string, page int64) (res ResNotifications, err error)
	OAuthGithub(code string, error string) (res ResOAuthGithub, err error)
	Passwd(token string, req ReqPasswd) (res ResPasswd, err error)
	PublicInfoGet(token string, uid int64) (res ResPublicInfoGet, err error)
	PublicInfoPut(token string, req ReqPublicInfoPut) (res ResPublicInfoPut, err error)
	RefreshToken(req ReqRefreshToken) (res ResRefreshToken, err error)
	Register(req ReqRegister) (res ResRegister, err error)
	UserAnswers(token string, uid int64, page int64) (res ResUserAnswers, err error)
	UserQuestions(token string, uid int64, page int64) (res ResUserQuestions, err error)
	VerificationCode(register bool, email string) (res ResVerificationCode, err error)
	Verify(email string, code int64) (res ResVerify, err error)
	WordBan(token string, req ReqWordBan) (res ResWordBan, err error)
	WordsBanned(token string, page int64) (res ResWordsBanned, err error)
}
