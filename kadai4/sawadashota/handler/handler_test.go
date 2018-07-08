package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo2/kadai4/sawadashota/fortune"
	"github.com/gopherdojo/dojo2/kadai4/sawadashota/handler"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Errorf("expect: %d but actual: %d", http.StatusOK, rw.StatusCode)
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal(err)
	}

	var hr handler.HandleResponse
	buf := bytes.NewReader([]byte(b))
	dec := json.NewDecoder(buf)

	if err := dec.Decode(&hr); err != nil {
		t.Fatal(err)
	}

	if _, err := fortune.Parse(hr.Result); err != nil {
		t.Error(err)
	}
}
