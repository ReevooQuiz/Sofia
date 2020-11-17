package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2/bson"
)

type QuestionsService interface {
	Init(questionsDao ...dao.QuestionsDao)

	FindByQid(qid string) (question entity.Questions, err error)
	Insert(question entity.Questions) (qid bson.ObjectId, err error)
}
