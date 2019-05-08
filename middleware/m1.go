package middleware

import (
	"net/http"

	"github.com/bookun/consideration-of-logger/logger"
)

func Process1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("m1", "m1", "hello m1")
		next.ServeHTTP(w, r)
	}
}
