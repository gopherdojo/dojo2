package converter

import (
	"path/filepath"
	"strings"
	"os"
	"image"
)

type Converter interface {
	Convert(outputFormat string) error
	AbsPath() string
	IsDir() bool
}

//拡張子名操作の便利機能をもつ, ファイルパスを表現する型.
type convertFile struct {
	absPath string
	isDir   bool
}

func (f convertFile) AbsPath() string {
	return f.absPath
}

func (f convertFile) IsDir() bool {
	return f.isDir
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
