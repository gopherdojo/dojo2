package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo2/kadai1/tokunaga"
	"os"
	"path/filepath"
)

type Usage string

var extTo string
var extFrom string
var usage Usage = "./imageConverter [-from=png|jpeg ] [-to=png|jpeg ] directory"
var permmitedExts []string = []string{"png", "jpeg", "jpg"}

func init() {
	flag.StringVar(&extFrom, "from", "jpeg", string(usage))
	flag.StringVar(&extTo, "to", "png", string(usage))
	flag.Parse()
}

func main() {
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: "+string(usage))
		os.Exit(1)
	}
	if !checkExtPermmited(extFrom, permmitedExts) {
		fmt.Fprintln(os.Stderr, "-from is only png, jpeg, jpg")
		os.Exit(1)
	}
	if !checkExtPermmited(extTo, permmitedExts) {
		fmt.Fprintln(os.Stderr, "-to is only png, jpeg, jpg")
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
	if !info.IsDir() && filepath.Ext(path) == "."+extFrom {
		tokunaga.ConvertImage(path, adaptExt(extFrom), adaptExt(extTo))
	}
	return nil
}

func adaptExt(ext string) tokunaga.ImageConverter {
	if ext == "jpeg" || ext == "jpg" {
		return tokunaga.JpegWrapper(ext)
	}
	if ext == "png" {
		return tokunaga.PngWrapper(ext)
	}
	return tokunaga.PngWrapper(ext)
}

func checkExtPermmited(ext string, permittedExts []string) bool {
	for _, permmitedExt := range permittedExts {
		if ext == permmitedExt {
			return true
		}
	}
	return false
}
