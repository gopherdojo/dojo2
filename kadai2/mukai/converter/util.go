package converter

import (
	"strings"
	"path/filepath"
)

//拡張子の取得(.なし)
func Ext(path string) string {
	list := strings.Split(filepath.Ext(path), ".")
	if 2 <= len(list) {
		return list[len(list)-1]
	}
	return ""
}

//拡張子が同じか判定.
func IsSameExt(path string, ext string) bool {
	return strings.ToLower(Ext(path)) == strings.ToLower(ext)
}
