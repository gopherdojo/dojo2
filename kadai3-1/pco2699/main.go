package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"time"

	"dojo2/kadai3-1/pco2699/lib"
)



func main(){
	question := &Question.Question{}
	file, err := question.Open("./kadai3-1/pco2699/question.txt");
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open the file!")
	}

	err = question.Read(file)
	if err == nil{
		fmt.Fprintf(os.Stderr, "Could not read the file!")
	}

	point := 0
	ch1 := input(os.Stdin)

	for {
		answer := question.Randomize()
		fmt.Println(answer)

		fmt.Print(">")

		select {
			case v1 := <- ch1:
				if v1 == answer {
					fmt.Println("Correct!")
					point++
				} else {
					fmt.Println("Wrong!")
				}
			case <-time.After(100 * time.Second):
				fmt.Println("Game End")
				fmt.Printf("Total Point : %v", point)
				return
		}
	}
}

func input(r io.Reader) <-chan string{
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