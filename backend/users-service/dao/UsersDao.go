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
	FindFollowByUidAndFollower(ctx TransactionContext, uid int64, follower int64) (follow entity.Follows, err error)
	FindLabelByTitle(ctx TransactionContext, title string) (label entity.Labels, err error)
	FindLabelsByUid(ctx TransactionContext, uid int64) (labels []entity.Labels, err error)
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
