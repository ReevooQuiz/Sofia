package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

var (
	mongoUrl string
	mysqlUrl string
)

type QaDaoImpl struct {
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

func (q *QaDaoImpl) FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error) {
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
}
