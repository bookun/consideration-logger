package handler

import (
	"fmt"
	"net/http"

	"github.com/bookun/consideration-of-logger/logger"
)

type Handler struct {
	//Logger *zap.Logger
	Hoge string
}

func (h *Handler) Sample(w http.ResponseWriter, r *http.Request) {
	logger.Info("handler", "handler", "hello sample")
	fmt.Fprintln(w, h.Hoge)
}
