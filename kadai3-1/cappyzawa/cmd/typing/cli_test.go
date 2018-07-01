package main

import (
	"bytes"
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
		args := strings.Split("mytyping -invalid", " ")
		actual := cli.Run(args)
		if actual != ExitCodeParseFlagError {
			t.Errorf("actual should be %d, actual is %d", ExitCodeParseFlagError, actual)
		}
	})
}
