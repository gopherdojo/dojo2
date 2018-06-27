package convert

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
)

type Args struct {
	dir    string
	srcFmt string
	dstFmt string
}

func TestConvert(t *testing.T) {

	argsList := []Args{
		{"../testdata/test", "jpg", "png"},
		{"../testdata/test", "jpeg", "png"},
		{"../testdata/test", "jpeg", "gif"},
		{"../testdata/test2", "png", "gif"},
		{"../testdata/test2", "png", "jpg"},
		{"../testdata/test3", "gif", "jpg"},
		{"../testdata/test3", "gif", "png"},
		{"../testdata/test4", "png", "jpg"},
	}

	for _, arg := range argsList {

		t.Run(fmt.Sprintf("convertTest %s to %s", arg.srcFmt, arg.dstFmt), func(t *testing.T) {
			err := Convert(arg.dir, arg.srcFmt, arg.dstFmt)
			if err != nil {
				t.Error(err)
			}
			removeTestOutput("./output")
		})
	}
}

func TestInvalidConvert(t *testing.T) {

	failArgsList := []Args{
		{"../testdata/test", "jpbg", "png"},
		{"../testdata/test", "jpg", "pdf"},
		{"../testdata/test2", "pag", "jpg"},
		{"../testdata/test3", "gof", "jpg"},
		{"../testdata/test2", "png", "gof"},
	}

	for _, arg := range failArgsList {

		t.Run(fmt.Sprintf("convertFailTest %s to %s", arg.srcFmt, arg.dstFmt), func(t *testing.T) {
			err := Convert(arg.dir, arg.srcFmt, arg.dstFmt)
			if err == nil {
				t.Error("Invalid Convert Result")
			}
			removeTestOutput("./output")
		})
	}
}

func TestFail_isFileOrDirExists(t *testing.T) {

	t.Run("NG isFileOrDirExists", func(t *testing.T) {
		result := isFileOrDirExists("../notfound/1.png")
		if result {
			t.Error("Invalid Result")
		}
	})
}

func TestFail_dirwalk(t *testing.T) {

	t.Run("NG dirwalk", func(t *testing.T) {
		_, err := dirwalk("../notfound", "png", "jpg")
		expectedError := errors.New("ioutil.ReadDir() with ../notfound: open ../notfound: no such file or directory")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})
}

func TestFail_convert(t *testing.T) {
	t.Run("NG convert", func(t *testing.T) {
		err := convert(FileInfo{name: "999", srcFmt: "png", dstFmt: "jpg"})
		expectedError := errors.New("os.Open() with 999.png: open 999.png: no such file or directory")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})
}

func TestFail_encode(t *testing.T) {
	t.Run("NG encode", func(t *testing.T) {

		f := testTempFile(t)
		img := testTempImage(t)
		err := encode("pdf", f, img)
		expectedError := errors.New("不正な画像形式を出力先に指定しています")
		if err.Error() != expectedError.Error() {
			t.Error(err)
		}
	})
}

func TestFail_decode(t *testing.T) {
	t.Run("NG decode", func(t *testing.T) {
		f, error := os.Open("../testdata/test4/1.txt")
		if error != nil {
			t.Error(error)
		}
		_, error = decode(f)
		if error != nil {
			t.Error(error)
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

func testTempImage(t *testing.T) image.Image {
	t.Helper()

	file, error := os.Open("../testdata/test/2.jpg")
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

func removeTestOutput(dir string) {
	os.RemoveAll(dir)
}
