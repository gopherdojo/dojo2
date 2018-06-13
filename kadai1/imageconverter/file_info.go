package imageconverter

import "path/filepath"

// FileInfo FilePathを基に、付属情報を取得できる関数が定義された型
type FileInfo struct {
	Path FilePath
}

// Ext .付き拡張子の文字列を返す
func (fi *FileInfo) Ext() string {
	return filepath.Ext(string(fi.Path))
}

// Format .なし拡張子の文字列を返す
func (fi *FileInfo) Format() Format {
	return Format(fi.Ext()[1:])
}
