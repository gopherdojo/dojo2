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

type Fortune struct {
	Content string `json:"content"`
}

var fortuneList = []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}

func getCurrentTime() time.Time {
	return time.Now()
}

func selectFortune() string {
	t := getCurrentTime()
	if t.Month() == 1 && (1 <= t.Day() && t.Day() <= 3) {
		return selectFortuneOnlyDaikichi()
	}
	return selectFortuneRandom()
}

func selectFortuneOnlyDaikichi() string {
	return "大吉"
}

func selectFortuneRandom() string {
	rand.Seed(time.Now().UnixNano())
	return fortuneList[rand.Int()%len(fortuneList)]
}

func handler(w http.ResponseWriter, r *http.Request) {
	sf := selectFortune()
	f := &Fortune{Content: sf}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(f); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
