package middleware

import (
	"net/http"

	"github.com/bookun/consideration-of-logger/util"
	"go.uber.org/zap"
)

//type M1 struct {
//	Logger *zap.Logger
//}

func Process1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.Logger.Info("m1", zap.String("m1", "hello m1"))
		next.ServeHTTP(w, r)
	}
}
