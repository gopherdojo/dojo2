package fortune

import (
	"fmt"
	"math/rand"
	"time"
)

type Fortune int

const (
	AmountItems = 10

	Daikichi Fortune = iota
	Kichi
	Chukichi
	Shokichi
	Suekichi
	Kyo
	Daikyo
)

var items = [AmountItems]Fortune{
	Daikichi,
	Kichi,
	Kichi,
	Chukichi,
	Chukichi,
	Shokichi,
	Shokichi,
	Suekichi,
	Kyo,
	Daikyo,
}

// Result returns fortune
func Result(s rand.Source, t time.Time) Fortune {
	if isNewYear(t) {
		return Daikichi
	}

	return items[rand.New(s).Intn(AmountItems)]
}

// Parse string to Fortune
func Parse(str string) (Fortune, error) {
	switch str {
	case "大吉":
		return Daikichi, nil
	case "吉":
		return Kichi, nil
	case "中吉":
		return Chukichi, nil
	case "小吉":
		return Shokichi, nil
	case "末吉":
		return Suekichi, nil
	case "凶":
		return Kyo, nil
	case "大凶":
		return Daikyo, nil
	default:
		return 0, fmt.Errorf("cannot parse %s", str)
	}
}

// String from Fortune
func (f Fortune) String() string {
	switch f {
	case Daikichi:
		return "大吉"
	case Kichi:
		return "吉"
	case Chukichi:
		return "中吉"
	case Shokichi:
		return "小吉"
	case Suekichi:
		return "末吉"
	case Kyo:
		return "凶"
	default:
		return "大凶"
	}
}

func isNewYear(t time.Time) bool {
	if t.Month() != time.January {
		return false
	}

	return 1 <= t.Day() && t.Day() <= 3
}
