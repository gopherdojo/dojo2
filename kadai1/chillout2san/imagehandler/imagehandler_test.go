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

	wrongFileName := dir + "/testpicture1.jpg"

	test.Run("ディレクトリが正しい場合", func(t *testing.T) {
		testPicture, _ := os.Open(correctFileName)
		defer testPicture.Close()
		expect, _, _ := image.Decode(testPicture)
		actual := imagehandler.Decode(wrongFileName)

		if !reflect.DeepEqual(expect, actual) {
			test.Errorf("テストは失敗に終わった")
		}
	})
}