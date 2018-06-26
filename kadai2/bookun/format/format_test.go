package format

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"reflect"
	"testing"
)

var (
	formatOptions = []struct {
		src string
		dst string
	}{
		{"jpg", "jpeg"}, {"jpg", "png"},
		{"jpeg", "jpg"}, {"jpeg", "png"},
		{"png", "jpg"}, {"png", "jpeg"},
	}
	formatOptionsNG = []struct {
		src string
		dst string
	}{
		{"hoge", "fuga"}, {"hoge", "jpg"},
		{"png", "hoge"}, {"", ""},
	}
)

func TestNewFormatOK(t *testing.T) {
	var expectedDecoderType reflect.Type
	var expectedEncoderType reflect.Type
	for _, formatOption := range formatOptions {
		switch formatOption.src {
		case "jpg", "jpeg":
			expectedDecoderType = reflect.ValueOf(&JPEG{}).Type()
		case "png":
			expectedDecoderType = reflect.ValueOf(&PNG{}).Type()
		}
		switch formatOption.dst {
		case "jpg", "jpeg":
			expectedEncoderType = reflect.ValueOf(&JPEG{}).Type()
		case "png":
			expectedEncoderType = reflect.ValueOf(&PNG{}).Type()
		}
		format, err := NewFormat(formatOption.src, formatOption.dst)
		if err != nil {
			fmt.Println(err)
		}
		resultDecoderType := reflect.ValueOf(format.Decoder).Type()
		resultEncoderType := reflect.ValueOf(format.Encoder).Type()
		if expectedDecoderType != resultDecoderType {
			t.Errorf("expected : %s\n", expectedDecoderType)
			t.Errorf("result : %s\n", resultDecoderType)
		}
		if expectedEncoderType != resultEncoderType {
			t.Errorf("expected : %s\n", expectedEncoderType)
			t.Errorf("result : %s\n", resultEncoderType)
		}

	}
}

func TestNewFormatNG(t *testing.T) {
	for _, formatOption := range formatOptionsNG {
		format, err := NewFormat(formatOption.src, formatOption.dst)
		expecedError := fmt.Errorf("Supported formats are only jpeg, jpg or png")
		if format != nil {
			t.Error(format)
		}
		if err.Error() != expecedError.Error() {
			t.Error(expecedError)
			t.Error(err)
		}
	}
}

func TestJpegDecode(t *testing.T) {
	format, _ := NewFormat("jpeg", "png")
	t.Run("OK", func(t *testing.T) {
		jpegImg, _ := os.Open("test_images/sample1.jpg")
		_, err := format.Decoder.Decode(jpegImg)
		if err != nil {
			t.Error(err)
		}
		jpegImg.Close()
	})
	t.Run("NG", func(t *testing.T) {
		pngImg, _ := os.Open("test_images/sample1.png")
		_, err := format.Decoder.Decode(pngImg)
		if err == nil {
			t.Error("expected: JPEG Decoder can not treat PNG image")
		}
		pngImg.Close()
	})
}

func TestJpegEncode(t *testing.T) {
	t.Helper()
	format, _ := NewFormat("png", "jpg")
	sampleDstFile, _ := os.Create("test_images/sampleDstOK.jpeg")
	sampleSrcFile, _ := os.Open("test_images/sample1.png")
	sampleSrcImage, _ := png.Decode(sampleSrcFile)
	err := format.Encoder.Encode(sampleDstFile, sampleSrcImage)
	sampleDstFile.Close()
	if err != nil {
		t.Error(err)
	}
	sampleDstFile, _ = os.Open("test_images/sampleDstOK.jpeg")
	_, formatStr, _ := image.Decode(sampleDstFile)
	if formatStr != "jpeg" {
		t.Error(formatStr)
	}
}

func TestPngDecode(t *testing.T) {
	format, _ := NewFormat("png", "jpeg")
	t.Run("OK", func(t *testing.T) {
		pngImg, _ := os.Open("test_images/sample1.png")
		_, err := format.Decoder.Decode(pngImg)
		if err != nil {
			t.Error(err)
		}
		pngImg.Close()
	})
	t.Run("NG", func(t *testing.T) {
		jpegImg, _ := os.Open("test_images/sample1.jpeg")
		_, err := format.Decoder.Decode(jpegImg)
		if err == nil {
			t.Error("expected: PNG Decoder can not treat JPEG image")
		}
		jpegImg.Close()
	})
}
func TestPngEncode(t *testing.T) {
	t.Helper()
	format, _ := NewFormat("jpg", "png")
	sampleDstFile, _ := os.Create("test_images/sampleDstOK.png")
	sampleSrcFile, _ := os.Open("test_images/sample1.jpg")
	sampleSrcImage, _ := jpeg.Decode(sampleSrcFile)
	err := format.Encoder.Encode(sampleDstFile, sampleSrcImage)
	sampleDstFile.Close()
	if err != nil {
		t.Error(err)
	}
	sampleDstFile, _ = os.Open("test_images/sampleDstOK.png")
	_, formatStr, _ := image.Decode(sampleDstFile)
	if formatStr != "png" {
		t.Error(formatStr)
	}
}
