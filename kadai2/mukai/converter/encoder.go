package converter

import (
	"strings"
	"io"
	"image"
	"image/jpeg"
	"image/gif"
	"image/png"
)

type encoder interface {
	Encode(io.Writer, image.Image) error
}

func getEncoder(format string) encoder {
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
	encoder
}

func (e jpegEncoder) Encode(writer io.Writer, image image.Image) error {
	return jpeg.Encode(writer, image, &jpeg.Options{Quality: 100})
}

type gifEncoder struct {
	encoder
}

func (e gifEncoder) Encode(writer io.Writer, image image.Image) error {
	return gif.Encode(writer, image, nil)
}

type pngEncoder struct {
	encoder
}

func (e pngEncoder) Encode(writer io.Writer, image image.Image) error {
	return png.Encode(writer, image)
}
