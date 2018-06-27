package main

import (
	"dojo2/kadai1/pco2699/lib"
	"os"
)

//
func main() {
	cli := &imgconverter.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
