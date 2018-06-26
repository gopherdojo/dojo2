package images

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

// Decoder transforms binary to an image.
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

// Encoder transforms an image to binary.
type Encoder interface {
	Encode(io.Writer, image.Image) error
}

// AutoDetect represents auto-detect.
type AutoDetect struct{}

// Decode automatically detects format and decodes the binary.
func (f *AutoDetect) Decode(r io.Reader) (image.Image, error) {
	m, _, err := image.Decode(r)
	return m, err
}

// JPEG represents JPEG format.
type JPEG struct {
	Options jpeg.Options
}

// Decode transforms the JPEG binary to image.
func (f *JPEG) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Encode transforms the JPEG image to binary.
func (f *JPEG) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &f.Options)
}

// PNG represents PNG format.
type PNG struct {
	Options png.Encoder
}

// Decode transforms the PNG binary to image.
func (f *PNG) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode transforms the PNG image to binary.
func (f *PNG) Encode(w io.Writer, m image.Image) error {
	return f.Options.Encode(w, m)
}

// GIF represents GIF format.
type GIF struct {
	Options gif.Options
}

// Decode transforms the GIF binary to image.
func (f *GIF) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

// Encode transforms the GIF image to binary.
func (f *GIF) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, &f.Options)
}
