package option

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"runtime"
	"time"
)

type Option struct {
	proc    int
	timeout int
	output  string
	writer  io.Writer
}

// DefaultTimeout is default value for request timeout
const DefaultTimeout = 10

// Parse CLI options
func Parse(w io.Writer) (*Option, *url.URL, error) {
	opts := &Option{
		writer: w,
	}

	flag.IntVar(&opts.proc, "p", runtime.NumCPU(), "Parallelism")
	flag.IntVar(&opts.timeout, "t", DefaultTimeout, "Timeout")
	flag.StringVar(&opts.output, "o", "", "output")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return nil, nil, fmt.Errorf("args need url")
	}

	u, err := url.Parse(flag.Arg(0))

	if err != nil {
		return nil, nil, err
	}

	if opts.output == "" {
		opts.output = filepath.Base(u.Path)
	}

	return opts, u, nil
}

// Proc is proc getter
func (o *Option) Proc() int {
	return o.proc
}

// Timeout is timeout getter
func (o *Option) Timeout() time.Duration {
	return time.Duration(o.timeout) * time.Second
}

// Writer
func (o *Option) Writer() io.Writer {
	return o.writer
}

// Output path
func (o *Option) Output() string {
	return o.output
}
