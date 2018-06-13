package imageconverter

import "path/filepath"

// FileInfo FilePathを基に、付属情報を取得できる関数が定義された型
type FileInfo struct {
	Path FilePath
}

// Ext hogehoge
func (fi *FileInfo) Ext() string {
	return filepath.Ext(string(fi.Path))
}

// Format hogehoge
func (fi *FileInfo) Format() Format {
	return Format(fi.Ext()[1:])
}
