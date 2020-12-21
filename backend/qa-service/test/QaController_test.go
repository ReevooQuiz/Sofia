package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/SKFE396/qa-service/controller"
	"github.com/SKFE396/qa-service/mock"
	"github.com/SKFE396/qa-service/service"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

const (
	Succeeded = iota
	Failed    = iota
	Expired   = iota
)

func TestControllerInit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)
	mockQaService.EXPECT().Init(nil, nil).AnyTimes()
	tests := []struct {
		name string
	}{
		{"Initialize"},
	}
	q := controller.QaController{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpServerExitDone := &sync.WaitGroup{}
			httpServerExitDone.Add(1)
			server := q.Init(httpServerExitDone, mockQaService)
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
	mockQaService := mock.NewMockQaService(mockCtrl)
	mockQaService.EXPECT().Destruct()
	var q controller.QaController
	q.SetQaService(mockQaService)
	q.Destruct()
}

func TestControllerQuestions(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/questions", q.Questions)

	/******************************************* GET *********************************************/
	t.Log("Testing GET")
	getTests := []struct {
		name       string
		page       string
		token      string
		mock       bool
		mockPage   int64
		mockToken  string
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{"Normal", "0", "token", true, 0, "token", Succeeded, "mock result", Succeeded, "mock result"},
		{"Invalid page", "234h45", "token", false, 0, "", 0, nil, Failed, nil},
		{"Expired", "0", "token", true, 0, "token", Expired, nil, Expired, nil},
	}
	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().MainPage(tt.mockToken, tt.mockPage).Return(tt.mockCode, tt.mockResult)
			}
			r, _ := http.NewRequest("GET", "/questions?page="+tt.page, nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != http.StatusOK {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, http.StatusOK)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			if res != want {
				t.Errorf("Actual: %v, expect: %v.", res, want)
			}
		})
	}

	/******************************************* POST *********************************************/
	t.Log("Testing POST")
	postTests := []struct {
		name       string
		token      string
		req        service.ReqQuestionsPost
		mock       bool
		mockToken  string
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"token",
			service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "category",
				Labels:   []string{"math"},
			},
			true,
			"token",
			Succeeded,
			"mock result",
			Succeeded,
			"mock result",
		},
		{
			"Expired",
			"token",
			service.ReqQuestionsPost{
				Title:    "title",
				Content:  "content",
				Category: "category",
				Labels:   []string{},
			},
			true,
			"token",
			Expired,
			nil,
			Expired,
			nil,
		},
	}
	for _, tt := range postTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().AddQuestion(tt.mockToken, tt.req).Return(tt.mockCode, tt.mockResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("POST", "/questions", bytes.NewReader(body))
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != http.StatusOK {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, http.StatusOK)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			if res != want {
				t.Errorf("Actual: %v, expect: %v.", res, want)
			}
		})
	}

	/******************************************* PUT *********************************************/
	t.Log("Testing PUT")
	putTests := []struct {
		name       string
		token      string
		req        service.ReqQuestionsPut
		mock       bool
		mockToken  string
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"token",
			service.ReqQuestionsPut{
				Qid: "234234",
				Title:    "title",
				Content:  "content",
				Category: "category",
				Labels:   []string{"math"},
			},
			true,
			"token",
			Succeeded,
			"mock result",
			Succeeded,
			"mock result",
		},
		{
			"Expired",
			"token",
			service.ReqQuestionsPut{
				Qid:"3434",
				Title:    "title",
				Content:  "content",
				Category: "category",
				Labels:   []string{},
			},
			true,
			"token",
			Expired,
			nil,
			Expired,
			nil,
		},
	}
	for _, tt := range putTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().ModifyQuestion(tt.mockToken, tt.req).Return(tt.mockCode, tt.mockResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("PUT", "/questions", bytes.NewReader(body))
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != http.StatusOK {
				t.Errorf("Actual: %v, expect: %v.", w.Result().StatusCode, http.StatusOK)
			}
			responseBody := make([]byte, w.Body.Len())
			_, _ = w.Body.Read(responseBody)
			var res controller.ServerResponse
			_ = json.Unmarshal(responseBody, &res)
			want := controller.ServerResponse{Code: tt.wantCode, Result: tt.wantResult}
			if res != want {
				t.Errorf("Actual: %v, expect: %v.", res, want)
			}
		})
	}
}
