package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var rescase = []string{
	`{"code":0,"result":"大吉"}`,
	`{"code":0,"result":"中吉"}`,
	`{"code":0,"result":"小吉"}`,
	`{"code":0,"result":"吉"}`,
	`{"code":0,"result":"凶"}`,
	`{"code":0,"result":"大凶"}`,
}

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuzi", nil)
	Omikuzi(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	for _, v := range rescase {
		if v == string(b) {
			return
		}
	}
	t.Fatalf("unexpected response: %s", string(b))
}
