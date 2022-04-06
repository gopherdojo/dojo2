package main

import (
	"dojo/kadai1/chillout2san/converter"
	"errors"
	"flag"
	"log"
)

/*
	ディレクトリが指定されていなかったり、拡張子がjpgもしくはpng以外で指定された際にerrorを返します。
*/
func validate(beforeExtension, afterExtension, targetDir *string) error {
	if *beforeExtension != "jpg" && *beforeExtension != "png" {
		return errors.New("変更前の拡張子はjpgもしくはpngのどちらかを指定してください。")
	}

	if *afterExtension != "jpg" && *afterExtension != "png" {
		return errors.New("変更後の拡張子はjpgもしくはpngのどちらかを指定してください。")
	}

	if *targetDir == "" {
		return errors.New("ディレクトリが指定されていません。")
	}
	return nil
}

func main() {
	beforeExtension := flag.String("before", "jpg", "変換前の拡張子")
	afterExtension := flag.String("after", "png", "変換後の拡張子")
	targetDir := flag.String("path", "", "変換する写真のあるディレクトリ")
	flag.Parse()

	if err := validate(beforeExtension, afterExtension, targetDir); err != nil {
		log.Fatal(err)
	}

	if err := converter.Convert(*beforeExtension, *afterExtension, *targetDir); err != nil {
		log.Fatal(err)
	}
}
