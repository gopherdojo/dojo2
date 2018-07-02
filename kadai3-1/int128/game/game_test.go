package game

import (
	"strings"
	"testing"
	"time"
)

func TestGameScanLines(t *testing.T) {
	v := Vocabulary([]string{"foo"})
	r := strings.NewReader("foo\nfoo\nbar")
	g := &Game{v, time.Second, r}
	s := &Score{}
	c := 0
	cancel := func() { c++ }
	g.scanLines(s, cancel, make(chan error))

	if c != 1 {
		t.Errorf("cancel should be called once but %d", c)
	}
	if s.GiveUp != true {
		t.Errorf("s.GiveUp wants true but %v", s.GiveUp)
	}
	if s.CorrectWords != 2 {
		t.Errorf("s.CorrectWords wants 2 but %d", s.CorrectWords)
	}
	if s.TotalWords != 3 {
		t.Errorf("s.TotalWords wants 3 but %d", s.TotalWords)
	}
}
