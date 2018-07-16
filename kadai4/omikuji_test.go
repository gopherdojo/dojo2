package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo2/kadai4/fortune"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	f := fortune.Fortune{}
	if err := json.Unmarshal(b, &f); err != nil {
		t.Fatal("failed json unmarshal")
	}
	// fortuneListのうちどれかが返ってきていればOK
	var fortuneList = []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}
	exists := false
	for _, v := range fortuneList {
		if v == f.Content {
			exists = true
			break
		}
	}
	if exists == false {
		t.Fatalf("unexpected f.Content: %s", f.Content)
	}
}
