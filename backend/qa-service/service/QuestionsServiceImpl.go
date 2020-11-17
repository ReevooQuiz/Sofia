package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2/bson"
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

func (q *QuestionsServiceImpl) FindByQid(qid bson.ObjectId) (question entity.Questions, err error) {
	return q.questionsDao.FindByQid(qid)
}

func (q *QuestionsServiceImpl) Insert(question entity.Questions) (qid bson.ObjectId, err error) {
	return q.questionsDao.Insert(question)
}
