package main

import (
	"bytes"
	"testing"
)

var buffer *bytes.Buffer

func init() {
	buffer = &bytes.Buffer{}
	writer = buffer
}

func TestStart(t *testing.T) {
	defer buffer.Reset()
	start()
	expected := "Game start!!\nThe limit time of each typing is 10.\n"
	actual := buffer.String()
	if actual != expected {
		t.Errorf(`expected=%s, actual=%s`, expected, actual)
	}
}

func TestEnd(t *testing.T) {
	defer buffer.Reset()
	end()
	expected := "Game finish!!\n"
	actual := buffer.String()
	if actual != expected {
		t.Errorf(`expected=%s, actual=%s`, expected, actual)
	}
}

func TestInputCh(t *testing.T) {
	input := "hello typing game."
	inputBuf := bytes.NewBufferString(input)
	ch := inputCh(inputBuf)
	received := <-ch
	if input != received {
		t.Errorf(`expected=%s, actual=%s`, input, received)
	}
}

func TestGame(t *testing.T) {
	// input := ""
}
