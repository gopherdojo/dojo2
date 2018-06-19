// Package converter Overview
package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// MyImageFile contains fileExt, filepath before converted, and after converted.
type MyImageFile struct {
	fileExt      string // file extension
	fromFilepath string // from filepath
	toFilepath   string // to filepath
}

// ConvertImagesFromJpgToPngInDir converts all jpg images to png in target directory.
func ConvertImagesFromJpgToPngInDir(dir string) {
	for _, path := range GetFilepathListByExt(dir, ".jpg") {
		myImageFile := MyImageFile{fileExt: "jpg", fromFilepath: path, toFilepath: ""}
		if err := ConvertJpgToPng(&myImageFile); err != nil && myImageFile.fileExt == "png" {
			fmt.Println("[ERROR] Convert failed from jpg to png.")
		} else {
			fmt.Println("Convert Success from " + myImageFile.fromFilepath + " to " + myImageFile.toFilepath)
		}
	}
}

// ConvertImagesFromPngToJpgInDir converts all png images to jpg in target directory.
func ConvertImagesFromPngToJpgInDir(dir string) {
	for _, path := range GetFilepathListByExt(dir, ".png") {
		myImageFile := MyImageFile{fileExt: "png", fromFilepath: path, toFilepath: ""}
		if err := ConvertPngToJpg(&myImageFile); err != nil && myImageFile.fileExt == "jpg" {
			fmt.Println("[ERROR] Convert failed from png to jpg.")
		} else {
			fmt.Println("Convert Success from " + myImageFile.fromFilepath + " to " + myImageFile.toFilepath)
		}
	}
}

// GetFilepathListByExt returns the filepath list by the extention in target directory.
func GetFilepathListByExt(dir string, ext string) []string {
	filePathList := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			filePathList = append(filePathList, path)
		}
		return nil
	})

	if err != nil {
		return nil
	}

	return filePathList
}

// ConvertJpgToPng convert target jpg image to png.
func ConvertJpgToPng(imageFile *MyImageFile) error {
	file, err := os.Open(imageFile.fromFilepath)
	if err != nil {
		fmt.Println("[ERROR] No Such File.")
		return err
	}
	defer file.Close()

	outputFilePath := "outputs/" + GetFileNameWithoutExt(imageFile.fromFilepath) + ".png"

	img, _, err := image.Decode(file)

	dstfile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("[ERROR] No Such FilePath.")
		return err
	}
	defer dstfile.Close()

	err = jpeg.Encode(dstfile, img, nil)
	if err != nil {
		fmt.Println("[ERROR] File Encode Failed.")
		return err
	}

	imageFile.fileExt = "png"
	imageFile.toFilepath = outputFilePath

	return nil
}

// ConvertPngToJpg convert target png image to jpg.
func ConvertPngToJpg(imageFile *MyImageFile) error {
	file, err := os.Open(imageFile.fromFilepath)
	if err != nil {
		fmt.Println("[ERROR] No Such File.")
		return err
	}
	defer file.Close()

	outputFilePath := "outputs/" + GetFileNameWithoutExt(imageFile.fromFilepath) + ".jpg"

	img, _, err := image.Decode(file)

	dstfile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("[ERROR] No Such FilePath.")
		return err
	}
	defer dstfile.Close()

	err = png.Encode(dstfile, img)
	if err != nil {
		fmt.Println("[ERROR] File Encode Failed.")
		return err
	}

	imageFile.fileExt = "jpg"
	imageFile.toFilepath = outputFilePath

	return nil
}

// GetFileNameWithoutExt returns the filename without file extension
func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
