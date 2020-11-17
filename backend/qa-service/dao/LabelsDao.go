package dao

import "github.com/zhanghanchong/qa-service/entity"

type LabelsDao interface {
	Init() (err error)
	Destruct()

	FindByTitle(title string) (label entity.Labels, err error)
	Insert(label entity.Labels) (lid int64, err error)
}
