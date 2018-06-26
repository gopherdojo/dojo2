package converter

import (
	"fmt"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/bookun/format"
)

var (
	formatJPEGToPNG       *format.Format
	formatPNGToJPEG       *format.Format
	ConverterJPEGToPNG    *Converter
	ConverterPNGToJPEG    *Converter
	ConverterNGNotFound   *Converter
	ConverterNGPermission *Converter
	ConverterNGDecode     *Converter
)

func before() {
	formatJPEGToPNG, _ = format.NewFormat("jpeg", "png")
	ConverterJPEGToPNG = NewConverter("../test_images/sample1.jpg", "../test_images/sample1.png", *formatJPEGToPNG)
	formatPNGToJPEG, _ = format.NewFormat("png", "jpeg")
	ConverterPNGToJPEG = NewConverter("../test_images/sample1.png", "../test_images/sample1.jpeg", *formatPNGToJPEG)
	ConverterNGNotFound = NewConverter("../test_images/unexist.png", "../test_images/sample1.jpeg", *formatPNGToJPEG)
	ConverterNGPermission = NewConverter("../test_images/sample1.png", "/etc/hogehoge.jpg", *formatPNGToJPEG)
	ConverterNGDecode = NewConverter("../test_images/sample1.png", "/etc/hogehoge.jpg", *formatJPEGToPNG)
}

func Test(t *testing.T) {
	before()
}

func TestConvert(t *testing.T) {
	t.Run("Jpeg to PNG", func(t *testing.T) {
		err := ConverterJPEGToPNG.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("PNG to Jpeg", func(t *testing.T) {
		err := ConverterPNGToJPEG.Convert()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("File not found", func(t *testing.T) {
		err := ConverterNGNotFound.Convert()

		expecedError := fmt.Errorf("open %s: no such file or directory", ConverterNGNotFound.srcFileName)
		if err.Error() != expecedError.Error() {
			t.Error(err.Error())
			t.Error(expecedError.Error())
		}
	})
	t.Run("Permission error", func(t *testing.T) {
		err := ConverterNGPermission.Convert()
		expecedError := fmt.Errorf("open %s: permission denied", ConverterNGPermission.dstFileName)
		if err.Error() != expecedError.Error() {
			t.Error(err.Error())
			t.Error(expecedError.Error())
		}
	})
	t.Run("Decode error", func(t *testing.T) {
		err := ConverterNGDecode.Convert()
		expecedError := fmt.Errorf("invalid JPEG format: missing SOI marker")
		if err.Error() != expecedError.Error() {
			t.Error(err.Error())
			t.Error(expecedError.Error())
		}
	})
}
