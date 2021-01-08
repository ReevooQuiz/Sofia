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

type Notifications struct {
	Type int8
	Id0  int64
	Id1  int64
}

type Pageable struct {
	Number int64
	Size   int64
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

func (u *UsersDaoImpl) FindAnswerByAid(ctx TransactionContext, aid int64) (answer entity.Answers, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from answers where aid = ?")
	if err != nil {
		return answer, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(aid).Scan(&answer.Aid, &answer.Answerer, &answer.Qid, &answer.CommentCount, &answer.CriticismCount, &answer.LikeCount, &answer.ApprovalCount, &answer.Time)
	return answer, err
}

func (u *UsersDaoImpl) FindAnswerDetailByAid(ctx TransactionContext, aid int64) (answerDetail entity.AnswerDetails, err error) {
	var res []entity.AnswerDetails
	err = ctx.session.DB("sofia").C("answer_details").Find(bson.M{"_id": aid}).All(&res)
	if err != nil {
		return answerDetail, err
	}
	if len(res) == 0 {
		return answerDetail, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindAnswersByAnswererOrderByTimeDescPageable(ctx TransactionContext, answerer int64, pageable Pageable) (answers []entity.Answers, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from answers where answerer = ? order by time desc limit ?, ?")
	if err != nil {
		return answers, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(answerer, pageable.Number*pageable.Size, pageable.Size)
	if err != nil {
		return answers, err
	}
	for res.Next() {
		var answer entity.Answers
		err = res.Scan(&answer.Aid, &answer.Answerer, &answer.Qid, &answer.CommentCount, &answer.CriticismCount, &answer.LikeCount, &answer.ApprovalCount, &answer.Time)
		if err != nil {
			return answers, err
		}
		answers = append(answers, answer)
	}
	return answers, err
}

func (u *UsersDaoImpl) FindApproveAnswerByUidAndAid(ctx TransactionContext, uid int64, aid int64) (approveAnswer entity.ApproveAnswers, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from approve_answers where uid = ? and aid = ?")
	if err != nil {
		return approveAnswer, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid, aid).Scan(&approveAnswer.Uid, &approveAnswer.Aid, &approveAnswer.Time)
	return approveAnswer, err
}

func (u *UsersDaoImpl) FindBanWords(ctx TransactionContext) (banWords []entity.BanWords, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from ban_words")
	if err != nil {
		return banWords, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query()
	if err != nil {
		return banWords, err
	}
	banWords = []entity.BanWords{}
	for res.Next() {
		var banWord entity.BanWords
		err = res.Scan(&banWord.Word)
		if err != nil {
			return banWords, err
		}
		banWords = append(banWords, banWord)
	}
	return banWords, err
}

func (u *UsersDaoImpl) FindBanWordsPageable(ctx TransactionContext, pageable Pageable) (banWords []entity.BanWords, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from ban_words limit ?, ?")
	if err != nil {
		return banWords, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(pageable.Number*pageable.Size, pageable.Size)
	if err != nil {
		return banWords, err
	}
	banWords = []entity.BanWords{}
	for res.Next() {
		var banWord entity.BanWords
		err = res.Scan(&banWord.Word)
		if err != nil {
			return banWords, err
		}
		banWords = append(banWords, banWord)
	}
	return banWords, err
}

func (u *UsersDaoImpl) FindCommentByCmid(ctx TransactionContext, cmid int64) (comment entity.Comments, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from comments where cmid = ?")
	if err != nil {
		return comment, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(cmid).Scan(&comment.Cmid, &comment.Uid, &comment.Aid, &comment.Time)
	return comment, err
}

func (u *UsersDaoImpl) FindCriticismByCtid(ctx TransactionContext, ctid int64) (criticism entity.Criticisms, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from criticisms where ctid = ?")
	if err != nil {
		return criticism, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ctid).Scan(&criticism.Ctid, &criticism.Uid, &criticism.Aid, &criticism.Time)
	return criticism, err
}

func (u *UsersDaoImpl) FindFavoriteByUidAndTitle(ctx TransactionContext, uid int64, title string) (favorite entity.Favorites, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from favorites where uid = ? and title = ?")
	if err != nil {
		return favorite, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid, title).Scan(&favorite.Fid, &favorite.Uid, &favorite.Title)
	return favorite, err
}

func (u *UsersDaoImpl) FindFavoriteItemByFidAndQid(ctx TransactionContext, fid int64, qid int64) (favoriteItem entity.FavoriteItems, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from favorite_items where fid = ? and qid = ?")
	if err != nil {
		return favoriteItem, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(fid, qid).Scan(&favoriteItem.Fid, &favoriteItem.Qid)
	return favoriteItem, err
}

func (u *UsersDaoImpl) FindFavoriteItemsByFidPageable(ctx TransactionContext, fid int64, pageable Pageable) (favoriteItems []entity.FavoriteItems, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from favorite_items where fid = ? limit ?, ?")
	if err != nil {
		return favoriteItems, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(fid, pageable.Number*pageable.Size, pageable.Size)
	if err != nil {
		return favoriteItems, err
	}
	favoriteItems = []entity.FavoriteItems{}
	for res.Next() {
		var favoriteItem entity.FavoriteItems
		err = res.Scan(&favoriteItem.Fid, &favoriteItem.Qid)
		if err != nil {
			return favoriteItems, err
		}
		favoriteItems = append(favoriteItems, favoriteItem)
	}
	return favoriteItems, err
}

func (u *UsersDaoImpl) FindFollowByUidAndFollower(ctx TransactionContext, uid int64, follower int64) (follow entity.Follows, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from follows where uid = ? and follower = ?")
	if err != nil {
		return follow, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid, follower).Scan(&follow.Uid, &follow.Follower, &follow.Time)
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
		err = res.Scan(&follow.Uid, &follow.Follower, &follow.Time)
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
		err = res.Scan(&follow.Uid, &follow.Follower, &follow.Time)
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

func (u *UsersDaoImpl) FindLabelsByQid(ctx TransactionContext, qid int64) (labels []entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select lid, title from question_labels natural join labels where qid = ?")
	if err != nil {
		return labels, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(qid)
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

func (u *UsersDaoImpl) FindLikeAnswerByUidAndAid(ctx TransactionContext, uid int64, aid int64) (likeAnswer entity.LikeAnswers, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from like_answers where uid = ? and aid = ?")
	if err != nil {
		return likeAnswer, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid, aid).Scan(&likeAnswer.Uid, &likeAnswer.Aid, &likeAnswer.Time)
	return likeAnswer, err
}

func (u *UsersDaoImpl) FindNotificationsByUidPageable(ctx TransactionContext, uid int64, pageable Pageable) (notifications []Notifications, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select type, id0, id1 from ((select 0 as type, aid as id0, 0 as id1, time from answers where qid in (select qid from questions where raiser = ?)) union (select 1 as type, uid as id0, aid as id1, time from like_answers where aid in (select aid from answers where answerer = ?)) union (select 2 as type, uid as id0, aid as id1, time from approve_answers where aid in (select aid from answers where answerer = ?)) union (select 3 as type, cmid as id0, 0 as id1, time from comments where aid in (select aid from answers where answerer = ?)) union (select 4 as type, ctid as id0, 0 as id1, time from criticisms where aid in (select aid from answers where answerer = ?)) union (select 5 as type, follower as id0, 0 as id1, time from follows where uid = ?)) as N order by time desc limit ?, ?")
	if err != nil {
		return notifications, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(uid, uid, uid, uid, uid, uid, pageable.Number*pageable.Size, pageable.Size)
	notifications = []Notifications{}
	for res.Next() {
		var notification Notifications
		err = res.Scan(&notification.Type, &notification.Id0, &notification.Id1)
		if err != nil {
			return notifications, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, err
}

func (u *UsersDaoImpl) FindQuestionByQid(ctx TransactionContext, qid int64) (question entity.Questions, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from questions where qid = ?")
	if err != nil {
		return question, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(qid).Scan(&question.Qid, &question.Raiser, &question.Category, &question.AcceptedAnswer, &question.AnswerCount, &question.ViewCount, &question.FavoriteCount, &question.Time, &question.Scanned, &question.Closed)
	return question, err
}

func (u *UsersDaoImpl) FindQuestionDetailByQid(ctx TransactionContext, qid int64) (questionDetail entity.QuestionDetails, err error) {
	var res []entity.QuestionDetails
	err = ctx.session.DB("sofia").C("question_details").Find(bson.M{"_id": qid}).All(&res)
	if err != nil {
		return questionDetail, err
	}
	if len(res) == 0 {
		return questionDetail, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindQuestionsByRaiserOrderByTimeDescPageable(ctx TransactionContext, raiser int64, pageable Pageable) (questions []entity.Questions, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from questions where raiser = ? order by time desc limit ?, ?")
	if err != nil {
		return questions, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(raiser, pageable.Number*pageable.Size, pageable.Size)
	if err != nil {
		return questions, err
	}
	for res.Next() {
		var question entity.Questions
		err = res.Scan(&question.Qid, &question.Raiser, &question.Category, &question.AcceptedAnswer, &question.AnswerCount, &question.ViewCount, &question.FavoriteCount, &question.Time, &question.Scanned, &question.Closed)
		if err != nil {
			return questions, err
		}
		questions = append(questions, question)
	}
	return questions, err
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
	err = ctx.session.DB("sofia").C("user_details").Find(bson.M{"_id": uid}).All(&res)
	if err != nil {
		return userDetail, err
	}
	if len(res) == 0 {
		return userDetail, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (u *UsersDaoImpl) FindUsersByRolePageable(ctx TransactionContext, role int8, pageable Pageable) (users []entity.Users, err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select * from users where role = ? limit ?, ?")
	if err != nil {
		return users, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(role, pageable.Number*pageable.Size, pageable.Size)
	if err != nil {
		return users, err
	}
	users = []entity.Users{}
	for res.Next() {
		var user entity.Users
		err = res.Scan(&user.Uid, &user.Oid, &user.Name, &user.Nickname, &user.Salt, &user.HashPassword, &user.Email, &user.Gender, &user.Profile, &user.Role, &user.AccountType, &user.ActiveCode, &user.PasswdCode, &user.Exp, &user.FollowerCount, &user.FollowingCount, &user.QuestionCount, &user.AnswerCount, &user.LikeCount, &user.ApprovalCount, &user.NotificationTime)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u *UsersDaoImpl) InsertApproveAnswer(ctx TransactionContext, approveAnswer entity.ApproveAnswers) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into approve_answers values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(approveAnswer.Uid, approveAnswer.Aid, approveAnswer.Time)
	return err
}

func (u *UsersDaoImpl) InsertBanWord(ctx TransactionContext, banWord entity.BanWords) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into ban_words values(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(banWord.Word)
	return err
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

func (u *UsersDaoImpl) InsertFavoriteItem(ctx TransactionContext, favoriteItem entity.FavoriteItems) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into favorite_items values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(favoriteItem.Fid, favoriteItem.Qid)
	return err
}

func (u *UsersDaoImpl) InsertFollow(ctx TransactionContext, follow entity.Follows) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into follows values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(follow.Uid, follow.Follower, follow.Time)
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

func (u *UsersDaoImpl) InsertLikeAnswer(ctx TransactionContext, likeAnswer entity.LikeAnswers) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("insert into like_answers values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(likeAnswer.Uid, likeAnswer.Aid, likeAnswer.Time)
	return err
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

func (u *UsersDaoImpl) RemoveApproveAnswerByUidAndAid(ctx TransactionContext, uid int64, aid int64) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("delete from approve_answers where uid = ? and aid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, aid)
	return err
}

func (u *UsersDaoImpl) RemoveBanWordByWord(ctx TransactionContext, word string) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("delete from ban_words where word = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(word)
	return err
}

func (u *UsersDaoImpl) RemoveFavoriteItemByFidAndQid(ctx TransactionContext, fid int64, qid int64) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("delete from favorite_items where fid = ? and qid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(fid, qid)
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

func (u *UsersDaoImpl) RemoveLikeAnswerByUidAndAid(ctx TransactionContext, uid int64, aid int64) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("delete from like_answers where uid = ? and aid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, aid)
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

func (u *UsersDaoImpl) UpdateAnswerByAid(ctx TransactionContext, answer entity.Answers) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("update answers set answerer = ?, qid = ?, comment_count = ?, criticism_count = ?, like_count = ?, approval_count = ?, time = ? where aid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(answer.Answerer, answer.Qid, answer.CommentCount, answer.CriticismCount, answer.LikeCount, answer.ApprovalCount, answer.Time, answer.Aid)
	return err
}

func (u *UsersDaoImpl) UpdateQuestionByQid(ctx TransactionContext, question entity.Questions) (err error) {
	var stmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("update questions set raiser = ?, category = ?, accepted_answer = ?, answer_count = ?, view_count = ?, favorite_count = ?, time = ?, scanned = ?, closed = ? where qid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(question.Raiser, question.Category, question.AcceptedAnswer, question.AnswerCount, question.ViewCount, question.FavoriteCount, question.Time, question.Scanned, question.Closed, question.Qid)
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
	return ctx.session.DB("sofia").C("user_details").Update(bson.M{"_id": userDetail.Uid}, userDetail)
}
