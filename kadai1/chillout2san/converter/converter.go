package converter

import (
	imagehandler "dojo/kadai1/chillout2san/imagehandler"
	"errors"
	"io/fs"
	"path/filepath"
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
			return err
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

		if presentExtension != ".jpg" && presentExtension != ".png" {
			return nil
		}

		if presentExtension != "."+beforeExtension {
			return errors.New("指定された拡張子が実際の拡張子と異なります。")
		}

		image, err := imagehandler.Decode(targetPath)
		if err != nil {
			return err
		}

		err = imagehandler.Encode(image, beforeExtension, afterExtension, targetPath)
		if err != nil {
			return err
		}
	}

	return nil
}
