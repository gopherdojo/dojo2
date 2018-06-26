// image converter.
package converter

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// convert image files from inputFormat to outputFormat in specific directory recursively.
// inputFormat and outputFormat is available only jpeg(jpg), gif, or png.
func RecursiveConvert(dir string, inputFormat string, outputFormat string) error {
	if !isAvailableFormat(inputFormat) || !isAvailableFormat(outputFormat) {
		return fmt.Errorf(
			"available formats are jpg(jpeg), png, gif ONLY. input parameter is %s, outpur parameter is %s",
			inputFormat, outputFormat)
	}
	infos, err := files(dir)
	if err != nil {
		return err
	}
	for _, file := range infos {
		inputConvertFile := convertFile{absPath: filepath.Join(dir, file.Name()), isDir: file.IsDir()}
		if file.IsDir() {
			RecursiveConvert(inputConvertFile.absPath, inputFormat, outputFormat)
		} else if inputConvertFile.isSameExt(inputFormat) {
			outPath := inputConvertFile.arbitraryExtAbsPath(outputFormat)
			err := internalConvert(inputConvertFile.absPath, outPath, outputFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to convert " + inputConvertFile.absPath)
			}
		}
	}
	return nil
}

func files(dir string) ([]os.FileInfo, error) {
	println(dir)
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return infos, e
	}
	return infos, nil
}

func internalConvert(inputFile string, outputFile string, outputFormat string) error {
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	input, err := os.Open(inputFile)
	defer input.Close()
	if err != nil {
		return err
	}
	decode, _, err := image.Decode(input)
	if err != nil {
		return err
	}
	if encoder := GetEncoder(outputFormat); encoder != nil {
		return encoder.Encode(out, decode)
	}
	return nil
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


