package main

import (
	"flag"
	"fmt"

	"./converter"
)

func main() {
	var targetFilepath string
	var fromFileExt string
	var toFileExt string
	flag.StringVar(&targetFilepath, "target", ".", "target filepath.")
	flag.StringVar(&fromFileExt, "from", "jpg", "from file extention")
	flag.StringVar(&toFileExt, "to", "png", "to file extention")
	flag.Parse()

	switch {
	case fromFileExt == "jpg" && toFileExt == "png":
		converter.ConvertImagesFromJpgToPngInDir(targetFilepath)
	case fromFileExt == "png" && toFileExt == "jpg":
		converter.ConvertImagesFromPngToJpgInDir(targetFilepath)
	default:
		fmt.Println("[ERROR]No supprot to convert from " + fromFileExt + " to " + toFileExt)
	}
}
