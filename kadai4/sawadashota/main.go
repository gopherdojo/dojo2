package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gopherdojo/dojo2/kadai4/sawadashota/handler"
)

const DefaultPort = "8080"

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
}

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
