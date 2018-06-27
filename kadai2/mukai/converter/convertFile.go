package converter

import (
	"os"
	"image"
	"fmt"
)

type converterFileInterface interface {
	convert(outputFormat string) (string, error)
	absolutePath() string
	isDirectory() bool
}

//拡張子名操作の便利機能をもつ, ファイルパスを表現する型.
type converterFile struct {
	absPath string
	isDir   bool
}

func (f converterFile) absolutePath() string {
	return f.absPath
}

func (f converterFile) isDirectory() bool {
	return f.isDir
}

func (f converterFile) convert(outputFormat string) (string, error) {
	outputFile := ArbitraryExtAbsPath(f.absPath, outputFormat)
	out, err := os.Create(outputFile)
	if err != nil {
		return "", err
	}
	defer out.Close()
	input, err := os.Open(f.absPath)
	defer input.Close()
	if err != nil {
		return "", err
	}
	decode, _, err := image.Decode(input)
	if err != nil {
		return "", err
	}
	if encoder := GetEncoder(outputFormat); encoder != nil {
		err = encoder.Encode(out, decode)
		if err != nil {
			return "", nil
		}
		return outputFile, nil
	}
	return "", fmt.Errorf("encoder is nil")
}
