package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
)

type LabelsService interface {
	Init(labelsDao ...dao.LabelsDao) (err error)
	Destruct()

	FindByTitle(title string) (label entity.Labels, err error)
	Insert(label entity.Labels) (lid int64, err error)
}
