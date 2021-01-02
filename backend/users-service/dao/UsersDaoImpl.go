package dao

import (
	"context"
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

type TransactionContext struct {
	sqlTx   *sql.Tx
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
}

func (u *UsersDaoImpl) Begin(read bool) (ctx TransactionContext, err error) {
	var tx *sql.Tx
	if read {
		tx, err = u.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	} else {
		tx, err = u.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	}
	if err != nil {
		return ctx, err
	}
	return TransactionContext{tx, u.session.New()}, nil
}

func (u *UsersDaoImpl) Commit(t *TransactionContext) (err error) {
	t.session.Close()
	return t.sqlTx.Commit()
}

func (u *UsersDaoImpl) Rollback(t *TransactionContext) (err error) {
	t.session.Close()
	return t.sqlTx.Rollback()
}

func (u *UsersDaoImpl) FindFollowByUidAndFollower(ctx TransactionContext, uid int64, follower int64) (follow entity.Follows, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from follows where uid = ? and follower = ?")
	if err != nil {
		return follow, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid, follower).Scan(&follow.Uid, &follow.Follower)
	return follow, err
}

func (u *UsersDaoImpl) FindFollowsByFollower(ctx TransactionContext, follower int64) (follows []entity.Follows, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from follows where follower = ?")
	if err != nil {
		return follows, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(follower)
	if err != nil {
		return follows, err
	}
	follows = []entity.Follows{}
	for res.Next() {
		var follow entity.Follows
		err = res.Scan(&follow.Uid, &follow.Follower)
		if err != nil {
			return follows, err
		}
		follows = append(follows, follow)
	}
	return follows, err
}

func (u *UsersDaoImpl) FindFollowsByUid(ctx TransactionContext, uid int64) (follows []entity.Follows, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from follows where uid = ?")
	if err != nil {
		return follows, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(uid)
	if err != nil {
		return follows, err
	}
	follows = []entity.Follows{}
	for res.Next() {
		var follow entity.Follows
		err = res.Scan(&follow.Uid, &follow.Follower)
		if err != nil {
			return follows, err
		}
		follows = append(follows, follow)
	}
	return follows, err
}

func (u *UsersDaoImpl) FindLabelByTitle(ctx TransactionContext, title string) (label entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from labels where title = ?")
	if err != nil {
		return label, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(title).Scan(&label.Lid, &label.Title)
	return label, err
}

func (u *UsersDaoImpl) FindLabelsByUid(ctx TransactionContext, uid int64) (labels []entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select lid, title from user_labels natural join labels where uid = ?")
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

func (u *UsersDaoImpl) FindUserByEmail(ctx TransactionContext, email string) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from users where email = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserByName(ctx TransactionContext, name string) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from users where name = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(name).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserByOidAndAccountType(ctx TransactionContext, oid string, accountType int8) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from users where oid = ? and account_type = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(oid, accountType).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserByUid(ctx TransactionContext, uid int64) (user entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from users where uid = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid).Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
	return user, err
}

func (u *UsersDaoImpl) FindUserDetailByUid(ctx TransactionContext, uid int64) (userDetail entity.UserDetails, err error) {
	var res []entity.UserDetails
	err = ctx.session.DB("sofia").C("user_details").Find(bson.M{"uid": uid}).All(&res)
	if err != nil {
		return userDetail, err
	}
	if len(res) == 0 {
		return userDetail, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) InsertFavorite(ctx TransactionContext, favorite entity.Favorites) (fid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into favorites(uid, title) values(?, ?)")
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

func (u *UsersDaoImpl) InsertFollow(ctx TransactionContext, follow entity.Follows) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into follows values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(follow.Uid, follow.Follower)
	return err
}

func (u *UsersDaoImpl) InsertLabel(ctx TransactionContext, label entity.Labels) (lid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into labels(title) values(?)")
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

func (u *UsersDaoImpl) InsertUser(ctx TransactionContext, user entity.Users) (uid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into users(oid, name, nickname, salt, hash_password, email, gender, profile, role, account_type, active_code, passwd_code, exp, follower_count, following_count, question_count, answer_count, like_count, approval_count, notification_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
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

func (u *UsersDaoImpl) InsertUserDetail(ctx TransactionContext, userDetail entity.UserDetails) (err error) {
	return ctx.session.DB("sofia").C("user_details").Insert(userDetail)
}

func (u *UsersDaoImpl) InsertUserLabel(ctx TransactionContext, userLabel entity.UserLabels) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into user_labels values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userLabel.Uid, userLabel.Lid)
	return err
}

func (u *UsersDaoImpl) RemoveFollowByUidAndFollower(ctx TransactionContext, uid int64, follower int64) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("delete from follows where uid = ? and follower = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, follower)
	return err
}

func (u *UsersDaoImpl) RemoveUserLabelsByUid(ctx TransactionContext, uid int64) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("delete from user_labels where uid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid)
	return err
}

func (u *UsersDaoImpl) UpdateUserByUid(ctx TransactionContext, user entity.Users) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("update users set oid = ?, name = ?, nickname = ?, salt = ?, hash_password = ?, email = ?, gender = ?, profile= ?, role = ?, account_type = ?, active_code = ?, passwd_code = ?, exp = ?, follower_count = ?, following_count = ?, question_count = ?, answer_count = ?, like_count = ?, approval_count = ?, notification_time = ? where uid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Oid, user.Name, user.Nickname, user.Salt, user.HashPassword, user.Email, user.Gender, user.Profile, user.Role, user.AccountType, user.ActiveCode, user.PasswdCode, user.Exp, user.FollowerCount, user.FollowingCount, user.QuestionCount, user.AnswerCount, user.LikeCount, user.ApprovalCount, user.NotificationTime, user.Uid)
	return err
}

func (u *UsersDaoImpl) UpdateUserDetailByUid(ctx TransactionContext, userDetail entity.UserDetails) (err error) {
	return ctx.session.DB("sofia").C("user_details").Update(bson.M{"uid": userDetail.Uid}, userDetail)
}
