package handler

import (
	"fmt"
	"net/http"

	"github.com/bookun/consideration-of-logger/util"
	"go.uber.org/zap"
)

type Handler struct {
	//Logger *zap.Logger
	Hoge string
}

func (h *Handler) Sample(w http.ResponseWriter, r *http.Request) {
	util.Logger.Info("handler", zap.String("handler", "hello sample"))
	fmt.Fprintln(w, h.Hoge)
}
