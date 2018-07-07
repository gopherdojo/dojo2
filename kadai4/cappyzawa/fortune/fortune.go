package fortune

import (
	"math/rand"
	"time"
)

const (
	DAIKICHI = "大吉"
	KICHI    = "吉"
	KYO      = "凶"
	DAIKYO   = "大凶"
)

type Fortune struct {
	Date   time.Time `json:"date"`
	Result string    `json:"result"`
}

func NewFortune(date time.Time) *Fortune {
	return &Fortune{
		Date: date,
	}
}

func (f *Fortune) Draw() Fortune {
	if f.isNewYear(f.Date) {
		return Fortune{Date: f.Date, Result: DAIKICHI}
	}
	results := []string{
		DAIKICHI,
		KICHI,
		KYO,
		DAIKYO,
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(results))
	return Fortune{Date: f.Date, Result: results[r]}
}

func (f *Fortune) isNewYear(t time.Time) bool {
	if t.Month() == 1 {
		if t.Day() == 1 || t.Day() == 2 || t.Day() == 3 {
			return true
		}
	}
	return false
}
