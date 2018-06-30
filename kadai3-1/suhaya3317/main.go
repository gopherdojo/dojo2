package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
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
	qt := 0
	n := 0

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("終了！！！")
		fmt.Printf("%v問中／%v問正解 \n", qt, n)
		os.Exit(0)
	}()

	words := [...]string{"go", "ruby", "java", "javascript", "swift", "python", "php"}

	for {
		rand.Seed(time.Now().UnixNano())
		word := words[rand.Intn(7)]
		fmt.Println(word)
		fmt.Print(">")
		qt++
		if <-ch == word {
			fmt.Println("TRUE!")
			n++
		} else {
			fmt.Println("FALSE!")
		}
	}
}
