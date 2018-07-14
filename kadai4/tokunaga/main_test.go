package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type testNower struct {
	now time.Time
}

type testIntner struct {
	randInt int
}

func (t testNower) Now() time.Time {
	return t.now
}
func (t testIntner) Intn(_ int) int {
	return t.randInt - 1
}

func TestOpen(t *testing.T) {
	shougatsu := testNower{now: time.Date(2018, 1, 1, 9, 11, 11, 11, time.UTC)}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuji", nil)
	o := omikuji{nower: shougatsu, intner: testIntner{randInt: 5}}
	o.open(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	expected := `{"result":"大吉"}`
	if s := strings.TrimSpace(string(b)); s != expected {
		t.Fatalf("want: %s, got: %s", expected, s)
	}
}

func TestEncodeJson(t *testing.T) {
	expected := `{"result":"大吉"}`
	o := omikuji{response: response{Result: "大吉"}}
	actual := o.encodeJson()
	if strings.TrimSpace(actual.String()) != expected {
		t.Errorf("want: o.encodeJson() = %s, got: %s ", expected, actual)
	}
}

func TestPickUp(t *testing.T) {
	nenmatsu := testNower{now: time.Date(2017, 12, 31, 23, 59, 11, 11, time.UTC)}
	shougatsu := testNower{now: time.Date(2018, 1, 1, 9, 11, 11, 11, time.UTC)}
	cases := []struct {
		name     string
		input    omikuji
		expected string
	}{
		{name: "お正月以外", input: omikuji{nower: nenmatsu, intner: testIntner{randInt: 5}}, expected: "大凶"},
		{name: "お正月", input: omikuji{nower: shougatsu, intner: testIntner{randInt: 5}}, expected: "大吉"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.input.pickUp()
			if actual := c.input.Result; actual != c.expected {
				t.Errorf("want o.Result = %v, got %v", c.expected, actual)
			}
		})
	}
}

func TestIsOsyougatu(t *testing.T) {
	cases := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{name: "12/31", input: time.Date(2017, 12, 31, 23, 59, 11, 11, time.UTC), expected: false},
		{name: "1/1", input: time.Date(2018, 1, 1, 9, 11, 11, 11, time.UTC), expected: true},
		{name: "1/2", input: time.Date(2018, 1, 2, 10, 12, 11, 11, time.UTC), expected: true},
		{name: "1/3", input: time.Date(2018, 1, 3, 11, 12, 11, 11, time.UTC), expected: true},
		{name: "1/4", input: time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC), expected: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := isOsyougatu(c.input); actual != c.expected {
				t.Errorf("want isOsyougatu(%s) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}

func TestGetDaikiti(t *testing.T) {
	expected := "大吉"
	actual := getDaikiti()
	if actual != expected {
		t.Errorf("want: getDaikiti() = %s, got: %s ", expected, actual)
	}
}
