package middleware

import (
	"net/http"

	"github.com/bookun/consideration-of-logger/logger"
)

func Process2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Error("m2", "m2", "hello m2")
		next.ServeHTTP(w, r)
	}
}
