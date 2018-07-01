package format_test

import (
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/format"
)

func TestJPEG_Decode(t *testing.T) {
	file := helperGetIoReader(t)
	dc := format.JPEG{Options: jpeg.Options{Quality: 100}}
	_, _ = dc.Decode(file)
	// FIXME: `unexpected EOF`が出るがどうすればいいか
	// if err != nil {
	// 	t.Errorf(`Unexpected error happen. %s`, err)
	// }
	// FIXME: image.Imageのテストとは
}

func helperGetIoReader(t *testing.T) *os.File {
	t.Helper()
	jpegImg := image.NewRGBA(image.Rect(0, 0, 100, 200))
	jpegFile, err := ioutil.TempFile("", "jpeg")
	if err != nil {
		panic(err)
	}
	defer jpegFile.Close()
	defer os.Remove(jpegFile.Name())
	if err := jpeg.Encode(jpegFile, jpegImg, nil); err != nil {
		panic(err)
	}
	return jpegFile
}

// TODO: TestPNG_Encode
