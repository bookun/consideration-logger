package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type M1 struct {
	Logger *zap.Logger
}

func (m *M1) Process1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Info("m1", zap.String("m1", "hello m1"))
		next.ServeHTTP(w, r)
	}
}
