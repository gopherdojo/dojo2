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
	// TODO: io.Readerを作る場所をテストヘルパーに切り出す
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

	dc := format.JPEG{Options: jpeg.Options{Quality: 100}}
	_, err = dc.Decode(jpegFile)
	// FIXME: `unexpected EOF`が出るがどうすればいいか
	// if err != nil {
	// 	t.Errorf(`Unexpected error happen. %s`, err)
	// }
	// FIXME: image.Imageのテストとは
}

// TODO: TestPNG_Encode
