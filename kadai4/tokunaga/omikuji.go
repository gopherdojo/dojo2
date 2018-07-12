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

const daikiti = 0

var box = map[int]string{
	0: "大吉",
	1: "中吉",
	2: "小吉",
	3: "凶",
	4: "大凶",
}

var syougatu = [...]string{
	"01-01",
	"01-02",
	"01-03",
}

type timer interface {
	Now() time.Time
}

type timeWrapper struct{}

func (t timeWrapper) Now() time.Time {
	return time.Now()
}

type Omikuji struct {
	Result string `json:"result"`
}

func (o Omikuji) open(w http.ResponseWriter, r *http.Request) {
	o.pickUp(timeWrapper{})
	buf := encodeJson(&o)
	fmt.Fprintf(w, buf.String())
}

func encodeJson(p *Omikuji) bytes.Buffer {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	return buf
}

func (o *Omikuji) pickUp(timer timer) {
	if isOsyougatu(timer.Now()) {
		o.Result = getDaikiti()
	} else {
		o.Result = box[rand.Intn(5)]
	}

}

func isOsyougatu(date time.Time) bool {
	day := date.Format("01-02")
	for _, sanganichi := range syougatu {
		if day == sanganichi {
			return true
		}
	}
	return false
}

func getDaikiti() string {
	return box[daikiti]
}
