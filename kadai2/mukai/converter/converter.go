// image converter.
package converter

import (
	"io/ioutil"
	"path/filepath"
	"os"
	"image"
	"io"
	"strings"
	"image/png"
	"image/jpeg"
	"image/gif"
	"fmt"
)

// convert image files from inputFormat to outputFormat in specific directory recursively.
// inputFormat and outputFormat is available only jpeg(jpg), gif, or png.
func Convert(dir string, inputFormat string, outputFormat string) (error) {
	if !isAvailableFormat(inputFormat) || !isAvailableFormat(outputFormat) {
		return fmt.Errorf(
			"available formats are jpg(jpeg), png, gif ONLY. input parameter is %s, outpur parameter is %s",
			inputFormat, outputFormat)
	}
	infos, _ := ioutil.ReadDir(dir)
	for _, file := range infos {
		inputConvertFile := convertFile{absPath: filepath.Join(dir, file.Name()), isDir: file.IsDir()}
		if file.IsDir() {
			Convert(inputConvertFile.absPath, inputFormat, outputFormat)
		} else if inputConvertFile.isSameExt(inputFormat) {
			outPath := inputConvertFile.arbitraryExtAbsPath(outputFormat)
			internalConvert(inputConvertFile.absPath, outPath, outputFormat)
		}
	}
	return nil
}

func internalConvert(inputFile string, outputFile string, outputFormat string) {
	out, _ := os.Create(outputFile)
	defer out.Close()
	input, err := os.Open(inputFile)
	defer input.Close()
	if err != nil {
		println(err)
	}
	decode, _, err := image.Decode(input)
	if err != nil {
		println(err)
	}
	encode(outputFormat, out, decode)
}

func isAvailableFormat(format string) bool {
	lowerFormat := strings.ToLower(format)
	switch lowerFormat {
	case "jpg", "jpeg", "gif", "png":
		return true
	default:
		return false
	}
}

func encode(format string, w io.Writer, m image.Image) {
	switch strings.ToLower(format) {
	case "png":
		png.Encode(w, m)
	case "jpg", "jpeg":
		jpeg.Encode(w, m, &jpeg.Options{Quality:100})
	case "gif":
		gif.Encode(w, m, nil)
	}
}
