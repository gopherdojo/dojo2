package main

import (
	"dojo2/kadai4/mukai/handler"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.Handle("/", handler.HeaderMiddleWare(handler.OmikujiHandler{DateProvider: handler.NowTimeProvider{}}))
	http.ListenAndServe(":8080", nil)
}
