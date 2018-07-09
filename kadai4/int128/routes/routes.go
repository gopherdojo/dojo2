package routes

import (
	"log"
	"net/http"

	"github.com/gopherdojo/dojo2/kadai4/int128/omikuji"
)

// New returns a handler with application routes.
func New() http.Handler {
	m := http.NewServeMux()
	m.Handle("/api/omikuji", &OmikujiHandler{omikuji.New()})
	return withLogging(m)
}

func withLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}
