package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo2/kadai1/cappyzawa"
	"io"
	"os"
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (c *CLI) Execute(args []string) int {
	var f, t string
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&f, "f", "jpeg", "enable to select format before conversion")
	flags.StringVar(&t, "t", "png", "enable to select format after conversion")

	flags.Parse(args[1:])

	if len(flags.Args()) == 0 {
		fmt.Fprintf(c.ErrStream, "Usage: %s [-n] dir_path...\n", os.Args[0])
		return 1
	}
	dir := flags.Arg(0)

	var decoder conv.Converter
	if f == "jpeg" {
		decoder = new(conv.Jpeg)
	} else if f == "png" {
		decoder = new(conv.Png)
	} else {
		fmt.Fprint(c.ErrStream, "invalid format\n")
	}

	var encoder conv.Converter
	if t == "jpeg" {
		encoder = new(conv.Jpeg)
	} else if t == "png" {
		encoder = new(conv.Png)
	} else {
		fmt.Fprint(c.ErrStream, "invalid format\n")
	}
	iFilePath := conv.NewIFilePath()
	command := conv.NewCommand(iFilePath, decoder, encoder)
	if err := command.Run(dir, f, t); err != nil {
		fmt.Fprintf(c.ErrStream, "%s\n", err.Error())
		return 1
	}
	return 0
}

func main() {
	cli := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Execute(os.Args))
}
