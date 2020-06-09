package main

import (
	"log"
	"net/http"
	"time"

	"github.com/linehk/gin-blog/config"
	"github.com/linehk/gin-blog/model"
	"github.com/linehk/gin-blog/router"
)

var sc = config.Cfg.Server

func main() {
	model.Setup()

	server := &http.Server{
		Addr:           sc.Addr,
		Handler:        router.Setup(),
		ReadTimeout:    time.Duration(sc.ReadTimeout * int(time.Second)), // 转换成时间数据结构
		WriteTimeout:   time.Duration(sc.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
