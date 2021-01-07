package main

import (
	"search-service/controller"
	"search-service/service"
	"sync"
)

func main() {
	w := &sync.WaitGroup{}
	w.Add(1)
	c := &controller.SearchController{}
	c.Init(w, &service.SearchServiceImpl{})
	defer c.Destruct()
	w.Add(1)
	w.Wait()
}
