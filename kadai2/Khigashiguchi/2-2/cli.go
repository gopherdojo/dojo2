package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/conversion"
	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/options"
)

func main() {
	opts, err := options.Parse(os.Args)
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
