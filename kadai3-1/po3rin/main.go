package main

import (
	"bufio"
	"context"
	"dojo2/kadai3-1/po3rin/question"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var questions = question.List
var num int

const sec = 30000

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

func makeQuetion() string {
	rand.Seed(time.Now().UnixNano())
	q := questions[rand.Intn(len(questions))]
	return q
}

func main() {
	fmt.Printf("typing game start ! %v millsecound !", sec)
	ch := input(os.Stdin)
	q := makeQuetion()
	fmt.Println("==============")
	fmt.Println(q)
	fmt.Print(">")

	bc := context.Background()
	tc := sec * time.Millisecond
	ctx, cancel := context.WithTimeout(bc, tc)
	defer cancel()

	for {
		select {
		case answer := <-ch:
			q := makeQuetion()
			fmt.Println("==============")
			fmt.Println(q)
			fmt.Print(">")
			if answer != q {
				num++
			}
		case <-ctx.Done():
			fmt.Printf("end! number of correct is %v\n", num)
			os.Exit(1)
		}
	}
}
