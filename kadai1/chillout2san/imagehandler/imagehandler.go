package imagehandler

import (
	errorhandle "dojo/kadai1/chillout2san/errorhandler"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

/*
	画像オブジェクトを返却します。
*/
func Decode(targetPath string) image.Image {
	file, err := os.Open(targetPath)
	errorhandle.AlertError(err)
	defer file.Close()

	image, _, err := image.Decode(file)
	errorhandle.AlertError(err)

	return image
}

/*
	画像を出力します。
*/
func Encode(image image.Image, beforeExtension string, afterExtension string, targetPath string) {
	dir, fileName := filepath.Split(targetPath)
	newFileName := strings.Replace(fileName, "."+beforeExtension, "."+afterExtension, 1)
	newPath := filepath.Join(dir, newFileName)

	outPath, err := os.Create(newPath)
	errorhandle.AlertError(err)
	defer outPath.Close()

	if afterExtension == "jpg" {
		err := jpeg.Encode(outPath, image, nil)
		errorhandle.AlertError(err)
	}

	if afterExtension == "png" {
		err := png.Encode(outPath, image)
		errorhandle.AlertError(err)
	}

	err = os.Remove(targetPath)
	errorhandle.AlertError(err)
}
