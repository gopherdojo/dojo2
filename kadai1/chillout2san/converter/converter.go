package converter

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Convert(beforeExtension string, afterExtension string, path string) {
	presentExtension := filepath.Ext(path)
	if presentExtension != "." + beforeExtension {
		log.Fatal("指定した拡張子とファイルの拡張子が異なります。")
	}

	dir, fileName := filepath.Split(path)

	newFileName := strings.Replace(fileName, "." + beforeExtension, "." + afterExtension, 1)

	newPath := filepath.Join(dir, newFileName)
	if err := os.Rename(path, newPath); err != nil {
		log.Fatal(err)
	}
}