package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gopherdojo/dojo2/kadai4/int128/routes"
)

func main() {
	addr := ":8000"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	s := http.Server{
		Addr:    addr,
		Handler: routes.New(),
	}
	log.Printf("Listening on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
