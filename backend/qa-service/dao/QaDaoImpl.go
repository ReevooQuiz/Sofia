package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type QaDaoImpl struct {
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

func (q *QaDaoImpl) Init() (err error) {
	q.db, err = sql.Open("mysql", mysqlUrl)
	return err
}

func (q *QaDaoImpl) Destruct() {
	_ = q.db.Close()
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
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return question, err
	}
	defer session.Close()
	c := session.DB("sofia").C("questions")
	err = c.Find(bson.M{"_id": qid}).All(&question)
	return question, err
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
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return qid, err
	}
	defer session.Close()
	question.Qid = bson.NewObjectId()
	c := session.DB("sofia").C("questions")
	err = c.Insert(question)
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
