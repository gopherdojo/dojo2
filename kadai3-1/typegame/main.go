package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func input(r io.Reader) <-chan string {
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

func main() {
	ch := input(os.Stdin)

	chanIsFinish := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		chanIsFinish <- true
	}()

	for {
		fmt.Print(">")
		var isFinished = false

		select {
		case isFinished = <-chanIsFinish:
			fmt.Println("FINISHED!")
		case res := <-ch:
			fmt.Println(res)
		}

		if isFinished {
			break
		}
	}
}
