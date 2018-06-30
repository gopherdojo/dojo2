package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	ch := input(os.Stdin)
	words, e := loadWords()
	if e != nil {
		os.Exit(1)
	}
	correct := 0
	const TimeLimit = 60
	timeLimitCh := time.After(TimeLimit * time.Second)
	rand.Seed(time.Now().UnixNano())
FOR:
	for {
		q := words[rand.Intn(len(words))]
		fmt.Println("> " + q)
		select {
		case <- timeLimitCh:
			break FOR
		case a := <-ch:
			if q == a {
				correct = correct + 1
			}
		}
	}
	fmt.Println()
	fmt.Println(correct)
}

func input(r io.Reader) <-chan string {
	ch := make(chan string, 0)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func loadWords() ([]string, error) {
	file, e := os.Open("dic.txt")
	if e != nil {
		return nil, e
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	words := make([]string, 1)
	for reader.Scan() {
		words = append(words, reader.Text())
	}
	return words, nil
}
