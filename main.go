package main

import (
	"fmt"
	"net/http"

	"github.com/linehk/GinBlog/config"
	"github.com/linehk/GinBlog/model"
	"github.com/linehk/GinBlog/router"
)

func init() {
	config.Init()
	model.Init()
}

func main() {
	addr := fmt.Sprintf("localhost:%d", config.Server.HttpPort)
	handler := router.Init()
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
	server.ListenAndServe()
}
