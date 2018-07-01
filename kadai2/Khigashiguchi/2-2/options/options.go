package options

import (
	"errors"
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/format"
)

// Options command line tool argument
type Options struct {
	Files []string
	In    string
	Out   string
}

var in *string
var out *string

// Parse return parsed option
func Parse(args []string) (*Options, error) {
	f := flag.NewFlagSet(args[0], flag.ContinueOnError)
	in = f.String("in", "jpg", "extension name of target convert file")
	out = f.String("out", "png", "extension name of converted file")
	f.Usage()
	if err := f.Parse(args[1:]); err != nil {
		return nil, err
	}
	if f.NArg() == 0 {
		return nil, errors.New("few argument")
	}
	opts := &Options{
		In:    *in,
		Out:   *out,
		Files: f.Args(),
	}
	return opts, nil
}

// Decoder return a decoder matched with extention
func (opts *Options) Decoder() (format.Decoder, error) {
	switch opts.In {
	case "jpg", "jpeg":
		return &format.JPEG{}, nil
	case "png":
		return &format.PNG{}, nil
	}
	return nil, fmt.Errorf("Unknown extention type: %s", opts.In)
}

// Encoder return a encoder matched with extention
func (opts *Options) Encoder() (format.Encoder, error) {
	switch opts.Out {
	case "jpg", "jpeg":
		return &format.JPEG{Options: jpeg.Options{Quality: jpeg.DefaultQuality}}, nil
	case "png":
		return &format.PNG{Options: png.Encoder{CompressionLevel: png.DefaultCompression}}, nil
	}
	return nil, fmt.Errorf("Unknown extention type: %s", opts.Out)
}
