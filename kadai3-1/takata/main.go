package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const limitTime = 5

var writer io.Writer

var words []string

func init() {
	writer = os.Stdout

	// TODO: ファイルから読み込むなどを検討する
	words = []string{
		"apple",
		"banana",
		"orange",
		"grape",
		"Melon",
		"Muscat",
		"strawberry",
		"persimmon",
		"kiwi fruit",
		"cherry",
	}
}

func start() {
	fmt.Fprintln(writer, "Game スタート")
	fmt.Fprintf(writer, "制限時間は %v秒です\n", limitTime)
}

func end() {
	fmt.Fprintln(writer, "Game 終了")
}

func createCh(input io.Reader) <-chan string {
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

func run(input io.Reader) int {

	var answer int

	ch := createCh(input)

	for i, w := range words {

		fmt.Fprintf(writer, "%d回目 文字を入力してください -> %s\n", i+1, w)
		fmt.Fprint(writer, "> ")
		select {
		case <-time.After(limitTime * time.Second):
			fmt.Fprintln(writer, "時間切れです")
		case tw := <-ch:
			if judge(w, tw) {
				fmt.Fprintln(writer, "正解")
				answer++
			} else {
				fmt.Fprintln(writer, "残念")
			}
		}
	}
	fmt.Fprintf(writer, "正解数: %d\n", answer)
	return answer
}

func judge(expected string, actual string) bool {

	if actual == expected {
		return true
	} else {
		return false
	}
}

func main() {
	start()
	run(os.Stdin)
	end()
}
