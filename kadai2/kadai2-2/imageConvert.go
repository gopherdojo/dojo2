package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

var (
	inExtOpt  = flag.String("i", "jpg", "変換対象の画像ファイルの種類")
	outExtOpt = flag.String("o", "png", "変換後の画像ファイルの種類")
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "対象パスが指定されていません")
		return
	}
	targetPath := imageconverter.FilePath(flag.Args()[0])

	inputFormat := imageconverter.Format(*inExtOpt)
	outputFormat := imageconverter.Format(*outExtOpt)

	var icCommandValidator imageconverter.CommandValidator
	if (!icCommandValidator.ExtValidate(inputFormat)) || (!icCommandValidator.ExtValidate(outputFormat)) {
		fmt.Fprintf(os.Stderr, "画像ファイルフォーマットが対応していません")
		return
	}

	var icFacade imageconverter.Facade
	icFacade.Run(targetPath, inputFormat, outputFormat)
}
