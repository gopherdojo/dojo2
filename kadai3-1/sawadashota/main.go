package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gopherdojo/dojo2/kadai3-1/sawadashota/faker"
)

// AmountTime is amount of game rounds
const AmountTime = 5

// Timeout is how long seconds waiting for typing
const Timeout = 5

// isHard game mode
var isHard bool

type Faker interface {
	Word() string
}

func init() {
	flag.BoolVar(&isHard, "hard", false, "Hard mode")
	flag.Parse()
}

func main() {
	var f Faker
	var word string

	if isHard {
		fmt.Println("Start Hard Mode!")
		f = new(faker.Hard)
	} else {
		fmt.Println("Start Easy Mode!")
		f = new(faker.Easy)
	}

	correctCount := 0

	ch := input(os.Stdin)

	fmt.Println("This is typing game!")
	fmt.Printf("Plase type appearing word in %d seconds\n", Timeout)

	for i := 0; i < AmountTime; i++ {
		word = f.Word()

		fmt.Printf("\nWord: %s\n", word)
		fmt.Print("----> ")

		select {
		case <-time.After(Timeout * time.Second):
			fmt.Println("Time Over!")
		case answer := <-ch:
			if word == answer {
				fmt.Println("Great!")
				correctCount++
			} else {
				fmt.Println("Uh-Oh")
			}
		}
	}

	if correctCount == AmountTime {
		fmt.Print("\nPerfect!\n")
		return
	}

	fmt.Printf("Correct: %d, Incorrect: %d\n", correctCount, AmountTime-correctCount)
}

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
