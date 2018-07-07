package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gopherdojo/dojo2/kadai4/cappyzawa/fortune"

	"time"

	"go.uber.org/zap"
)

type FortuneResponse struct {
	fortune.Fortune
}

type handler struct {
	logger *zap.Logger
	date   time.Time
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.logger.Sync()
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.logger.Error("not allowed",
			zap.String("method", r.Method),
			zap.Int("code", http.StatusMethodNotAllowed))
		return
	}

	service := fortune.NewFortune(h.date)
	f := service.Draw()
	response := &FortuneResponse{f}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("encode error",
			zap.Int("code", http.StatusMethodNotAllowed))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	h.logger.Info("draw a fortune",
		zap.Int("code", http.StatusOK),
		zap.String("result", response.Result))
	w.Write([]byte(buf.String()))
}

func main() {
	logger, _ := zap.NewProduction()
	handler := handler{
		logger: logger,
		date:   time.Now(),
	}
	http.HandleFunc("/", handler.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
