package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"time"

	"github.com/gopherdojo/dojo2/kadai4/cappyzawa/fortune"
	"go.mercari.io/go-httpdoc"
	"go.uber.org/zap"
)

func TestHandler(t *testing.T) {
	t.Helper()
	doc := &httpdoc.Document{
		Name:           "Omikuji API",
		ExcludeHeaders: []string{"Content-Length", "Accept-Encoding"},
	}
	defer func() {
		if err := doc.Generate("doc/omikuji.md"); err != nil {
			t.Fatalf("err: %s", err)
		}
	}()

	date := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	logger, _ := zap.NewProduction()
	mux := http.NewServeMux()
	mux.Handle("/", httpdoc.Record(&handler{
		date:   date,
		logger: logger,
	}, doc, &httpdoc.RecordOption{
		Description: "Draw a fortune",
		ExcludeHeaders: []string{
			"User-Agent",
			"Content-Length",
		},

		WithValidate: func(validator *httpdoc.Validator) {
			validator.ResponseBody(t, []httpdoc.TestCase{
				httpdoc.NewTestCase("Date", date, "Today"),
				httpdoc.NewTestCase("Result", fortune.DAIKICHI, "Fortune Result")},
				&FortuneResponse{},
			)
		},
	}))

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	req := testNewRequest(t, testServer.URL+"/")
	if _, err := http.DefaultClient.Do(req); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testNewRequest(t *testing.T, urlStr string) *http.Request {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	req.Header.Set("Content-type", "application/json")
	return req
}
