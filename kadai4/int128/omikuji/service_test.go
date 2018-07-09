package omikuji

import (
	"math/rand"
	"testing"
	"time"
)

func TestDefaultService_IsShogatsu(t *testing.T) {
	tokyoTime, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	matrix := []struct {
		now      time.Time
		expected bool
	}{
		{time.Date(2017, 12, 31, 0, 0, 0, 0, tokyoTime), false},
		{time.Date(2017, 12, 31, 23, 59, 59, 0, tokyoTime), false},
		{time.Date(2018, 1, 1, 0, 0, 0, 0, tokyoTime), true},
		{time.Date(2018, 1, 2, 0, 0, 0, 0, tokyoTime), true},
		{time.Date(2018, 1, 3, 0, 0, 0, 0, tokyoTime), true},
		{time.Date(2018, 1, 3, 23, 59, 59, 0, tokyoTime), true},
		{time.Date(2018, 1, 4, 0, 0, 0, 0, tokyoTime), false},
		{time.Date(2018, 2, 1, 0, 0, 0, 0, tokyoTime), false},
	}
	for _, m := range matrix {
		t.Run(m.now.String(), func(t *testing.T) {
			s := &DefaultService{
				random: rand.New(rand.NewSource(0)),
				time:   func() time.Time { return m.now },
			}
			actual := s.isShogatsu()
			if m.expected != actual {
				t.Errorf("shogatsu wants %v but %v", m.expected, actual)
			}
		})
	}
}
