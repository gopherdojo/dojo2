package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func getFortune() string {
	var result string
	switch rand.Intn(6) {
	case 0:
		result = "凶"
	case 1, 2:
		result = "小吉"
	case 3, 4:
		result = "中吉"
	case 5:
		result = "大吉"
	}
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getFortune())
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
