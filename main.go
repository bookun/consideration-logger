package main

import (
	"log"
	"net/http"

	"github.com/bookun/consideration-of-logger/handler"
	"github.com/bookun/consideration-of-logger/middleware"
	"go.uber.org/zap"
)

func run() error {
	myZap, err := zap.NewProduction()
	if err != nil {
		return err
	}
	m1 := middleware.M1{Logger: myZap}
	m2 := middleware.M2{Logger: myZap}
	h := handler.Handler{Logger: myZap}
	http.HandleFunc("/", m2.Process2(m1.Process1(h.Sample)))

	return http.ListenAndServe(":8080", nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
