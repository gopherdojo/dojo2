// Package options provides parsing command line options.
package options

import (
	"flag"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/gopherdojo/dojo2/kadai1/int128/images"
)

// Options represents command line options.
type Options struct {
	Paths          []string
	From           *string
	To             *string
	JPEGQuality    *int
	PNGCompression *string
	GIFColors      *int
}

// Parse returns the command line arguments which includes the command name.
// No argument or any unknown flag will show the usage and return an error.
// Caller should exit on an error.
func Parse(args []string) (*Options, error) {
	f := flag.NewFlagSet(args[0], flag.ContinueOnError)
	opts := &Options{
		From:           f.String("from", "jpg", "Source image format: auto, jpg, png, gif"),
		To:             f.String("to", "png", "Destination image format: jpg, png, gif"),
		JPEGQuality:    f.Int("jpeg-quality", jpeg.DefaultQuality, "JPEG quality"),
		PNGCompression: f.String("png-compression", "default", "PNG compression level: default, no, best-speed, best-compression"),
		GIFColors:      f.Int("gif-colors", 256, "GIF number of colors"),
	}
	f.Usage = func() {
		fmt.Fprintf(f.Output(), "Usage: %s FILE or DIRECTORY...\n", f.Name())
		f.PrintDefaults()
	}
	if err := f.Parse(args[1:]); err != nil {
		return nil, err
	}
	if f.NArg() == 0 {
		f.Usage()
		return nil, fmt.Errorf("too few argument")
	}
	opts.Paths = f.Args()
	return opts, nil
}

// Decoder returns a decoder configured with the options.
func (opts *Options) Decoder() (images.Decoder, error) {
	switch *opts.From {
	case "auto":
		return &images.AutoDetect{}, nil
	case "jpg":
		return &images.JPEG{}, nil
	case "png":
		return &images.PNG{}, nil
	case "gif":
		return &images.GIF{}, nil
	}
	return nil, fmt.Errorf("Unknown source image format: %s", *opts.From)
}

// Encoder returns a encoder configured with the options.
func (opts *Options) Encoder() (images.Encoder, error) {
	switch *opts.To {
	case "jpg":
		return &images.JPEG{Options: jpeg.Options{Quality: *opts.JPEGQuality}}, nil
	case "png":
		c, err := opts.pngCompression()
		if err != nil {
			return nil, err
		}
		return &images.PNG{Options: png.Encoder{CompressionLevel: c}}, nil
	case "gif":
		return &images.GIF{Options: gif.Options{NumColors: *opts.GIFColors}}, nil
	}
	return nil, fmt.Errorf("Unknown destination image format: %s", *opts.To)
}

func (opts *Options) pngCompression() (png.CompressionLevel, error) {
	switch *opts.PNGCompression {
	case "default":
		return png.DefaultCompression, nil
	case "no":
		return png.NoCompression, nil
	case "best-speed":
		return png.BestSpeed, nil
	case "best-compression":
		return png.BestCompression, nil
	}
	return png.DefaultCompression, fmt.Errorf("Unknown PNG compression level: %s", *opts.PNGCompression)
}
