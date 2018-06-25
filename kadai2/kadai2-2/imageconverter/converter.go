package imageconverter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
)

// Converter 画像ファイル変換器
type Converter struct{}

// Run 画像を変換する。対象フォーマットではないファイルは無視する。
func (Converter) Run(f FileInfo, in, out Format) {

	fFormat := f.Format()

	if fFormat.NormalizedFormat() != in {
		return
	}
	// 画像を変換する処理
	reg, err := regexp.Compile(`\` + f.Ext() + `$`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "正規表現解析に失敗しました")
		return
	}
	newPath := reg.ReplaceAllString(string(f.Path), out.Ext())
	fmt.Println(string(f.Path) + " => " + newPath)

	orgFile, err := os.Open(string(f.Path))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイルが開けません")
		return
	}
	defer orgFile.Close()

	img, _, err := image.Decode(orgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "画像のDecodeに失敗しました")
		return
	}

	newFile, err := os.Create(newPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "画像のDecodeに失敗しました")
		return
	}
	defer newFile.Close()

	switch out {
	case "png":
		png.Encode(newFile, img)
	case "jpg":
		jpeg.Encode(newFile, img, nil)
	}

}
