package service

import (
	"github.com/zhanghanchong/users-service/dao"
)

type UsersService interface {
	Init(usersDAO ...dao.UsersDao) (err error)
	Destruct()
	Login(req ReqLogin) (res ResLogin, err error)
	OAuthGithub(code string, error string) (res ResOAuthGithub, err error)
	Passwd(token string, req ReqPasswd) (res ResPasswd, err error)
	Register(req ReqRegister) (res ResRegister, err error)
	VerificationCode(register bool, email string) (res ResVerificationCode, err error)
	Verify(email string, code int64) (res ResVerify, err error)
}
