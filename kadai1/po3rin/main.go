package main

import (
	"dojo2/kadai1/po3rin/extension"
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	from := flag.String("f", "jpg", "what encode from")
	to := flag.String("t", "png", "what encode to")
	dir := flag.String("d", ".", "target dir")
	flag.Parse()

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+*from {
			arg := extension.Arg{
				From: *from,
				To:   *to,
				Path: path,
			}
			return arg.Convert()
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
