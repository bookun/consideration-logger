package middleware

import (
	"net/http"

	"github.com/bookun/consideration-of-logger/util"
	"go.uber.org/zap"
)

//type M2 struct {
//	Logger *zap.Logger
//}

func Process2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.Logger.Info("m2", zap.String("m2", "hello m2"))
		next.ServeHTTP(w, r)
	}
}
