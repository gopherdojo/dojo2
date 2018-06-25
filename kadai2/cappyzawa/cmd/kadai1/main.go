package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gopherdojo/dojo2/kadai2/cappyzawa"
)

// CLI - struct having stream fields
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Execute - execute conversion
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
		return 1
	}

	var encoder conv.Converter
	if t == "jpeg" {
		encoder = new(conv.Jpeg)
	} else if t == "png" {
		encoder = new(conv.Png)
	} else {
		fmt.Fprint(c.ErrStream, "invalid format\n")
		return 1
	}
	command := conv.NewCommand(decoder, encoder)
	files, err := command.Run(dir, f, t)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "%s\n", err.Error())
		return 1
	}
	for _, file := range files {
		fmt.Fprintf(c.OutStream, "created file: %s\n", file)
	}
	return 0
}

func main() {
	cli := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Execute(os.Args))
}
