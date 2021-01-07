package service

import (
	"search-service/dao"
	"search-service/rpc"
)

type SearchService interface {
	Init(qaDao dao.SearchDao, usersRPC rpc.UsersRPC) (err error)
	Destruct()
	SearchQuestions(token string, page int64, text string) (code int8, result interface{})
	SearchAnswers(token string, page int64, text string) (code int8, result interface{})
	SearchUsers(token string, page int64, text string) (code int8, result interface{})
	HotList(token string) (code int8, result interface{})
	Search(token string, text string) (code int8, result interface{})
}
