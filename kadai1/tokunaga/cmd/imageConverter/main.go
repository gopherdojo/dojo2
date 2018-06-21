package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai1/tokunaga"
)

type Usage string

var extTo string
var extFrom string
var usage Usage = "./imageConverter [-from=png|jpeg ] [-to=png|jpeg ] directory"
var permmitedExts []string = []string{"png", "jpeg", "jpg"}

// 引数の初期化
func init() {
	flag.StringVar(&extFrom, "from", "jpeg", string(usage))
	flag.StringVar(&extTo, "to", "png", string(usage))
}

// メイン関数
func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: "+string(usage))
		os.Exit(1)
	}
	if !checkExtPermmited(extFrom, permmitedExts) {
		fmt.Fprintln(os.Stderr, "-from is only png, jpeg, jpg")
		os.Exit(1)
	}
	if !checkExtPermmited(extTo, permmitedExts) {
		fmt.Fprintln(os.Stderr, "-to is only png, jpeg, jpg")
		os.Exit(1)
	}
	directory := flag.Args()[0]
	if err := filepath.Walk(directory, delegateFileOperation); err != nil {
		os.Exit(1)
	}
}

// 引数のディレクトリ以下を引数の関数で再帰的に処理する
func delegateFileOperation(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && filepath.Ext(path) == "."+extFrom {
		tokunaga.ConvertImage(path, tokunaga.AdaptExt(extFrom), tokunaga.AdaptExt(extTo))
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
