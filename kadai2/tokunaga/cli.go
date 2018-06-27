package tokunaga

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type CLI struct {
	OutStream, ErrStream io.Writer
	to, from             string
}

var permmitedExts = []string{"png", "jpeg", "jpg"}

const usage = "./imageConverter [-from=png|jpeg ] [-to=png|jpeg ] directory"

// 引数処理を含めた具体的な処理
func (c *CLI) Run(args []string) int {
	// オプション引数のパース
	flags := flag.NewFlagSet("imageConverter", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&c.from, "from", "jpeg", usage)
	flags.StringVar(&c.to, "to", "png", usage)

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.ErrStream, "Usage: "+string(usage))
		return ExitCodeParseFlagError
	}
	if len(flags.Args()) != 1 {
		fmt.Fprintln(c.ErrStream, "Usage: "+string(usage))
		return ExitCodeParseFlagError
	}
	if !checkExtPermmited(c.from, permmitedExts) {
		fmt.Fprintln(c.ErrStream, "-from is only png, jpeg, jpg")
		return ExitCodeParseFlagError
	}
	if !checkExtPermmited(c.to, permmitedExts) {
		fmt.Fprintln(c.ErrStream, "-to is only png, jpeg, jpg")
		return ExitCodeParseFlagError
	}
	directory := flags.Args()[0]
	if err := filepath.Walk(directory, c.delegateFileOperation); err != nil {
		fmt.Fprintln(c.ErrStream, err)
		return ExitCodeParseFlagError
	}

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	fmt.Fprintln(c.OutStream, "conver is succeed")
	return ExitCodeOK
}

// 引数のディレクトリ以下を引数の関数で再帰的に処理する
func (c *CLI) delegateFileOperation(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && filepath.Ext(path) == "."+c.from {
		converter := Converter{AdaptExt(c.from), AdaptExt(c.to)}
		converter.ConvertImage(path)
	}
	return nil
}

// 引数の拡張子が許可されているものならばtrue, それ以外なら false を返す
func checkExtPermmited(ext string, permittedExts []string) bool {
	for _, permmitedExt := range permittedExts {
		if ext == permmitedExt {
			return true
		}
	}
	return false
}
