package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"time"
	"math/rand"
	"strconv"
)

var wordList = []string{
	"apple",
	"brother",
	"cross",
	"dinner",
	"evening",
	"final",
	"great",
	"happy",
	"idea",
	"judge",
	"knight",
	"least",
	"mountain",
	"new",
	"open",
	"pen",
	"question",
	"response",
	"star",
	"teacher",
	"unique",
	"victory",
	"winner",
	"xbox",
	"yahoo",
	"zero",
}

func main() {
	chTyping := input(os.Stdin)
	chTimer := time.After(30 * time.Second)
	var correct int

	for {
		q := question()
		fmt.Println("Let's input : " + q)
		fmt.Print(">")
		select {
		case in := <-chTyping:
			if q == in {
				correct += 1
				fmt.Println("Yeah!!  (count: " + strconv.Itoa(correct) + ")")
			} else {
				fmt.Println("Wrong.  (count: " + strconv.Itoa(correct) + ")")
			}
		case <-chTimer:
			fmt.Println("")
			fmt.Println("Time Up!!!")
			fmt.Println("Correct Count is ... " + strconv.Itoa(correct))
			return
		}
	}

}

func question() string {
	rand.Seed(time.Now().UnixNano())
	return wordList[rand.Int() % len(wordList)]
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

