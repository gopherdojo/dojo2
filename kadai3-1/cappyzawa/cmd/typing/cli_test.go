package main

import (
	"bytes"
	"strings"
	"testing"
	"os"
)

func TestCLI_Run(t *testing.T) {
	t.Run("parseFlagError", func(t *testing.T) {
		t.Helper()
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{
			OutStream: outStream,
			ErrStream: errStream,
		}
		args := strings.Split("mytyping -invalid", " ")
		actual := cli.Run(args)
		if actual != ExitCodeParseFlagError {
			t.Errorf("actual should be %d, actual is %d", ExitCodeParseFlagError, actual)
		}
	})
	t.Run("codeOk", func(t *testing.T) {
		t.Helper()
		file, err := os.Open("../../testdata/answer.txt")
		if err != nil {
			t.Error("file does not ")
		}
		inStream := file
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{
			InStream: inStream,
			OutStream: outStream,
			ErrStream: errStream,
		}
		args := strings.Split("mytyping -s 5", " ")
		actual := cli.Run(args)
		if actual != ExitCodeOK {
			t.Errorf("actual should be %d, actual is %d", ExitCodeOK, actual)
		}
	})
}
