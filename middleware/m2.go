package middleware

import (
	"net/http"

	"github.com/bookun/consideration-of-logger/logger"
)

func Process2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logMap := make(map[string]string)
		logMap["m2"] = "hello p2"
		logger.Error("m2", logMap)
		next.ServeHTTP(w, r)
	}
}
