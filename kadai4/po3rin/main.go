package main

import (
	"dojo2/kadai4/po3rin/handler"
	"dojo2/kadai4/po3rin/middleware"
	"net/http"
)

func main() {
	http.HandleFunc("/omikuzi", middleware.CheckNewYear(handler.Omikuzi))
	http.ListenAndServe(":8080", nil)
}
