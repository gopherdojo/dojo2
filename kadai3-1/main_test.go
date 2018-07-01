package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var q = question{}

func TestChoose(t *testing.T) {
	testWords := []string{"dog"}
	result := q.choose(testWords)
	expected := testWords[0]
	if result != expected {
		t.Error(result)
		t.Error(expected)
	}
}

func TestInput(t *testing.T) {
	t.Helper()
	file, err := os.Open("test_input/input_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	result := q.input(file)
	expected := "dog"
	r := <-result
	if r != expected {
		t.Error(r)
		t.Error(expected)
	}
}

type mockQuestion struct{}

func (m *mockQuestion) choose(candidates []string) string {
	return "mock"
}

func (m *mockQuestion) input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		ch <- "mock"
	}()
	return ch
}

func TestRun(t *testing.T) {
	t.Helper()
	m := &mockQuestion{}
	result := Run(words, m, m)
	if result != ExitCodeOK {
		t.Error(result)
	}
}
