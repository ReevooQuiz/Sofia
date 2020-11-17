package dao

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type QuestionsDaoImpl struct {
}

var mongoUrl string

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	mongoUrl = os.Getenv("MONGO_URL")
}

func (q *QuestionsDaoImpl) FindByQid(qid string) (question entity.Questions, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return question, err
	}
	defer session.Close()
	c := session.DB("sofia").C("questions")
	err = c.Find(bson.M{"qid": qid}).All(&question)
	return question, err
}

func (q *QuestionsDaoImpl) Insert(question entity.Questions) (qid bson.ObjectId, err error) {
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
