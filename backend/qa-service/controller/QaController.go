package controller

import (
	"encoding/json"
	"github.com/SKFE396/qa-service/service"
	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", q.Questions)
	mux.HandleFunc("/question", q.QuestionDetail)
	mux.HandleFunc("/answers", q.Answers)
	mux.HandleFunc("/answer", q.AnswerDetail)
	mux.HandleFunc("/comments", q.Comments)
	mux.HandleFunc("/criticisms", q.Criticisms)
	mux.HandleFunc("/disable_question", q.DisableQuestion)
	mux.HandleFunc("/delete_answer", q.DeleteAnswer)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		Debug:            true,
	})
	handler := c.Handler(mux)
	go func() {
		defer group.Done()
		if err := http.ListenAndServe(":9093", handler); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}

func (q *QaController) Destruct() {
	q.qaService.Destruct()
}

func (q *QaController) Questions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			var page int64
			page, err = strconv.ParseInt(r.FormValue("page"), 10, 32)
			category := r.FormValue("category")
			token := r.Header.Get("Authorization")
			if err == nil {
				code, result := q.qaService.MainPage(token, category, page)
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

func (q *QaController) DisableQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "POST" {
		var req service.ReqQuestionsDelete
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.DeleteQuestion(token, req)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

func (q *QaController) Criticisms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			aid, aidErr := strconv.ParseInt(r.FormValue("aid"), 10, 64)
			page, pageErr := strconv.ParseInt(r.FormValue("page"), 10, 64)
			token := r.Header.Get("Authorization")
			if aidErr == nil && pageErr == nil {
				code, result := q.qaService.GetCriticisms(token, aid, page)
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
		var req service.ReqCriticismsPost
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.AddCriticism(token, req)
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

func (q *QaController) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "POST" {
		var req service.ReqAnswersDelete
		err = json.NewDecoder(r.Body).Decode(&req)
		if err == nil {
			token := r.Header.Get("Authorization")
			code, result := q.qaService.DeleteAnswer(token, req)
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
