package handler

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type OmikujiHandler struct {
	TimeProvider TimeProvider
}
func (o OmikujiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := result{Data: o.omikuji()}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(result); err != nil {
		fmt.Fprintln(w, "error")
	}
	fmt.Fprintln(w, buf.String())
}

func (o OmikujiHandler) omikuji() string {
	time := o.TimeProvider.Time()
	if time.Month() == 1 && (1 <= time.Day() && time.Day() <= 3) {
		return "大吉"
	}
	i := rand.Intn(4)
	switch i {
	case 0: return "大吉"
	case 1: return "中吉"
	case 2: return "小吉"
	case 3: return "凶"
	default: return "凶"
	}
}

type result struct {
	Data string `json:"data"`
}

