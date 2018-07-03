package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var questions = []string{
	"ruby",
	"python",
	"node",
	"go",
	"php",
	"nim",
	"docker",
	"rustc",
	"docker-compose",
	"scala",
	"rails",
	"git",
	"npm",
	"vim",
	"brew",
	"yarn",
	"pip",
	"dep",
	"aws",
	"gcloud",
	"kubectl",
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

func makeQuetion() string {
	rand.Seed(time.Now().UnixNano())
	q := questions[rand.Intn(len(questions))]
	return q
}

func main() {
	ch := input(os.Stdin)
	for {
		q := makeQuetion()
		fmt.Println("==============")
		fmt.Println(q)
		fmt.Print(">")
		answer := <-ch
		if answer != q {
			fmt.Println("×! miss")
			continue
		}
		fmt.Println("◯! correct!")
	}
}
