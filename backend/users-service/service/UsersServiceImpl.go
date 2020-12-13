package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/zhanghanchong/users-service/dao"
	"github.com/zhanghanchong/users-service/entity"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

const (
	SaltSize = 16
)

type UsersServiceImpl struct {
	usersDao dao.UsersDao
}

func (u *UsersServiceImpl) Init(usersDao ...dao.UsersDao) (err error) {
	if len(usersDao) == 0 {
		usersDao = append(usersDao, &dao.UsersDaoImpl{})
		rand.Seed(time.Now().UnixNano())
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

func (u *UsersServiceImpl) FindUserByName(name string) (user entity.Users, err error) {
	return u.usersDao.FindUserByName(name)
}

func (u *UsersServiceImpl) FindUserByOidAndAccountType(oid string, accountType int8) (user entity.Users, err error) {
	return u.usersDao.FindUserByOidAndAccountType(oid, accountType)
}

func (u *UsersServiceImpl) FindUserByUid(uid bson.ObjectId) (user entity.Users, err error) {
	return u.usersDao.FindUserByUid(uid)
}

func (u *UsersServiceImpl) InsertFavorite(favorite entity.Favorites) (fid int64, err error) {
	return u.usersDao.InsertFavorite(favorite)
}

func (u *UsersServiceImpl) InsertUser(user entity.Users) (uid bson.ObjectId, err error) {
	return u.usersDao.InsertUser(user)
}

func (u *UsersServiceImpl) UpdateUser(user entity.Users) (err error) {
	return u.usersDao.UpdateUser(user)
}

func (u *UsersServiceImpl) HashPassword(password string, salt string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password + salt)))
}

func (u *UsersServiceImpl) generateSalt() string {
	b := make([]byte, SaltSize)
	for i := range b {
		b[i] = byte(rand.Uint32() & 0xFF)
	}
	return fmt.Sprintf("%x", b)
}
