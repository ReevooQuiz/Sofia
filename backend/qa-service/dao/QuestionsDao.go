package dao

import "github.com/zhanghanchong/qa-service/entity"

type QuestionsDao interface {
	FindByQid(qid string) (question entity.Questions, err error)
}
