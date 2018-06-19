package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai1/int128/images"
	"github.com/gopherdojo/dojo2/kadai1/int128/options"
)

func main() {
	opts, err := options.Parse(os.Args)
	if err != nil {
		os.Exit(1)
	}
	decoder, err := opts.Decoder()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	encoder, err := opts.Encoder()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	conversion := &images.Conversion{
		Decoder:        decoder,
		Encoder:        encoder,
		DestinationExt: *opts.To,
	}
	for _, parent := range opts.Paths {
		if err := filepath.Walk(parent, func(path string, info os.FileInfo, err error) error {
			switch {
			case err != nil:
				return err
			case !info.IsDir():
				destination := conversion.ReplaceExt(path)
				log.Printf("%s -> %s", path, destination)
				if err := conversion.Do(path, destination); err != nil {
					log.Printf("Skipped %s: %s", path, err)
				}
				return nil
			default:
				return nil
			}
		}); err != nil {
			log.Printf("Skipped %s: %s", parent, err)
		}
	}
}
