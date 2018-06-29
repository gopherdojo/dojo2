package main

import (
	"io"
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
