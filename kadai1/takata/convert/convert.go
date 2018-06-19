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
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Convert images with srcFmt in dir to images with dstFmt
func Convert(dir string, srcFmt string, dstFmt string) {

	if !exists(dir) {
		panic("ディレトリは存在しません")
	}

	fInfo, _ := os.Stat(dir)
	if !fInfo.IsDir() {
		panic("ディレクトリを指定してください")
	}

	files := dirwalk(dir, srcFmt, dstFmt)

	for _, file := range files {
		convertImage(file)
	}
}

// find files with some extension in a directory recursively
func dirwalk(dir string, srcFmt string, dstFmt string) []FileInformation {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []FileInformation
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()), srcFmt, dstFmt)...)
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
	return paths
}

// convert file with some extension to file with other extension
func convertImage(src FileInformation) {

	if !isAvailableFormat(src.srcFmt) {
		log.Fatal("指定した変換元の画像形式は無効です")
		return
	}

	if !isAvailableFormat(src.dstFmt) {
		log.Fatal("指定した変換先の画像形式は無効です")
		return
	}

	startConvert(src)
}

// Encode the targer image
func startConvert(src FileInformation) {

	file, err := os.Open(src.name + "." + src.srcFmt)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	img, _, err := decode(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	dstFile := makeDstFile(src)

	out, err := os.Create(dstFile)
	defer out.Close()

	if !encode(src.dstFmt, out, img) {
		log.Fatal("encodeに失敗")
	}
}

// decode file
func decode(file *os.File) (image.Image, string, error) {
	return image.Decode(file)
}

// encode image to dstFormat
func encode(format string, out *os.File, img image.Image) bool {

	switch format {
	case "jpeg", "jpg":
		jpeg.Encode(out, img, nil)
		return true
	case "gif":
		gif.Encode(out, img, nil)
		return true
	case "png":
		png.Encode(out, img)
		return true
	default:
		return false
	}
}

// make destination file
func makeDstFile(src FileInformation) string {
	dstDir := "output"
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		os.Mkdir(dstDir, 0777)
	}
	return dstDir + "/" + fmt.Sprintf("%s.%s", getFileNameWithoutExt(src.name), src.dstFmt)
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
func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// FileInformation contains Name and srcFmt and dstFmt.
type FileInformation struct {
	name   string // name of a file
	srcFmt string // original format of a file
	dstFmt string // format to convert
}
