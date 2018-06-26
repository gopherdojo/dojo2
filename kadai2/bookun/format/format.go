package format

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// Format - Decoder, Encoderを保持する構造体
type Format struct {
	Decoder Decoder
	Encoder Encoder
}

// NewFormat - 画像を変換するためのDecoderとEncoderを選択してFormatを返す
func NewFormat(srcFormat, dstFormat string) (*Format, error) {
	format := Format{}
	if err := format.setDecoder(srcFormat); err != nil {
		return nil, err
	}
	if err := format.setEncoder(dstFormat); err != nil {
		return nil, err
	}
	return &format, nil
}

// Decoder - Decodeの実装が必要なインターフェース
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

// Encoder - Encodeの実装が必要なインターフェース
type Encoder interface {
	Encode(io.Writer, image.Image) error
}

// JPEG - 空フィールドを持つ構造体
type JPEG struct{}

// Decode - jpeg.Decode
func (j *JPEG) Decode(srcFile io.Reader) (image.Image, error) {
	img, err := jpeg.Decode(srcFile)
	return img, err
}

// Encode - jpeg.Encode
func (j *JPEG) Encode(dstFile io.Writer, img image.Image) error {
	err := jpeg.Encode(dstFile, img, nil)
	return err
}

// PNG - 空フィールドを持つ構造体
type PNG struct{}

// Decode - png.Decode
func (p *PNG) Decode(srcFile io.Reader) (image.Image, error) {
	img, err := png.Decode(srcFile)
	return img, err
}

// Encode - png.Encode
func (p *PNG) Encode(dstFile io.Writer, img image.Image) error {
	err := png.Encode(dstFile, img)
	return err
}

func (f *Format) setDecoder(srcFormat string) error {
	switch srcFormat {
	case "jpeg", "jpg":
		f.Decoder = &JPEG{}
	case "png":
		f.Decoder = &PNG{}
	default:
		return fmt.Errorf("Supported formats are only jpeg, jpg or png")
	}
	return nil
}

func (f *Format) setEncoder(dstFormat string) error {
	switch dstFormat {
	case "jpeg", "jpg":
		f.Encoder = &JPEG{}
	case "png":
		f.Encoder = &PNG{}
	default:
		return fmt.Errorf("Supported formats are only jpeg, jpg or png")
	}
	return nil
}
