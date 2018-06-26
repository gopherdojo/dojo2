package converter

import (
	"path/filepath"
	"strings"
	"os"
	"image"
)

type Converter interface {
	Convert(inputFile string, outputFormat string) error
}

//拡張子名操作の便利機能をもつ, ファイルパスを表現する型.
type convertFile struct {
	absPath string
	isDir   bool
}

//拡張子の取得(.なし)
func (f convertFile) ext() string {
	list := strings.Split(filepath.Ext(f.absPath), ".")
	if 2 <= len(list) && !f.isDir {
		return list[len(list)-1]
	}
	return ""
}

//任意の拡張子に変換したパスを取得.
func (f convertFile) arbitraryExtAbsPath(ext string) string {
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
func (f convertFile) isSameExt(ext string) bool {
	if f.isDir {
		return false
	}
	return strings.ToLower(f.ext()) == strings.ToLower(ext)
}

func (f convertFile) Convert(outputFormat string) error {
	outputFile := f.arbitraryExtAbsPath(outputFormat)
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	input, err := os.Open(f.absPath)
	defer input.Close()
	if err != nil {
		return err
	}
	decode, _, err := image.Decode(input)
	if err != nil {
		return err
	}
	if encoder := GetEncoder(outputFormat); encoder != nil {
		return encoder.Encode(out, decode)
	}
	return nil
}
