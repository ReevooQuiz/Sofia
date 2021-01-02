package service

import (
	"github.com/zhanghanchong/users-service/dao"
)

type UsersService interface {
	Init(usersDAO ...dao.UsersDao) (err error)
	Destruct()
	CheckToken(token string) (res ResCheckToken, err error)
	Follow(token string, uid int64, follow bool) (res ResFollow, err error)
	InfoList(req ReqInfoList) (res ResInfoList, err error)
	Login(req ReqLogin) (res ResLogin, err error)
	OAuthGithub(code string, error string) (res ResOAuthGithub, err error)
	Passwd(token string, req ReqPasswd) (res ResPasswd, err error)
	PublicInfoGet(token string, uid int64) (res ResPublicInfoGet, err error)
	PublicInfoPut(token string, req ReqPublicInfoPut) (res ResPublicInfoPut, err error)
	RefreshToken(req ReqRefreshToken) (res ResRefreshToken, err error)
	Register(req ReqRegister) (res ResRegister, err error)
	VerificationCode(register bool, email string) (res ResVerificationCode, err error)
	Verify(email string, code int64) (res ResVerify, err error)
}
