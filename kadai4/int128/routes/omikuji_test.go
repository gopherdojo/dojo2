package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo2/kadai4/int128/omikuji"
)

func TestOmikujiHandler_GET(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/omikuji", nil)
	h := &OmikujiHandler{omikuji.New()}
	h.ServeHTTP(w, r)
	res := w.Result()
	defer res.Body.Close()
	if http.StatusOK != res.StatusCode {
		t.Fatalf("res.StatusCode wants %d but %d", http.StatusOK, res.StatusCode)
	}
	d := json.NewDecoder(res.Body)
	var json OmikujiGetResponse
	if err := d.Decode(&json); err != nil {
		t.Fatal(err)
	}
}

func TestOmikujiHandler_POST(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/omikuji", nil)
	h := &OmikujiHandler{omikuji.New()}
	h.ServeHTTP(w, r)
	res := w.Result()
	defer res.Body.Close()
	if http.StatusMethodNotAllowed != res.StatusCode {
		t.Fatalf("res.StatusCode wants %d but %d", http.StatusMethodNotAllowed, res.StatusCode)
	}
}
