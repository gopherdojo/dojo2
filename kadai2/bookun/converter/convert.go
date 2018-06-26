package converter

import (
	"os"

	"github.com/gopherdojo/dojo2/kadai2/bookun/format"
)

// Converter - ファイル名を持つ構造体
type Converter struct {
	srcFileName, dstFileName string
	Format                   format.Format
}

//NewConverter - srcFileの名前とdstFileの名前を受け取りConverterを生成
func NewConverter(srcFileName, dstFileName string, format format.Format) *Converter {
	c := Converter{}
	c.srcFileName = srcFileName
	c.dstFileName = dstFileName
	c.Format = format
	return &c
}

//Convert - srcFile から dstFile へ変換を行う
func (c *Converter) Convert() error {
	srcFile, err := os.Open(c.srcFileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	img, err := c.Format.Decoder.Decode(srcFile)
	if err != nil {
		return err
	}
	dstFile, err := os.Create(c.dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	err = c.Format.Encoder.Encode(dstFile, img)
	return err
}
