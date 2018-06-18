package tokunaga

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"io"
	"image"
)

type Ext interface {
	Decede(r io.Reader) (image.Image, error)
	Encode(w io.Writer, img image.Image) error
}

func ConvertImage(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open error: %v\n", err)
		return err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
		return err
	}
	newImage, err := os.Create(FullBasename(filename) + ".jpeg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file create error: %v\n", err)
		return err
	}
	return jpeg.Encode(newImage, img, nil)
}