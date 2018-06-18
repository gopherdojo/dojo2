package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"github.com/gopherdojo/dojo2/kadai1/tokunaga"
)

var ext string

type Usage string

var usage Usage = "./imageConverter [-ext=png|jpeg ] directory"

func init() {
	flag.StringVar(&ext, "ext", "png", string(usage))
	flag.Parse()
}

func main() {
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: "+string(usage))
		os.Exit(1)
	}
	directory := flag.Args()[0]
	if err := filepath.Walk(directory, delegateFileOperation); err != nil {
		os.Exit(1)
	}
}

func delegateFileOperation(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && filepath.Ext(path) == "."+ext {
		tokunaga.ConvertImage(path)
	}
	return nil
}
