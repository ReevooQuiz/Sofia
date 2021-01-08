package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SKFE396/qa-service/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

const (
	PageSize        = 5
	CommentPageSize = 10
	answerFields    = "aid,answerer,qid,view_count,comment_count,criticism_count,like_count,approval_count,time,scanned"
	questionFields  = "qid,closed,raiser,title,category,accepted_answer,answer_count,view_count,favorite_count,time,scanned"
)

var (
	mongoUrl string
	mysqlUrl string
)

type QaDaoImpl struct {
	db      *sql.DB
	session *mgo.Session
}

type TransactionContext struct {
	sqlTx   *sql.Tx
	session *mgo.Session
}

type AnswerActionInfo struct {
	Liked      bool
	Approved   bool
	Approvable bool
}

func (q *QaDaoImpl) Commit(t *TransactionContext) {
	t.session.Close()
	e := t.sqlTx.Commit()
	if e != nil {
		log.Warn(e)
	}
}

func (q *QaDaoImpl) Rollback(t *TransactionContext) {
	t.session.Close()
	e := t.sqlTx.Rollback()
	if e != nil {
		log.Warn(e)
	}
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	mongoUrl = os.Getenv("MONGO_URL")
	mysqlUrl = os.Getenv("MYSQL_URL")
}

func (q *QaDaoImpl) Init() (err error) {
	q.db, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}
	q.session, err = mgo.Dial(mongoUrl)
	q.session.SetPoolLimit(100)
	return err
}

func (q *QaDaoImpl) Destruct() {
	_ = q.db.Close()
}

func (q *QaDaoImpl) Begin(read bool) (ctx TransactionContext, err error) {
	var tx *sql.Tx
	if read {
		tx, err = q.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	} else {
		tx, err = q.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	}
	if err != nil {
		return
	}
	ss := q.session.New()
	if ss == nil {
		e := tx.Rollback()
		if e != nil {
			log.Warn(e)
		}
		return ctx, errors.New("failed to create mongo session")
	}
	return TransactionContext{tx, ss}, nil
}

func (q *QaDaoImpl) FindCommentDetails(ctx TransactionContext, comments []entity.Comments) (details []entity.CommentDetails) {
	var findErr error
	var current entity.CommentDetails
	for _, v := range comments {
		findErr = ctx.session.DB("sofia").C("comment_details").FindId(v.Cmid).One(&current)
		if findErr != nil {
			log.Warn(findErr)
			details = append(details, current)
		} else {
			details = append(details, entity.CommentDetails{})
		}
	}
	return
}

func (q *QaDaoImpl) FindCriticismDetails(ctx TransactionContext, criticisms []entity.Criticisms) (details []entity.CriticismDetails) {
	var findErr error
	var current entity.CriticismDetails
	for _, v := range criticisms {
		findErr = ctx.session.DB("sofia").C("criticism_details").FindId(v.Ctid).One(&current)
		if findErr != nil {
			log.Warn(findErr)
			details = append(details, current)
		} else {
			details = append(details, entity.CriticismDetails{})
		}
	}
	return
}

func (q *QaDaoImpl) FindQuestionDetails(ctx TransactionContext, questions []entity.Questions) (questionDetails []entity.QuestionDetails) {
	var findErr error
	var current entity.QuestionDetails
	for _, v := range questions {
		findErr = ctx.session.DB("sofia").C("question_details").FindId(v.Qid).One(&current)
		if findErr != nil {
			log.Warn(findErr)
			questionDetails = append(questionDetails, current)
		} else {
			questionDetails = append(questionDetails, entity.QuestionDetails{})
		}
	}
	return
}

func (q *QaDaoImpl) FindAnswerDetails(ctx TransactionContext, answers []entity.Answers) (answerDetails []entity.AnswerDetails) {
	var findErr error
	var current entity.AnswerDetails
	for _, v := range answers {
		findErr = ctx.session.DB("sofia").C("answer_details").FindId(v.Aid).One(&current)
		if findErr != nil {
			log.Info(findErr)
			answerDetails = append(answerDetails, current)
		} else {
			answerDetails = append(answerDetails, entity.AnswerDetails{})
		}
	}
	return
}

func (q *QaDaoImpl) ParseQuestions(rows *sql.Rows) (result []entity.Questions, err error) {
	var it entity.Questions
	for rows.Next() {
		err = rows.Scan(
			&it.Qid,
			&it.Closed,
			&it.Raiser,
			&it.Category,
			&it.AcceptedAnswer,
			&it.AnswerCount,
			&it.ViewCount,
			&it.FavoriteCount,
			&it.Time)
		if err != nil {
			return
		}
		result = append(result, it)
	}
	return result, nil
}

func (q *QaDaoImpl) ParseAnswers(rows *sql.Rows) (result []entity.Answers, err error) {
	var it entity.Answers
	for rows.Next() {
		err = rows.Scan(
			&it.Aid,
			&it.Answerer,
			&it.Qid,
			&it.CommentCount,
			&it.CriticismCount,
			&it.LikeCount,
			&it.ApprovalCount,
			&it.Time)
		if err != nil {
			return
		}
		result = append(result, it)
	}
	return result, nil
}

func (q *QaDaoImpl) AssignLabels(ctx TransactionContext, questions []entity.Questions) (err error) {
	var rows *sql.Rows
	for _, v := range questions {
		rows, err = ctx.sqlTx.Query("select title from labels natural join(select lid from question_labels where qid=?)as L", v.Qid)
		if err != nil {
			return
		}
		v.Labels = make([]string, 0)
		var current string
		for rows.Next() {
			err = rows.Scan(&current)
			if err != nil {
				return
			}
			v.Labels = append(v.Labels, current)
		}
	}
	return nil
}

func (q *QaDaoImpl) IncQuestionCount(ctx TransactionContext, uid int64) (err error) {
	res, err := ctx.sqlTx.Exec("update users set question_count=question_count+1 where uid=?", uid)
	if err != nil {
		return
	}
	affected, e := res.RowsAffected()
	if e == nil && affected != 1 {
		log.Warn("IncQuestionCount: uid = ", uid, ", but rows affected = ", affected)
	}
	return nil
}

func (q *QaDaoImpl) IncUserAnswerCount(ctx TransactionContext, uid int64) (err error) {
	res, err := ctx.sqlTx.Exec("update users set answer_count=answer_count+1 where uid=?", uid)
	if err != nil {
		return
	}
	affected, e := res.RowsAffected()
	if e == nil && affected != 1 {
		log.Warn("IncAnswerCount: uid = ", uid, ", but rows affected = ", affected)
	}
	return nil
}

func (q *QaDaoImpl) IncCommentCount(ctx TransactionContext, aid int64) (err error) {
	res, err := ctx.sqlTx.Exec("update answers set comment_count=comment_count+1 where aid=?", aid)
	if err != nil {
		return
	}
	affected, e := res.RowsAffected()
	if e == nil && affected != 1 {
		log.Warn("IncCommentCount: aid = ", aid, ", but rows affected = ", affected)
	}
	return nil
}

func (q *QaDaoImpl) IncCriticismCount(ctx TransactionContext, aid int64) (err error) {
	res, err := ctx.sqlTx.Exec("update answers set criticism_count=criticism_count+1 where aid=?", aid)
	if err != nil {
		return
	}
	affected, e := res.RowsAffected()
	if e == nil && affected != 1 {
		log.Warn("IncCriticismCount: aid = ", aid, ", but rows affected = ", affected)
	}
	return nil
}

func (q *QaDaoImpl) CheckQuestionOwner(ctx TransactionContext, qid int64, uid int64) (result bool, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select exists(select*from questions where qid=? and raiser=?)", qid, uid)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return
		}
	} else {
		return false, errors.New("sql error - no rows returned from a `select exists(...)`")
	}
	return
}

func (q *QaDaoImpl) CheckAnswerOwner(ctx TransactionContext, aid int64, uid int64) (result bool, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select exists(select*from answers where aid=? and answerer=?)", aid, uid)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return
		}
	} else {
		return false, errors.New("sql error - no rows returned from a `select exists(...)`")
	}
	return
}

func (q *QaDaoImpl) GetAnswerActionInfos(ctx TransactionContext, uid int64, qids []int64, aids []int64) (infos []AnswerActionInfo, err error) {
	infos = make([]AnswerActionInfo, len(aids))
	var userLabels = make(map[int64]bool)
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select * from user_labels where uid=?", uid)
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		userLabels[id] = true
	}
	for i, v := range aids {
		rows, err = ctx.sqlTx.Query(`select
			exists(select*from like_answers where uid=? and aid=?)as liked,
			exists(select*from approve_answers where uid=? and aid=?)as approved`,
			uid, v,
			uid, v)
		if err != nil {
			return
		}
		if rows.Next() {
			err = rows.Scan(&infos[i].Liked, &infos[i].Approved)
			if err != nil {
				e := rows.Close()
				if e != nil {
					log.Warn(e)
				}
				return
			}
		}
		e := rows.Close()
		if e != nil {
			log.Warn(e)
		}
		rows, err = ctx.sqlTx.Query("select lid from question_labels where qid=?", qids[i])
		if err != nil {
			return
		}
		infos[i].Approvable = true
		for rows.Next() {
			var id int64
			err = rows.Scan(&id)
			_, ok := userLabels[id]
			if !ok {
				infos[i].Approvable = false
				break
			}
		}
		e = rows.Close()
		if e != nil {
			log.Warn(e)
		}
	}
	return infos, nil
}

func (q *QaDaoImpl) MakeLabels(ctx TransactionContext, titles []string) (labels []int64, err error) {
	labels = make([]int64, len(titles))
	var stmt *sql.Stmt
	var insertStmt *sql.Stmt
	stmt, err = ctx.sqlTx.Prepare("select lid from labels where title=?")
	if err != nil {
		return
	}
	insertStmt, err = ctx.sqlTx.Prepare("insert into labels(title)values(?)")
	if err != nil {
		return
	}
	var rows *sql.Rows
	var res sql.Result
	for i, title := range titles {
		rows, err = stmt.Query(title)
		if err != nil {
			return
		}
		if rows.Next() {
			// exists
			err = rows.Scan(&labels[i])
			if err != nil {
				return
			}
		} else {
			// doesn't exist
			res, err = insertStmt.Exec(title)
			if err != nil {
				return
			}
			labels[i], err = res.LastInsertId()
			if err != nil {
				return
			}
		}
	}
	return labels, nil
}

func (q *QaDaoImpl) GetBannedWords(ctx TransactionContext) (words []string, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select word from ban_words")
	if err != nil {
		return
	}
	for rows.Next() {
		var it string
		err = rows.Scan(&it)
		words = append(words, it)
	}
	return words, nil
}

func (q *QaDaoImpl) AddQuestion(ctx TransactionContext, uid int64, title string, content string, category string, labels []string, pictureUrl string, head string) (resultQid int64, err error) {
	// insert into questions
	var res sql.Result
	res, err = ctx.sqlTx.Exec(
		"insert into questions(raiser,category,answer_count,view_count,favorite_count,time,scanned)values(?,?,0,0,0,?,0)",
		uid, category, time.Now().Unix())
	if err != nil {
		return
	}
	var qid int64
	qid, err = res.LastInsertId()
	// mark labels
	if len(labels) > 0 {
		var labelIds []int64
		labelIds, err = q.MakeLabels(ctx, labels)
		var query = "insert into question_labels(qid,lid) values(?,?)"
		var params = append([]interface{}{}, labelIds[0])
		for i := 1; i < len(labels); i++ {
			query += ",(?,?)"
			params = append(params, qid, labelIds[i])
		}
		_, err = ctx.sqlTx.Exec(query, params...)
		if err != nil {
			return
		}
	}
	// insert into question_details
	var questionDetail entity.QuestionDetails
	questionDetail.Qid = qid
	questionDetail.Title = title
	questionDetail.Content = content
	questionDetail.PictureUrl = pictureUrl
	questionDetail.Head = head
	err = ctx.session.DB("sofia").C("question_details").Insert(questionDetail)
	if err != nil {
		return
	}
	return qid, nil
}

func (q *QaDaoImpl) ModifyQuestion(ctx TransactionContext, qid int64, title string, content string, category string, labels []string, pictureUrl string, head string) (err error) {
	// remove old labels
	remove := "delete from question_labels where qid=?"
	_, err = ctx.sqlTx.Exec(remove, qid)
	if err != nil {
		return err
	}
	// mark labels
	if len(labels) > 0 {
		var labelIds []int64
		labelIds, err = q.MakeLabels(ctx, labels)
		var update = "insert into question_labels(qid,lid) values(?,?)"
		var params = append([]interface{}{}, labelIds[0])
		for i := 1; i < len(labels); i++ {
			update += ",(?,?)"
			params = append(params, qid, labelIds[i])
		}
		_, err = ctx.sqlTx.Exec(update, params...)
		if err != nil {
			return
		}
	}
	// modify
	_, err = ctx.sqlTx.Exec("update questions set category=?where qid=?", category, qid)
	if err != nil {
		return
	}
	return ctx.session.DB("sofia").C("question_details").UpdateId(
		qid,
		bson.D{{"$set", bson.D{
			{"content", content},
			{"pictureUrl", pictureUrl},
			{"head", head},
			{"title", title}}}})
}

func (q *QaDaoImpl) AddAnswer(ctx TransactionContext, uid int64, qid int64, content string, pictureUrl string, head string) (aid int64, err error) {
	// insert into answers
	var res sql.Result
	res, err = ctx.sqlTx.Exec(
		"insert into answers(answerer,qid,comment_count,criticism_count,like_count,approval_count,time,scanned)values(?,?,0,0,0,0,?,0)",
		uid, qid, time.Now().Unix())
	if err != nil {
		return
	}
	aid, err = res.LastInsertId()
	// insert into answer_details
	var answerDetail entity.AnswerDetails
	answerDetail.Aid = aid
	answerDetail.Content = content
	answerDetail.PictureUrl = pictureUrl
	answerDetail.Head = head
	err = ctx.session.DB("sofia").C("answer_details").Insert(answerDetail)
	if err != nil {
		return
	}
	return aid, nil
}

func (q *QaDaoImpl) ModifyAnswer(ctx TransactionContext, aid int64, content string, pictureUrl string, head string) (err error) {
	return ctx.session.DB("sofia").C("answer_details").UpdateId(
		aid,
		bson.D{{"$set", bson.D{{"content", content}, {"pictureUrl", pictureUrl}, {"head", head}}}})
}

func (q *QaDaoImpl) ParseComments(rows *sql.Rows) (comments []entity.Comments, err error) {
	var res []entity.Comments
	for rows.Next() {
		var it entity.Comments
		err = rows.Scan(&it.Cmid, &it.Uid, &it.Time)
		if err != nil {
			return
		}
		res = append(res, it)
	}
	return res, nil
}

func (q *QaDaoImpl) ParseCriticisms(rows *sql.Rows) (comments []entity.Criticisms, err error) {
	var res []entity.Criticisms
	for rows.Next() {
		var it entity.Criticisms
		err = rows.Scan(&it.Ctid, &it.Uid, &it.Time)
		if err != nil {
			return
		}
		res = append(res, it)
	}
	return res, nil
}

func (q *QaDaoImpl) GetComments(ctx TransactionContext, aid int64, page int64) (comments []entity.Comments, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query(
		"select cmid,uid,time from comments where aid=? limit ?,?", aid, page*CommentPageSize, CommentPageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	comments, err = q.ParseComments(rows)
	return
}

func (q *QaDaoImpl) GetCriticisms(ctx TransactionContext, aid int64, page int64) (comments []entity.Criticisms, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query(
		"select ctid,uid,time from criticisms where aid=? limit ?,?", aid, page*CommentPageSize, CommentPageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	comments, err = q.ParseCriticisms(rows)
	return
}

func (q *QaDaoImpl) MainPage(ctx TransactionContext, uid int64, category string, page int64) (questions []entity.Questions, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query(
		"select "+questionFields+" from questions natural join(select distinct qid from question_labels where lid in(select lid from user_labels where uid=?))as Q where ?='all' or category=? order by time desc limit ?,?",
		uid, category, category, page*PageSize, PageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	questions, err = q.ParseQuestions(rows)
	if err != nil {
		return
	}
	err = q.AssignLabels(ctx, questions)
	return questions, err
}

func (q *QaDaoImpl) FindQuestionById(ctx TransactionContext, qid int64) (question []entity.Questions, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select "+questionFields+" from questions where qid=?", qid)
	if err != nil {
		return
	}
	defer rows.Close()
	question, err = q.ParseQuestions(rows)
	if err != nil {
		return
	}
	err = q.AssignLabels(ctx, question)
	return question, err
}

func (q *QaDaoImpl) FindAnswerById(ctx TransactionContext, aid int64) (answer []entity.Answers, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select "+questionFields+" from answers where aid=?", aid)
	if err != nil {
		return
	}
	defer rows.Close()
	answer, err = q.ParseAnswers(rows)
	if err != nil {
		return
	}
	return answer, err
}

func (q *QaDaoImpl) SaveQuestionSkeleton(ctx TransactionContext, question entity.Questions) (err error) {
	res, err := ctx.sqlTx.Exec(
		"update questions set category=?,accepted_answer=?,answer_count=?,view_count=?,favorite_count=?,scanned=?where qid=?",
		question.Category, question.AcceptedAnswer, question.AnswerCount, question.ViewCount,
		question.FavoriteCount, question.Scanned, question.Qid)
	if err != nil {
		return
	}
	var affected int64
	affected, err = res.RowsAffected()
	if err == nil && affected != 1 {
		log.Warn("SaveQuestionSkeleton: qid = ", question.Qid, ", but affected = ", affected)
	}
	return nil
}

func (q *QaDaoImpl) SaveAnswerSkeleton(ctx TransactionContext, answer entity.Answers) (err error) {
	res, err := ctx.sqlTx.Exec(
		"update answers set comment_count=?,criticism_count=?,like_count=?,approval_count=?,scanned=?where aid=?",
		answer.CommentCount, answer.CriticismCount, answer.LikeCount, answer.ApprovalCount, answer.Scanned, answer.Aid)
	if err != nil {
		return
	}
	var affected int64
	affected, err = res.RowsAffected()
	if err == nil && affected != 1 {
		log.Warn("SaveAnswerSkeleton: aid = ", answer.Aid, ", but affected = ", affected)
	}
	return nil
}

func (q *QaDaoImpl) FindQuestionAnswers(ctx TransactionContext, qid int64, page int64, sort int8) (answers []entity.Answers, err error) {
	var rows *sql.Rows
	query := "select " + answerFields + " from answers where qid=? order by "
	if sort == 1 {
		query += "like_count desc"
	} else {
		query += "time desc"
	}
	query += " limit ?,?"
	rows, err = ctx.sqlTx.Query(query, qid, page*PageSize, PageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	answers, err = q.ParseAnswers(rows)
	return answers, err
}

func (q *QaDaoImpl) AddComment(ctx TransactionContext, uid int64, aid int64, content string) (cmid int64, err error) {
	// insert into answers
	var res sql.Result
	res, err = ctx.sqlTx.Exec(
		"insert into comments(uid, aid,time)values(?,?,?)",
		uid, aid, time.Now().Unix())
	if err != nil {
		return
	}
	cmid, err = res.LastInsertId()
	// insert into answer_details
	var commentDetail entity.CommentDetails
	commentDetail.Cmid = cmid
	commentDetail.Content = content
	err = ctx.session.DB("sofia").C("comment_details").Insert(commentDetail)
	if err != nil {
		return
	}
	return cmid, nil
}

func (q *QaDaoImpl) AddCriticism(ctx TransactionContext, uid int64, aid int64, content string) (ctid int64, err error) {
	// insert into answers
	var res sql.Result
	res, err = ctx.sqlTx.Exec(
		"insert into criticisms(uid,aid,time)values(?,?,?)",
		uid, aid, time.Now().Unix())
	if err != nil {
		return
	}
	ctid, err = res.LastInsertId()
	// insert into answer_details
	var criticismDetail entity.CriticismDetails
	criticismDetail.Ctid = ctid
	criticismDetail.Content = content
	err = ctx.session.DB("sofia").C("criticism_details").Insert(criticismDetail)
	if err != nil {
		return
	}
	return ctid, nil
}

func (q *QaDaoImpl) DeleteQuestion(ctx TransactionContext, qid int64) (err error) {
	_, err = ctx.sqlTx.Exec("delete from questions where qid=?", qid)
	if err != nil {
		return
	}
	return ctx.session.DB("sofia").C("question_details").RemoveId(qid)
}

func (q *QaDaoImpl) DeleteAnswer(ctx TransactionContext, aid int64) (err error) {
	_, err = ctx.sqlTx.Exec("delete from answers where aid=?", aid)
	if err != nil {
		return
	}
	return ctx.session.DB("sofia").C("answer_details").RemoveId(aid)
}
