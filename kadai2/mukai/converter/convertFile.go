package converter

import (
	"strings"
	"path/filepath"
)

//拡張子名操作の便利機能をもつ, ファイルパスを表現する型.
type convertFile struct {
	absPath string
	isDir bool
}

//拡張子の取得(.なし)
func (f convertFile) ext() (string) {
	list := strings.Split(filepath.Ext(f.absPath), ".")
	if 2 <= len(list) && !f.isDir {
		return list[len(list) - 1]
	}
	return ""
}
//任意の拡張子に変換したパスを取得.
func (f convertFile) arbitraryExtAbsPath(ext string) (string) {
	dir, file := filepath.Split(f.absPath)
	if f.isDir {
		return dir
	}
	split := strings.Split(file, ".")
	if len(split) < 2 {
		return f.absPath
	}
	return filepath.Join(dir, split[0]) + "." + ext
}
//拡張子が同じか判定.
func (f convertFile) isSameExt(ext string) (bool) {
	if f.isDir {
		return false
	}
	return strings.ToLower(f.ext()) == strings.ToLower(ext)
}
