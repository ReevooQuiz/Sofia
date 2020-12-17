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

var (
	mongoUrl string
	mysqlUrl string
)

type UsersDaoImpl struct {
	db      *sql.DB
	session *mgo.Session
}

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

func (u *UsersDaoImpl) FindLabelByTitle(title string) (label entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from labels where title = ?")
	if err != nil {
		return label, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(title).Scan(&label.Lid, &label.Title)
	return label, err
}

func (u *UsersDaoImpl) FindLabelsByUid(uid int64) (labels []entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select lid, title from user_labels natural join labels where uid = ?")
	if err != nil {
		return labels, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(uid)
	if err != nil {
		return labels, err
	}
	labels = []entity.Labels{}
	for res.Next() {
		var label entity.Labels
		err = res.Scan(&label.Lid, &label.Title)
		if err != nil {
			return labels, err
		}
		labels = append(labels, label)
	}
	return labels, err
}

func (u *UsersDaoImpl) FindUserByEmail(email string) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where email = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserByName(name string) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where name = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(name).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserByOidAndAccountType(oid string, accountType int8) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where oid = ? and account_type = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(oid, accountType).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserByUid(uid int64) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where uid = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserDetailByUid(uid int64) (userDetail entity.UserDetails, err error) {
	var res []entity.UserDetails
	err = u.session.DB("sofia").C("user_details").Find(bson.M{"uid": uid}).All(&res)
	if err != nil {
		return userDetail, err
	}
	if len(res) == 0 {
		return userDetail, errors.New("mongo: no rows in result set")
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

func (u *UsersDaoImpl) InsertLabel(label entity.Labels) (lid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("insert into labels(title) values(?)")
	if err != nil {
		return lid, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(label.Title)
	if err != nil {
		return lid, err
	}
	lid, err = res.LastInsertId()
	return lid, err
}

func (u *UsersDaoImpl) InsertUser(user entity.Users) (uid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("insert into users(oid, name, nickname, salt, hash_password, email, gender, profile, role, account_type, active_code, passwd_code, exp, follower_count, following_count, question_count, answer_count, like_count, approval_count, notification_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return uid, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(user.Oid, user.Name, user.Nickname, user.Salt, user.HashPassword, user.Email, user.Gender, user.Profile, user.Role, user.AccountType, user.ActiveCode, user.PasswdCode, user.Exp, user.FollowerCount, user.FollowingCount, user.QuestionCount, user.AnswerCount, user.LikeCount, user.ApprovalCount, user.NotificationTime)
	if err != nil {
		return uid, err
	}
	uid, err = res.LastInsertId()
	return uid, err
}

func (u *UsersDaoImpl) InsertUserDetail(userDetail entity.UserDetails) (err error) {
	return u.session.DB("sofia").C("user_details").Insert(userDetail)
}

func (u *UsersDaoImpl) InsertUserLabel(userLabel entity.UserLabels) (err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("insert into user_labels values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userLabel.Uid, userLabel.Lid)
	return err
}

func (u *UsersDaoImpl) RemoveUserLabelsByUid(uid int64) (err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("delete from user_labels where uid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid)
	return err
}

func (u *UsersDaoImpl) UpdateUserByUid(user entity.Users) (err error) {
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("update users set oid = ?, name = ?, nickname = ?, salt = ?, hash_password = ?, email = ?, gender = ?, profile= ?, role = ?, account_type = ?, active_code = ?, passwd_code = ?, exp = ?, follower_count = ?, following_count = ?, question_count = ?, answer_count = ?, like_count = ?, approval_count = ?, notification_time = ? where uid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Oid, user.Name, user.Nickname, user.Salt, user.HashPassword, user.Email, user.Gender, user.Profile, user.Role, user.AccountType, user.ActiveCode, user.PasswdCode, user.Exp, user.FollowerCount, user.FollowingCount, user.QuestionCount, user.AnswerCount, user.LikeCount, user.ApprovalCount, user.NotificationTime, user.Uid)
	return err
}

func (u *UsersDaoImpl) UpdateUserDetailByUid(userDetail entity.UserDetails) (err error) {
	return u.session.DB("sofia").C("user_details").Update(bson.M{"uid": userDetail.Uid}, userDetail)
}
