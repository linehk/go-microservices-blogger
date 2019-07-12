package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/linehk/gin-blog/config"
	"github.com/linehk/gin-blog/model"
	"github.com/linehk/gin-blog/router"
)

func main() {
	config.Setup()
	model.Setup()

	addr := fmt.Sprintf("localhost:%d", config.Server.HttpPort)
	handler := router.Setup()
	readTimeout := config.Server.ReadTimeout
	writeTimeout := config.Server.WriteTimeout
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Fatal(server.ListenAndServe())
}
