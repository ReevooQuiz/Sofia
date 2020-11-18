package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/users-service/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type UsersDaoImpl struct {
	db *sql.DB
}

var mongoUrl string
var mysqlUrl string

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	mongoUrl = os.Getenv("MONGO_URL")
	mysqlUrl = os.Getenv("MYSQL_URL")
}

func (u *UsersDaoImpl) Init() (err error) {
	u.db, err = sql.Open("mysql", mysqlUrl)
	return err
}

func (u *UsersDaoImpl) Destruct() {
	_ = u.db.Close()
}

func (u *UsersDaoImpl) FindUserByEmail(email string) (user entity.Users, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return user, err
	}
	defer session.Close()
	var res []entity.Users
	err = session.DB("sofia").C("users").Find(bson.M{"email": email}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindUserByNickname(nickname string) (user entity.Users, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return user, err
	}
	defer session.Close()
	var res []entity.Users
	err = session.DB("sofia").C("users").Find(bson.M{"nickname": nickname}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindUserByUid(uid bson.ObjectId) (user entity.Users, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return user, err
	}
	defer session.Close()
	var res []entity.Users
	err = session.DB("sofia").C("users").Find(bson.M{"_id": uid}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) InsertUser(user entity.Users) (uid bson.ObjectId, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return uid, err
	}
	defer session.Close()
	user.Uid = bson.NewObjectId()
	err = session.DB("sofia").C("users").Insert(user)
	return user.Uid, err
}

func (u *UsersDaoImpl) UpdateUser(user entity.Users) (err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return err
	}
	defer session.Close()
	return session.DB("sofia").C("users").Update(bson.M{"_id": user.Uid}, user)
}
