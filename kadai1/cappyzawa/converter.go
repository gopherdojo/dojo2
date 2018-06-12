package conv

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type Jpeg struct {
}

func (j *Jpeg) Encode(w io.Writer, m image.Image) error {
	o := &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	}
	return jpeg.Encode(w, m, o)
}

func (j *Jpeg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

type Png struct {
}

func (p *Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func (p *Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

type Converter interface {
	Encode(w io.Writer, m image.Image) error
	Decode(r io.Reader) (image.Image, error)
}
