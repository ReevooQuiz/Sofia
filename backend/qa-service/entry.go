package main

import (
	"github.com/SKFE396/qa-service/controller"
	"github.com/SKFE396/qa-service/service"
	"sync"
)

func main() {
	w := &sync.WaitGroup{}
	w.Add(1)
	c := &controller.QaController{}
	c.Init(w, &service.QaServiceImpl{})
	defer c.Destruct()
	w.Add(1)
	w.Wait()
}
