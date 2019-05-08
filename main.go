package main

import (
	"log"
	"net/http"

	"github.com/bookun/consideration-of-logger/handler"
	"github.com/bookun/consideration-of-logger/logger"
	"github.com/bookun/consideration-of-logger/middleware"
	"go.uber.org/zap"
)

func run() error {
	h := handler.Handler{Hoge: "sample"}
	http.HandleFunc("/", middleware.Process2(middleware.Process1(h.Sample)))

	return http.ListenAndServe(":8080", nil)
}

func main() {
	myZap, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger = myZap
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
