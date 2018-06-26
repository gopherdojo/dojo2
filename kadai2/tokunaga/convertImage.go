package tokunaga

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type PngWrapper string
type JpegWrapper string

type Decoder interface {
	Decode(reader io.Reader) (image.Image, error)
}
type Encoder interface {
	Encode(writer io.Writer, image image.Image) error
}

// ファイルのデコード, エンコード, 拡張子文字列返却用インターフェース
type DecodeEncoder interface {
	Decoder
	Encoder
	Ext() string
}

// 引数の文字列の拡張子を表すラッパークラスを返す
func AdaptExt(ext string) DecodeEncoder {
	if ext == "jpeg" || ext == "jpg" {
		return JpegWrapper(ext)
	}
	if ext == "png" {
		return PngWrapper(ext)
	}
	return PngWrapper(ext)
}

// filenameで指定されたファイルを extFrom から extFrom に変換する 例) png -> jepg
func ConvertImage(filename string, extFrom DecodeEncoder, extTo DecodeEncoder) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open error: %v\n", err)
		return err
	}
	defer file.Close()
	img, err := extFrom.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
		return err
	}
	newImage, err := os.Create(FullBasename(filename) + "." + extTo.Ext())
	if err != nil {
		fmt.Fprintf(os.Stderr, "file create error: %v\n", err)
		return err
	}
	return extTo.Encode(newImage, img)
}

func (p PngWrapper) Encode(writer io.Writer, image image.Image) error {
	return png.Encode(writer, image)
}

func (p PngWrapper) Decode(reader io.Reader) (image.Image, error) {
	return png.Decode(reader)
}

func (p PngWrapper) Ext() string {
	return string(p)
}

func (j JpegWrapper) Encode(writer io.Writer, image image.Image) error {
	return jpeg.Encode(writer, image, nil)
}

func (j JpegWrapper) Decode(reader io.Reader) (image.Image, error) {
	return jpeg.Decode(reader)
}

func (j JpegWrapper) Ext() string {
	return string(j)
}
