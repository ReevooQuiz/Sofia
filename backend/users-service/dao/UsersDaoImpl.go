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
	db      *sql.DB
	session *mgo.Session
}

var (
	mongoUrl string
	mysqlUrl string
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	mongoUrl = os.Getenv("MONGO_URL")
	mysqlUrl = os.Getenv("MYSQL_URL")
}

func (u *UsersDaoImpl) Init() (err error) {
	u.db, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}
	u.session, err = mgo.Dial(mongoUrl)
	return err
}

func (u *UsersDaoImpl) Destruct() {
	_ = u.db.Close()
	u.session.Close()
}

func (u *UsersDaoImpl) FindUserByEmail(email string) (user entity.Users, err error) {
	var res []entity.Users
	err = u.session.DB("sofia").C("users").Find(bson.M{"email": email}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindUserByName(name string) (user entity.Users, err error) {
	var res []entity.Users
	err = u.session.DB("sofia").C("users").Find(bson.M{"name": name}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindUserByOidAndAccountType(oid string, accountType int8) (user entity.Users, err error) {
	var res []entity.Users
	err = u.session.DB("sofia").C("users").Find(bson.M{"oid": oid, "account_type": accountType}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindUserByUid(uid bson.ObjectId) (user entity.Users, err error) {
	var res []entity.Users
	err = u.session.DB("sofia").C("users").Find(bson.M{"_id": uid}).All(&res)
	if err != nil {
		return user, err
	}
	if len(res) == 0 {
		return user, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) InsertFavorite(favorite entity.Favorites) (fid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("insert into favorites(uid, title) values(?, ?)")
	if err != nil {
		return fid, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(favorite.Uid, favorite.Title)
	if err != nil {
		return fid, err
	}
	fid, err = res.LastInsertId()
	return fid, err
}

func (u *UsersDaoImpl) InsertUser(user entity.Users) (uid bson.ObjectId, err error) {
	user.Uid = bson.NewObjectId()
	err = u.session.DB("sofia").C("users").Insert(user)
	return user.Uid, err
}

func (u *UsersDaoImpl) UpdateUser(user entity.Users) (err error) {
	return u.session.DB("sofia").C("users").Update(bson.M{"_id": user.Uid}, user)
}
