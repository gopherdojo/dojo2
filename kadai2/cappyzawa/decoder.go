package conv

import (
	"image"
	"io"
)

// Decoder - interface defining a method for decode
type Decoder interface {
	Decode(r io.Reader) (image.Image, string, error)
}

type decoder struct{}

func NewDecoder() Decoder {
	return &decoder{}
}

// Decode - wrap https://godoc.org/image#Decode
func (decoder) Decode(r io.Reader) (image.Image, string, error) {
	return image.Decode(r)
}
