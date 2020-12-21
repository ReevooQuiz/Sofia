package test

import (
	"github.com/SKFE396/qa-service/dao"
	"github.com/SKFE396/qa-service/mock"
	"github.com/SKFE396/qa-service/rpc"
	"github.com/SKFE396/qa-service/service"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestServiceInit(t *testing.T) {
	q := service.QaServiceImpl{}
	tests := []struct {
		name string
		qaDao dao.QaDao
		usersRPC rpc.UsersRPC
	}{
		{"Initialize", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = q.Init(tt.qaDao, tt.usersRPC)
		})
	}
}

func TestServiceDestruct(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockQaDao.EXPECT().Init()
	mockQaDao.EXPECT().Destruct()
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	q.Destruct()
}

func TestServiceQuestionListResponse(t *testing.T) {
}