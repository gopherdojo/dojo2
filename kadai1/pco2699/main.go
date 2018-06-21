package main

import (
	"os"
	"dojo2/kadai1/pco2699/lib"
)

//
func main() {
	cli := &imgconverter.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
