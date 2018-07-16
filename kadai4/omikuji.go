package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gopherdojo/dojo2/kadai4/fortune"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var dq fortune.DefaultQlock
	fs := fortune.FortuneSelector{Qlock: dq}
	f := fs.SelectFortune()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(f); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())
}

func main() {
	fmt.Println("Web Server Starting...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
