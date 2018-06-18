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
func dirwalk(dir string, srcFmt string, dstFmt string) []File {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []File
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

		path := File{filepath.Join(dir, name[:pos-1]), srcFmt, dstFmt}
		paths = append(paths, path)
	}
	return paths
}

// convert file with some extension to file with other extension
func convertImage(src File) {
	file, err := os.Open(src.name + "." + src.srcFmt)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	dstDir := "output"
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		os.Mkdir(dstDir, 0777)
	}
	dstFile := dstDir + "/" + fmt.Sprintf("%s.%s", getFileNameWithoutExt(src.name), src.dstFmt)
	out, err := os.Create(dstFile)
	defer out.Close()

	switch src.dstFmt {
	case "jpeg", "jpg":
		err = jpeg.Encode(out, img, nil)
	case "gif":
		err = gif.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	}
	if err != nil {
		log.Fatal(err)
		return
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

// File contains Name and srcFmt and dstFmt.
type File struct {
	name   string // name of a file
	srcFmt string // original format of a file
	dstFmt string // format to convert
}
