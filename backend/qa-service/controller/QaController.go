package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/qa-service/entity"
	"github.com/zhanghanchong/qa-service/service"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"sync"
	"time"
)

type QaController struct {
	labelsService    service.LabelsService
	questionsService service.QuestionsService
}

func (q *QaController) Init(group *sync.WaitGroup, labelsService service.LabelsService, questionsService service.QuestionsService) (server *http.Server) {
	q.labelsService = labelsService
	q.questionsService = questionsService
	server = &http.Server{Addr: ":9090"}
	http.HandleFunc("/questions", q.Questions)
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}

func (q *QaController) Questions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req struct {
			Raiser   bson.ObjectId `json:"raiser"`
			Title    string        `json:"title"`
			Content  string        `json:"content"`
			Category string        `json:"category"`
			Tags     []string      `json:"tags"`
		}
		var res struct {
			Code   int8 `json:"code"`
			Result struct {
				Qid bson.ObjectId `json:"qid"`
			} `json:"result"`
		}
		err := q.labelsService.Init()
		defer q.labelsService.Destruct()
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Qid = ""
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		q.questionsService.Init()
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Qid = ""
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		var question entity.Questions
		question.Raiser = req.Raiser
		question.Title = req.Title
		question.Content = req.Content
		question.Category = req.Category
		question.AcceptedAnswer = bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		question.AnswerCount = 0
		question.ViewCount = 0
		question.FavoriteCount = 0
		question.Time = time.Now()
		question.Qid, err = q.questionsService.Insert(question)
		if err != nil {
			log.Info(err)
			res.Code = 1
			res.Result.Qid = bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		for _, tag := range req.Tags {
			var label entity.Labels
			label.Title = tag
			label.Lid, err = q.labelsService.Insert(label)
			if err != nil {
				label, err = q.labelsService.FindByTitle(tag)
				if err != nil {
					log.Info(err)
					res.Code = 1
					res.Result.Qid = bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
					object, _ := json.Marshal(res)
					_, _ = w.Write(object)
					return
				}
			}
			log.Info(label.Lid)
		}
	}
}
