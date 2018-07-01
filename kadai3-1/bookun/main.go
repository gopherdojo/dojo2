package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	// ExitCodeOK - 正常終了したときの値
	ExitCodeOK = iota
)

var words = []string{"dog", "cat", "pony", "mouse", "rat", "sheep", "camel", "snake", "tiger", "lion"}

type chooser interface {
	choose([]string) string
}

type inputer interface {
	input(io.Reader) <-chan string
}

type question struct {
}

func (q *question) choose(candidates []string) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	word := candidates[r.Intn(len(candidates))]
	return word
}

func (q *question) input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

// Run - 処理の結果のコードを返す
func Run(candidates []string, c chooser, i inputer) int {
	var answerCount int
	ch := i.input(os.Stdin)
	timeout := time.After(5 * time.Second)
	for {
		answer := c.choose(candidates)
		fmt.Println(answer)
		fmt.Print("> ")
		select {
		case v1 := <-ch:
			if answer == v1 {
				answerCount++
			}
		case <-timeout:
			fmt.Printf("\ntime up\n")
			fmt.Printf("%d問正解\n", answerCount)
			return ExitCodeOK
		}
	}
}

func main() {
	q := &question{}
	os.Exit(Run(words, q, q))
}
