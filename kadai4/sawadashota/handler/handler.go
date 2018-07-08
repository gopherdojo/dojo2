package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo2/kadai4/sawadashota/fortune"
)

var randSource rand.Source

type HandleResponse struct {
	Result string `json:"result"`
}

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f := fortune.Result(randSource, time.Now().Local())
	hr := &HandleResponse{Result: f.String()}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	if err := enc.Encode(hr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, buf.String())
}
