package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	t.Run("parseFlagError", func(t *testing.T) {
		t.Helper()
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{
			OutStream: outStream,
			ErrStream: errStream,
		}
		args := strings.Split("typing -invalid", " ")
		actual := cli.Run(args)
		if actual != ExitCodeParseFlagError {
			t.Errorf("actual should be %d, actual is %d", ExitCodeParseFlagError, actual)
		}
	})
	t.Run("codeOk", func(t *testing.T) {
		t.Helper()
		file, err := os.Open("../../testdata/answer.txt")
		defer file.Close()
		if err != nil {
			t.Error("file does not exist")
		}
		inStream := file
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{
			InStream:  inStream,
			OutStream: outStream,
			ErrStream: errStream,
		}
		args := strings.Split("typing -s 1", " ")
		actual := cli.Run(args)
		if actual != ExitCodeOK {
			t.Errorf("actual should be %d, actual is %d", ExitCodeOK, actual)
		}
		isCorrectContain := strings.Contains(outStream.String(), "correct")
		isPerfectContain := strings.Contains(outStream.String(), "perfect!!")
		if !isCorrectContain || !isPerfectContain {
			t.Error("outStream should contain correct")
		}
	})
}
