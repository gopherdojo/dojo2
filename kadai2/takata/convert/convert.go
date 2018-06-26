/*
Package convert provides convert function
to some extension to other extension
*/
package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Convert images with srcFmt in dir to images with dstFmt
func Convert(dir string, srcFmt string, dstFmt string) error {

	if !isFileOrDirExists(dir) {
		return errors.New("ディレトリは存在しません")
	}

	fInfo, error := os.Stat(dir)
	if error != nil {
		return error
	}

	if !fInfo.IsDir() {
		return errors.New("ディレクトリを指定してください")
	}

	files, error := dirwalk(dir, srcFmt, dstFmt)
	if error != nil {
		return error
	}
	for _, file := range files {
		if error := convertImage(file); error != nil {
			return error
		}
	}
	return nil
}

// find files with some extension in a directory recursively
func dirwalk(dir string, srcFmt string, dstFmt string) ([]FileInformation, error) {

	files, error := ioutil.ReadDir(dir)
	if error != nil {
		return nil, error
	}

	var paths []FileInformation
	for _, file := range files {
		if file.IsDir() {
			files, error := dirwalk(filepath.Join(dir, file.Name()), srcFmt, dstFmt)
			if error != nil {
				return nil, error
			}
			paths = append(paths, files...)
			continue
		}
		name := file.Name()
		pos := strings.LastIndex(name, ".") + 1
		if name[pos:] != srcFmt {
			continue
		}

		path := FileInformation{filepath.Join(dir, name[:pos-1]), srcFmt, dstFmt}
		paths = append(paths, path)
	}
	return paths, nil
}

// convert file with some extension to file with other extension
func convertImage(src FileInformation) error {

	if !isAvailableFormat(src.srcFmt) {
		return errors.New("指定した変換元の画像形式は無効です")
	}

	if !isAvailableFormat(src.dstFmt) {
		return errors.New("指定した変換元の画像形式は無効です")
	}

	if error := startConvert(src); error != nil {
		return error
	}
	return nil
}

// Encode the targer image
func startConvert(src FileInformation) error {

	file, error := os.Open(src.name + "." + src.srcFmt)
	if error != nil {
		return error
	}
	defer file.Close()

	img, _, error := decode(file)
	if error != nil {
		return error
	}

	dstFile, error := makeDstFile(src)
	if error != nil {
		return error
	}

	out, error := os.Create(dstFile)
	if error != nil {
		return error
	}
	defer out.Close()

	if error := encode(src.dstFmt, out, img); error != nil {
		return error
	}
	return nil
}

// decode file
func decode(file *os.File) (image.Image, string, error) {
	return image.Decode(file)
}

// encode image to dstFormat
func encode(format string, out *os.File, img image.Image) error {

	switch format {
	case "jpeg", "jpg":
		if error := jpeg.Encode(out, img, nil); error != nil {
			return error
		}
		return nil
	case "gif":
		if error := gif.Encode(out, img, nil); error != nil {
			return error
		}
		return nil
	case "png":
		if error := png.Encode(out, img); error != nil {
			return error
		}
		return nil
	default:
		return errors.New("不正な画像形式を出力先に指定しています")
	}
}

// make destination file
func makeDstFile(src FileInformation) (string, error) {
	dstDir := "output"
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		if error := os.Mkdir(dstDir, 0777); error != nil {
			return "", error
		}
	}
	return filepath.Join(dstDir, fmt.Sprintf("%s.%s", getFileNameWithoutExt(src.name), src.dstFmt)), nil
}

// check if this format is avalable
func isAvailableFormat(format string) bool {
	lowerFormat := strings.ToLower(format)
	switch lowerFormat {
	case "jpg", "jpeg", "png", "gif":
		return true
	default:
		return false

	}
}

// Get file name withour extension
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// Check if file exists
func isFileOrDirExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

// FileInformation contains Name and srcFmt and dstFmt.
type FileInformation struct {
	name   string // name of a file
	srcFmt string // original format of a file
	dstFmt string // format to convert
}
