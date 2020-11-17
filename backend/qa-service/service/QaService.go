package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2/bson"
)

type QaService interface {
	Init(qaDao ...dao.QaDao) (err error)
	Destruct()

	FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error)
	FindLabelByLid(lid int64) (label entity.Labels, err error)
	FindLabelByTitle(title string) (label entity.Labels, err error)
	FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error)
	FindQuestionLabelsByQid(qid string) (questionLabels []entity.QuestionLabels, err error)
	InsertLabel(label entity.Labels) (lid int64, err error)
	InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error)
	InsertQuestionLabel(questionLabel entity.QuestionLabels) (err error)
}
