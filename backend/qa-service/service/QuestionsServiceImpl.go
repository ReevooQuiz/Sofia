package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
)

type QuestionsServiceImpl struct {
	questionsDao dao.QuestionsDao
}

func (q *QuestionsServiceImpl) Init(questionsDao ...dao.QuestionsDao) {
	if len(questionsDao) == 0 {
		questionsDao = append(questionsDao, &dao.QuestionsDaoImpl{})
	}
	q.questionsDao = questionsDao[0]
}

func (q *QuestionsServiceImpl) FindByQid(qid string) (question entity.Questions, err error) {
	return q.questionsDao.FindByQid(qid)
}
