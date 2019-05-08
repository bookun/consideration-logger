package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type M2 struct {
	Logger *zap.Logger
}

func (m *M2) Process2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Info("m2", zap.String("m2", "hello m2"))
		next.ServeHTTP(w, r)
	}
}
