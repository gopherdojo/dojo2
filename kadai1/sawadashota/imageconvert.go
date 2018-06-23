package sawadashota

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// Image object
type Image struct {
	Src image.Image
}

// New Converter interface
func New(path string) (*Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	return &Image{Src: img}, nil
}

// ToJpeg convert jpeg image format
func (i *Image) ToJpeg(dest string) error {

	out, err := os.Create(dest)

	if err != nil {
		return err
	}
	defer out.Close()

	options := &jpeg.Options{Quality: 100}

	jpeg.Encode(out, i.Src, options)

	return nil
}

// ToPng convert png image format
func (i *Image) ToPng(dest string) error {
	out, err := os.Create(dest)

	if err != nil {
		return err
	}
	defer out.Close()

	png.Encode(out, i.Src)

	return nil
}

// ToGif convert gif image format
func (i *Image) ToGif(dest string) error {
	out, err := os.Create(dest)

	if err != nil {
		return err
	}
	defer out.Close()

	gif.Encode(out, i.Src, nil)

	return nil
}
