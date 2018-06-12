package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/gopherdojo/dojo2/kadai1/sawadashota"
)

const (
	Png  = "png"
	Jpeg = "jpg"
	Gif  = "gif"
)

func main() {
	from := flag.String("f", "jpg", "Image type from: png, jpg or gif")
	to := flag.String("t", "png", "Image type to: png, jpg or gif")
	flag.Parse()

	if len(flag.Args()) < 1 {
		exitWithError("Please designation target directory")
	}

	r := regexImageExtension(*from)

	for _, target := range flag.Args() {

		if _, err := os.Stat(target); err != nil {
			exitWithError(err.Error())
		}

		imagePaths := targetImages(target, r)
		convertAllImages(&imagePaths, *to)
	}
}

// exitWithError print error message and exit error status
func exitWithError(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func targetImages(dir string, r *regexp.Regexp) []string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		rel := path.Join(dir, file.Name())

		if r.MatchString(file.Name()) {
			paths = append(paths, path.Join(pwd, rel))
			continue
		}

		if file.IsDir() {
			paths = append(paths, targetImages(rel, r)...)
		}
	}

	return paths
}

func regexImageExtension(convertFrom string) *regexp.Regexp {
	var r *regexp.Regexp
	switch convertFrom {
	case Png:
		r = regexp.MustCompile(`\.png`)
	case Jpeg:
		r = regexp.MustCompile(`\.jpe?g`)
	case Gif:
		r = regexp.MustCompile(`\.gif`)
	default:
		exitWithError(fmt.Sprintf("Image type should be %s, %s or %s", Png, Jpeg, Gif))
	}

	return r
}

func convertAllImages(imagePaths *[]string, convertTo string) {
	c := convert(convertTo)

	for _, imagePath := range *imagePaths {
		i, err := sawadashota.New(imagePath)

		if err != nil {
			exitWithError(err.Error())
		}

		c(i, newFilename(imagePath, convertTo))
	}
}

func convert(convertTo string) func(i sawadashota.Converter, dest string) {
	var f func(i sawadashota.Converter, dest string)
	switch convertTo {
	case Png:
		f = func(i sawadashota.Converter, dest string) {
			i.ToPng(dest)
		}
	case Jpeg:
		f = func(i sawadashota.Converter, dest string) {
			i.ToJpeg(dest)
		}
	case Gif:
		f = func(i sawadashota.Converter, dest string) {
			i.ToGif(dest)
		}
	default:
		exitWithError(fmt.Sprintf("Image type should be %s, %s or %s", Png, Jpeg, Gif))
	}

	return f
}

func newFilename(imagePath, convertTo string) string {
	var filename string
	r := regexp.MustCompile(`\.\w+$`)
	switch convertTo {
	case Png:
		filename = r.ReplaceAllString(imagePath, fmt.Sprintf(".%s", Png))
	case Jpeg:
		filename = r.ReplaceAllString(imagePath, fmt.Sprintf(".%s", Jpeg))
	case Gif:
		filename = r.ReplaceAllString(imagePath, fmt.Sprintf(".%s", Gif))
	default:
		exitWithError(fmt.Sprintf("Image type should be %s, %s or %s", Png, Jpeg, Gif))
	}
	return filename
}
