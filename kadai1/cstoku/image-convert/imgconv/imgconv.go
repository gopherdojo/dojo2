package imgconv

import (
	"path/filepath"
	"os"
	"image/jpeg"
	"image/png"
	"image/gif"
	"errors"
	"image"
	"strings"
	"regexp"
	"fmt"
	"flag"
	"io"
)

type imgType struct {
	Type string
	Ext  string
}

// 画像のタイプと拡張子のスライス
var supportImgTypes = []imgType{
	{"jpg", ".jpg"},
	{"jpg", ".jpeg"},
	{"png", ".png"},
	{"gif", ".gif"},
}

// 画像の読み込み
func ReadImg(reader io.Reader, inputType string) (img image.Image, err error) {
	switch strings.ToLower(inputType) {
	case "jpg":
		img, err = jpeg.Decode(reader)
	case "png":
		img, err = png.Decode(reader)
	case "gif":
		img, err = gif.Decode(reader)
	default:
		err = errors.New("Illegal input type.")
	}
	return
}

// 画像の書き出し
func WriteImg(writer io.Writer, img image.Image,outputType string) (err error) {
	switch strings.ToLower(outputType) {
	case "jpg":
		err = jpeg.Encode(writer, img, nil)
	case "png":
		err = png.Encode(writer, img)
	case "gif":
		err = gif.Encode(writer, img, nil)
	default:
		err = errors.New("Illegal output type.")
	}
	return
}

// 画像の変換
func ConvertImg(inputType, outputType, path string) error {
	// 入力のタイプとファイルの拡張子をチェック
	if !IsSupportImgType(inputType, path) {
		return errors.New("Unsupport input type.")
	}

	// 画像ファイルをファイルを開く
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// 画像ファイルを読み込み
	img, err := ReadImg(f, inputType)

	// 拡張子を拾う正規表現
	re := regexp.MustCompile(`\.(jpg|png|gif)$`)

	// 入力画像のパスをもとに拡張子だけ置き換えてファイルを作成
	outputFile, err := os.Create(re.ReplaceAllString(path, "." + outputType))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 出力タイプを指定してファイルを書き出し
	return WriteImg(outputFile, img, outputType)
}

// 画像タイプとファイルの拡張子をチェック
func IsSupportImgType(imgType, path string) bool {
	imgtype := strings.ToLower(imgType)
	for _, i := range supportImgTypes {
		// 画像のタイプとファイルの拡張子をチェック
		if i.Type == imgtype && i.Ext == filepath.Ext(path) {
			return true
		}
	}
	return false
}

func WalkConvert(inputType, outputType, path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// 対応ファイルだった場合変換を実行
		if  !info.IsDir() && IsSupportImgType(inputType, path) {
			if err := ConvertImg(inputType, outputType, path); err == nil {
				fmt.Printf("Converted %s -> %s\n", path, outputType)
			} else {
				return err
			}
		}
		return nil
	})
}

func RunCli() {
	var (
		exitCode = 0
		inputType = flag.String("i", "jpg", "Input image type")
		outputType = flag.String("o", "png", "Output image type")
	)

	flag.Parse()

	for _, dir := range flag.Args() {
		if f, err := os.Stat(dir); err == nil && f.IsDir() {
			if err := WalkConvert(*inputType, *outputType, dir); err != nil {
				fmt.Println(err)
				exitCode = 1
				break
			}
		} else {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(dir, "is not a directory.")
			}

			exitCode = 1
			break
		}
	}

	os.Exit(exitCode)
}
