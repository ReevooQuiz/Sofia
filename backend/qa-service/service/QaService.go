package service

import (
	"github.com/SKFE396/qa-service/dao"
	"github.com/SKFE396/qa-service/rpc"
)

type QaService interface {
	Init(qaDao dao.QaDao, usersRPC rpc.UsersRPC) (err error)
	Destruct()

	/*FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error)
	FindLabelByLid(lid int64) (label entity.Labels, err error)
	FindLabelByTitle(title string) (label entity.Labels, err error)
	FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error)
	FindQuestionLabelsByQid(qid string) (questionLabels []entity.QuestionLabels, err error)
	InsertKcard(kcard entity.Kcards) (kid int64, err error)
	InsertLabel(label entity.Labels) (lid int64, err error)
	InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error)
	InsertQuestionLabel(questionLabel entity.QuestionLabels) (err error)*/

	AddQuestion(token string, req ReqQuestionsPost) (int8, interface{})
	ModifyQuestion(token string, req ReqQuestionsPut) (int8, interface{})
	AddAnswer(token string, req ReqAnswersPost) (int8, interface{})
	ModifyAnswer(token string, req ReqAnswersPut) (int8, interface{})

	MainPage(token string, page int64) (int8, interface{})
	QuestionDetail(token string, qid int64) (int8, interface{})
	ListAnswers(token string, qid int64, page int64, sort int8) (int8, interface{})
	AnswerDetail(token string, aid int64) (int8, interface{})

	GetComments(token string, aid int64, page int64) (int8, interface{})
	AddComment(token string, req ReqCommentsPost) (int8, interface{})

	GetCriticisms(token string, aid int64, page int64) (int8, interface{})
	AddCriticism(token string, req ReqCriticismsPost) (int8, interface{})
}
