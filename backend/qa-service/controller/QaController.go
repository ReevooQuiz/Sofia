package controller

import (
	"bytes"
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
	qaService service.QaService
}

func (q *QaController) Init(group *sync.WaitGroup, qaService service.QaService) (server *http.Server) {
	q.qaService = qaService
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
	if r.Method == "GET" {
		type Owner struct {
			UserId   bson.ObjectId `json:"user_id"`
			UserName string        `json:"user_name"`
			UserIcon string        `json:"user_icon"`
		}
		type Tag struct {
			Tid   int64  `json:"tid"`
			Title string `json:"title"`
		}
		type Answer struct {
			Aid            bson.ObjectId `json:"aid"`
			Owner          Owner         `json:"owner"`
			LikeCount      int64         `json:"like_count"`
			CriticismCount int64         `json:"criticism_count"`
			ApprovalCount  int64         `json:"approval_count"`
			CommentCount   int64         `json:"comment_count"`
			Content        string        `json:"content"`
		}
		var res struct {
			Code        int8     `json:"code"`
			Owner       Owner    `json:"owner"`
			Title       string   `json:"title"`
			Content     string   `json:"content"`
			AnswerCount int64    `json:"answer_count"`
			FollowCount int64    `json:"follow_count"`
			ViewCount   int64    `json:"view_count"`
			Tags        []Tag    `json:"tags"`
			AnswerList  []Answer `json:"answer_list"`
		}
		err := q.qaService.Init()
		defer q.qaService.Destruct()
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		err = r.ParseForm()
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		qid := bson.ObjectIdHex(r.FormValue("qid"))
		var question entity.Questions
		question, err = q.qaService.FindQuestionByQid(qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		var answers []entity.Answers
		answers, err = q.qaService.FindAnswersByQid(qid)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		var questionLabels []entity.QuestionLabels
		questionLabels, err = q.qaService.FindQuestionLabelsByQid(qid.Hex())
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		res.Code = 0
		res.Owner.UserId = question.Raiser
		res.Owner.UserName = "root"
		res.Owner.UserIcon = "root"
		res.Title = question.Title
		res.Content = question.Content
		res.AnswerCount = question.AnswerCount
		res.FollowCount = question.FavoriteCount
		res.ViewCount = question.ViewCount
		for _, questionLabel := range questionLabels {
			var label entity.Labels
			label, err = q.qaService.FindLabelByLid(questionLabel.Lid)
			if err != nil {
				log.Info(err)
				res.Code = 1
				object, _ := json.Marshal(res)
				_, _ = w.Write(object)
				return
			}
			res.Tags = append(res.Tags, Tag{label.Lid, label.Title})
		}
		for _, answer := range answers {
			res.AnswerList = append(res.AnswerList, Answer{answer.Aid, Owner{answer.Answerer, "root", "root"}, answer.LikeCount, answer.CriticismCount, answer.ApprovalCount, answer.CommentCount, answer.Content})
		}
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
	}
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
		err := q.qaService.Init()
		defer q.qaService.Destruct()
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Info(err)
			res.Code = 1
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
		question.Qid, err = q.qaService.InsertQuestion(question)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		for _, tag := range req.Tags {
			var label entity.Labels
			label.Title = tag
			label.Lid, err = q.qaService.InsertLabel(label)
			if err != nil {
				label, err = q.qaService.FindLabelByTitle(tag)
				if err != nil {
					log.Info(err)
					res.Code = 1
					object, _ := json.Marshal(res)
					_, _ = w.Write(object)
					return
				}
			}
			var questionLabel entity.QuestionLabels
			questionLabel.Qid = question.Qid.Hex()
			questionLabel.Lid = label.Lid
			err = q.qaService.InsertQuestionLabel(questionLabel)
			if err != nil {
				log.Info(err)
				res.Code = 1
				object, _ := json.Marshal(res)
				_, _ = w.Write(object)
				return
			}
		}
		var requestBody struct {
			Content string `json:"content"`
		}
		requestBody.Content = question.Title
		var requestBodyJson []byte
		requestBodyJson, err = json.Marshal(requestBody)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		var request *http.Request
		request, err = http.NewRequest("POST", "http://192.168.126.131:3000/keyword", bytes.NewReader(requestBodyJson))
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		request.Header.Set("Content-Type", "application/json")
		client := http.Client{}
		var response *http.Response
		response, err = client.Do(request)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		responseBodyJson := make([]byte, response.ContentLength)
		_, err = response.Body.Read(responseBodyJson)
		var responseBody struct {
			Status string   `json:"status"`
			Data   []string `json:"data"`
		}
		err = json.Unmarshal(responseBodyJson, &responseBody)
		if err != nil {
			log.Info(err)
			res.Code = 1
			object, _ := json.Marshal(res)
			_, _ = w.Write(object)
			return
		}
		for _, keyword := range responseBody.Data {
			var kcard entity.Kcards
			kcard.Title = keyword
			kcard.Kid, err = q.qaService.InsertKcard(kcard)
			if err != nil {
				log.Info(err)
				res.Code = 1
				object, _ := json.Marshal(res)
				_, _ = w.Write(object)
				return
			}
		}
		res.Code = 0
		res.Result.Qid = question.Qid
		object, _ := json.Marshal(res)
		_, _ = w.Write(object)
	}
}
