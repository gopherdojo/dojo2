package main

import (
	"bytes"
	"testing"
)

func ExampleFinish() {
	finish(3)
	// Output: -----------FINISH!-----------
	// 正解数:3
}

func TestInput(t *testing.T) {
	inputWord := "Hello Worlds!"
	inputBuf := bytes.NewBufferString(inputWord)
	ch := input(inputBuf)
	received := <-ch
	if inputWord != received {
		t.Errorf("want: receive %s from chanel, got: receive %s from chanel", inputWord, received)
	}
}

func TestRead(t *testing.T) {
	input := "Hello Worlds!"
	ch := make(chan string)
	inputBuf := bytes.NewBufferString(input)
	go read(inputBuf, ch)
	received := <-ch
	if input != received {
		t.Errorf("want: receive %s from chanel, got: receive %s from chanel", input, received)
	}
}
