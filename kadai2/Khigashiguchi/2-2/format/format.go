package format

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// Decoder transform binary to image
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

// Encoder transfrom image to binary
type Encoder interface {
	Encode(io.Writer, image.Image) error
}

// JPEG represents JPEG format
type JPEG struct {
	Options jpeg.Options
}

// Decode transform JPEG binary to image
func (f *JPEG) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Encode transform binary to JPEG image
func (f *JPEG) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &f.Options)
}

// PNG represents PNG format
type PNG struct {
	Options png.Encoder
}

// Decode transform PNG binary to image
func (f *PNG) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode transform PNG image to binary
func (f *PNG) Encode(w io.Writer, m image.Image) error {
	return f.Options.Encode(w, m)
}
