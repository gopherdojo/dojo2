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
func main() {
	fmt.Fprintf(writer, "The limit time of each typing is %v.\n", limitTime)

	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	var correct int
	for i, w := range words {
		fmt.Fprintf(writer, "%d st trial! Type it! -> %s\n", i+1, w)
		fmt.Fprint(writer, "> ")
		select {
		case <-time.After(limitTime * time.Second):
			fmt.Fprintln(writer, "Time over! Next!")
		case tw := <-ch:
			if tw == w {
				fmt.Fprintln(writer, "Your typing is correct!")
				correct++
			} else {
				fmt.Fprintln(writer, "Your typing is wrong. Next!")
			}
		}
	}
	if len(words) == correct {
		fmt.Fprintln(writer, "Complete! Finish!")
	}
	fmt.Fprintf(writer, "Correct count: %d\n", correct)
}
