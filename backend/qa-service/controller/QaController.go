package controller

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/qa-service/service"
	"net/http"
	"sync"
)

type QaController struct {
	questionsService service.QuestionsService
}

func (q *QaController) Init(group *sync.WaitGroup, questionsService service.QuestionsService) (server *http.Server) {
	q.questionsService = questionsService
	server = &http.Server{Addr: ":9090"}
	go func() {
		defer group.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}
