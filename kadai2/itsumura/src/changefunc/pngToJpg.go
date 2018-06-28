package changefunc

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"
)

//PngToJpg :Encode Png file to Jpg file and return number of files which are successed to encode.
func PngToJpg(paths []string, filetype string) int {
	success := 0
	fmt.Println("PngToJpg実行")

	for _, inputPath := range paths {
		newFileNameSlice := strings.Split(inputPath, ".")

		var err error

		//try
		if newFileNameSlice[len(newFileNameSlice)-1] != "png" {
			fmt.Println(inputPath + "はpngファイルでないためスキップ")
			continue
		}

		//ファイル名を変更
		outputPath := ""
		for i := 0; i < len(newFileNameSlice)-1; i++ {
			outputPath += newFileNameSlice[i]
		}
		outputPath = outputPath + ".jpg"
		fmt.Println(inputPath + "から" + outputPath + "を作成しています")

		//出力ファイル生成
		out, err := os.Create(outputPath)

		//入力ファイルを開く
		file, err := os.Open(inputPath)

		//画像のバイナリデータ
		img, _, err := image.Decode(file)

		//変換のオプション　初期値nil
		opts := &jpeg.Options{Quality: 100}

		if err == nil {
			err := jpeg.Encode(out, img, opts)
			if err == nil {
				success++
			}
		}
		if err != nil {
			continue
		}
	}
	return success
}
