package dao

import (
	"github.com/zhanghanchong/users-service/entity"
)

type UsersDao interface {
	Init() (err error)
	Destruct()
	FindUserByEmail(email string) (user entity.Users, err error)
	FindUserByName(name string) (user entity.Users, err error)
	FindUserByOidAndAccountType(oid string, accountType int8) (user entity.Users, err error)
	FindUserByUid(uid int64) (user entity.Users, err error)
	FindUserDetailByUid(uid int64) (userDetail entity.UserDetails, err error)
	InsertFavorite(favorite entity.Favorites) (fid int64, err error)
	InsertUser(user entity.Users) (uid int64, err error)
	InsertUserDetail(userDetail entity.UserDetails) (err error)
	UpdateUser(user entity.Users) (err error)
	UpdateUserDetail(userDetail entity.UserDetails) (err error)
}
