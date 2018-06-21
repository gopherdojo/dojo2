package imgconverter

import (
	"image"
	"image/jpeg"
	"os"
	"io"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"image/png"
	"image/gif"
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

// ステータスの終了コード
//  ExitCodeOK: 正常終了
//  ExitCodeNG: 正常終了
const(
	ExitCodeOK = iota // 0
	ExitCodeNG	      // 1
)

// Runでメインの処理を記述する
func (c* CLI) Run(args []string) int {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.SetOutput(c.ErrStream)

	// 読み込み元フォーマットの設定
	// TODO: 短いフラグも定義する
	const defaultSrcFormat = "jpg"
	var srcFormat string
	flags.StringVar(&srcFormat, "src", defaultSrcFormat, "specify source format to convert")

	// 変換後フォーマットの設定
	// TODO: 短いフラグも定義する
	const defaultTargetFormat = "png"
	var targetFormat string
	flags.StringVar(&targetFormat, "dst", defaultTargetFormat, "specify target format to convert")

	// --helpや-hを出力した際にstatuscodeを0にするようにUsageを書き換える
	flags.Usage = func() {
		fmt.Println("Usage:")
		flags.PrintDefaults()
		os.Exit(0)
	}

	flags.Parse(args[1:])

	// 入力されたフォーマットに.を足す
	srcFormat = "." + srcFormat
	targetFormat = "." + targetFormat

	// 変換対象外のフォーマットを指定していたらエラー終了
	if checkUnacceptableFormat(srcFormat) || checkUnacceptableFormat(targetFormat) {
		fmt.Fprint(c.ErrStream, "Please specify the convertable format")
		return ExitCodeNG
	}

	// ディレクトリ名が指定されていなかったらエラー終了
	if flags.NArg() != 1 {
		fmt.Fprint(c.ErrStream, "Please specify the directory name to convert")
		return ExitCodeNG
	}

	// ディレクトリ名を引数から取得
	dirName := flags.Arg(0)

	// 変換対象ファイルを入力されたディレクトリ名から探し出す
	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		// 拡張子が小文字でも大文字でも検索できるように検索された拡張子を小文字化する
		fileExtension := filepath.Ext(path)
		fileExtensionLower := strings.ToLower(fileExtension)
		if fileExtensionLower == srcFormat {
			if err := ConvertImage(path, fileExtension, targetFormat); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(c.ErrStream, "%v", err)
		return ExitCodeNG
	}

	return ExitCodeOK
}

// checkUnacceptableFormatは本プログラムが受け付ける画像変換のフォーマットをチェックする関数
//  input:  format(string) -> 変換元画像の形式(.pngなど)
//  output: bool           -> true: 入力されたフォーマットは本プログラムでは受付できません
//                         -> false: 入力されたフォーマットは本プログラムでは受付可能
func checkUnacceptableFormat(format string) bool {
	acceptableFormat := [...]string{".jpg", ".png", ".gif"}
	for _, acc := range acceptableFormat {
		if acc == format {
			return false
		}
	}
	return true
}

// convertImageで画像の変換を行う
//
//  input:  src(string)           -> 変換元画像ファイルのフルパス
//          sourceFormat(string)  -> 変換元画像の形式(.png,.gifなど)
//          targetFormat(string)  -> 変換先画像の形式(.png,.gifなど)
//
//  output: error                 -> エラーの場合の情報、正常終了の場合 nil
func ConvertImage(src string, sourceFormat string, targetFormat string) error {
	file, err := os.Open(src);
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dest := strings.Replace(src, sourceFormat, targetFormat, 1)
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	switch targetFormat {
	case ".jpg":
		if err := jpeg.Encode(out, img, nil); err != nil {
			return err
		}
	case ".png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	case ".gif":
		if err := gif.Encode(out, img, nil); err != nil {
			return err
		}
	}

	return nil
}

