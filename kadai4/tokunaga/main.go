// handler は、 HTTP リクエストの情報を返します。
package main

import (
	crand "crypto/rand"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"time"
)

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

type timeWrapper struct{}
type randWrapper struct{}

func (t timeWrapper) Now() time.Time {
	return time.Now()
}

func (r randWrapper) Intn(n int) int {
	return rand.Intn(n)
}

func main() {
	omikuji := omikuji{nower: timeWrapper{}, intner: randWrapper{}}
	http.HandleFunc("/omikuji", omikuji.open)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
