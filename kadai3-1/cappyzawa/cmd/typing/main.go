package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type CLI struct {
	InStream             io.Reader
	OutStream, ErrStream io.Writer
}

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeTimeoutError
)

var (
	qaMap map[string]string
)

func init() {
	tmpMap := make(map[string]string)
	questions := []string{"strawberry", "pineapple", "banana", "pear", "apple", "cherry", "grapefruit", "grape", "peach", "papaya"}
	for _, q := range questions {
		tmpMap[q] = q
	}
	qaMap = tmpMap
}

func (c *CLI) Input(r io.Reader) <-chan string {
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

func (c *CLI) TimeAfter(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (c *CLI) Judge(count int, question, answer string) int {
	if qaMap[question] == answer {
		count++
	}
	return count
}

func (c *CLI) Run(args []string) int {
	var s int64
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.Int64Var(&s, "s", 60, "enable to select limit time")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	var count int
	limit := time.Duration(s)
	timerCh := c.TimeAfter(limit * time.Second)
	inputCh := c.Input(c.InStream)
	for q := range qaMap {
		fmt.Fprintf(c.OutStream, "> %s\n", q)
		select {
		case in := <-inputCh:
			count = c.Judge(count, q, in)
		case <-timerCh:
			fmt.Fprintf(c.OutStream, "\ncorrect: %d\n", count)
			return ExitCodeOK
		}
	}
	return ExitCodeTimeoutError
}

func main() {
	cli := &CLI{InStream: os.Stdin, OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
