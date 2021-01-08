package dao

import "github.com/SKFE396/search-service/entity"

type SearchDao interface {
	Init() (err error)
	Destruct()
	Begin(read bool) (ctx TransactionContext, err error)
	Commit(ctx *TransactionContext)
	Rollback(ctx *TransactionContext)
	FindQuestionDetails(ctx TransactionContext, questions []entity.Questions) (questionDetails []entity.QuestionDetails)
	SearchQuestions(ctx TransactionContext, page int64, text string) (questions []entity.Questions, err error)
	GetBannedWords(ctx TransactionContext) (words []string, err error)

	SearchAnswers(ctx TransactionContext, page int64, text string) (details []entity.AnswerDetails, err error)
	FindAnswerSkeletons(ctx TransactionContext, details []entity.AnswerDetails) (answers []entity.Answers)

	GetAnswerActionInfos(ctx TransactionContext, uid int64, qids []int64, aids []int64) (infos []AnswerActionInfo, err error)

	SearchUsers(ctx TransactionContext, page int64, text string) (result []SearchUserResult, err error)
	HotList(ctx TransactionContext) (questions []entity.Questions, err error)

	Search(ctx TransactionContext, text string) (result []KListItem, err error)
}
