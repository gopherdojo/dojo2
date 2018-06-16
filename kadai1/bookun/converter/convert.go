package converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Converter - ファイル名を持つ構造体
type Converter struct {
	srcFileName, dstFileName string
}

//NewConverter - srcFileの名前とdstFileの名前を受け取りConverterを生成
func NewConverter(srcFileName, dstFileName string) *Converter {
	c := Converter{}
	c.srcFileName = srcFileName
	c.dstFileName = dstFileName
	return &c
}

//Convert - srcFile から dstFile へ変換を行う
func (c *Converter) Convert() bool {
	srcFile, err := os.Open(c.srcFileName)
	if err != nil {
		return false
	}
	defer srcFile.Close()
	img, _, err := image.Decode(srcFile)
	if err != nil {
		return false
	}
	dstFile, err := os.Create(c.dstFileName)
	if err != nil {
		return false
	}
	switch filepath.Ext(c.dstFileName) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(dstFile, img, nil)
	case ".png":
		err = png.Encode(dstFile, img)
	}
	if err != nil {
		return false
	}
	return true
}
