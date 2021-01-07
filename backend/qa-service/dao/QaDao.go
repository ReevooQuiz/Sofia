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

	FindQuestionDetails(ctx TransactionContext, questions []entity.Questions) (questionDetails []entity.QuestionDetails)
	FindAnswerDetails(ctx TransactionContext, answers []entity.Answers) (answerDetails []entity.AnswerDetails)

	FindQuestionById(ctx TransactionContext, qid int64) (question []entity.Questions, err error)
	FindAnswerById(ctx TransactionContext, aid int64) (answer []entity.Answers, err error)
	SaveQuestionSkeleton(ctx TransactionContext, question entity.Questions) (err error)
	SaveAnswerSkeleton(ctx TransactionContext, answer entity.Answers) (err error)

	/*FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error)
	FindLabelByLid(lid int64) (label entity.Labels, err error)
	FindLabelByTitle(title string) (label entity.Labels, err error)
	FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error)
	FindQuestionLabelsByQid(qid string) (questionLabels []entity.QuestionLabels, err error)
	InsertKcard(kcard entity.Kcards) (kid int64, err error)
	InsertLabel(label entity.Labels) (lid int64, err error)
	InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error)
	InsertQuestionLabel(questionLabel entity.QuestionLabels) (err error)*/

	IncQuestionCount(ctx TransactionContext, uid int64) (err error)
	IncUserAnswerCount(ctx TransactionContext, uid int64) (err error)

	CheckQuestionOwner(ctx TransactionContext, qid int64, uid int64) (result bool, err error)
	CheckAnswerOwner(ctx TransactionContext, aid int64, uid int64) (result bool, err error)

	GetAnswerActionInfos(ctx TransactionContext, uid int64, qids []int64, aids []int64) (infos []AnswerActionInfo, err error)

	MakeLabels(ctx TransactionContext, titles []string) (labels []int64, err error)
	GetBannedWords(ctx TransactionContext) (words []string, err error)
	AddQuestion(ctx TransactionContext, uid int64, title string, content string, category string, labels []string, pictureUrl string, head string) (qid int64, err error)
	ModifyQuestion(ctx TransactionContext, qid int64, title string, content string, category string, labels []string, pictureUrl string, head string) (err error)

	AddAnswer(ctx TransactionContext, uid int64, qid int64, content string, pictureUrl string, head string) (aid int64, err error)
	ModifyAnswer(ctx TransactionContext, aid int64, content string, pictureUrl string, head string) (err error)

	MainPage(ctx TransactionContext, uid int64, page int64) (questions []entity.Questions, err error)
	FindQuestionAnswers(ctx TransactionContext, qid int64, page int64, sort int8) (answers []entity.Answers, err error)

	GetComments(ctx TransactionContext, aid int64, page int64) (comments []entity.Comments, err error)
	FindCommentDetails(ctx TransactionContext, comments []entity.Comments) (details []entity.CommentDetails)
	AddComment(ctx TransactionContext, uid int64, aid int64, content string) (cmid int64, err error)
}
