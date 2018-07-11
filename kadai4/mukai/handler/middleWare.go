package handler

import "net/http"

func HeaderMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	})
}
