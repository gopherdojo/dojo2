package main

import (
	"os"

	"github.com/gopherdojo/dojo2/kadai2/tokunaga"
)

// メイン関数
func main() {
	cli := &tokunaga.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
