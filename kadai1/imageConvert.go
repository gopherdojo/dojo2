package main

import (
	"flag"

	"github.com/gopherdojo/dojo2/kadai1/imageconverter"
)

var (
	inExtOpt  = flag.String("i", "jpg", "変換対象の画像ファイルの種類")
	outExtOpt = flag.String("o", "png", "変換後の画像ファイルの種類")
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("対象パスが指定されていません")
	}
	targetPath := imageconverter.FilePath(flag.Args()[0])

	inputFormat := imageconverter.Format(*inExtOpt)
	outputFormat := imageconverter.Format(*outExtOpt)

	var icCommandValidator imageconverter.CommandValidator
	if (!icCommandValidator.ExtValidate(inputFormat)) || (!icCommandValidator.ExtValidate(outputFormat)) {
		panic("画像ファイルフォーマットが対応していません")
	}

	var icFacade imageconverter.Facade
	icFacade.Run(targetPath, inputFormat, outputFormat)
}
