package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

// Game represents a typing game.
type Game struct {
	Vocabulary Vocabulary
	Timeout    time.Duration
	Reader     io.Reader
}

// Score represents a game score.
type Score struct {
	CorrectWords int
	TotalWords   int
	GiveUp       bool
}

// New returns a new game.
func New(vocabulary Vocabulary, timeout time.Duration) *Game {
	return &Game{vocabulary, timeout, os.Stdin}
}

// Start starts a game and waits for timeout or cancel (ctrl+D).
// If timeout or cancel occurred, this returns the score.
// If any error occurred, this returns the error.
func (g *Game) Start(ctx context.Context) (*Score, error) {
	ctx, cancel := context.WithTimeout(ctx, g.Timeout)
	defer cancel()
	errCh := make(chan error)
	defer close(errCh)
	score := &Score{}

	go g.scanLines(score, cancel, errCh)

	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != context.DeadlineExceeded && err != context.Canceled {
			return nil, err
		}
		return score, nil
	case err := <-errCh:
		return nil, err
	}
}

func (g *Game) scanLines(score *Score, cancel func(), errCh chan<- error) {
	s := bufio.NewScanner(g.Reader)
	for {
		expected := g.Vocabulary.NextWord()
		fmt.Printf("%s\n>> ", expected)
		if !s.Scan() {
			score.GiveUp = true
			cancel()
			return
		}
		if err := s.Err(); err != nil {
			errCh <- err
			return
		}
		if expected == s.Text() {
			score.CorrectWords++
		}
		score.TotalWords++
	}
}
