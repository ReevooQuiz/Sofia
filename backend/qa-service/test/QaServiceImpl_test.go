package test

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/SKFE396/qa-service/dao"
	"github.com/SKFE396/qa-service/entity"
	"github.com/SKFE396/qa-service/mock"
	"github.com/SKFE396/qa-service/rpc"
	"github.com/SKFE396/qa-service/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestServiceInit(t *testing.T) {
	q := service.QaServiceImpl{}
	tests := []struct {
		name     string
		qaDao    dao.QaDao
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
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockQaDao.EXPECT().Init()
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	tests := []struct {
		name            string
		uid             int64
		role            int8
		questions       []entity.Questions
		questionDetails []entity.QuestionDetails
		keywords        []string
		mock            bool
		mockIds         []int64
		mockError       error
		mockResult      []rpc.UserInfo
	}{
		{
			"Normal",
			456,
			service.USER,
			[]entity.Questions{{
				15,
				5,
				"title",
				"math",
				sql.NullInt64{Valid: false},
				5,
				4,
				7,
				time.Now().Unix(),
				[]string{"gradient"},
				false,
			}},
			[]entity.QuestionDetails{{
				15,
				"What is gradient?",
				"pictureUrl",
				"What is gradient?",
			}},
			[]string{"xxx"},
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
			456,
			service.USER,
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
			result, err := q.QuestionListResponse(tt.uid, tt.role, tt.questions, tt.questionDetails, &tt.keywords)
			a.Equal(err, tt.mockError)
			if tt.mockResult != nil {
				a.NotNil(result)
				questions = result.([]service.QuestionListItem)
				a.Nil(err)
				a.Equal(len(questions), len(tt.mockResult))
				if len(tt.mockResult) > 0 {
					a.Equal(questions[0].Title, tt.questions[0].Title)
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

func TestServiceMatchKeywords(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	tests := []struct {
		name  string
		text  string
		words []string
		want  bool
	}{
		{"Match", "aaabbbcccdefggg", []string{"bcc", "asd"}, true},
		{"Mismatch", "aaabbbcccdefhhh", []string{"aaaa", "cbb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a.Equal(tt.want, service.MatchKeywords(&tt.text, &tt.words))
		})
	}
}

func TestServiceFindTextAndPicture(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	tests := []struct {
		name         string
		html         string
		words        []string
		wantPicture  string
		wantHead     []rune
		foundPicture bool
		foundHead    bool
		hasKeywords  bool
	}{
		{
			"Normal",
			"<div><p>This is a sample.</p><p>Let's begin.</p><img src=\"testUrl\" /></div>",
			[]string{"sampler"},
			"testUrl",
			[]rune("This is a sample. Let's begin."),
			true,
			false,
			false,
		},
		{
			"No Picture",
			"<div><p>No pictures</p></div>",
			[]string{"sample"},
			"",
			[]rune{},
			false,
			false,
			false,
		},
		{
			"Has keywords",
			"<div><p>This is a keyWord</p></div>",
			[]string{"Keyword"},
			"",
			[]rune{},
			false,
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var picture string
			var head []rune
			node, err := html.Parse(strings.NewReader(tt.html))
			a.Nil(err)
			foundPicture, foundHead, hasKeywords := service.FindTextAndPicture(&tt.words, &picture, &head, node, true, true)
			a.Equal(tt.foundPicture, foundPicture)
			a.Equal(tt.foundHead, foundHead)
			a.Equal(tt.hasKeywords, hasKeywords)
			if tt.foundPicture {
				a.Equal(tt.wantPicture, picture)
			}
			if tt.foundHead {
				a.Equal(tt.wantHead, head)
			}
		})
	}
}

func TestServiceParseContent(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	var q service.QaServiceImpl
	tests := []struct {
		name        string
		content     string
		words       []string
		pictureUrl  string
		head        string
		hasKeywords bool
	}{
		{
			"Normal",
			"Where the north wind meets the sea",
			[]string{"river"},
			"",
			"Where the north wind meets the sea ",
			false,
		},
		{
			"Has Keyword",
			"There's a river full of memory",
			[]string{"river"},
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pictureUrl, head, hasKeyword := q.ParseContent(&tt.content, &tt.words)
			a.Equal(tt.pictureUrl, pictureUrl)
			a.Equal(tt.head, head)
			a.Equal(tt.hasKeywords, hasKeyword)
		})
	}
}

func TestServiceAddQuestion(t *testing.T) {
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	var (
		BannedWords = []string{"magic"}
	)
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockQaDao.EXPECT().Init()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	longContent := string(make([]byte, service.QuestionContentLengthMax+1))

	tests := []struct {
		name                 string
		token                string
		mockToken            bool
		mockSuc              bool
		mockUid              int64
		mockRole             int8
		mockBannedWords      bool
		mockBannedResult     []string
		mockBannedErr        error
		mockAddQuestion      bool
		mockQid              int64
		mockAddErr           error
		mockIncQuestionCount bool
		incQuestionCountErr  error
		req                  service.ReqQuestionsPost
		expectedCode         int8
		expectedResult       interface{}
		rollback             bool
		commit               bool
	}{
		{
			name:                 "Normal",
			token:                "token",
			mockToken:            true,
			mockSuc:              true,
			mockUid:              1,
			mockRole:             service.ADMIN,
			mockBannedWords:      true,
			mockBannedResult:     BannedWords,
			mockBannedErr:        nil,
			mockAddQuestion:      true,
			mockIncQuestionCount: true,
			mockAddErr:           nil,
			mockQid:              456,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"food"},
			},
			expectedCode:   service.Succeeded,
			expectedResult: map[string]string{"qid": "456"},
			rollback:       false,
			commit:         true,
		},
		{
			name:            "Failed token",
			token:           "token",
			mockToken:       true,
			mockSuc:         false,
			mockUid:         0,
			mockRole:        0,
			mockBannedWords: false,
			mockAddQuestion: false,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{},
			},
			expectedCode:   service.Expired,
			expectedResult: nil,
			rollback:       false,
			commit:         false,
		},
		{
			name:            "Long Title",
			token:           "token",
			mockToken:       true,
			mockSuc:         true,
			mockUid:         0,
			mockRole:        service.ADMIN,
			mockBannedWords: false,
			mockAddQuestion: false,
			req: service.ReqQuestionsPost{
				Title:    string(make([]byte, service.QuestionTitleLengthMax+1)),
				Content:  "this is the content",
				Category: "life",
				Labels:   []string{},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:            "Excessive Labels",
			token:           "token",
			mockToken:       true,
			mockSuc:         true,
			mockUid:         0,
			mockRole:        service.ADMIN,
			mockBannedWords: false,
			mockAddQuestion: false,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "this is the content",
				Category: "life",
				Labels:   []string{"1", "2", "3", "4", "5", "6"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:            "Long Label",
			token:           "token",
			mockToken:       true,
			mockSuc:         true,
			mockUid:         0,
			mockRole:        service.ADMIN,
			mockBannedWords: false,
			mockAddQuestion: false,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "this is the content",
				Category: "life",
				Labels:   []string{"short label", "123456789012345678901234567890123"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:            "Long Content",
			token:           "token",
			mockToken:       true,
			mockSuc:         true,
			mockUid:         0,
			mockRole:        service.ADMIN,
			mockBannedWords: false,
			mockAddQuestion: false,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  longContent,
				Category: "life",
				Labels:   []string{"math"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:             "Failed to Get Banned Words",
			token:            "token",
			mockToken:        true,
			mockSuc:          true,
			mockUid:          0,
			mockRole:         service.ADMIN,
			mockBannedWords:  true,
			mockBannedErr:    errors.New("rpc error"),
			mockBannedResult: nil,
			mockAddQuestion:  false,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       true,
			commit:         false,
		},
		{
			name:             "Has Keywords",
			token:            "token",
			mockToken:        true,
			mockSuc:          true,
			mockUid:          0,
			mockRole:         service.ADMIN,
			mockBannedWords:  true,
			mockBannedResult: BannedWords,
			mockBannedErr:    nil,
			mockAddQuestion:  false,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "And in her song, all magic flows",
				Category: "life",
				Labels:   []string{},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": HasKeyword},
			rollback:       true,
			commit:         false,
		},
		{
			name:             "Failed to Add Question",
			token:            "token",
			mockToken:        true,
			mockSuc:          true,
			mockUid:          0,
			mockRole:         service.ADMIN,
			mockBannedWords:  true,
			mockBannedResult: BannedWords,
			mockBannedErr:    nil,
			mockAddQuestion:  true,
			mockAddErr:       errors.New("failed to add"),
			mockQid:          0,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       true,
			commit:         false,
		},
		{
			name:             "Failed to IncQuestionCount",
			token:            "token",
			mockToken:        true,
			mockSuc:          true,
			mockUid:          0,
			mockRole:         service.ADMIN,
			mockBannedWords:  true,
			mockBannedResult: BannedWords,
			mockBannedErr:    nil,
			mockAddQuestion:  true,
			mockAddErr:       nil,
			mockQid:          0,
			req: service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{},
			},
			mockIncQuestionCount: true,
			incQuestionCountErr:  errors.New("err"),
			expectedCode:         service.Failed,
			expectedResult:       map[string]int8{"type": UnknownError},
			rollback:             true,
			commit:               false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockToken {
				mockUsersRPC.EXPECT().ParseToken(tt.token).Return(tt.mockSuc, tt.mockUid, tt.mockRole)
			}
			if tt.mockBannedWords {
				mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(tt.mockBannedResult, tt.mockBannedErr)
			}
			if tt.mockAddQuestion {
				mockQaDao.EXPECT().AddQuestion(gomock.Any(), tt.mockUid, tt.req.Title, tt.req.Content, tt.req.Category, tt.req.Labels, gomock.Any(), gomock.Any()).Return(tt.mockQid, tt.mockAddErr)
			}
			if tt.mockIncQuestionCount {
				mockQaDao.EXPECT().IncQuestionCount(gomock.Any(), tt.mockUid).Return(tt.incQuestionCountErr)
			}
			if tt.rollback {
				mockQaDao.EXPECT().Rollback(gomock.Any())
			}
			if tt.commit {
				mockQaDao.EXPECT().Commit(gomock.Any())
			}
			code, result := q.AddQuestion(tt.token, tt.req)
			a.Equal(tt.expectedCode, code)
			a.Equal(tt.expectedResult, result)
		})
	}
}

func TestServiceModifyQuestion(t *testing.T) {
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	a := assert.New(t)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var (
		longContent = string(make([]byte, service.QuestionContentLengthMax+1))
		BannedWords = []string{"magic"}
	)
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	tests := []struct {
		name                    string
		token                   string
		qid                     int64
		mockParseToken          bool
		parseTokenSuc           bool
		parseTokenUid           int64
		parseTokenRole          int8
		mockCheckQuestionOwner  bool
		checkQuestionOwnerOwner bool
		checkQuestionOwnerErr   error
		mockGetBannedWords      bool
		getBannedWordsWords     []string
		getBannedWordsErr       error
		mockModifyQuestion      bool
		modifyQuestionErr       error
		req                     service.ReqQuestionsPut
		expectedCode            int8
		expectedResult          interface{}
		rollback                bool
		commit                  bool
	}{
		{
			name:                    "Normal",
			token:                   "token",
			qid:                     456,
			mockParseToken:          true,
			parseTokenSuc:           true,
			parseTokenUid:           0,
			parseTokenRole:          service.USER,
			mockCheckQuestionOwner:  true,
			checkQuestionOwnerOwner: true,
			checkQuestionOwnerErr:   nil,
			mockGetBannedWords:      true,
			getBannedWordsWords:     []string{"magic"},
			getBannedWordsErr:       nil,
			mockModifyQuestion:      true,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"math"},
			},
			expectedCode:   service.Succeeded,
			expectedResult: nil,
			rollback:       false,
			commit:         true,
		},
		{
			name:                   "Invalid Qid",
			token:                  "token",
			mockParseToken:         false,
			mockCheckQuestionOwner: false,
			mockGetBannedWords:     false,
			mockModifyQuestion:     false,
			req: service.ReqQuestionsPut{
				Qid:      "234ihu",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"main"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       false,
			commit:         false,
		},
		{
			name:                   "Failed Token",
			token:                  "token",
			qid:                    234,
			mockParseToken:         true,
			parseTokenSuc:          false,
			parseTokenUid:          0,
			parseTokenRole:         service.USER,
			mockCheckQuestionOwner: false,
			mockGetBannedWords:     false,
			mockModifyQuestion:     false,
			req: service.ReqQuestionsPut{
				Qid:      "234",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"main"},
			},
			expectedCode:   service.Expired,
			expectedResult: nil,
			rollback:       false,
			commit:         false,
		},
		{
			name:                   "Long Title",
			token:                  "token",
			qid:                    456,
			mockParseToken:         true,
			parseTokenSuc:          true,
			parseTokenUid:          4,
			parseTokenRole:         service.ADMIN,
			mockCheckQuestionOwner: false,
			mockGetBannedWords:     false,
			mockModifyQuestion:     false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    string(make([]byte, service.QuestionTitleLengthMax+1)),
				Content:  "content",
				Category: "life",
				Labels:   []string{"main"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:                   "Long Content",
			token:                  "token",
			qid:                    456,
			mockParseToken:         true,
			parseTokenSuc:          true,
			parseTokenUid:          4,
			parseTokenRole:         service.ADMIN,
			mockCheckQuestionOwner: false,
			mockGetBannedWords:     false,
			mockModifyQuestion:     false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  longContent,
				Category: "life",
				Labels:   []string{"main"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:                   "Excessive Labels",
			token:                  "token",
			qid:                    456,
			mockParseToken:         true,
			parseTokenSuc:          true,
			parseTokenUid:          4,
			parseTokenRole:         service.ADMIN,
			mockCheckQuestionOwner: false,
			mockGetBannedWords:     false,
			mockModifyQuestion:     false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"1", "2", "3", "4", "5", "6"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:                   "Long Label",
			token:                  "token",
			qid:                    456,
			mockParseToken:         true,
			parseTokenSuc:          true,
			parseTokenUid:          4,
			parseTokenRole:         service.ADMIN,
			mockCheckQuestionOwner: false,
			mockGetBannedWords:     false,
			mockModifyQuestion:     false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"123456789012345678901234567890123"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": ConstraintsViolated},
			rollback:       false,
			commit:         false,
		},
		{
			name:                    "Failed CheckOwner",
			token:                   "token",
			qid:                     456,
			mockParseToken:          true,
			parseTokenSuc:           true,
			parseTokenUid:           5,
			parseTokenRole:          service.USER,
			mockCheckQuestionOwner:  true,
			checkQuestionOwnerOwner: true,
			checkQuestionOwnerErr:   errors.New("check owner error"),
			mockGetBannedWords:      false,
			mockModifyQuestion:      false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"math"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       true,
			commit:         false,
		},
		{
			name:                    "Not Owner",
			token:                   "token",
			qid:                     456,
			mockParseToken:          true,
			parseTokenSuc:           true,
			parseTokenUid:           5,
			parseTokenRole:          service.USER,
			mockCheckQuestionOwner:  true,
			checkQuestionOwnerOwner: false,
			checkQuestionOwnerErr:   nil,
			mockGetBannedWords:      false,
			mockModifyQuestion:      false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"math"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       true,
			commit:         false,
		},
		{
			name:                    "Failed to Get Banned Words",
			token:                   "token",
			qid:                     456,
			mockParseToken:          true,
			parseTokenSuc:           true,
			parseTokenUid:           5,
			parseTokenRole:          service.USER,
			mockCheckQuestionOwner:  true,
			checkQuestionOwnerOwner: true,
			checkQuestionOwnerErr:   nil,
			mockGetBannedWords:      true,
			getBannedWordsWords:     nil,
			getBannedWordsErr:       errors.New("error"),
			mockModifyQuestion:      false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "content",
				Category: "life",
				Labels:   []string{"math"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       true,
			commit:         false,
		},
		{
			name:                    "Has Keywords",
			token:                   "token",
			qid:                     456,
			mockParseToken:          true,
			parseTokenSuc:           true,
			parseTokenUid:           5,
			parseTokenRole:          service.USER,
			mockCheckQuestionOwner:  true,
			checkQuestionOwnerOwner: true,
			checkQuestionOwnerErr:   nil,
			mockGetBannedWords:      true,
			getBannedWordsWords:     BannedWords,
			getBannedWordsErr:       nil,
			mockModifyQuestion:      false,
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "It's not magic, it's just your fear.",
				Category: "life",
				Labels:   []string{"magic"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": HasKeyword},
			rollback:       true,
			commit:         false,
		},
		{
			name:                    "Failed to Modify",
			token:                   "token",
			qid:                     456,
			mockParseToken:          true,
			parseTokenSuc:           true,
			parseTokenUid:           5,
			parseTokenRole:          service.USER,
			mockCheckQuestionOwner:  true,
			checkQuestionOwnerOwner: true,
			checkQuestionOwnerErr:   nil,
			mockGetBannedWords:      true,
			getBannedWordsWords:     BannedWords,
			getBannedWordsErr:       nil,
			mockModifyQuestion:      true,
			modifyQuestionErr:       errors.New("error"),
			req: service.ReqQuestionsPut{
				Qid:      "456",
				Title:    "title",
				Content:  "What to eat",
				Category: "life",
				Labels:   []string{"food"},
			},
			expectedCode:   service.Failed,
			expectedResult: map[string]int8{"type": UnknownError},
			rollback:       true,
			commit:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Test case " + tt.name + ":")
			if tt.mockParseToken {
				mockUsersRPC.EXPECT().ParseToken(tt.token).Return(tt.parseTokenSuc, tt.parseTokenUid, tt.parseTokenRole)
			}
			if tt.mockCheckQuestionOwner {
				mockQaDao.EXPECT().CheckQuestionOwner(gomock.Any(), tt.qid, tt.parseTokenUid).Return(tt.checkQuestionOwnerOwner, tt.checkQuestionOwnerErr)
			}
			if tt.mockGetBannedWords {
				mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(tt.getBannedWordsWords, tt.getBannedWordsErr)
			}
			if tt.mockModifyQuestion {
				mockQaDao.EXPECT().ModifyQuestion(gomock.Any(), tt.qid, tt.req.Title, tt.req.Content, tt.req.Category, tt.req.Labels, gomock.Any(), gomock.Any()).Return(tt.modifyQuestionErr)
			}
			if tt.rollback {
				mockQaDao.EXPECT().Rollback(gomock.Any())
			}
			if tt.commit {
				mockQaDao.EXPECT().Commit(gomock.Any())
			}
			code, result := q.ModifyQuestion(tt.token, tt.req)
			a.Equal(tt.expectedCode, code)
			a.Equal(tt.expectedResult, result)
		})
	}
}

func TestServiceMainPage(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockQaDao.EXPECT().Init()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	mockQaDao.EXPECT().Commit(gomock.Any()).AnyTimes()
	mockQaDao.EXPECT().Rollback(gomock.Any()).AnyTimes()
	a := assert.New(t)
	var (
		MainPageResult = []entity.Questions{
			{
				17,
				5,
				"title",
				"life",
				sql.NullInt64{Valid: false},
				4,
				9,
				2,
				time.Now().Unix(),
				[]string{"math"},
				false,
			},
		}
		DetailsResult = []entity.QuestionDetails{
			{
				17,
				"What to eat?\n![](picture URL)",
				"picture URL",
				"What to eat",
			},
		}
		UserInfosResult = []rpc.UserInfo{
			{
				"sk",
				"SK",
				"icon URL",
			},
		}
	)
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	tests := []struct {
		name               string
		token              string
		category           string
		page               int64
		mockParseToken     bool
		parseTokenSuc      bool
		parseTokenUid      int64
		parseTokenRole     int8
		mockMainPage       bool
		mainPageResult     []entity.Questions
		mainPageErr        error
		mockFindDetails    bool
		findDetailsResult  []entity.QuestionDetails
		mockGetBannedWords bool
		getBannedWordsRes  []string
		getBannedWordsErr  error
		mockGetUserInfos   bool
		getUserInfosUids   []int64
		getUserInfosResult []rpc.UserInfo
		getUserInfosErr    error
		expectedCode       int8
		expectedResult     interface{}
	}{
		{
			name:               "Normal",
			token:              "token",
			category:           "life",
			page:               1,
			mockParseToken:     true,
			parseTokenSuc:      true,
			parseTokenUid:      5,
			parseTokenRole:     service.USER,
			mockMainPage:       true,
			mainPageResult:     MainPageResult,
			mainPageErr:        nil,
			mockFindDetails:    true,
			findDetailsResult:  DetailsResult,
			mockGetBannedWords: true,
			getBannedWordsRes:  []string{},
			getBannedWordsErr:  nil,
			mockGetUserInfos:   true,
			getUserInfosUids:   []int64{MainPageResult[0].Raiser},
			getUserInfosResult: UserInfosResult,
			getUserInfosErr:    nil,
			expectedCode:       service.Succeeded,
			expectedResult: []service.QuestionListItem{
				{
					Qid: strconv.FormatInt(MainPageResult[0].Qid, 10),
					Owner: service.Owner{
						Uid:      strconv.FormatInt(MainPageResult[0].Raiser, 10),
						Name:     UserInfosResult[0].Name,
						Nickname: UserInfosResult[0].Nickname,
						Icon:     UserInfosResult[0].Icon,
					},
					Title:         MainPageResult[0].Title,
					Time:          fmt.Sprint(time.Unix(MainPageResult[0].Time, 0)),
					AnswerCount:   MainPageResult[0].AnswerCount,
					ViewCount:     MainPageResult[0].ViewCount,
					FavoriteCount: MainPageResult[0].FavoriteCount,
					Category:      MainPageResult[0].Category,
					Labels:        MainPageResult[0].Labels,
					Head:          DetailsResult[0].Head,
					PictureUrls:   []string{DetailsResult[0].PictureUrl},
				},
			},
		},
		{
			name:             "Failed Token",
			token:            "token",
			category:         "life",
			page:             1,
			mockParseToken:   true,
			parseTokenSuc:    false,
			parseTokenUid:    0,
			parseTokenRole:   service.USER,
			mockMainPage:     false,
			mockFindDetails:  false,
			mockGetUserInfos: false,
			expectedCode:     service.Expired,
			expectedResult:   nil,
		},
		{
			name:             "Negative Page",
			token:            "token",
			category:         "life",
			page:             -1,
			mockParseToken:   true,
			parseTokenSuc:    true,
			parseTokenUid:    1,
			parseTokenRole:   service.USER,
			mockMainPage:     false,
			mockFindDetails:  false,
			mockGetUserInfos: false,
			expectedCode:     service.Failed,
			expectedResult:   nil,
		},
		{
			name:             "Failed to Get Main Page",
			token:            "token",
			category:         "life",
			page:             1,
			mockParseToken:   true,
			parseTokenSuc:    true,
			parseTokenUid:    1,
			parseTokenRole:   service.USER,
			mockMainPage:     true,
			mainPageResult:   nil,
			mainPageErr:      errors.New("error"),
			mockFindDetails:  false,
			mockGetUserInfos: false,
			expectedCode:     service.Failed,
			expectedResult:   nil,
		},
		{
			name:               "Failed to Get Banned Words",
			token:              "token",
			category:           "life",
			page:               1,
			mockParseToken:     true,
			parseTokenSuc:      true,
			parseTokenUid:      1,
			parseTokenRole:     service.USER,
			mockMainPage:       true,
			mainPageResult:     MainPageResult,
			mainPageErr:        nil,
			mockFindDetails:    true,
			findDetailsResult:  DetailsResult,
			mockGetBannedWords: true,
			getBannedWordsRes:  []string{},
			getBannedWordsErr:  errors.New("error"),
			expectedCode:       service.Failed,
			expectedResult:     nil,
		},
		{
			name:               "Failed to Get User Infos",
			token:              "token",
			category:           "life",
			page:               1,
			mockParseToken:     true,
			parseTokenSuc:      true,
			parseTokenUid:      1,
			parseTokenRole:     service.USER,
			mockMainPage:       true,
			mainPageResult:     MainPageResult,
			mainPageErr:        nil,
			mockFindDetails:    true,
			findDetailsResult:  DetailsResult,
			mockGetBannedWords: true,
			getBannedWordsRes:  []string{},
			getBannedWordsErr:  nil,
			mockGetUserInfos:   true,
			getUserInfosUids:   []int64{MainPageResult[0].Raiser},
			getUserInfosResult: nil,
			getUserInfosErr:    errors.New("error"),
			expectedCode:       service.Failed,
			expectedResult:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Test case " + tt.name + ":")
			if tt.mockParseToken {
				mockUsersRPC.EXPECT().ParseToken(tt.token).Return(tt.parseTokenSuc, tt.parseTokenUid, tt.parseTokenRole)
			}
			if tt.mockMainPage {
				mockQaDao.EXPECT().MainPage(gomock.Any(), tt.parseTokenUid, tt.category, tt.page).Return(tt.mainPageResult, tt.mainPageErr)
			}
			if tt.mockFindDetails {
				mockQaDao.EXPECT().FindQuestionDetails(gomock.Any(), tt.mainPageResult).Return(tt.findDetailsResult)
			}
			if tt.mockGetBannedWords {
				mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(tt.getBannedWordsRes, tt.getBannedWordsErr)
			}
			if tt.mockGetUserInfos {
				mockUsersRPC.EXPECT().GetUserInfos(tt.getUserInfosUids).Return(tt.getUserInfosResult, tt.getUserInfosErr)
			}
			code, result := q.MainPage(tt.token, tt.category, tt.page)
			a.Equal(tt.expectedCode, code)
			a.Equal(tt.expectedResult, result)
		})
	}
}

func TestServiceQuestionDetail(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	question := []entity.Questions{{
		Qid:            234,
		Raiser:         78,
		Title:          "title",
		Category:       "life",
		AcceptedAnswer: sql.NullInt64{Valid: false},
	}}
	questionWithAccepted := []entity.Questions{{
		Qid:            234,
		Raiser:         78,
		Title:          "title",
		Category:       "life",
		AcceptedAnswer: sql.NullInt64{Valid: true, Int64: 567},
	}}
	details := []entity.QuestionDetails{{
		234,
		"content",
		"",
		"content",
	}}
	userInfos := []rpc.UserInfo{{
		"skfe",
		"skfe2",
		"icon data",
	}}
	tests := []struct {
		name                     string
		token                    string
		qid                      int64
		parseTokenSuc            bool
		parseTokenUid            int64
		parseTokenRole           int8
		mockFindQuestion         bool
		findQuestionQuestion     []entity.Questions
		findQuestionErr          error
		mockGetBannedWords       bool
		getBannedWordsRes        []string
		getBannedWordsErr        error
		mockFindDetails          bool
		findDetailsDetails       []entity.QuestionDetails
		mockGetUserInfos         bool
		getUserInfosRes          []rpc.UserInfo
		getUserInfosErr          error
		mockSaveQuestionSkeleton bool
		commit                   bool
		rollback                 bool
		wantCode                 int8
		wantResult               interface{}
	}{
		{
			name:                     "Normal",
			token:                    "token",
			qid:                      234,
			parseTokenSuc:            true,
			parseTokenUid:            56,
			parseTokenRole:           service.USER,
			mockFindQuestion:         true,
			findQuestionQuestion:     question,
			mockGetBannedWords:       true,
			getBannedWordsRes:        []string{},
			getBannedWordsErr:        nil,
			mockFindDetails:          true,
			findDetailsDetails:       details,
			mockGetUserInfos:         true,
			getUserInfosRes:          userInfos,
			mockSaveQuestionSkeleton: true,
			commit:                   true,
			wantCode:                 service.Succeeded,
			wantResult:               service.QuestionInfo{Qid: "234", Owner: service.Owner{Uid: "78", Name: "skfe", Nickname: "skfe2", Icon: "icon data"}, Title: "title", Category: "life", Accepted: "", Content: "content", Time: "0"},
		},
		{
			name:                     "Normal2",
			token:                    "token",
			qid:                      234,
			parseTokenSuc:            true,
			parseTokenUid:            56,
			parseTokenRole:           service.USER,
			mockFindQuestion:         true,
			findQuestionQuestion:     questionWithAccepted,
			mockGetBannedWords:       true,
			getBannedWordsRes:        []string{},
			getBannedWordsErr:        nil,
			mockFindDetails:          true,
			findDetailsDetails:       details,
			mockGetUserInfos:         true,
			getUserInfosRes:          userInfos,
			mockSaveQuestionSkeleton: true,
			commit:                   true,
			wantCode:                 service.Succeeded,
			wantResult:               service.QuestionInfo{Qid: "234", Owner: service.Owner{Uid: "78", Name: "skfe", Nickname: "skfe2", Icon: "icon data"}, Title: "title", Category: "life", Accepted: "567", Content: "content", Time: "0"},
		},
		{
			name:          "Token expired",
			parseTokenSuc: false,
			wantCode:      service.Expired,
		},
		{
			name:                 "Failed to Find Question by Id",
			parseTokenSuc:        true,
			parseTokenUid:        46,
			parseTokenRole:       service.USER,
			mockFindQuestion:     true,
			findQuestionQuestion: nil,
			findQuestionErr:      errors.New("xxx"),
			rollback:             true,
			wantCode:             service.Failed,
			wantResult:           nil,
		},
		{
			name:                 "Find Question not Found",
			parseTokenSuc:        true,
			parseTokenUid:        345,
			parseTokenRole:       service.USER,
			mockFindQuestion:     true,
			findQuestionQuestion: []entity.Questions{},
			rollback:             true,
			wantCode:             service.Failed,
			wantResult:           nil,
		},
		{
			name:                 "Failed to Get Banned Words",
			parseTokenSuc:        true,
			parseTokenUid:        345,
			parseTokenRole:       service.USER,
			mockFindQuestion:     true,
			findQuestionQuestion: question,
			mockGetBannedWords:   true,
			getBannedWordsErr:    errors.New("xxx"),
			rollback:             true,
			wantCode:             service.Failed,
			wantResult:           nil,
		},
		{
			name:                 "Failed to Get User Infos",
			parseTokenSuc:        true,
			parseTokenUid:        345,
			parseTokenRole:       service.USER,
			mockFindQuestion:     true,
			findQuestionQuestion: question,
			mockGetBannedWords:   true,
			getBannedWordsRes:    []string{"river"},
			mockFindDetails:      true,
			findDetailsDetails:   details,
			mockGetUserInfos:     true,
			getUserInfosErr:      errors.New("xx"),
			rollback:             true,
			wantCode:             service.Failed,
			wantResult:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsersRPC.EXPECT().ParseToken(tt.token).Return(tt.parseTokenSuc, tt.parseTokenUid, tt.parseTokenRole)
			if tt.mockFindQuestion {
				mockQaDao.EXPECT().FindQuestionById(gomock.Any(), tt.qid).Return(tt.findQuestionQuestion, tt.findQuestionErr)
			}
			if tt.mockFindDetails {
				mockQaDao.EXPECT().FindQuestionDetails(gomock.Any(), tt.findQuestionQuestion).Return(tt.findDetailsDetails)
			}
			if tt.mockGetUserInfos {
				mockUsersRPC.EXPECT().GetUserInfos(gomock.Any()).Return(tt.getUserInfosRes, tt.getUserInfosErr)
			}
			if tt.mockSaveQuestionSkeleton {
				mockQaDao.EXPECT().SaveQuestionSkeleton(gomock.Any(), gomock.Any())
			}
			if tt.mockGetBannedWords {
				mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(tt.getBannedWordsRes, tt.getBannedWordsErr)
			}
			if tt.commit {
				mockQaDao.EXPECT().Commit(gomock.Any())
			}
			if tt.rollback {
				mockQaDao.EXPECT().Rollback(gomock.Any())
			}
			code, result := q.QuestionDetail(tt.token, tt.qid)
			a.Equal(tt.wantCode, code)
			a.Equal(tt.wantResult, result)
		})
	}
}

func TestServiceAnswerListResponse(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	answers := []entity.Answers{{Aid: 56, Answerer: 7, Qid: 234}}
	answerDetails := []entity.AnswerDetails{{Aid: 56, Content: "content", PictureUrl: "pic url", Head: "Head"}}
	keywords := []string{"key"}
	userInfoResult := []rpc.UserInfo{{Name: "sk", Nickname: "nick", Icon: "icon"}}
	var ctx dao.TransactionContext
	var (
		uid  int64 = 5
		role int8  = service.USER
	)

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos([]int64{7}).Return(userInfoResult, nil)
		mockQaDao.EXPECT().GetAnswerActionInfos(ctx, uid, []int64{234}, []int64{56}).Return([]dao.AnswerActionInfo{{Liked: true, Approved: false, Approvable: true}}, nil)
		result, err := q.AnswerListResponse(ctx, uid, role, answers, answerDetails, &keywords)
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
		_, err := q.AnswerListResponse(ctx, uid, role, answers, answerDetails, &keywords)
		a.NotNil(err)
	})

	t.Run("Failed to Get Action Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos([]int64{7}).Return(userInfoResult, nil)
		mockQaDao.EXPECT().GetAnswerActionInfos(ctx, uid, []int64{234}, []int64{56}).Return(nil, errors.New("gg"))
		_, err := q.AnswerListResponse(ctx, uid, role, answers, answerDetails, &keywords)
		a.NotNil(err)
	})
}

func TestServiceAddAnswer(t *testing.T) {
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	req := service.ReqAnswersPost{
		Qid:     "456",
		Content: "content",
	}
	var uid int64 = 5
	var role int8 = service.USER
	banned := []string{"river"}
	pictureUrl := ""
	questions := []entity.Questions{{Qid: 456, Raiser: 89, Title: "title", Category: "life", Labels: []string{}}}
	unknownFailure := map[string]int8{"type": UnknownError}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().AddAnswer(gomock.Any(), uid, int64(456), req.Content, pictureUrl, "content ").Return(int64(47), nil)
		mockQaDao.EXPECT().IncUserAnswerCount(gomock.Any(), uid).Return(nil)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), int64(456)).Return(questions, nil)
		mockQaDao.EXPECT().SaveQuestionSkeleton(gomock.Any(), gomock.Any()).Return(nil)
		mockQaDao.EXPECT().Commit(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(int8(service.Succeeded), code)
		a.Equal("47", result.(map[string]string)["aid"])
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, int64(0), int8(0))
		code, _ := q.AddAnswer(token, req)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Long Content", func(t *testing.T) {
		long := string(make([]byte, service.AnswerContentLengthMax+1))
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		nreq := req
		nreq.Content = long
		code, result := q.AddAnswer(token, nreq)
		a.Equal(map[string]int8{"type": ConstraintsViolated}, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, errors.New("xx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Has Keywords", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return([]string{"content"}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(map[string]int8{"type": HasKeyword}, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Add Answer", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().AddAnswer(gomock.Any(), uid, int64(456), req.Content, pictureUrl, "content ").Return(int64(0), errors.New("xx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Inc User Answer Count", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().AddAnswer(gomock.Any(), uid, int64(456), req.Content, pictureUrl, "content ").Return(int64(47), nil)
		mockQaDao.EXPECT().IncUserAnswerCount(gomock.Any(), uid).Return(errors.New("xx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Find Question By Id", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().AddAnswer(gomock.Any(), uid, int64(456), req.Content, pictureUrl, "content ").Return(int64(47), nil)
		mockQaDao.EXPECT().IncUserAnswerCount(gomock.Any(), uid).Return(nil)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), int64(456)).Return(nil, errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Find Question not Exist", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().AddAnswer(gomock.Any(), uid, int64(456), req.Content, pictureUrl, "content ").Return(int64(47), nil)
		mockQaDao.EXPECT().IncUserAnswerCount(gomock.Any(), uid).Return(nil)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), int64(456)).Return([]entity.Questions{}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Save Question Skeleton", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().AddAnswer(gomock.Any(), uid, int64(456), req.Content, pictureUrl, "content ").Return(int64(47), nil)
		mockQaDao.EXPECT().IncUserAnswerCount(gomock.Any(), uid).Return(nil)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), int64(456)).Return(questions, nil)
		mockQaDao.EXPECT().SaveQuestionSkeleton(gomock.Any(), gomock.Any()).Return(errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.AddAnswer(token, req)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceModifyAnswer(t *testing.T) {
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	req := service.ReqAnswersPut{
		Aid:     "346",
		Content: "content",
	}
	var aid int64 = 346
	var uid int64 = 76
	var role int8 = service.USER
	answer := []entity.Answers{{Aid: 345, Answerer: 76, Qid: 123}}
	banned := []string{"tiger"}
	unknownFailure := map[string]int8{"type": UnknownError}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answer, nil)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().ModifyAnswer(gomock.Any(), aid, "content", "", "content ").Return(nil)
		mockQaDao.EXPECT().Commit(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Nil(result)
		a.Equal(int8(service.Succeeded), code)
	})

	t.Run("Failed Aid", func(t *testing.T) {
		nreq := req
		nreq.Aid = "234x"
		code, result := q.ModifyAnswer(token, nreq)
		a.Equal(unknownFailure, result)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Long Content", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		long := string(make([]byte, service.QuestionContentLengthMax+1))
		nreq := req
		nreq.Content = long
		code, result := q.ModifyAnswer(token, nreq)
		a.Equal(int8(service.Failed), code)
		a.Equal(map[string]int8{"type": ConstraintsViolated}, result)
	})

	t.Run("Failed to Find Answer By Id", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(nil, errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, result)
	})

	t.Run("Find Answer By Id not Found", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return([]entity.Answers{}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, result)
	})

	t.Run("Not Owner", func(t *testing.T) {
		newAnswer := []entity.Answers{{Aid: 345, Answerer: 77, Qid: 123}}
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(newAnswer, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, result)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answer, nil)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, result)
	})

	t.Run("Has Keywords", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answer, nil)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return([]string{"content"}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(map[string]int8{"type": HasKeyword}, result)
	})

	t.Run("Failed to Modify Answer", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answer, nil)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(banned, nil)
		mockQaDao.EXPECT().ModifyAnswer(gomock.Any(), aid, "content", "", "content ").Return(errors.New("xxxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ModifyAnswer(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, result)
	})
}

func TestServiceListAnswers(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		qid  int64 = 234
		page int64 = 2
		sort int8  = 1
		uid  int64 = 76
		role int8  = service.USER
	)
	answers := []entity.Answers{{Aid: 345, Answerer: 36, Qid: 234}}
	details := []entity.AnswerDetails{{Aid: 345, Content: "content", PictureUrl: "", Head: "content "}}
	userInfos := []rpc.UserInfo{{Name: "tsw", Nickname: "sk", Icon: "icon"}}
	actionInfos := []dao.AnswerActionInfo{{Liked: false, Approved: false, Approvable: true}}
	bannedWords := []string{"xxx"}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionAnswers(gomock.Any(), qid, page, sort).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos([]int64{36}).Return(userInfos, nil)
		mockQaDao.EXPECT().GetAnswerActionInfos(gomock.Any(), uid, []int64{234}, []int64{345}).Return(actionInfos, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.ListAnswers(token, qid, page, sort)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]service.AnswerListItem)[0]
		a.Equal("36", res.Owner.Uid)
		a.Equal("sk", res.Owner.Nickname)
		a.True(res.Approvable)
		a.Equal("content ", res.Head)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.ListAnswers(token, qid, page, sort)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Find Question Answers", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionAnswers(gomock.Any(), qid, page, sort).Return(nil, errors.New("xx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.ListAnswers(token, qid, page, sort)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionAnswers(gomock.Any(), qid, page, sort).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.ListAnswers(token, qid, page, sort)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get User Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionAnswers(gomock.Any(), qid, page, sort).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos([]int64{36}).Return(nil, errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.ListAnswers(token, qid, page, sort)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Answer Action Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionAnswers(gomock.Any(), qid, page, sort).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos([]int64{36}).Return(userInfos, nil)
		mockQaDao.EXPECT().GetAnswerActionInfos(gomock.Any(), uid, []int64{234}, []int64{345}).Return(nil, errors.New("xxx"))
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.ListAnswers(token, qid, page, sort)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceAnswerDetail(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		aid  int64 = 345
		uid  int64 = 76
		role int8  = service.USER
	)
	answers := []entity.Answers{{Aid: 345, Answerer: 36, Qid: 234}}
	details := []entity.AnswerDetails{{Aid: 345, Content: "content", PictureUrl: "", Head: "content "}}
	userInfos := []rpc.UserInfo{{Name: "name", Nickname: "nick", Icon: "icon"}}
	actionInfos := []dao.AnswerActionInfo{{Liked: true, Approved: false, Approvable: false}}
	bannedWords := []string{"xxx"}
	err := errors.New("xxx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos([]int64{36}).Return(userInfos, nil)
		mockQaDao.EXPECT().GetAnswerActionInfos(gomock.Any(), uid, []int64{234}, []int64{345}).Return(actionInfos, nil)
		mockQaDao.EXPECT().Commit(gomock.Any())
		code, result := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Succeeded), code)
		res := result.(service.AnswerInfo)
		a.Equal("nick", res.Owner.Nickname)
		a.Equal("icon", res.Owner.Icon)
		a.Equal("content", res.Content)
		a.True(res.Liked)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Find Answer by Id", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Find Answer by Id not Found", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return([]entity.Answers{}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Find Answer by Id not Found", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return([]entity.Answers{}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get User Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos([]int64{36}).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Answer Action Infos", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindAnswerById(gomock.Any(), aid).Return(answers, nil)
		mockQaDao.EXPECT().FindAnswerDetails(gomock.Any(), answers).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos([]int64{36}).Return(userInfos, nil)
		mockQaDao.EXPECT().GetAnswerActionInfos(gomock.Any(), uid, []int64{234}, []int64{345}).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.AnswerDetail(token, aid)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceCommentListResponse(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	var (
		cmid int64 = 345
	)
	comments := []entity.Comments{{Cmid: cmid, Uid: 11, Aid: 97}}
	details := []entity.CommentDetails{{Cmid: 456, Content: "ha ha ha"}}
	keywords := []string{"test"}
	uids := []int64{11}
	err := errors.New("xxx")
	userInfos := []rpc.UserInfo{{Name: "name", Nickname: "nickname", Icon: "icon"}}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		result, err := q.CommentListResponse(comments, details, &keywords)
		res := result.([]service.CommentListItem)[0]
		a.Nil(err)
		a.Equal("nickname", res.Nickname)
		a.Equal("345", res.Cmid)
		a.Equal("icon", res.Icon)
		a.Equal("ha ha ha", res.Content)
	})

	t.Run("Failed to Get User Info", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		_, err := q.CommentListResponse(comments, details, &keywords)
		a.NotNil(err)
	})
}

func TestServiceGetComments(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		aid  int64 = 345
		page int64 = 5
		uid  int64 = 76
		role int8  = service.USER
	)
	comments := []entity.Comments{{Cmid: 98, Uid: 9, Aid: aid}}
	details := []entity.CommentDetails{{Cmid: 98, Content: "ha ha ha"}}
	uids := []int64{9}
	userInfos := []rpc.UserInfo{{Name: "name", Nickname: "nick", Icon: "icon"}}
	bannedWords := []string{"xxx"}
	err := errors.New("xxx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetComments(gomock.Any(), aid, page).Return(comments, nil)
		mockQaDao.EXPECT().FindCommentDetails(gomock.Any(), comments).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.GetComments(token, aid, page)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]service.CommentListItem)[0]
		a.Equal("ha ha ha", res.Content)
		a.False(res.HasKeywords)
		a.Equal("nick", res.Nickname)
		a.Equal("98", res.Cmid)
		a.Equal("9", res.Uid)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.GetComments(token, aid, page)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Get Comments", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetComments(gomock.Any(), aid, page).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.GetComments(token, aid, page)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetComments(gomock.Any(), aid, page).Return(comments, nil)
		mockQaDao.EXPECT().FindCommentDetails(gomock.Any(), comments).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.GetComments(token, aid, page)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetComments(gomock.Any(), aid, page).Return(comments, nil)
		mockQaDao.EXPECT().FindCommentDetails(gomock.Any(), comments).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.GetComments(token, aid, page)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceAddComment(t *testing.T) {
	const (
		ConstraintsViolated = 0
		HasKeywords         = 1
		UnknownError        = 2
	)
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		aid  int64 = 345
		uid  int64 = 76
		role int8  = service.USER
		cmid int64 = 234
	)
	req := service.ReqCommentsPost{
		Aid:     "345",
		Content: "ha ha ha",
	}
	bannedWords := []string{"xxx"}
	err := errors.New("xxx")
	unknownFailure := map[string]int8{"type": UnknownError}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockQaDao.EXPECT().AddComment(gomock.Any(), uid, aid, req.Content).Return(cmid, nil)
		mockQaDao.EXPECT().Commit(gomock.Any())
		code, result := q.AddComment(token, req)
		res := result.(map[string]string)["cmid"]
		a.Equal(int8(service.Succeeded), code)
		a.Equal("234", res)
	})

	t.Run("Failed Parse", func(t *testing.T) {
		newReq := service.ReqCommentsPost{
			Aid:     "345sdf",
			Content: "ha ha ha",
		}
		code, _ := q.AddComment(token, newReq)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.AddComment(token, req)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Long Content", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		newReq := service.ReqCommentsPost{
			Aid:     "345",
			Content: string(make([]byte, service.CommentLengthMax+1)),
		}
		code, result := q.AddComment(token, newReq)
		a.Equal(int8(service.Failed), code)
		a.Equal(map[string]int8{"type": ConstraintsViolated}, result)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, res := q.AddComment(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, res)
	})

	t.Run("Failed to Add Comment", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockQaDao.EXPECT().AddComment(gomock.Any(), uid, aid, req.Content).Return(int64(0), err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, res := q.AddComment(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, res)
	})

	t.Run("Has Keywords", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return([]string{"ha"}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, res := q.AddComment(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(map[string]int8{"type": HasKeywords}, res)
	})
}

func TestServiceCriticismListResponse(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	var (
		ctid int64 = 345
	)
	criticisms := []entity.Criticisms{{Ctid: ctid, Uid: 11, Aid: 97}}
	details := []entity.CriticismDetails{{Ctid: 456, Content: "ha ha ha"}}
	keywords := []string{"test"}
	uids := []int64{11}
	err := errors.New("xxx")
	userInfos := []rpc.UserInfo{{Name: "name", Nickname: "nickname", Icon: "icon"}}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		result, err := q.CriticismListResponse(criticisms, details, &keywords)
		res := result.([]service.CriticismListItem)[0]
		a.Nil(err)
		a.Equal("nickname", res.Nickname)
		a.Equal("345", res.Ctid)
		a.Equal("icon", res.Icon)
		a.Equal("ha ha ha", res.Content)
	})

	t.Run("Failed to Get User Info", func(t *testing.T) {
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		_, err := q.CriticismListResponse(criticisms, details, &keywords)
		a.NotNil(err)
	})
}

func TestServiceGetCriticisms(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		aid  int64 = 345
		page int64 = 5
		uid  int64 = 76
		role int8  = service.USER
	)
	criticisms := []entity.Criticisms{{Ctid: 98, Uid: 9, Aid: aid}}
	details := []entity.CriticismDetails{{Ctid: 98, Content: "ha ha ha"}}
	uids := []int64{9}
	userInfos := []rpc.UserInfo{{Name: "name", Nickname: "nick", Icon: "icon"}}
	bannedWords := []string{"xxx"}
	err := errors.New("xxx")

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetCriticisms(gomock.Any(), aid, page).Return(criticisms, nil)
		mockQaDao.EXPECT().FindCriticismDetails(gomock.Any(), criticisms).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(userInfos, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, result := q.GetCriticisms(token, aid, page)
		a.Equal(int8(service.Succeeded), code)
		res := result.([]service.CriticismListItem)[0]
		a.Equal("ha ha ha", res.Content)
		a.False(res.HasKeywords)
		a.Equal("nick", res.Nickname)
		a.Equal("98", res.Ctid)
		a.Equal("9", res.Uid)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.GetCriticisms(token, aid, page)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed to Get Criticisms", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetCriticisms(gomock.Any(), aid, page).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.GetCriticisms(token, aid, page)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetCriticisms(gomock.Any(), aid, page).Return(criticisms, nil)
		mockQaDao.EXPECT().FindCriticismDetails(gomock.Any(), criticisms).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.GetCriticisms(token, aid, page)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetCriticisms(gomock.Any(), aid, page).Return(criticisms, nil)
		mockQaDao.EXPECT().FindCriticismDetails(gomock.Any(), criticisms).Return(details)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockUsersRPC.EXPECT().GetUserInfos(uids).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.GetCriticisms(token, aid, page)
		a.Equal(int8(service.Failed), code)
	})
}

func TestServiceAddCriticism(t *testing.T) {
	const (
		ConstraintsViolated = 0
		HasKeywords         = 1
		UnknownError        = 2
	)
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)

	token := "token"
	var (
		aid  int64 = 345
		uid  int64 = 76
		role int8  = service.USER
		ctid int64 = 234
	)
	req := service.ReqCriticismsPost{
		Aid:     "345",
		Content: "ha ha ha",
	}
	bannedWords := []string{"xxx"}
	err := errors.New("xxx")
	unknownFailure := map[string]int8{"type": UnknownError}

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockQaDao.EXPECT().AddCriticism(gomock.Any(), uid, aid, req.Content).Return(ctid, nil)
		mockQaDao.EXPECT().Commit(gomock.Any())
		code, result := q.AddCriticism(token, req)
		res := result.(map[string]string)["ctid"]
		a.Equal(int8(service.Succeeded), code)
		a.Equal("234", res)
	})

	t.Run("Failed Parse", func(t *testing.T) {
		newReq := service.ReqCriticismsPost{
			Aid:     "345sdf",
			Content: "ha ha ha",
		}
		code, _ := q.AddCriticism(token, newReq)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.AddCriticism(token, req)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Long Content", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		newReq := service.ReqCriticismsPost{
			Aid:     "345",
			Content: string(make([]byte, service.CriticismLengthMax+1)),
		}
		code, result := q.AddCriticism(token, newReq)
		a.Equal(int8(service.Failed), code)
		a.Equal(map[string]int8{"type": ConstraintsViolated}, result)
	})

	t.Run("Failed to Get Banned Words", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, res := q.AddCriticism(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, res)
	})

	t.Run("Failed to Add Criticism", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return(bannedWords, nil)
		mockQaDao.EXPECT().AddCriticism(gomock.Any(), uid, aid, req.Content).Return(int64(0), err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, res := q.AddCriticism(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(unknownFailure, res)
	})

	t.Run("Has Keywords", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().GetBannedWords(gomock.Any()).Return([]string{"ha"}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, res := q.AddCriticism(token, req)
		a.Equal(int8(service.Failed), code)
		a.Equal(map[string]int8{"type": HasKeywords}, res)
	})
}

func TestServiceDeleteQuestion(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaDao := mock.NewMockQaDao(mockCtrl)
	mockUsersRPC := mock.NewMockUsersRPC(mockCtrl)
	mockQaDao.EXPECT().Init().AnyTimes()
	mockQaDao.EXPECT().Begin(gomock.Any()).Return(dao.TransactionContext{}, nil).AnyTimes()
	var q service.QaServiceImpl
	_ = q.Init(mockQaDao, mockUsersRPC)
	a := assert.New(t)
	var (
		qid int64  = 345
		uid int64 = 23
		role int8 = service.USER
	)
	req := service.ReqQuestionsDelete{Qid: "345"}
	question := []entity.Questions{{Qid: 345, Raiser: uid}}
	err := errors.New("xx")

	token := "token"

	t.Run("Normal", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), qid).Return(question, nil)
		mockQaDao.EXPECT().DeleteQuestion(gomock.Any(), qid).Return(nil)
		mockQaDao.EXPECT().Commit(gomock.Any())
		code, _ := q.DeleteQuestion(token, req)
		a.Equal(int8(service.Succeeded), code)
	})

	t.Run("Failed Parse", func(t *testing.T) {
		newReq := service.ReqQuestionsDelete{Qid: "2342uih"}
		code, _ := q.DeleteQuestion(token, newReq)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Failed Token", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(false, uid, role)
		code, _ := q.DeleteQuestion(token, req)
		a.Equal(int8(service.Expired), code)
	})

	t.Run("Failed Find Question By Id", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), qid).Return(nil, err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.DeleteQuestion(token, req)
		a.Equal(int8(service.Failed), code)
	})

	t.Run("Find Question By Id not Found", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), qid).Return([]entity.Questions{}, nil)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.DeleteQuestion(token, req)
		a.Equal(int8(service.Succeeded), code)
	})

	t.Run("Failed to Delete Question", func(t *testing.T) {
		mockUsersRPC.EXPECT().ParseToken(token).Return(true, uid, role)
		mockQaDao.EXPECT().FindQuestionById(gomock.Any(), qid).Return(question, nil)
		mockQaDao.EXPECT().DeleteQuestion(gomock.Any(), qid).Return(err)
		mockQaDao.EXPECT().Rollback(gomock.Any())
		code, _ := q.DeleteQuestion(token, req)
		a.Equal(int8(service.Failed), code)
	})
}
