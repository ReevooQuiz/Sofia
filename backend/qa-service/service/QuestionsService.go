package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
)

type QuestionsService interface {
	Init(questionsDao ...dao.QuestionsDao)

	FindByQid(qid string) (question entity.Questions, err error)
}
