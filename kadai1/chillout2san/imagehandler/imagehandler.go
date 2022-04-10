package imagehandler

import (
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
func Decode(targetPath string) (image.Image, error) {
	file, err := os.Open(targetPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return image, nil
}

/*
	画像を出力します。
*/
func Encode(image image.Image, beforeExtension string, afterExtension string, targetPath string)error {
	dir, fileName := filepath.Split(targetPath)
	newFileName := strings.Replace(fileName, "."+beforeExtension, "."+afterExtension, 1)
	newPath := filepath.Join(dir, newFileName)

	outPath, err := os.Create(newPath)
	if err != nil {
		return nil
	}
	defer outPath.Close()

	if afterExtension == "jpg" {
		err := jpeg.Encode(outPath, image, nil)
		if err != nil {
			return err
		}
	}

	if afterExtension == "png" {
		err := png.Encode(outPath, image)
		if err != nil {
			return err
		}
	}

	err = os.Remove(targetPath)
	if err != nil {
		return err
	}

	return nil
}
