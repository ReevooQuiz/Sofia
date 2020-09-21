package service

import (
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
)

type UsersService interface {
	Init(usersDAO ...dao.UsersDao) (err error)
	Destruct()

	FindByEmail(email string) (user entity.Users, err error)
	FindById(id int64) (user entity.Users, err error)
	FindByUsername(username string) (user entity.Users, err error)
	Insert(user entity.Users) (id int64, err error)
	Update(user entity.Users) (err error)
}
