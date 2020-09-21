package dao

import "github.com/zhanghanchong/users-service/entity"

type UsersDao interface {
	Init() (err error)
	Destruct()

	FindByEmail(email string) (user entity.Users, err error)
	FindById(id int64) (user entity.Users, err error)
	FindByUsername(username string) (user entity.Users, err error)
	Insert(user entity.Users) (id int64, err error)
	Update(user entity.Users) (err error)
}
