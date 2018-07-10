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

func init() {
	if err := serRandSeed(); err != nil {
		log.Fatal(err)
	}
}

type timer interface {
	Now() time.Time
}

type timeWrapper struct{}

type omikuji struct {
	Result string `json:"result"`
}

func (t timeWrapper) Now() time.Time {
	return time.Now()
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

func (o omikuji) open(w http.ResponseWriter, r *http.Request) {
	o.pickUp(timeWrapper{})
	buf := encodeJson(&o)
	fmt.Fprintf(w, buf.String())
}

func encodeJson(p *omikuji) bytes.Buffer {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	return buf
}

func (o *omikuji) pickUp(timer timer) {
	if isOsyougatu(time.Now()) {
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
