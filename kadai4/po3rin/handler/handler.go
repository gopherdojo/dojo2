package handler

import (
	"dojo2/kadai4/po3rin/domain"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// Omikuzi - write fortune at random
func Omikuzi(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(6)
	var fortune string
	switch i {
	case 0:
		fortune = "大吉"
	case 1:
		fortune = "中吉"
	case 2:
		fortune = "小吉"
	case 3:
		fortune = "吉"
	case 4:
		fortune = "凶"
	case 5:
		fortune = "大凶"
	}
	res := &domain.Response{
		Code:   0,
		Result: fortune,
	}
	resjson, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resjson)
}
