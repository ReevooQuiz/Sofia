package dao

import (
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2/bson"
)

type QuestionsDao interface {
	FindByQid(qid string) (question entity.Questions, err error)
	Insert(question entity.Questions) (qid bson.ObjectId, err error)
}
