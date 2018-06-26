package main

import (
	"./converter"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var i, o string

func init() {
	flag.StringVar(&i, "i", "", "input image ext")
	flag.StringVar(&o, "o", "", "output image ext")
}

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	if len(i) == 0 || len(o) == 0 || len(dir) == 0 {
		fmt.Println(usage())
		os.Exit(1)
	}
	s, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	abs, err := filepath.Abs(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := converter.RecursiveConvert(filepath.Join(abs, dir), i, o, converter.Path{}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func usage() string {
	return "usage: imgconv -i jpg -o png dir"
}
