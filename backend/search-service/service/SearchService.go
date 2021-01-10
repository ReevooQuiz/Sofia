package service

import (
	"github.com/SKFE396/search-service/dao"
	"github.com/SKFE396/search-service/rpc"
)

type SearchService interface {
	Init(qaDao dao.SearchDao, usersRPC rpc.UsersRPC) (err error)
	Destruct()
	SearchQuestions(page int64, text string) (code int8, result interface{})
	SearchAnswers(page int64, text string) (code int8, result interface{})
	SearchUsers(page int64, text string) (code int8, result interface{})
	HotList() (code int8, result interface{})
	Search(text string) (code int8, result interface{})
}
