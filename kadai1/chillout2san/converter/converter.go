package converter

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/*
	第三引数のディレクトリ以下を再起的に処理します。
	第一引数に与えられた拡張子と現在の拡張子が異なる場合、errorを返します。
	第二引数に与えられた拡張子の画像に変更します。
*/
func Convert(beforeExtension string, afterExtension string, targetDir string) error {
	targetPathList := []string{}

	err := filepath.WalkDir(targetDir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		targetPathList = append(targetPathList, path)

		return nil
	})

	if err != nil {
		return err
	}

	for _, targetPath := range targetPathList {
		presentExtension := filepath.Ext(targetPath)

		if presentExtension != "."+beforeExtension {
			return errors.New("指定された拡張子が実際の拡張子と異なります。")
		}

		dir, fileName := filepath.Split(targetPath)

		newFileName := strings.Replace(fileName, "."+beforeExtension, "."+afterExtension, 1)

		newPath := filepath.Join(dir, newFileName)
		if err := os.Rename(targetPath, newPath); err != nil {
			return err
		}
	}

	return nil
}
