package fortune

import (
	"math/rand"
	"time"
)

var fortuneList = []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}

// Fortune はおみくじの構造体です
type Fortune struct {
	Content string `json:"content"`
}

// FortuneSelector はおみくじ選択処理を行う構造体です
type FortuneSelector struct {
	Qlock Qlock
}

// DefaultQlock はGetCurrenttimeで普通に現在時刻を返す時計です
type DefaultQlock struct{}

// GetCurrentTime は現在時刻を返す関数です
func (d DefaultQlock) GetCurrentTime() time.Time {
	return time.Now()
}

// Qlock は時計のインターフェースです
type Qlock interface {
	GetCurrentTime() time.Time
}

// SelectFortune はおみくじリストの中からひとつ選んでFortuneオブジェクトを返す関数です
func (f FortuneSelector) SelectFortune() Fortune {
	t := f.Qlock.GetCurrentTime()
	if t.Month() == 1 && (1 <= t.Day() && t.Day() <= 3) {
		return f.selectFortuneOnlyDaikichi()
	}
	return f.selectFortuneRandom()
}

func (f FortuneSelector) selectFortuneOnlyDaikichi() Fortune {
	return Fortune{Content: "大吉"}
}

func (f FortuneSelector) selectFortuneRandom() Fortune {
	rand.Seed(time.Now().UnixNano())
	content := fortuneList[rand.Int()%len(fortuneList)]
	return Fortune{Content: content}
}
