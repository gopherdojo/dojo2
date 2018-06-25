package converter

import (
	"strings"
	"io"
	"image"
	"image/jpeg"
	"image/gif"
	"image/png"
)

type Encoder interface {
	Encode(io.Writer, image.Image)
}

func GetEncoder(format string) Encoder {
	switch strings.ToLower(format) {
	case "png":
		return pngEncoder{}
	case "jpg", "jpeg":
		return jpegEncoder{}
	case "gif":
		return gifEncoder{}
	default:
		return nil
	}
}

type jpegEncoder struct {
	Encoder
}

func (e jpegEncoder) Encode(writer io.Writer, image image.Image) {
	jpeg.Encode(writer, image, &jpeg.Options{Quality: 100})
}

type gifEncoder struct {
	Encoder
}

func (e gifEncoder) Encode(writer io.Writer, image image.Image) {
	gif.Encode(writer, image, nil)
}

type pngEncoder struct {
	Encoder
}

func (e pngEncoder) Encode(writer io.Writer, image image.Image) {
	png.Encode(writer, image)
}
