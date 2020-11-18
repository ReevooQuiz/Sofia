package service

import (
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
	"gopkg.in/mgo.v2/bson"
)

type UsersService interface {
	Init(usersDAO ...dao.UsersDao) (err error)
	Destruct()

	FindUserByEmail(email string) (user entity.Users, err error)
	FindUserByNickname(nickname string) (user entity.Users, err error)
	FindUserByUid(uid bson.ObjectId) (user entity.Users, err error)
	InsertUser(user entity.Users) (uid bson.ObjectId, err error)
	UpdateUser(user entity.Users) (err error)
}
