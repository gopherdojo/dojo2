package main

import (
	"dojo/kadai1/chillout2san/converter"
	"flag"
)

func main() {
	beforeExtension := string(*flag.String("before", "JPG", "変換前の拡張子"))
	afterExtension := string(*flag.String("after", "PNG", "変換後の拡張子"))
	path := string(*flag.String("path", "", "変換する写真のあるディレクトリ"))
	flag.Parse()
	converter.Convert(beforeExtension, afterExtension, path)
}