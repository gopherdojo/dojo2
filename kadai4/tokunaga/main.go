// handler は、 HTTP リクエストの情報を返します。
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
)

var box = map[int]string{
	0: "大吉",
	1: "中吉",
	2: "小吉",
	3: "凶",
	4: "大凶",
}

func init() {
	if err := serRandSeed(); err != nil {
		log.Fatal(err)
	}
}

func serRandSeed() error {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return err
}

func main() {
	omikuji := omikuji{}
	http.HandleFunc("/omikuji", omikuji.open)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type omikuji struct {
	Result string `json:"result"`
}

func (o omikuji) open(w http.ResponseWriter, r *http.Request) {
	o.pickUp()
	p := &o
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, buf.String())
}

func (o *omikuji)pickUp() {
	o.Result = box[rand.Intn(5)]
}