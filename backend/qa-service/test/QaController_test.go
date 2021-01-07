package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/SKFE396/qa-service/controller"
	"github.com/SKFE396/qa-service/mock"
	"github.com/SKFE396/qa-service/service"
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
			time.Sleep(200 * time.Microsecond)
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

func TestControllerAnswers(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)
	a := assert.New(t)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/answers", q.Answers)

	/******************************************* GET *********************************************/
	t.Log("Testing GET")
	getTests := []struct {
		name              string
		token             string
		qid               string
		page              string
		sort              string
		mockListAnswers   bool
		mockQid           int64
		mockPage          int64
		mockSort          int8
		listAnswersCode   int8
		listAnswersResult interface{}
		wantCode          int8
		wantResult        interface{}
	}{
		{
			"Normal",
			"token",
			"234",
			"5",
			"1",
			true,
			234,
			5,
			1,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Wrong Parameters",
			"token",
			"234",
			"5sd",
			"1",
			false,
			0,
			0,
			0,
			service.Succeeded,
			"result",
			service.Failed,
			nil,
		},
	}
	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("GET - ", tt.name)
			if tt.mockListAnswers {
				mockQaService.EXPECT().ListAnswers(tt.token, tt.mockQid, tt.mockPage, tt.mockSort).Return(tt.listAnswersCode, tt.listAnswersResult)
			}
			r, _ := http.NewRequest("GET", "/answers?qid="+tt.qid+"&page="+tt.page+"&sort="+tt.sort, nil)
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

	/******************************************* POST *********************************************/
	t.Log("Testing POST")
	postTests := []struct {
		name            string
		token           string
		req             service.ReqAnswersPost
		mockAddAnswer   bool
		addAnswerCode   int8
		addAnswerResult interface{}
		wantCode        int8
		wantResult      interface{}
	}{
		{
			name:  "Normal",
			token: "token",
			req: service.ReqAnswersPost{
				Qid:     "234",
				Content: "content",
			},
			mockAddAnswer:   true,
			addAnswerCode:   service.Succeeded,
			addAnswerResult: "result",
			wantCode:        service.Succeeded,
			wantResult:      "result",
		},
		{
			name:  "Failed to modify",
			token: "token",
			req: service.ReqAnswersPost{
				Qid:     "234",
				Content: "content",
			},
			mockAddAnswer:   true,
			addAnswerCode:   service.Failed,
			addAnswerResult: nil,
			wantCode:        service.Failed,
			wantResult:      nil,
		},
	}
	for _, tt := range postTests {
		t.Run("Normal", func(t *testing.T) {
			if tt.mockAddAnswer {
				mockQaService.EXPECT().AddAnswer(tt.token, tt.req).Return(tt.addAnswerCode, tt.addAnswerResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("POST", "/answers", bytes.NewReader(body))
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

	/******************************************* PUT *********************************************/
	t.Log("Testing PUT")
	putTests := []struct {
		name               string
		token              string
		req                service.ReqAnswersPut
		mockModifyAnswer   bool
		modifyAnswerCode   int8
		modifyAnswerResult interface{}
		wantCode           int8
		wantResult         interface{}
	}{
		{
			name:  "Normal",
			token: "token",
			req: service.ReqAnswersPut{
				Aid:     "2346",
				Content: "new content",
			},
			mockModifyAnswer:   true,
			modifyAnswerCode:   service.Succeeded,
			modifyAnswerResult: nil,
			wantCode:           service.Succeeded,
			wantResult:         nil,
		},
		{
			name:  "Failed to modify",
			token: "token",
			req: service.ReqAnswersPut{
				Aid:     "2346",
				Content: "new content",
			},
			mockModifyAnswer:   true,
			modifyAnswerCode:   service.Failed,
			modifyAnswerResult: nil,
			wantCode:           service.Failed,
			wantResult:         nil,
		},
	}
	for _, tt := range putTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockModifyAnswer {
				mockQaService.EXPECT().ModifyAnswer(tt.token, tt.req).Return(tt.modifyAnswerCode, tt.modifyAnswerResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("PUT", "/answers", bytes.NewReader(body))
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

func TestControllerDeleteAnswer(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/delete_answer", q.DeleteAnswer)

	/******************************************* POST *********************************************/
	t.Log("Testing POST")
	delTests := []struct {
		name       string
		token      string
		aid        int64
		req        service.ReqAnswersDelete
		mock       bool
		mockToken  string
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			name:       "Normal",
			token:      "token",
			aid:        345,
			req:        service.ReqAnswersDelete{Aid: "345"},
			mock:       true,
			mockToken:  "token",
			mockCode:   service.Succeeded,
			mockResult: "mock result",
			wantCode:   service.Succeeded,
			wantResult: "mock result",
		},
		{
			name:      "Expired",
			token:     "token",
			aid:       345,
			req:       service.ReqAnswersDelete{Aid: "345"},
			mock:      true,
			mockToken: "token",
			mockCode:  service.Expired,
			wantCode:  service.Expired,
		},
	}
	for _, tt := range delTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().DeleteAnswer(tt.mockToken, tt.req).Return(tt.mockCode, tt.mockResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("POST", "/delete_answer", bytes.NewReader(body))
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
		category   string
		mock       bool
		mockPage   int64
		mockToken  string
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"0",
			"token",
			"life",
			true,
			0,
			"token",
			service.Succeeded,
			"mock result",
			service.Succeeded,
			"mock result",
		},
		{
			"Invalid page",
			"234h45",
			"token",
			"life",
			false,
			0,
			"",
			0,
			nil,
			service.Failed,
			nil,
		},
		{
			"Expired",
			"0",
			"token",
			"life",
			true,
			0,
			"token",
			service.Expired,
			nil,
			service.Expired,
			nil,
		},
	}
	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().MainPage(tt.mockToken, tt.category, tt.mockPage).Return(tt.mockCode, tt.mockResult)
			}
			r, _ := http.NewRequest("GET", "/questions?page="+tt.page+"&category="+tt.category, nil)
			r.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			if w.Result().StatusCode != http.StatusOK {
				t.Errorf("Actual: %v, expect : %v.", w.Result().StatusCode, http.StatusOK)
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
			service.Succeeded,
			"mock result",
			service.Succeeded,
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
			service.Expired,
			nil,
			service.Expired,
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
				Qid:      "234234",
				Title:    "title",
				Content:  "content",
				Category: "category",
				Labels:   []string{"math"},
			},
			true,
			"token",
			service.Succeeded,
			"mock result",
			service.Succeeded,
			"mock result",
		},
		{
			"Expired",
			"token",
			service.ReqQuestionsPut{
				Qid:      "3434",
				Title:    "title",
				Content:  "content",
				Category: "category",
				Labels:   []string{},
			},
			true,
			"token",
			service.Expired,
			nil,
			service.Expired,
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

func TestControllerDisableQuestion(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/disable_question", q.DisableQuestion)

	/******************************************* POST *********************************************/
	t.Log("Testing POST")
	delTests := []struct {
		name       string
		token      string
		qid        int64
		req        service.ReqQuestionsDelete
		mock       bool
		mockToken  string
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			name:       "Normal",
			token:      "token",
			qid:        345,
			req:        service.ReqQuestionsDelete{Qid: "345"},
			mock:       true,
			mockToken:  "token",
			mockCode:   service.Succeeded,
			mockResult: "mock result",
			wantCode:   service.Succeeded,
			wantResult: "mock result",
		},
		{
			name:      "Expired",
			token:     "token",
			qid:       345,
			req:       service.ReqQuestionsDelete{Qid: "345"},
			mock:      true,
			mockToken: "token",
			mockCode:  service.Expired,
			wantCode:  service.Expired,
		},
	}
	for _, tt := range delTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().DeleteQuestion(tt.mockToken, tt.req).Return(tt.mockCode, tt.mockResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("POST", "/disable_question", bytes.NewReader(body))
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

func TestControllerAnswerDetail(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)
	a := assert.New(t)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/answer", q.AnswerDetail)

	tests := []struct {
		name       string
		aid        string
		token      string
		mock       bool
		mockAid    int64
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"234",
			"token",
			true,
			234,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Failed to get",
			"234",
			"token",
			true,
			234,
			service.Failed,
			nil,
			service.Failed,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().AnswerDetail(tt.token, tt.mockAid).Return(tt.mockCode, tt.mockResult)
			}
			r, _ := http.NewRequest("GET", "/answer?aid="+tt.aid, nil)
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

func TestControllerQuestionDetail(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)
	a := assert.New(t)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/question", q.QuestionDetail)

	tests := []struct {
		name       string
		qid        string
		token      string
		mock       bool
		mockQid    int64
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"234",
			"token",
			true,
			234,
			service.Succeeded,
			nil,
			service.Succeeded,
			nil,
		},
		{
			"Invalid qid string",
			"234hu",
			"token",
			false,
			0,
			service.Succeeded,
			nil,
			service.Failed,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().QuestionDetail(tt.token, tt.mockQid)
			}
			r, _ := http.NewRequest("GET", "/question?qid="+tt.qid, nil)
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

func TestControllerComments(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)
	a := assert.New(t)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/comments", q.Comments)

	/******************************************* GET *********************************************/
	getTests := []struct {
		name       string
		token      string
		aid        string
		page       string
		mock       bool
		mockAid    int64
		mockPage   int64
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"token",
			"234",
			"5",
			true,
			234,
			5,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Wrong Parameters",
			"token",
			"234df",
			"5",
			false,
			234,
			5,
			service.Failed,
			nil,
			service.Failed,
			nil,
		},
	}

	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().GetComments(tt.token, tt.mockAid, tt.mockPage).Return(tt.mockCode, tt.mockResult)
			}
			r, _ := http.NewRequest("GET", "/comments?aid="+tt.aid+"&page="+tt.page, nil)
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

	/******************************************* POST *********************************************/
	postTests := []struct {
		name       string
		token      string
		req        service.ReqCommentsPost
		mock       bool
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"token",
			service.ReqCommentsPost{
				Aid:     "3456",
				Content: "content",
			},
			true,
			service.Succeeded,
			"mock result",
			service.Succeeded,
			"mock result",
		},
	}
	for _, tt := range postTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().AddComment(tt.token, tt.req).Return(tt.mockCode, tt.mockResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("POST", "/comments", bytes.NewReader(body))
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

func TestControllerCriticisms(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQaService := mock.NewMockQaService(mockCtrl)
	a := assert.New(t)

	var q controller.QaController
	q.SetQaService(mockQaService)
	mux.HandleFunc("/criticisms", q.Criticisms)

	/******************************************* GET *********************************************/
	getTests := []struct {
		name       string
		token      string
		aid        string
		page       string
		mock       bool
		mockAid    int64
		mockPage   int64
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"token",
			"234",
			"5",
			true,
			234,
			5,
			service.Succeeded,
			"result",
			service.Succeeded,
			"result",
		},
		{
			"Wrong Parameters",
			"token",
			"234df",
			"5",
			false,
			234,
			5,
			service.Failed,
			nil,
			service.Failed,
			nil,
		},
	}

	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().GetCriticisms(tt.token, tt.mockAid, tt.mockPage).Return(tt.mockCode, tt.mockResult)
			}
			r, _ := http.NewRequest("GET", "/criticisms?aid="+tt.aid+"&page="+tt.page, nil)
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

	/******************************************* POST *********************************************/
	postTests := []struct {
		name       string
		token      string
		req        service.ReqCriticismsPost
		mock       bool
		mockCode   int8
		mockResult interface{}
		wantCode   int8
		wantResult interface{}
	}{
		{
			"Normal",
			"token",
			service.ReqCriticismsPost{
				Aid:     "3456",
				Content: "content",
			},
			true,
			service.Succeeded,
			"mock result",
			service.Succeeded,
			"mock result",
		},
	}
	for _, tt := range postTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock {
				mockQaService.EXPECT().AddCriticism(tt.token, tt.req).Return(tt.mockCode, tt.mockResult)
			}
			body, _ := json.Marshal(tt.req)
			r, _ := http.NewRequest("POST", "/criticisms", bytes.NewReader(body))
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
