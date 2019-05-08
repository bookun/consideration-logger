package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bookun/consideration-of-logger/logger"
)

type Handler struct {
	//Logger *zap.Logger
	Hoge string
}

func (h *Handler) Sample(w http.ResponseWriter, r *http.Request) {
	logMap := make(map[string]string)
	logMap["handler"] = "hello handler"
	logMap["time"] = time.Now().String()
	logger.Info("handler", logMap)
	fmt.Fprintln(w, h.Hoge)
}
