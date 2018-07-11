package handler

import "time"

type DateProvider interface {
	Now() time.Time
}
type NowTimeProvider struct { }
func (p NowTimeProvider) Now() time.Time {
	return time.Now()
}

type ArbitraryTimeProvider struct {
	time time.Time
}
func (p ArbitraryTimeProvider) Now() time.Time {
	return p.time
}
