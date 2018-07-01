package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/conversion"
	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/options"
)

// CLI represents in/out
type CLI struct {
	outStream, errStream io.Writer
}

// Run execute cli flow
func (c *CLI) Run(args []string) {
	opts, err := options.Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse errors caused. %s.\n", err)
		os.Exit(1)
	}
	decoder, err := opts.Decoder()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get decoder")
		os.Exit(1)
	}
	encoder, err := opts.Encoder()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get encoder")
	}
	conversion := &conversion.Conversion{
		Decoder: decoder,
		Encoder: encoder,
		ToExt:   opts.Out,
	}
	for _, file := range opts.Files {
		if err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
			switch {
			case err != nil:
				return err
			case !info.IsDir():
				destination := conversion.ReplaceExt(path)
				log.Printf("%s -> %s", path, destination)
				return nil
			default:
				return nil
			}
		}); err != nil {
			log.Printf("Skipped %s: %s", file, err)
		}
	}
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	cli.Run(os.Args)
	os.Exit(0)
}
