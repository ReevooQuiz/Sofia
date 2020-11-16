package main

import (
	"github.com/zhanghanchong/qa-service/controller"
	"github.com/zhanghanchong/qa-service/service"
	"sync"
)

func main() {
	w := &sync.WaitGroup{}
	w.Add(1)
	c := &controller.QaController{}
	c.Init(w, &service.QuestionsServiceImpl{})
	w.Add(1)
	w.Wait()
}
