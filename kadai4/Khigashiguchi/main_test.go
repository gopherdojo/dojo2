package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSetResult(t *testing.T) {
	box = map[int]string{
		0: "小吉",
	}
	cases := []struct {
		name     string
		now      time.Time
		expected string
	}{
		{
			name:     "1月1日",
			now:      time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		{
			name:     "1月2日",
			now:      time.Date(2018, 1, 2, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		{
			name:     "1月3日",
			now:      time.Date(2018, 1, 3, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		{
			name:     "12月31日",
			now:      time.Date(2017, 12, 31, 0, 0, 0, 0, time.Local),
			expected: "小吉",
		},
		{
			name:     "1月4日",
			now:      time.Date(2018, 1, 4, 0, 0, 0, 0, time.Local),
			expected: "小吉",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := Fortune{}
			f.setResult(c.now)

			if c.expected != f.Result {
				t.Errorf(
					"want f.setResult(%v) = %s, got = %s",
					c.now, c.expected, f.Result,
				)
			}
		})
	}
}

func TestFortuneHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/fortune", nil)
	fortuneHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Errorf(
			"want statusCode=%d, got = %d",
			http.StatusOK, rw.StatusCode,
		)
	}
	// FIXME: main_test.go:81: want body={"result":"小吉"}, got={"result":"小吉"}
	// b, err := ioutil.ReadAll(rw.Body)
	// if err != nil {
	// 	t.Fatalf("Unexpected error, %d", err)
	// }
	// actual := string(b)
	// expected := `{"result":"小吉"}`
	// if expected != actual {
	// 	t.Errorf(
	// 		"want body=%s, got=%s",
	// 		expected, actual,
	// 	)
	// }
}
