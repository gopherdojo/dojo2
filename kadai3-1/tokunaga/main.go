package main

import (
	"bufio"
	crand "crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"
)

var word = [...]string{
	"toukyoutokkyokyokakyoku",
	"akamakigamiaomakigamikimakigami",
	"sushi",
	"tenpura",
	"kaiken",
	"nisshingeppo",
	"hyappatuhyakutyu",
	"kumiai",
	"taiiku",
	"kome",
}

const (
	limitSec  = 15
	ExitAbort = 1
)

func init() {
	if err := serRandSeed(); err != nil {
		os.Exit(ExitAbort)
	}
}

func serRandSeed() error {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return err
}

func main() {
	fmt.Printf("制限時間は%d秒です\n", limitSec)

	ch := input(os.Stdin)
	var correctCount int
	go output(ch, &correctCount)

	<-time.After(limitSec * time.Second)
	finish(correctCount)
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go read(r, ch)
	return ch
}

func read(reader io.Reader, ch chan string) {
	s := bufio.NewScanner(reader)
	for s.Scan() {
		ch <- s.Text()
	}
	close(ch)
}

func output(typeWord <-chan string, correctCount *int) {
	wordCount := len(word)
	for {
		answerWord := word[rand.Intn(wordCount)]
		fmt.Println(answerWord)
		fmt.Print("> ")
		typeWord := <-typeWord
		if typeWord == answerWord {
			fmt.Println("correct answer!")
			*correctCount++
		} else {
			fmt.Println("incorrect answer...")
		}
		fmt.Println("-----------------------------")
	}
}

func finish(correctCount int) {
	fmt.Println("\n-----------FINISH!-----------")
	fmt.Printf("正解数:%d\n", correctCount)
}
