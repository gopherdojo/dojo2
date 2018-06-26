package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai2/bookun/converter"
	"github.com/gopherdojo/dojo2/kadai2/bookun/format"
)

const (
	// ExitCodeOK - Runメソッドが成功した場合のの返り値
	ExitCodeOK = iota
	// ExitCodeParseFlagError - コマンドラインオプションのパースのエラー
	ExitCodeParseFlagError
	// ExitCodeCreateFormat - Format型の作成失敗時のエラー
	ExitCodeCreateFormat
	// ExitCodeSearchError - ディレクトリ内探索のエラー
	ExitCodeSearchError
	// ExitCodeConvertError - 変換失敗時のエラー
	ExitCodeConvertError
)

//CLI - io.Writer型の変数を2つ持つ構造体
type CLI struct {
	outStream, errStream io.Writer
}

//Run - 変換のための一連の処理を実行
func (c *CLI) Run(args []string) int {
	// Parse
	var srcFormat string
	var dstFormat string
	flags := flag.NewFlagSet("convert-cli", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.StringVar(&srcFormat, "s", "jpg", "src format")
	flags.StringVar(&dstFormat, "d", "png", "dest format")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if flags.NArg() == 0 {
		fmt.Fprintf(c.errStream, "./convert-cli directoryPath\n")
		return ExitCodeParseFlagError
	}

	format, err := format.NewFormat(srcFormat, dstFormat)
	if err != nil {
		return ExitCodeCreateFormat
	}

	// ----

	// ディレクトリ探索
	dirName := flags.Arg(0)
	targetFiles, err := c.searchFiles(dirName, "."+srcFormat)
	if err != nil {
		fmt.Fprintf(c.errStream, "searching err\n")
		return ExitCodeSearchError
	}
	// ---

	// 1ファイル毎に処理
	for _, v := range targetFiles {
		dstFileName := v[:len(v)-len(filepath.Ext(v))] + "." + dstFormat
		converter := converter.NewConverter(v, dstFileName, *format)
		if err := converter.Convert(); err != nil {
			return ExitCodeConvertError
		}
		fmt.Fprintf(c.outStream, "%s was converted to %s\n", v, dstFileName)
	}
	// ----
	return ExitCodeOK
}

func (c *CLI) searchFiles(targetDirName, targetFormat string) ([]string, error) {
	var targetFiles []string
	err := filepath.Walk(targetDirName, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == targetFormat {
			targetFiles = append(targetFiles, path)
		}
		return err
	})
	return targetFiles, err
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
