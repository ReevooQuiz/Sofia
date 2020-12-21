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
