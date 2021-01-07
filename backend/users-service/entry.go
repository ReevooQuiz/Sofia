package main

import (
	"github.com/zhanghanchong/users-service/controller"
	"github.com/zhanghanchong/users-service/service"
	"sync"
)

func main() {
	w := &sync.WaitGroup{}
	w.Add(1)
	c := &controller.UsersController{}
	c.Init(w, &service.UsersServiceImpl{})
	defer c.Destruct()
	w.Add(1)
	w.Wait()
}
