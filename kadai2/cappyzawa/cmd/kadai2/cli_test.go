package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI_Execute(t *testing.T) {

	cases := []struct{name string; input string; expected int}{
		{name: "jpegToPng", input: "myconv -f png -t jpeg ../../testdata/png", expected: ExitCodeOK},
		{name: "pngToJpeg", input: "myconv -f jpeg -t png ../../testdata/jpeg", expected: ExitCodeOK},
		{name: "parseFlagError", input: "myconv -invalid jpeg -t png ../../testdata/jpeg", expected: ExitCodeParseFlagError},
	}
	for _, c := range cases {
		t.Helper()
		t.Run(c.name, func(t *testing.T) {
			outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
			cli := &CLI{outStream, errStream}
			args := strings.Split(c.input, " ")
			actual := cli.Execute(args)
			if actual != c.expected {
				t.Errorf("actual is %d, wanted %d", actual, c.expected)
			}
		})
	}
}
