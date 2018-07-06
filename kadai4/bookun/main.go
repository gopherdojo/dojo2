package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var results = []string{"大吉", "吉", "凶", "大凶"}

// Omikuji have Date and Result
type Omikuji struct {
	Date   time.Time
	Result string `json:"result"`
}

// NewOmikuji generate Omikuji struct
func NewOmikuji(date time.Time) *Omikuji {
	o := Omikuji{Date: date}
	return &o
}

// Adapter is function
type Adapter func(http.Handler) http.Handler

// Adapt function takes the handler you want to adapt
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// SetHeader set "content-type" header
func SetHeader() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			h.ServeHTTP(w, r)
		})
	}
}

func (o *Omikuji) handle(w http.ResponseWriter, r *http.Request) {
	if o.Date.Month() == 1 && (o.Date.Day() >= 1 && o.Date.Day() <= 3) {
		o.Result = "大吉"
	} else {
		o.Result = results[rand.Int()%len(results)]
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(o); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, buf.String())
}

func main() {
	o := NewOmikuji(time.Now())
	handler := http.HandlerFunc(o.handle)
	http.Handle("/", Adapt(handler, SetHeader()))
	http.ListenAndServe(":3000", nil)
}
