package dao

import (
	"github.com/SKFE396/qa-service/entity"
)

type QaDao interface {
	Init() (err error)
	Destruct()
	Begin(read bool) (ctx TransactionContext, err error)
	Commit(ctx *TransactionContext) (err error)
	Rollback(ctx *TransactionContext) (err error)

	FindDetails(ctx TransactionContext, questions []entity.Questions) (questionDetails []entity.QuestionDetails)

	/*FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error)
	FindLabelByLid(lid int64) (label entity.Labels, err error)
	FindLabelByTitle(title string) (label entity.Labels, err error)
	FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error)
	FindQuestionLabelsByQid(qid string) (questionLabels []entity.QuestionLabels, err error)
	InsertKcard(kcard entity.Kcards) (kid int64, err error)
	InsertLabel(label entity.Labels) (lid int64, err error)
	InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error)
	InsertQuestionLabel(questionLabel entity.QuestionLabels) (err error)*/

	CheckQuestionOwner(ctx TransactionContext, qid int64, uid int64) (result bool, err error)

	MakeLabels(ctx TransactionContext, titles []string) (labels []int64, err error)
	GetBannedWords(ctx TransactionContext, ) (words []string, err error)
	AddQuestion(ctx TransactionContext, uid int64, title string, content string, category string, labels []string, pictureUrl string, head string) (qid int64, err error)
	ModifyQuestion(ctx TransactionContext, qid int64, title string, content string, category string, labels []string, pictureUrl string, head string) (err error)

	MainPage(ctx TransactionContext, uid int64, page int64) (questions []entity.Questions, err error)
}
