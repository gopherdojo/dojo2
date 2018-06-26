package test

import (
	"os"
	"path"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/tokunaga"
)

// 引数の文字列の拡張子を表す, stringを基にしたユーザー定義型を返す
func TestAdaptExt(t *testing.T) {
	cases := []struct {
		input    string
		expected tokunaga.DecodeEncoder
	}{
		{input: "jpeg", expected: tokunaga.JpegWrapper("jpeg")},
		{input: "jpg", expected: tokunaga.JpegWrapper("jpg")},
		{input: "png", expected: tokunaga.PngWrapper("png")},
		{input: "", expected: tokunaga.PngWrapper("")},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			if actual := tokunaga.AdaptExt(c.input); actual != c.expected {
				t.Errorf("want AdaptExt(%s) = %s, got %s", c.input, c.expected, actual)
			}
		})
	}

}

// 画像変換テスト
func TestConvertImage(t *testing.T) {
	t.Helper()
	testdataDir := os.Getenv("GOPATH") + "/src/github.com/gopherdojo/dojo2/kadai2/test/testdata/"
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{name: "png to jepg", input: testdataDir + "image_shallow_A.png", output: testdataDir + "image_shallow_A.jpeg"},
		{name: "jepg to png", input: testdataDir + "image_shallow_B.jpeg", output: testdataDir + "image_shallow_B.png"},
	}
	var files []string
	for _, c := range cases {
		files = append(files, c.output)
	}

	TearDown := SetupTestConvertImage(t, files)
	defer TearDown()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fromExt, toExt := tokunaga.AdaptExt(path.Ext(c.input)[1:]), tokunaga.AdaptExt(path.Ext(c.output)[1:])
			if err := tokunaga.ConvertImage(c.input, fromExt, toExt); err != nil {
				t.Errorf("%s can't convart %s", c.input, c.name)
			}
		})
		t.Run(c.name, func(t *testing.T) {
			if _, err := os.Stat(c.output); err != nil {
				t.Errorf("%s is not exist", c.output)
			}
		})
	}
}

// 画像変換テストの前後で変換後のファイルを削除する
func SetupTestConvertImage(t *testing.T, files []string) func() {
	deleteFiles(t, files)
	return func() {
		deleteFiles(t, files)
	}
}

func deleteFiles(t *testing.T, files []string) {
	for _, file := range files {
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			if err := os.Remove(file); err != nil {
				t.Fatal(err)
			}
		}
	}
}
