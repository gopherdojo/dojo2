package converter

import (
	"strings"
	"path/filepath"
)

//拡張子の取得(.なし)
func extension(path string) string {
	list := strings.Split(filepath.Ext(path), ".")
	if 2 <= len(list) {
		return list[len(list)-1]
	}
	return ""
}

//拡張子が同じか判定.
func isSameExt(path string, ext string) bool {
	return strings.ToLower(extension(path)) == strings.ToLower(ext)
}

//任意の拡張子に変換したパスを取得.
func arbitraryExtAbsPath(filePath string, ext string) string {
	dir, file := filepath.Split(filePath)
	split := strings.Split(file, ".")
	if len(split) < 2 {
		return filePath
	}
	return filepath.Join(dir, split[0]) + "." + ext
}
