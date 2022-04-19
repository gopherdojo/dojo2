package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("タイピングゲーム！制限時間は30秒以内！")
	timeLimit := time.After(time.Duration(30) * time.Second)
	counter := 0
	questionCounter := 0
	outputChannel := make(chan string)
	completeChannel := make(chan string)
	wordList := []string{"guitar", "piano", "drum", "bass"}

	// 問題を注入するgo-routine
	go func() {
		for _, word := range wordList {
			outputChannel <- word
		}
	}()

	// 問題を出題して回答を受け付けるgo-routine
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			output := <-outputChannel
			fmt.Println(output)
			scanner.Scan()
			questionCounter++
			if output == scanner.Text() {
				counter++
			}
			if questionCounter == len(wordList) {
				completeChannel <- "done"
			}
		}
	}()

	select {
	case <- completeChannel:
		fmt.Println("全問題終了しました")
		fmt.Println("正解数は" + strconv.Itoa(counter) + "個です")
	case <-timeLimit:
		fmt.Println("制限時間です。")
		fmt.Println("正解数は" + strconv.Itoa(counter) + "個です")
	}
}
