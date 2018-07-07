package fortune_test

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo2/kadai4/cappyzawa/fortune"
)

func TestFortune_Draw(t *testing.T) {
	t.Helper()
	cases := []struct {
		name   string
		input  time.Time
		expect string
	}{
		{name: "1/1", input: time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local), expect: fortune.DAIKICHI},
		{name: "1/2", input: time.Date(2018, 1, 2, 0, 0, 0, 0, time.Local), expect: fortune.DAIKICHI},
		{name: "1/3", input: time.Date(2018, 1, 3, 0, 0, 0, 0, time.Local), expect: fortune.DAIKICHI},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := fortune.NewFortune(c.input)
			actual := f.Draw()
			if actual.Result != c.expect {
				t.Errorf("result should be %s, actual %s", c.expect, actual.Result)
			}
		})
	}
}
