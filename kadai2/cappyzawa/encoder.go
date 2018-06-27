package conv

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// Encoder - interface defining a method for encode
type Encoder interface {
	Encode(w io.Writer, m image.Image) error
}

// Jpeg - struct for operating https://godoc.org/image/jpeg
type Jpeg struct {
	Options *jpeg.Options
}

// Png - struct for operation https://godoc.org/image/png
type Png struct {
}

// Encode - wrap https://godoc.org/image/jpeg#Encode
func (j *Jpeg) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, j.Options)
}

// Encode - wrap https://godoc.org/image/png#Encode
func (p *Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
