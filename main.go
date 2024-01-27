package main

import (
	"log"
	"net/http"
	"time"

	"github.com/linehk/gin-blog/config"
	"github.com/linehk/gin-blog/model"
	"github.com/linehk/gin-blog/router"
)

func main() {
	model.Setup()

	server := &http.Server{
		Addr:           config.Raw.String("ADDR"),
		Handler:        router.Setup(),
		ReadTimeout:    time.Duration(config.Raw.Int("READTIMEOUT") * int(time.Second)),
		WriteTimeout:   time.Duration(config.Raw.Int("WRITETIMEOUT") * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
