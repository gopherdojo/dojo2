package options

import (
	"errors"
	"flag"
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
