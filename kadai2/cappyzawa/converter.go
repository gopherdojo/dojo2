package conv

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// Converter - interface defining methods for conversion
type Converter interface {
	Encode(w io.Writer, m image.Image) error
	Decode(r io.Reader) (image.Image, error)
}

// Jpeg - struct for operating https://godoc.org/image/jpeg
type Jpeg struct {
}

// Encode - wrap https://godoc.org/image/jpeg#Encode
func (j *Jpeg) Encode(w io.Writer, m image.Image) error {
	o := &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	}
	return jpeg.Encode(w, m, o)
}

// Decode - wrap https://godoc.org/image/jpeg#Decode
func (j *Jpeg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Png - struct for operation https://godoc.org/image/png
type Png struct {
}

// Encode - wrap https://godoc.org/image/png#Encode
func (p *Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

// Decode - wrap https://godoc.org/image/png#Decode
func (p *Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}
