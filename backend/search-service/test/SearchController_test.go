package test

import (
	"context"
	"encoding/json"
	"github.com/SKFE396/search-service/controller"
	"github.com/SKFE396/search-service/mock"
	"github.com/SKFE396/search-service/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestControllerInit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	mockSearchService.EXPECT().Init(nil, nil).AnyTimes()
	tests := []struct {
		name string
	}{
		{"Initialize"},
	}
	q := controller.SearchController{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpServerExitDone := &sync.WaitGroup{}
			httpServerExitDone.Add(1)
			server := q.Init(httpServerExitDone, mockSearchService)
			time.Sleep(500 * time.Microsecond)
			if err := server.Shutdown(context.Background()); err != nil {
				t.Error(err)
			}
			httpServerExitDone.Wait()
		})
	}
}

func TestControllerDestruct(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	mockSearchService.EXPECT().Destruct()
	var s controller.SearchController
	s.SetSearchService(mockSearchService)
	s.Destruct()
}

func TestControllerSearchQuestions(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	a := assert.New(t)

	var s controller.SearchController
	s.SetSearchService(mockSearchService)
	mux.HandleFunc("/searchQuestions", s.SearchQuestions)
	tests := []struct {
		name string
		token string
		page string
		text string
		mockSearchQuestions bool
		mockPage int64
		searchQuestionsCode int8
		searchQuestionsResult interface{}
		wantCode int8
		wantResult interface{}
	} {
		{
			"Normal",
			"token",
			"6",
			"frozen",
			true,
			6,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Failed Parse",
			"token",
			"6sdf",
			"frozen",
			false,
			6,
			service.Succeeded,
			"result",
			service.Failed,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSearchQuestions {
				mockSearchService.EXPECT().SearchQuestions(tt.token, tt.mockPage, tt.text).Return(tt.searchQuestionsCode, tt.searchQuestionsResult)
			}
			r, _ := http.NewRequest("GET", "/searchQuestions?page=" + tt.page + "&text=" + tt.text, nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			a.Equal(http.StatusOK, w.Result().StatusCode)
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			a.Equal(want, res)
		})
	}
}

func TestControllerSearchAnswers(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	a := assert.New(t)

	var s controller.SearchController
	s.SetSearchService(mockSearchService)
	mux.HandleFunc("/searchAnswers", s.SearchAnswers)
	tests := []struct {
		name string
		token string
		page string
		text string
		mockSearchAnswers bool
		mockPage int64
		searchAnswersCode int8
		searchAnswersResult interface{}
		wantCode int8
		wantResult interface{}
	} {
		{
			"Normal",
			"token",
			"6",
			"frozen",
			true,
			6,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Failed Parse",
			"token",
			"6sdf",
			"frozen",
			false,
			6,
			service.Succeeded,
			"result",
			service.Failed,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSearchAnswers {
				mockSearchService.EXPECT().SearchAnswers(tt.token, tt.mockPage, tt.text).Return(tt.searchAnswersCode, tt.searchAnswersResult)
			}
			r, _ := http.NewRequest("GET", "/searchAnswers?page=" + tt.page + "&text=" + tt.text, nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			a.Equal(http.StatusOK, w.Result().StatusCode)
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			a.Equal(want, res)
		})
	}
}

func TestControllerSearchUsers(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	a := assert.New(t)

	var s controller.SearchController
	s.SetSearchService(mockSearchService)
	mux.HandleFunc("/searchUsers", s.SearchUsers)
	tests := []struct {
		name string
		token string
		page string
		text string
		mockSearchUsers bool
		mockPage int64
		searchUsersCode int8
		searchUsersResult interface{}
		wantCode int8
		wantResult interface{}
	} {
		{
			"Normal",
			"token",
			"6",
			"frozen",
			true,
			6,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Failed Parse",
			"token",
			"6sdf",
			"frozen",
			false,
			6,
			service.Succeeded,
			"result",
			service.Failed,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSearchUsers {
				mockSearchService.EXPECT().SearchUsers(tt.token, tt.mockPage, tt.text).Return(tt.searchUsersCode, tt.searchUsersResult)
			}
			r, _ := http.NewRequest("GET", "/searchUsers?page=" + tt.page + "&text=" + tt.text, nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			a.Equal(http.StatusOK, w.Result().StatusCode)
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			a.Equal(want, res)
		})
	}
}

func TestControllerHotList(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	a := assert.New(t)

	var s controller.SearchController
	s.SetSearchService(mockSearchService)
	mux.HandleFunc("/hotlist", s.HotList)
	tests := []struct {
		name string
		token string
		mockHotList bool
		hotListCode int8
		hotListResult interface{}
		wantCode int8
		wantResult interface{}
	} {
		{
			"Normal",
			"token",
			true,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockHotList {
				mockSearchService.EXPECT().HotList(tt.token).Return(tt.hotListCode, tt.hotListResult)
			}
			r, _ := http.NewRequest("GET", "/hotlist", nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			a.Equal(http.StatusOK, w.Result().StatusCode)
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			a.Equal(want, res)
		})
	}
}

func TestControllerSearch(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSearchService := mock.NewMockSearchService(mockCtrl)
	a := assert.New(t)

	var s controller.SearchController
	s.SetSearchService(mockSearchService)
	mux.HandleFunc("/search", s.Search)

	tests := []struct {
		name string
		token string
		text string
		mockSearch bool
		searchCode int8
		searchResult interface{}
		wantCode int8
		wantResult interface{}
	} {
		{
			"Normal",
			"token",
			"text",
			true,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSearch {
				mockSearchService.EXPECT().Search(tt.token, tt.text).Return(tt.searchCode, tt.searchResult)
			}
			r, _ := http.NewRequest("GET", "/search?text=" + tt.text, nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			a.Equal(http.StatusOK, w.Result().StatusCode)
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			a.Equal(want, res)
		})
	}
}