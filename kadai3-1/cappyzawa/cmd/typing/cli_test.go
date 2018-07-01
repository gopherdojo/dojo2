package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	t.Helper()
	t.Run("parseFlagError", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{
			OutStream: outStream,
			ErrStream: errStream,
		}
		args := strings.Split("typing -invalid", " ")
		expect := ExitCodeParseFlagError
		actual := cli.Run(args)
		if actual != expect {
			t.Errorf("actual should be %d, actual is %d", expect, actual)
		}
	})
	t.Run("codeOk", func(t *testing.T) {
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
		expect := ExitCodeOK
		actual := cli.Run(args)
		if actual != expect {
			t.Errorf("actual should be %d, actual is %d", expect, actual)
		}
		isCorrectContain := strings.Contains(outStream.String(), "correct")
		isPerfectContain := strings.Contains(outStream.String(), "perfect!!")
		if !isCorrectContain || !isPerfectContain {
			t.Error("outStream should contain correct")
		}
	})
}

func TestCLI_Input(t *testing.T) {
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
	expects := []string{"strawberry", "pineapple", "banana", "pear", "apple", "cherry", "grapefruit", "grape", "peach", "papaya"}
	for _, expect := range expects {
		actual := cli.Input(file)
		in, ok := <-actual
		if !ok {
			break
		} else {
			if in != expect {
				t.Errorf("%s should be %s", in, expect)
			}
		}
	}

}
