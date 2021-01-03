package dao

import (
	"github.com/zhanghanchong/users-service/entity"
)

type UsersDao interface {
	Init() (err error)
	Destruct()
	Begin(read bool) (ctx TransactionContext, err error)
	Commit(ctx *TransactionContext) (err error)
	Rollback(ctx *TransactionContext) (err error)
	FindAnswerByAid(ctx TransactionContext, aid int64) (answer entity.Answers, err error)
	FindAnswerDetailByAid(ctx TransactionContext, aid int64) (answerDetail entity.AnswerDetails, err error)
	FindApproveAnswerByUidAndAid(ctx TransactionContext, uid int64, aid int64) (approveAnswer entity.ApproveAnswers, err error)
	FindCommentByCmid(ctx TransactionContext, cmid int64) (comment entity.Comments, err error)
	FindCriticismByCtid(ctx TransactionContext, ctid int64) (criticism entity.Criticisms, err error)
	FindFollowByUidAndFollower(ctx TransactionContext, uid int64, follower int64) (follow entity.Follows, err error)
	FindFollowsByFollower(ctx TransactionContext, follower int64) (follows []entity.Follows, err error)
	FindFollowsByUid(ctx TransactionContext, uid int64) (follows []entity.Follows, err error)
	FindLabelByTitle(ctx TransactionContext, title string) (label entity.Labels, err error)
	FindLabelsByQid(ctx TransactionContext, qid int64) (labels []entity.Labels, err error)
	FindLabelsByUid(ctx TransactionContext, uid int64) (labels []entity.Labels, err error)
	FindLikeAnswerByUidAndAid(ctx TransactionContext, uid int64, aid int64) (likeAnswer entity.LikeAnswers, err error)
	FindNotificationsByUidPageable(ctx TransactionContext, uid int64, pageable Pageable) (notifications []Notifications, err error)
	FindQuestionDetailByQid(ctx TransactionContext, qid int64) (questionDetail entity.QuestionDetails, err error)
	FindQuestionsByRaiserOrderByTimeDescPageable(ctx TransactionContext, raiser int64, pageable Pageable) (questions []entity.Questions, err error)
	FindUserByEmail(ctx TransactionContext, email string) (user entity.Users, err error)
	FindUserByName(ctx TransactionContext, name string) (user entity.Users, err error)
	FindUserByOidAndAccountType(ctx TransactionContext, oid string, accountType int8) (user entity.Users, err error)
	FindUserByUid(ctx TransactionContext, uid int64) (user entity.Users, err error)
	FindUserDetailByUid(ctx TransactionContext, uid int64) (userDetail entity.UserDetails, err error)
	InsertFavorite(ctx TransactionContext, favorite entity.Favorites) (fid int64, err error)
	InsertFollow(ctx TransactionContext, follow entity.Follows) (err error)
	InsertLabel(ctx TransactionContext, label entity.Labels) (lid int64, err error)
	InsertUser(ctx TransactionContext, user entity.Users) (uid int64, err error)
	InsertUserDetail(ctx TransactionContext, userDetail entity.UserDetails) (err error)
	InsertUserLabel(ctx TransactionContext, userLabel entity.UserLabels) (err error)
	RemoveFollowByUidAndFollower(ctx TransactionContext, uid int64, follower int64) (err error)
	RemoveUserLabelsByUid(ctx TransactionContext, uid int64) (err error)
	UpdateUserByUid(ctx TransactionContext, user entity.Users) (err error)
	UpdateUserDetailByUid(ctx TransactionContext, userDetail entity.UserDetails) (err error)
}
