package extension

import (
	"os/exec"
	"testing"
)

func TestConvert(t *testing.T) {
	arg := Arg{
		From: "jpg",
		To:   "png",
		Path: "../images/gopher.jpg",
	}
	err := arg.Convert()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	exec.Command("rm", "../images/gopher.png").Run()
}
