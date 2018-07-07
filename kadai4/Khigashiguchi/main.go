package main

import (
	"bytes"
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"time"
)

var box = map[int]string{
	0: "大吉",
	1: "中吉",
	2: "中吉",
	3: "小吉",
	4: "小吉",
	5: "小吉",
	6: "凶",
	7: "凶",
	8: "大凶",
}

// Fortune おみくじ結果
type Fortune struct {
	Result string `json:"result"`
}

func (f *Fortune) setResult(t time.Time) {
	var result string
	if t.Month() == 1 && (t.Day() >= 1 && t.Day() <= 3) {
		result = "大吉"
	} else {
		result = box[rand.Intn(len(box))]
	}
	f.Result = result
}

func init() {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		log.Fatalf("Failed to init package. %s", err)
	}
	rand.Seed(seed.Int64())
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	f := &Fortune{}

	t := time.Now()
	f.setResult(t)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(f); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	fmt.Fprint(w, buf.String())
}

func main() {
	http.HandleFunc("/fortune", fortuneHandler)
	http.ListenAndServe(":8080", nil)
}
