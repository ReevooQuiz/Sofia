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
	PageSize = 5
	answerFields = "aid,answerer,qid,view_count,comment_count,criticism_count,like_count,approval_count,time,scanned"
	questionFields = "qid,raiser,title,category,accepted_answer,answer_count,view_count,favorite_count,time,scanned"
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
	sqlTx *sql.Tx
	session *mgo.Session
}

type AnswerActionInfo struct {
	Liked      bool
	Approved   bool
	Approvable bool
}

func (q *QaDaoImpl) Commit(t *TransactionContext) (err error) {
	t.session.Close()
	return t.sqlTx.Commit()
}

func (q *QaDaoImpl) Rollback(t *TransactionContext) (err error) {
	t.session.Close()
	return t.sqlTx.Rollback()
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
	return TransactionContext{tx, q.session.New()}, nil
}

/*func (q *QaDaoImpl) FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error) {
	err = q.session.DB("sofia").C("answers").Find(bson.M{"qid": qid}).All(&answers)
	return answers, err
}

func (q *QaDaoImpl) FindLabelByLid(lid int64) (label entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = q.db.Prepare("select * from labels where lid = ?")
	if err != nil {
		return label, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(lid).Scan(&label.Lid, &label.Title)
	return label, err
}

func (q *QaDaoImpl) FindLabelByTitle(title string) (label entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = q.db.Prepare("select * from labels where title = ?")
	if err != nil {
		return label, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(title).Scan(&label.Lid, &label.Title)
	return label, err
}

func (q *QaDaoImpl) FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error) {
	var res []entity.Questions
	err = q.session.DB("sofia").C("questions").Find(bson.M{"_id": qid}).All(&res)
	if err != nil {
		return question, err
	}
	if len(res) == 0 {
		return question, errors.New("mongo: no rows in result set")
	}
	return res[0], err
}

func (q *QaDaoImpl) FindQuestionLabelsByQid(qid string) (questionLabels []entity.QuestionLabels, err error) {
	var stmt *sql.Stmt
	stmt, err = q.db.Prepare("select * from question_labels where qid = ?")
	if err != nil {
		return questionLabels, err
	}
	defer stmt.Close()
	var res *sql.Rows
	res, err = stmt.Query(qid)
	if err != nil {
		return questionLabels, err
	}
	for res.Next() {
		var questionLabel entity.QuestionLabels
		err = res.Scan(&questionLabel.Qid, &questionLabel.Lid)
		if err != nil {
			return questionLabels, err
		}
		questionLabels = append(questionLabels, questionLabel)
	}
	return questionLabels, err
}

func (q *QaDaoImpl) InsertKcard(kcard entity.Kcards) (kid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = q.db.Prepare("insert into kcards(title) values(?)")
	if err != nil {
		return kid, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(kcard.Title)
	if err != nil {
		return kid, err
	}
	kid, err = res.LastInsertId()
	return kid, err
}

func (q *QaDaoImpl) InsertLabel(label entity.Labels) (lid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = q.db.Prepare("insert into labels(title) values(?)")
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

func (q *QaDaoImpl) InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error) {
	question.Qid = bson.NewObjectId()
	err = q.session.DB("sofia").C("questions").Insert(question)
	return question.Qid, err
}

func (q *QaDaoImpl) InsertQuestionLabel(questionLabel entity.QuestionLabels) (err error) {
	var stmt *sql.Stmt
	stmt, err = q.db.Prepare("insert into question_labels values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(questionLabel.Qid, questionLabel.Lid)
	return err
}*/

func (q *QaDaoImpl) FindQuestionDetails(ctx TransactionContext, questions []entity.Questions) (questionDetails []entity.QuestionDetails) {
	var findErr error
	var current entity.QuestionDetails
	for _, v := range questions {
		findErr = ctx.session.DB("sofia").C("question_details").FindId(v.Qid).One(current)
		if findErr != nil {
			log.Info(findErr)
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
		findErr = ctx.session.DB("sofia").C("answer_details").FindId(v.Aid).One(current)
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
			&it.ViewCount,
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
	var userLabels map[int64]bool = make(map[int64]bool)
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

func (q *QaDaoImpl) GetBannedWords(ctx TransactionContext, ) (words []string, err error) {
	var result struct {
		UselessId int64    `bson:"_id"`
		Words     []string `bson:"words"`
	}
	err = ctx.session.DB("sofia").C("banned_words").FindId(0).One(result)
	if err != nil {
		return
	}
	return result.Words, nil
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
	_, err = ctx.sqlTx.Exec("update questions set title=?,category=?where qid=?", title, category, qid)
	if err != nil {
		return
	}
	return ctx.session.DB("sofia").C("question_details").UpdateId(
		qid,
		bson.D{{"$set", bson.D{{"content", content}, {"pictureUrl", pictureUrl}, {"head", head}}}})
}

func (q *QaDaoImpl) AddAnswer(ctx TransactionContext, uid int64, qid int64, content string, pictureUrl string, head string) (aid int64, err error) {
	// insert into answers
	var res sql.Result
	res, err = ctx.sqlTx.Exec(
		"insert into answers(answerer,qid,comment_count,criticism_count,like_count,approval_count,view_count,time,scanned)values(?,?,0,0,0,0,0,?,0)",
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

func (q *QaDaoImpl) MainPage(ctx TransactionContext, uid int64, page int64) (questions []entity.Questions, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query(
		"select " + questionFields + " from questions natural join(select distinct qid from question_labels where lid in(select lid from user_labels where uid=?))as Q order by time desc limit ?,?",
		uid, page*PageSize, PageSize)
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
	rows, err = ctx.sqlTx.Query("select " + questionFields + " from questions where qid=?", qid)
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
	rows, err = ctx.sqlTx.Query("select " + questionFields + " from answers where aid=?", aid)
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
		"update questions set title=?,category=?,accepted_answer=?,answer_count=?,view_count=?,favorite_count=?,scanned=?",
		question.Title, question.Category, question.AcceptedAnswer, question.AnswerCount, question.ViewCount, question.FavoriteCount, question.Scanned)
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
		"update answers set view_count=?,comment_count_count=?,criticism_count=?like_count=?,approval_count=?,scanned=?",
		answer.ViewCount, answer.CommentCount, answer.CriticismCount, answer.LikeCount, answer.ApprovalCount, answer.Scanned)
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
	rows, err = ctx.sqlTx.Query(query, qid, page * PageSize, PageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	answers, err = q.ParseAnswers(rows)
	return answers, err
}
