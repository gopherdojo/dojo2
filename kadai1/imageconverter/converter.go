package imageconverter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
)

// Converter 画像ファイル変換器
type Converter struct{}

// Run 画像を変換する。対象フォーマットではないファイルは無視する。
func (c *Converter) Run(
	f FileInfo,
	inputFormat Format,
	outputFormat Format) {

	fFormat := f.Format()

	// 画像を変換する処理
	if fFormat.NormalizedFormat() == inputFormat {
		newPath := regexp.MustCompile(`\`+f.Ext()+`$`).ReplaceAllString(string(f.Path), outputFormat.Ext())
		println(string(f.Path) + " => " + newPath)

		orgFile, err := os.Open(string(f.Path))
		if err != nil {
			panic(err)
		}
		defer orgFile.Close()

		img, _, err := image.Decode(orgFile)
		if err != nil {
			panic(err)
		}

		newFile, err := os.Create(newPath)
		if err != nil {
			panic(err)
		}
		defer newFile.Close()

		switch string(outputFormat) {
		case "png":
			png.Encode(newFile, img)
		case "jpg":
			jpeg.Encode(newFile, img, nil)
		}
	}

}
