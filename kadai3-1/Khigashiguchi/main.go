package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const limitTime = 10

var words = [...]string{
	"I have a pen",
	"I have a apple",
	"ApplePen",
	"I have a pen",
	"I have a pineapple",
	"PenPineappleApplePen",
}

var writer io.Writer

func init() {
	writer = os.Stdout
}

func start() {
	fmt.Fprintln(writer, "Game start!!")
	fmt.Fprintf(writer, "The limit time of each typing is %v.\n", limitTime)
}

func end() {
	fmt.Fprintln(writer, "Game finish!!")
}

func inputCh(input io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(input)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func game(input io.Reader) {
	ch := inputCh(input)
	var correct int
	for i, w := range words {
		fmt.Fprintf(writer, "%d st trial! Type it! -> %s\n", i+1, w)
		fmt.Fprint(writer, "> ")
		select {
		case <-time.After(limitTime * time.Second):
			fmt.Fprintln(writer, "Time over!")
		case tw := <-ch:
			if tw == w {
				fmt.Fprintln(writer, "Your typing is correct!")
				correct++
			} else {
				fmt.Fprintln(writer, "Your typing is wrong.")
			}
		}
	}
	if len(words) == correct {
		fmt.Fprintln(writer, "Complete! Finish!")
	}
	fmt.Fprintf(writer, "Correct count: %d\n", correct)
}

func main() {
	start()
	game(os.Stdin)
}
