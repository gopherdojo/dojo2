package imagehandler_test

import (
	"dojo/kadai1/chillout2san/imagehandler"
	"image"
	"os"
	"reflect"
	"testing"
)

func TestDecode(test *testing.T) {
	dir, _ := os.Getwd()

	correctFileName := dir + "/testpicture.jpg"

	wrongFileName := dir + "/testpicture.png"

	test.Run("ディレクトリが正しい場合", func(t *testing.T) {
		testPicture, _ := os.Open(correctFileName)
		defer testPicture.Close()
		expect, _, _ := image.Decode(testPicture)
		actual, _ := imagehandler.Decode(correctFileName)

		if !reflect.DeepEqual(expect, actual) {
			test.Errorf("エラーが発生しました。")
		}
	})

	test.Run("ディレクトリが誤っている場合", func(t *testing.T) {
		_, err := imagehandler.Decode(wrongFileName)

		if err == nil {
			test.Errorf("ディレクトリが誤っているのにテストがパスしている。")
		}
	})
}