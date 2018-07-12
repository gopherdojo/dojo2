package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

type nower interface {
	Now() time.Time
}

type intner interface {
	Intn(int) int
}

type response struct {
	Result string `json:"result"`
}

type omikuji struct {
	nower
	intner
	response
}

func (o omikuji) open(w http.ResponseWriter, r *http.Request) {
	o.pickUp()
	buf := encodeJson(&o)
	fmt.Fprintf(w, buf.String())
}

func encodeJson(p *omikuji) bytes.Buffer {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p.response); err != nil {
		log.Fatal(err)
	}
	return buf
}

func (o *omikuji) pickUp() {
	if isOsyougatu(o.Now()) {
		o.Result = getDaikiti()
	} else {
		o.Result = box[o.Intn(5)]
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
