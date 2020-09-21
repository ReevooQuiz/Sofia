package service

import (
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
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

func (u *UsersServiceImpl) FindByEmail(email string) (user entity.Users, err error) {
	return u.usersDao.FindByEmail(email)
}

func (u *UsersServiceImpl) FindById(id int64) (user entity.Users, err error) {
	return u.usersDao.FindById(id)
}

func (u *UsersServiceImpl) FindByUsername(username string) (user entity.Users, err error) {
	return u.usersDao.FindByUsername(username)
}

func (u *UsersServiceImpl) Insert(user entity.Users) (id int64, err error) {
	return u.usersDao.Insert(user)
}

func (u *UsersServiceImpl) Update(user entity.Users) (err error) {
	return u.usersDao.Update(user)
}
