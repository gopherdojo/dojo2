// Package extension create & change extention of images.
package extension

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Arg define comand line argument
type Arg struct {
	From string
	To   string
	Path string
}

// Convert create & convert extension of images. png jpg jpeg gif are only supported.
func (a Arg) Convert() error {
	if a.From != "jpg" && a.From != "png" && a.From != "gif" {
		return fmt.Errorf("%s is not supported", a.From)
	}
	if a.To != "jpg" && a.To != "png" && a.To != "gif" {
		return fmt.Errorf("%s is not supported", a.To)
	}

	file, err := os.Open(a.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	extention := a.To
	output := strings.Replace(a.Path, a.From, "", 1)
	dstfile, err := os.Create(output + extention)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	switch extention {
	case "jpeg", "jpg":
		err = jpeg.Encode(dstfile, img, nil)
	case "gif":
		err = gif.Encode(dstfile, img, nil)
	case "png":
		err = png.Encode(dstfile, img)
	}
	if err != nil {
		return err
	}
	return nil
}
