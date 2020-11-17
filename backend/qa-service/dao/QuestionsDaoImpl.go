package dao

import (
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
	mongoUrl = "test:test@localhost"
}

func (q *QuestionsDaoImpl) FindByQid(qid string) (question entity.Questions, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return question, err
	}
	defer session.Close()
	c := session.DB("mydb").C("questions")
	err = c.Find(bson.M{"_id": qid}).All(&question)
	return question, err
}

func (q *QuestionsDaoImpl) Insert(question entity.Questions) (qid string, err error) {
	var session *mgo.Session
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return qid, err
	}
	defer session.Close()
	question.Qid = string(bson.NewObjectId())
	c := session.DB("mydb").C("questions")
	err = c.Insert(question)
	return question.Qid, err
}
