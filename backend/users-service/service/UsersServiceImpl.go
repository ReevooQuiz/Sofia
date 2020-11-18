package service

import (
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
	"gopkg.in/mgo.v2/bson"
)

type UsersServiceImpl struct {
	usersDao dao.UsersDao
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

func (u *UsersServiceImpl) FindUserByEmail(email string) (user entity.Users, err error) {
	return u.usersDao.FindUserByEmail(email)
}

func (u *UsersServiceImpl) FindUserByNickname(nickname string) (user entity.Users, err error) {
	return u.usersDao.FindUserByNickname(nickname)
}

func (u *UsersServiceImpl) FindUserByUid(uid bson.ObjectId) (user entity.Users, err error) {
	return u.usersDao.FindUserByUid(uid)
}

func (u *UsersServiceImpl) InsertUser(user entity.Users) (uid bson.ObjectId, err error) {
	return u.usersDao.InsertUser(user)
}

func (u *UsersServiceImpl) UpdateUser(user entity.Users) (err error) {
	return u.usersDao.UpdateUser(user)
}
