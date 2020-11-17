package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
	"gopkg.in/mgo.v2/bson"
)

type QaServiceImpl struct {
	qaDao dao.QaDao
}

func (q *QaServiceImpl) Init(qaDao ...dao.QaDao) (err error) {
	if len(qaDao) == 0 {
		qaDao = append(qaDao, &dao.QaDaoImpl{})
	}
	q.qaDao = qaDao[0]
	return q.qaDao.Init()
}

func (q *QaServiceImpl) Destruct() {
	q.qaDao.Destruct()
}

func (q *QaServiceImpl) FindLabelByTitle(title string) (label entity.Labels, err error) {
	return q.qaDao.FindLabelByTitle(title)
}

func (q *QaServiceImpl) FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error) {
	return q.qaDao.FindQuestionByQid(qid)
}

func (q *QaServiceImpl) InsertLabel(label entity.Labels) (lid int64, err error) {
	return q.qaDao.InsertLabel(label)
}

func (q *QaServiceImpl) InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error) {
	return q.qaDao.InsertQuestion(question)
}
