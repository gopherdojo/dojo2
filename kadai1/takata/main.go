package main

import (
	"convert"
	"flag"
)

// main
func main() {
	var (
		dir    = flag.String("d", "", "directory to convert imagefiles")
		srcFmt = flag.String("sf", "jpg", "src file format")
		dstFmt = flag.String("df", "png", "dest file format")
	)

	flag.Parse()

	convert.Convert(*dir, *srcFmt, *dstFmt)
}
