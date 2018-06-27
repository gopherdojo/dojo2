package convert

import (
	"image"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestConvert(t *testing.T) {
	t.Run("jpg to png", func(t *testing.T) {
		err := Convert("../test", "jpg", "png")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("png to jpg", func(t *testing.T) {
		err := Convert("../test2", "png", "jpg")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("gif to jpg", func(t *testing.T) {
		err := Convert("../test3", "gif", "jpg")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("NG isFileOrDirExists", func(t *testing.T) {
		err := Convert("../notfound", "png", "jpg")
		expectedError := errors.New("ディレクトリは存在しません")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})

	t.Run("NG dirwalk", func(t *testing.T) {
		_, err := dirwalk("../notfound", "png", "jpg")
		expectedError := errors.New("ioutil.ReadDir() with ../notfound: open ../notfound: no such file or directory")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})

	t.Run("NG convertImage srcFmt", func(t *testing.T) {
		err := convertImage(FileInformation{name: "1", srcFmt: "pdf", dstFmt: "png"})
		expectedError := errors.New("指定した変換元の画像形式は無効です")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})

	t.Run("NG convertImage dstFmt", func(t *testing.T) {
		err := convertImage(FileInformation{name: "1", srcFmt: "png", dstFmt: "pdf"})
		expectedError := errors.New("指定した変換先の画像形式は無効です")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})

	t.Run("NG startConvert", func(t *testing.T) {
		err := startConvert(FileInformation{name: "999", srcFmt: "png", dstFmt: "jpg"})
		expectedError := errors.New("os.Open() with 999.png: open 999.png: no such file or directory")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})

	t.Run("NG isAvailableFormat", func(t *testing.T) {
		result := isAvailableFormat("pdf")
		if result != false {
			t.Error("wrong result")
		}
	})

	t.Run("NG encode", func(t *testing.T) {

		f := testTempFile(t)
		img := testImage(t)
		err := encode("pdf", f, img)
		expectedError := errors.New("不正な画像形式を出力先に指定しています")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})
}

func testTempFile(t *testing.T) *os.File {
	t.Helper()
	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
	}
	tf.Close()
	return tf
}

func testImage(t *testing.T) image.Image {
	t.Helper()

	file, error := os.Open("1.jpg")
	if error != nil {
		t.Error(error)
	}
	defer file.Close()

	img, _, error := image.Decode(file)
	if error != nil {
		t.Error(error)
	}
	return img
}
