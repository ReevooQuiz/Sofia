package service

import (
	"github.com/zhanghanchong/qa-service/dao"
	"github.com/zhanghanchong/qa-service/entity"
)

type LabelsServiceImpl struct {
	labelsDao dao.LabelsDao
}

func (l *LabelsServiceImpl) Init(labelsDao ...dao.LabelsDao) (err error) {
	if len(labelsDao) == 0 {
		labelsDao = append(labelsDao, &dao.LabelsDaoImpl{})
	}
	l.labelsDao = labelsDao[0]
	return l.labelsDao.Init()
}

func (l *LabelsServiceImpl) Destruct() {
	l.labelsDao.Destruct()
}

func (l *LabelsServiceImpl) FindByTitle(title string) (label entity.Labels, err error) {
	return l.labelsDao.FindByTitle(title)
}

func (l *LabelsServiceImpl) Insert(label entity.Labels) (lid int64, err error) {
	return l.labelsDao.Insert(label)
}
