package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"io"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai1/sawadashota"
)

const (
	Png  = "png"
	Jpeg = "jpg"
	Gif  = "gif"
)

var (
	errorWriter      io.Writer
	newFilenameRegex *regexp.Regexp
)

func init() {
	errorWriter = os.Stderr
	newFilenameRegex = regexp.MustCompile(`\.\w+$`)
}

func main() {
	from := flag.String("f", "jpg", "Image type from: png, jpg or gif")
	to := flag.String("t", "png", "Image type to: png, jpg or gif")
	flag.Parse()

	if len(flag.Args()) < 1 {
		exitWithError("Please designation target directory")
	}

	r, err := regexImageExtension(*from)

	if err != nil {
		exitWithError(err.Error())
	}

	for _, target := range flag.Args() {

		err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
			switch {
			case err != nil:
				return err
			case r.MatchString(filepath.Ext(path)):
				i, err := sawadashota.New(path)

				if err != nil {
					return err
				}

				if err := convert(i, newFilename(path, *to), *to); err != nil {
					return err
				}

			default:
				return nil
			}

			if r.MatchString(filepath.Ext(path)) {

			}
			return nil
		})

		if err != nil {
			exitWithError(err.Error())
		}
	}
}

// exitWithError print error message and exit error status
func exitWithError(message string) {
	fmt.Fprintln(errorWriter, message)
	os.Exit(1)
}

func regexImageExtension(convertFrom string) (*regexp.Regexp, error) {
	var r *regexp.Regexp
	var err error

	switch convertFrom {
	case Png:
		r, err = regexp.Compile(`\.png`)
	case Jpeg:
		r, err = regexp.Compile(`\.jpe?g`)
	case Gif:
		r, err = regexp.Compile(`\.gif`)
	default:
		exitWithError(fmt.Sprintf("Image type should be %s, %s or %s", Png, Jpeg, Gif))
	}

	return r, err
}

func convert(i *sawadashota.Image, dest string, convertTo string) error {
	var err error

	switch convertTo {
	case Png:
		err = i.ToPng(dest)
	case Jpeg:
		err = i.ToJpeg(dest)
	case Gif:
		err = i.ToGif(dest)
	default:
		err = fmt.Errorf("image type should be %s, %s or %s", Png, Jpeg, Gif)
	}

	return err
}

func newFilename(imagePath, convertTo string) string {
	var filename string
	switch convertTo {
	case Png:
		filename = newFilenameRegex.ReplaceAllString(imagePath, fmt.Sprintf(".%s", Png))
	case Jpeg:
		filename = newFilenameRegex.ReplaceAllString(imagePath, fmt.Sprintf(".%s", Jpeg))
	case Gif:
		filename = newFilenameRegex.ReplaceAllString(imagePath, fmt.Sprintf(".%s", Gif))
	default:
		exitWithError(fmt.Sprintf("Image type should be %s, %s or %s", Png, Jpeg, Gif))
	}
	return filename
}
