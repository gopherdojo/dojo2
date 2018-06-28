package main

import (
	"dojo2/kadai2/pco2699/02_refactorForTest/lib"
	"os"
)

//
func main() {
	cli := &imgconverter.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
