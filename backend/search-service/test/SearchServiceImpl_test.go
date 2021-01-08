package test

import (
	"database/sql"
	"errors"
	"github.com/SKFE396/search-service/dao"
	"github.com/SKFE396/search-service/entity"
	"github.com/SKFE396/search-service/mock"
	"github.com/SKFE396/search-service/rpc"
	"github.com/SKFE396/search-service/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestServiceInit(t *testing.T) {
	s := service.SearchServiceImpl{}
	tests := []struct {
		name     string
		searchDao    dao.SearchDao
		usersRPC rpc.UsersRPC
	}{
		{"Initialize", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = s.Init(tt.searchDao, tt.usersRPC)
		})
	}
}

func TestServiceDestruct(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockSearchDao.EXPECT().Init()
	mockSearchDao.EXPECT().Destruct()
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	s.Destruct()
}

func TestServiceQuestionListResponse(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockSearchDao.EXPECT().Init()
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	tests := []struct {
		name            string
		questions       []entity.Questions
		questionDetails []entity.QuestionDetails
		keywords []string
		mock            bool
		mockIds         []int64
		mockError       error
		mockResult      []rpc.UserInfo
	}{
		{
			"Normal",
			[]entity.Questions{{
				15,
				5,
				"math",
				sql.NullInt64{Valid: false},
				5,
				4,
				7,
				time.Now().Unix(),
				[]string{"gradient"},
				false,
				false,
			}},
			[]entity.QuestionDetails{{
				15,
				"title",
				"What is gradient?",
				"pictureUrl",
				"What is gradient?",
			}},
			[]string {"xxx"},
			true,
			[]int64{5},
			nil,
			[]rpc.UserInfo{{
				"name",
				"nickname",
				"icon data",
			}},
		},
		{
			"RPC Failure",
			[]entity.Questions{},
			[]entity.QuestionDetails{},
			[]string{},
			true,
			[]int64{},
			errors.New("test error"),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockUsersRPC.EXPECT().GetUserInfos(tt.mockIds).Return(tt.mockResult, tt.mockError)
			}
			var questions []service.QuestionListItem
			result, err := s.QuestionListResponse(tt.questions, tt.questionDetails, &tt.keywords)
			a.Equal(err, tt.mockError)
			if tt.mockResult != nil {
				a.NotNil(result)
				questions = result.([]service.QuestionListItem)
				a.Nil(err)
				a.Equal(len(questions), len(tt.mockResult))
				if len(tt.mockResult) > 0 {
					a.Equal(questions[0].Title, tt.questionDetails[0].Title)
					shouldHave := tt.questionDetails[0].PictureUrl != ""
					has := len(questions[0].PictureUrls) > 0
					a.Equal(shouldHave, has)
				}
			} else {
				a.Equal(nil, result)
			}
		})
	}
}

func TestServiceSearchQuestions(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockSearchDao.EXPECT().Init().AnyTimes()
	mockSearchDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	text := "frozen"
	var (
		page int64 = 6
		uid int64 = 234
		role int8 = service.USER
	)
	questions := []entity.Questions {{Qid: 645, Raiser: 23, Category: "life"}}
	details := []entity.QuestionDetails {{Qid: 645, Content: "frozen code", Title: "title", PictureUrl: "pic url", Head: "frozen code "}}
	keywords := []string{"test"}
	userInfos := []rpc.UserInfo{{Name: "tsw", Nickname: "sk", Icon: "icon"}}
	uids := []int64{23}
	err := errors.New("xx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchQuestions(gomock.Any(), page, text).Return(questions, nil)
		mockSearchDao.EXPECT().FindQuestionDetails(gomock.Any(), questions).Return(details)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, result := s.SearchQuestions(token, page, text)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]service.QuestionListItem)[0]
		a.Equal("title", res.Title)
		a.False(res.HasKeywords)
		a.Equal("life", res.Category)
		a.Equal("23", res.Owner.Uid)
		a.Equal("sk", res.Owner.Nickname)
		a.Equal("icon", res.Owner.Icon)
		a.Equal("645", res.Qid)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := s.SearchQuestions(token, page, text)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Search Questions", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchQuestions(gomock.Any(), page, text).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchQuestions(token, page, text)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchQuestions(gomock.Any(), page, text).Return(questions, nil)
		mockSearchDao.EXPECT().FindQuestionDetails(gomock.Any(), questions).Return(details)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchQuestions(token, page, text)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get User Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchQuestions(gomock.Any(), page, text).Return(questions, nil)
		mockSearchDao.EXPECT().FindQuestionDetails(gomock.Any(), questions).Return(details)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchQuestions(token, page, text)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceAnswerListResponse(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockSearchDao.EXPECT().Init().AnyTimes()
	mockSearchDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.SearchServiceImpl
	_ = q.Init(mockSearchDao, mockUsersRPC)
	a := assert.New(t)

	answers := []entity.Answers {{Aid: 56, Answerer: 7, Qid: 234}}
	answerDetails := []entity.AnswerDetails {{Aid: 56, Content: "content", PictureUrl: "pic url", Head: "Head"}}
	keywords := []string{"key"}
	userInfoResult := []rpc.UserInfo{{Name: "sk", Nickname: "nick", Icon: "icon"}}
	var ctx dao.TransactionContext
	var uid int64 = 5

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos([]int64{7}).Return(userInfoResult, nil)
		mockSearchDao.EXPECT().GetAnswerActionInfos(ctx, uid, []int64{234}, []int64{56}).Return([]dao.AnswerActionInfo{{Liked: true, Approved: false, Approvable: true}}, nil)
		result, err := q.AnswerListResponse(ctx, uid, answers, answerDetails, &keywords)
		a.Nil(err)
		res := result.([]service.AnswerListItem)[0]
		a.True(res.Liked)
		a.False(res.Approved)
		a.True(res.Approvable)
		a.Equal("7", res.Owner.Uid)
		a.Equal("pic url", res.PictureUrls[0])
		a.Equal("icon", res.Owner.Icon)
		a.Equal("nick", res.Owner.Nickname)
	})

	t.Run("Failed to Get User Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos([]int64{7}).Return(nil, errors.New("err"))
		_, err := q.AnswerListResponse(ctx, uid, answers, answerDetails, &keywords)
		a.NotNil(err)
	})

	t.Run("Failed to Get Action Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos([]int64{7}).Return(userInfoResult, nil)
		mockSearchDao.EXPECT().GetAnswerActionInfos(ctx, uid, []int64{234}, []int64{56}).Return(nil, errors.New("gg"))
		_, err := q.AnswerListResponse(ctx, uid, answers, answerDetails, &keywords)
		a.NotNil(err)
	})
}

func TestServiceSearchAnswers(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockSearchDao.EXPECT().Init().AnyTimes()
	mockSearchDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	text := "frozen"
	var (
		page int64 = 6
		uid int64 = 234
		role int8 = service.USER
	)
	answers := []entity.Answers {{Aid: 645, Answerer: 23, Qid: 675}}
	details := []entity.AnswerDetails {{Aid: 645, Content: "frozen code", PictureUrl: "pic url", Head: "frozen code "}}
	keywords := []string{"test"}
	userInfos := []rpc.UserInfo{{Name: "tsw", Nickname: "sk", Icon: "icon"}}
	uids := []int64{23}
	aids := []int64{645}
	qids := []int64{675}
	answerActionInfos := []dao.AnswerActionInfo{{Liked: true, Approved: true, Approvable: false}}
	err := errors.New("xx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchAnswers(gomock.Any(), page, text).Return(details, nil)
		mockSearchDao.EXPECT().FindAnswerSkeletons(gomock.Any(), details).Return(answers)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		mockSearchDao.EXPECT().GetAnswerActionInfos(gomock.Any(), uid, qids, aids).Return(answerActionInfos, nil)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, result := s.SearchAnswers(token, page, text)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]service.AnswerListItem)[0]
		a.False(res.HasKeywords)
		a.Equal("23", res.Owner.Uid)
		a.Equal("sk", res.Owner.Nickname)
		a.Equal("icon", res.Owner.Icon)
		a.Equal("645", res.Aid)
		a.Equal("pic url", res.PictureUrls[0])
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := s.SearchAnswers(token, page, text)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Search Answers", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchAnswers(gomock.Any(), page, text).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchAnswers(token, page, text)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchAnswers(gomock.Any(), page, text).Return(details, nil)
		mockSearchDao.EXPECT().FindAnswerSkeletons(gomock.Any(), details).Return(answers)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchAnswers(token, page, text)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get User Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchAnswers(gomock.Any(), page, text).Return(details, nil)
		mockSearchDao.EXPECT().FindAnswerSkeletons(gomock.Any(), details).Return(answers)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchAnswers(token, page, text)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Answer Action Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchAnswers(gomock.Any(), page, text).Return(details, nil)
		mockSearchDao.EXPECT().FindAnswerSkeletons(gomock.Any(), details).Return(answers)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		mockSearchDao.EXPECT().GetAnswerActionInfos(gomock.Any(), uid, qids, aids).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchAnswers(token, page, text)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceSearchUsers(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockSearchDao.EXPECT().Init().AnyTimes()
	mockSearchDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	text := "frozen"
	var (
		page int64 = 6
		uid int64 = 234
		role int8 = service.USER
	)
	searchUserResults := []dao.SearchUserResult{{Uid: 234, Banned: true, Name: "name", Nickname: "nick", Profile: "profile"}}
	err := errors.New("xx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchUsers(gomock.Any(), page, text).Return(searchUserResults, nil)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, result := s.SearchUsers(token, page, text)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]dao.SearchUserResult)[0]
		a.Equal(int64(234), res.Uid)
		a.Equal("nick", res.Nickname)
		a.Equal("profile", res.Profile)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := s.SearchUsers(token, page, text)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Search Users", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().SearchUsers(gomock.Any(), page, text).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.SearchUsers(token, page, text)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceHotList(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockSearchDao.EXPECT().Init().AnyTimes()
	mockSearchDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		uid int64 = 234
		role int8 = service.USER
	)
	questions := []entity.Questions {{Qid: 645, Raiser: 23, Category: "life"}}
	details := []entity.QuestionDetails {{Qid: 645, Content: "frozen code", Title: "title", PictureUrl: "pic url", Head: "frozen code "}}
	keywords := []string{"test"}
	userInfos := []rpc.UserInfo{{Name: "tsw", Nickname: "sk", Icon: "icon"}}
	uids := []int64{23}
	err := errors.New("xx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().HotList(gomock.Any()).Return(questions, nil)
		mockSearchDao.EXPECT().FindQuestionDetails(gomock.Any(), questions).Return(details)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, result := s.HotList(token)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]service.QuestionListItem)[0]
		a.Equal("title", res.Title)
		a.False(res.HasKeywords)
		a.Equal("life", res.Category)
		a.Equal("23", res.Owner.Uid)
		a.Equal("sk", res.Owner.Nickname)
		a.Equal("icon", res.Owner.Icon)
		a.Equal("645", res.Qid)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := s.HotList(token)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Get Hot List", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().HotList(gomock.Any()).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.HotList(token)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().HotList(gomock.Any()).Return(questions, nil)
		mockSearchDao.EXPECT().FindQuestionDetails(gomock.Any(), questions).Return(details)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.HotList(token)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get User Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().HotList(gomock.Any()).Return(questions, nil)
		mockSearchDao.EXPECT().FindQuestionDetails(gomock.Any(), questions).Return(details)
		mockSearchDao.EXPECT().GetBannedWords(gomock.Any()).Return(keywords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.HotList(token)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceSearch(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchDao := mock.NewMockSearchDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockSearchDao.EXPECT().Init().AnyTimes()
	mockSearchDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var s service.SearchServiceImpl
	_ = s.Init(mockSearchDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	text := "test"
	var (
		uid int64 = 234
		role int8 = service.USER
	)
	cards := []dao.KListItem{{Title: "title"}}
	cards[0].Attr.Name = "name"
	cards[0].Attr.Origin = "123"
	err := errors.New("xx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().Search(gomock.Any(), text).Return(cards, nil)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, result := s.Search(token, text)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]dao.KListItem)[0]
		a.Equal("title", res.Title)
		a.Equal("name", res.Attr.Name)
		a.Equal("123", res.Attr.Origin)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := s.Search(token, text)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Search", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockSearchDao.EXPECT().Search(gomock.Any(), text).Return(nil, err)
		mockSearchDao.EXPECT().Rollback(gomock.Any())
		code, _ := s.Search(token, text)
		a.Equal(int8(service.Failed), code)
	})
}
