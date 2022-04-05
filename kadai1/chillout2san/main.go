package main

import (
	"dojo/kadai1/chillout2san/converter"
	"flag"
	"log"
)

func main() {
	beforeExtension := flag.String("before", "jpg", "変換前の拡張子")
	afterExtension := flag.String("after", "png", "変換後の拡張子")
	targetDir := flag.String("path", "", "変換する写真のあるディレクトリ")
	flag.Parse()

	isPathEmpty := *targetDir == ""
	if isPathEmpty {
		log.Fatal("ディレクトリが指定されていません。")
	}

	converter.Convert(*beforeExtension, *afterExtension, *targetDir)
}