package controller

import (
	"encoding/json"
	"github.com/SKFE396/qa-service/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
)

type ServerResponse struct {
	Code   int8        `json:"code"`
	Result interface{} `json:"result"`
}

type QaController struct {
	qaService service.QaService
}

func (q *QaController) SetQaService(qaService service.QaService) {
	q.qaService = qaService
}

func (q *QaController) Init(group *sync.WaitGroup, qaService service.QaService) (server *http.Server) {
	q.qaService = qaService
	err := q.qaService.Init(nil, nil)
	if err != nil {
		log.Info(err)
	}
	server = &http.Server{Addr: ":9093"}
	http.HandleFunc("/questions", q.Questions)
	http.HandleFunc("/question", q.QuestionDetail)
	http.HandleFunc("/answers", q.Answers)
	http.HandleFunc("/answer", q.AnswerDetail)
	http.HandleFunc("/comments", q.Comments)
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}

func (q *QaController) Destruct() {
	q.qaService.Destruct()
}

func (q *QaController) Questions(w http.ResponseWriter, r *http.Request) {
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			var page int64
			page, err = strconv.ParseInt(r.FormValue("page"), 10, 32)
			token := r.Header.Get("Authorization")
			if err == nil {
				code, result := q.qaService.MainPage(token, page)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
	if r.Method == "POST" {
		var req service.ReqQuestionsPost
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.AddQuestion(token, req)
			response.Code = code
			response.Result = result
			object, _ := json.Marshal(response)
			_, _ = w.Write(object)
			return
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
	if r.Method == "PUT" {
		var req service.ReqQuestionsPut
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.ModifyQuestion(token, req)
			response.Code = code
			response.Result = result
			object, _ := json.Marshal(response)
			_, _ = w.Write(object)
			return
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (q *QaController) QuestionDetail(w http.ResponseWriter, r *http.Request) {
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			var qid int64
			qid, err = strconv.ParseInt(r.FormValue("qid"), 10, 32)
			token := r.Header.Get("Authorization")
			if err == nil {
				code, result := q.qaService.QuestionDetail(token, qid)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (q *QaController) Answers(w http.ResponseWriter, r *http.Request) {
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			log.Warn("1 passed")
			qid, qidErr := strconv.ParseInt(r.FormValue("qid"), 10, 64)
			page, pageErr := strconv.ParseInt(r.FormValue("page"), 10, 64)
			sort, sortErr := strconv.ParseInt(r.FormValue("sort"), 10, 8)
			token := r.Header.Get("Authorization")
			if qidErr == nil && pageErr == nil && sortErr == nil {
				code, result := q.qaService.ListAnswers(token, qid, page, int8(sort))
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
	if r.Method == "POST" {
		var req service.ReqAnswersPost
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.AddAnswer(token, req)
			response.Code = code
			response.Result = result
			object, _ := json.Marshal(response)
			_, _ = w.Write(object)
			return
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
	if r.Method == "PUT" {
		var req service.ReqAnswersPut
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.ModifyAnswer(token, req)
			response.Code = code
			response.Result = result
			object, _ := json.Marshal(response)
			_, _ = w.Write(object)
			return
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (q *QaController) AnswerDetail(w http.ResponseWriter, r *http.Request) {
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			var aid int64
			aid, err = strconv.ParseInt(r.FormValue("aid"), 10, 32)
			token := r.Header.Get("Authorization")
			if err == nil {
				code, result := q.qaService.AnswerDetail(token, aid)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (q *QaController) Comments(w http.ResponseWriter, r *http.Request) {
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			aid, aidErr := strconv.ParseInt(r.FormValue("aid"), 10, 64)
			page, pageErr := strconv.ParseInt(r.FormValue("page"), 10, 64)
			token := r.Header.Get("Authorization")
			if aidErr == nil && pageErr == nil {
				code, result := q.qaService.GetComments(token, aid, page)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
	if r.Method == "POST" {
		var req service.ReqCommentsPost
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.AddComment(token, req)
			response.Code = code
			response.Result = result
			object, _ := json.Marshal(response)
			_, _ = w.Write(object)
			return
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}
