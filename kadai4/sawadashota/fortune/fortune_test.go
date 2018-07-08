package fortune_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gopherdojo/dojo2/kadai4/sawadashota/fortune"
)

var randSource rand.Source

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
}

func TestResult(t *testing.T) {
	tokyo, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		t.Fatal(err)
	}

	utc, err := time.LoadLocation("UTC")

	if err != nil {
		t.Fatal(err)
	}

	cases := map[string]struct {
		now    time.Time
		expect fortune.Fortune
	}{
		"tokyo 1.1": {
			now:    time.Date(2018, 1, 1, 0, 0, 0, 0, tokyo),
			expect: fortune.Daikichi,
		},
		"utc 1.1": {
			now:    time.Date(2018, 1, 1, 0, 0, 0, 0, utc),
			expect: fortune.Daikichi,
		},
		"tokyo 1.3": {
			now:    time.Date(2020, 1, 3, 23, 59, 59, 59, tokyo),
			expect: fortune.Daikichi,
		},
		"utc 1.3": {
			now:    time.Date(2020, 1, 3, 23, 59, 59, 59, utc),
			expect: fortune.Daikichi,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			f := fortune.Result(randSource, c.now)
			if f != c.expect {
				t.Errorf("expect: %s but actual: %s", c.expect, f)
			}
		})
	}
}

func TestParse(t *testing.T) {
	cases := map[string]struct {
		rawStr   string
		expect   fortune.Fortune
		hasError bool
	}{
		"大吉": {
			rawStr:   "大吉",
			expect:   fortune.Daikichi,
			hasError: false,
		},
		"中吉": {
			rawStr:   "中吉",
			expect:   fortune.Chukichi,
			hasError: false,
		},
		"大凶": {
			rawStr:   "大凶",
			expect:   fortune.Daikyo,
			hasError: false,
		},
		"dummy string": {
			rawStr:   "てすと",
			expect:   0,
			hasError: true,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			f, err := fortune.Parse(c.rawStr)
			hasError := err != nil

			if hasError != c.hasError {
				if hasError {
					t.Error(err)
				}

				t.Error("expect has error but has no errors")
			}

			// Finish test if expects error
			if hasError {
				return
			}

			if f != c.expect {
				t.Errorf("expect %s but actual %s", c.expect, f)
			}
		})
	}
}
