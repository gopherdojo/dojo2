package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"image/jpeg"

	"github.com/gopherdojo/dojo2/kadai2/cappyzawa"
)

// CLI - struct having stream fields
type CLI struct {
	OutStream, ErrStream io.Writer
}

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeInvalidFormat
	ExitCodeFailedRun
)

// Execute - execute conversion
func (c *CLI) Execute(args []string) int {
	var f, t string
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&f, "f", "jpeg", "enable to select format before conversion")
	flags.StringVar(&t, "t", "png", "enable to select format after conversion")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if len(flags.Args()) == 0 {
		fmt.Fprintf(c.ErrStream, "Usage: %s [-n] dir_path...\n", args[0])
		return ExitCodeParseFlagError
	}

	dirIndex := 0
	dir := flags.Arg(dirIndex)

	decoder := conv.NewDecoder()
	var encoder conv.Encoder
	if t == "jpeg" {
		encoder = &conv.Jpeg{Options: &jpeg.Options{Quality: jpeg.DefaultQuality}}
	} else if t == "png" {
		encoder = &conv.Png{}
	} else {
		fmt.Fprint(c.ErrStream, "invalid format\n")
		return ExitCodeInvalidFormat
	}
	command := conv.NewCommand(decoder, encoder)

	files, err := command.Run(dir, f, t)
	if err != nil {
		fmt.Fprintf(c.ErrStream, "%s\n", err.Error())
		return ExitCodeFailedRun
	}
	for _, file := range files {
		fmt.Fprintf(c.OutStream, "created file: %s\n", file)
	}
	return ExitCodeOK
}

func main() {
	cli := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Execute(os.Args))
}
