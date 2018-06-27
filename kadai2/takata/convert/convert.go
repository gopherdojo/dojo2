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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Converter is interface to convert all files in dir.
type Converter interface {
	Convert(dir string) error
}

// ImageConverter for images contains from format and to format.
type ImageConverter struct {
	From string // original format of a file
	To   string // format to convert
}

// FileInfo contains Name and srcFmt and dstFmt.
type FileInfo struct {
	name string // name of a file
}

// Convert images with srcFmt in dir to images with dstFmt
func (i ImageConverter) Convert(dir string) error {

	if !isFileOrDirExists(dir) {
		return errors.New("ディレクトリは存在しません")
	}

	fInfo, error := os.Stat(dir)
	if error != nil {
		return errors.Wrapf(error, "os.Stat() with %s", dir)
	}

	if !fInfo.IsDir() {
		return errors.New("ディレクトリを指定してください")
	}

	if !isAvailableFormat(i.From) {
		return errors.New("指定した変換元の画像形式は無効です")
	}

	if !isAvailableFormat(i.To) {
		return errors.New("指定した変換先の画像形式は無効です")
	}

	files, error := i.dirwalk(dir)
	if error != nil {
		return error
	}
	for _, file := range files {
		if error := i.convertImage(file); error != nil {
			return error
		}
	}
	return nil
}

// find files with some extension in a directory recursively
func (i ImageConverter) dirwalk(dir string) ([]FileInfo, error) {

	files, error := ioutil.ReadDir(dir)
	if error != nil {
		return nil, errors.Wrapf(error, "ioutil.ReadDir() with %s", dir)
	}

	var paths []FileInfo
	for _, file := range files {
		if file.IsDir() {
			files, error := i.dirwalk(filepath.Join(dir, file.Name()))
			if error != nil {
				return nil, error
			}
			paths = append(paths, files...)
			continue
		}
		name := file.Name()
		pos := strings.LastIndex(name, ".") + 1
		if name[pos:] != i.From {
			continue
		}

		path := FileInfo{filepath.Join(dir, name[:pos-1])}
		paths = append(paths, path)
	}
	return paths, nil
}

// convert file with some extension to file with other extension
func (i ImageConverter) convertImage(src FileInfo) error {
	if error := i.convert(src); error != nil {
		return error
	}
	return nil
}

// Convert the targer image
func (i ImageConverter) convert(src FileInfo) error {

	fileName := src.name + "." + i.From
	file, error := os.Open(fileName)
	if error != nil {
		return errors.Wrapf(error, "os.Open() with %s", fileName)
	}
	defer file.Close()

	img, error := i.decode(file)
	if error != nil {
		return error
	}

	dstFile, error := i.makeDstFile(src)
	if error != nil {
		return error
	}

	out, error := os.Create(dstFile)
	if error != nil {
		return errors.Wrapf(error, "os.Create() with %s", dstFile)
	}
	defer out.Close()

	if error := i.encode(out, img); error != nil {
		return error
	}
	return nil
}

// encode image to dstFormat
func (i ImageConverter) encode(out io.Writer, img image.Image) error {

	switch i.To {
	case "jpeg", "jpg":
		if error := jpeg.Encode(out, img, nil); error != nil {
			return errors.Wrapf(error, "jpeg.Encode() with %s", i.To)
		}
		return nil
	case "gif":
		if error := gif.Encode(out, img, nil); error != nil {
			return errors.Wrapf(error, "gif.Encode() with %s", i.To)
		}
		return nil
	case "png":
		if error := png.Encode(out, img); error != nil {
			return errors.Wrapf(error, "png.Encode() with %s", i.To)
		}
		return nil
	default:
		return errors.New("不正な画像形式を出力先に指定しています")
	}
}

// decode image
func (i ImageConverter) decode(r io.Reader) (image.Image, error) {
	img, _, error := image.Decode(r)
	if error != nil {
		return nil, error
	}
	return img, nil
}

// make destination file
func (i ImageConverter) makeDstFile(src FileInfo) (string, error) {
	dstDir := "output"
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		os.Mkdir(dstDir, 0777)
	}
	return filepath.Join(dstDir, fmt.Sprintf("%s.%s", getFileNameWithoutExt(src.name), i.From)), nil
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
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true

}
