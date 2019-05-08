package handler

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	Logger *zap.Logger
}

func (h *Handler) Sample(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("handler", zap.String("handler", "hello sample"))
	fmt.Fprintln(w, "hello")
}
