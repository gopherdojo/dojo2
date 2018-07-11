package handler

import "time"

type TimeProvider interface {
	Time() time.Time
}

type NowTimeProvider struct { }
// 実行時の日時を返却する.
func (p NowTimeProvider) Time() time.Time {
	return time.Now()
}

type ArbitraryTimeProvider struct {
	time time.Time
}
// ArbitraryTimeProviderに設定した任意の日時を返却する.
func (p ArbitraryTimeProvider) Time() time.Time {
	return p.time
}
