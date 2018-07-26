package middleware

import (
	"dojo2/kadai4/po3rin/domain"
	"encoding/json"
	"net/http"
	"time"
)

// CheckNewYear - if now is new year, return Daikichi
func CheckNewYear(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		if t.Year() == 1 && t.Day() <= 3 {
			res := &domain.Response{
				Code:   0,
				Result: "大吉",
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
		next.ServeHTTP(w, r)
	}
}
