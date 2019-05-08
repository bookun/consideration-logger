package middleware

import (
	"net/http"

	"github.com/bookun/consideration-of-logger/logger"
)

func Process1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logMap := make(map[string]string)
		logMap["m1"] = "hello p1"
		logger.Info("m1", logMap)
		next.ServeHTTP(w, r)
	}
}
