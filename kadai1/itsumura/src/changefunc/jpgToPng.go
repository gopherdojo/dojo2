/*
Package changefunc :Encode jpg or png data to other type of image.

jpgまたはpngファイルを、他方のファイル形式に変換します。

jpg => png

png => jpg

*/
package changefunc

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

//If error happen, fire panic.
//エラーが起きた時にpanicを起こす関数
func ec(err *error) {
	if *err != nil {
		fmt.Println(*err)
		panic(*err)
	}
}

// JpgToPng :Encode Jpg file to Png file and return number of files which are successed to encode.JpgをPngに変換する関数
func JpgToPng(paths []string, filetype string) int {
	success := 0
	fmt.Println("jpgToPng実行")

	for _, inputPath := range paths {
		newFileNameSlice := strings.Split(inputPath, ".")

		var err error

		//try
		if newFileNameSlice[len(newFileNameSlice)-1] != "jpg" &&
			newFileNameSlice[len(newFileNameSlice)-1] != "jpeg" {
			fmt.Println(inputPath + "はjpgファイルでないためスキップ")
			continue
		}

		//ファイル名を変更
		outputPath := ""
		for i := 0; i < len(newFileNameSlice)-1; i++ {
			outputPath += newFileNameSlice[i]
		}
		outputPath = outputPath + ".png"
		fmt.Println(inputPath + "から" + outputPath + "を作成しています")

		//出力ファイル生成
		out, err := os.Create(outputPath)

		//入力ファイルを開く
		file, err := os.Open(inputPath)

		//画像のバイナリデータ
		img, _, err := image.Decode(file)

		if err == nil {
			err := png.Encode(out, img)
			if err == nil {
				success++
			}
		}
		ec(&err)
	}
	return success
}
