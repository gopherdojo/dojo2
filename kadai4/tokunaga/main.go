// handler は、 HTTP リクエストの情報を返します。
package main

import (
	crand "crypto/rand"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
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

func main() {
	omikuji := Omikuji{}
	http.HandleFunc("/omikuji", omikuji.open)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
