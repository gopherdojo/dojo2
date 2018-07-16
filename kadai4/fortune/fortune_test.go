package fortune

import (
	"testing"
	"time"
)

func TestSelectFortuneDefault(t *testing.T) {
	var dq DefaultQlock
	fs := FortuneSelector{Qlock: dq}
	f := fs.SelectFortune()

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

type ManualQlock struct {
	Time time.Time
}

func (d ManualQlock) GetCurrentTime() time.Time {
	return d.Time
}

func TestSelectFortuneNewYear(t *testing.T) {

	timeList := []time.Time{
		time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2000, time.January, 2, 12, 0, 0, 0, time.UTC),
		time.Date(2000, time.January, 3, 23, 59, 0, 0, time.UTC),
	}

	for _, ti := range timeList {
		mq := ManualQlock{Time: ti}
		fs := FortuneSelector{Qlock: mq}
		f := fs.SelectFortune()

		if f.Content != "大吉" {
			t.Fatalf("unexpected f.Content when: %s", ti)
		}
	}
}
