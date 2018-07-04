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
	expected := "Game スタート\n制限時間は 5秒です\n"
	actual := buffer.String()
	if actual != expected {
		t.Errorf(`expected=%s, actual=%s`, expected, actual)
	}
}

func TestEnd(t *testing.T) {
	defer buffer.Reset()
	end()
	expected := "Game 終了\n"
	actual := buffer.String()
	if actual != expected {
		t.Errorf(`expected=%s, actual=%s`, expected, actual)
	}
}

func TestCreateCh(t *testing.T) {
	input := "hello world."
	inputBuf := bytes.NewBufferString(input)
	ch := createCh(inputBuf)
	received := <-ch
	if input != received {
		t.Errorf(`expected=%s, actual=%s`, input, received)
	}
}

func TestRunNoAnswer(t *testing.T) {
	defer buffer.Reset()
	expected := 0
	actual := run(buffer)
	if actual != expected {
		t.Errorf(`expected=%d, actual=%d`, expected, actual)
	}
}

func TestRunOK(t *testing.T) {
	defer buffer.Reset()
	expected := 1
	// TODO: どうしたら標準入力のテストができるのか?
	actual := run(buffer)
	if actual != expected {
		t.Errorf(`expected=%d, actual=%d`, expected, actual)
	}
}
