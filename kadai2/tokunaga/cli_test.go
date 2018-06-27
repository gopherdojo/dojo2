package tokunaga

import (
	"errors"
	"os"
	"testing"
)

func TestCliDelegateFileOperation(t *testing.T) {
	testdataDirPath := os.Getenv("GOPATH") + "/src/github.com/gopherdojo/dojo2/kadai2/test/testdata/"
	testdirInfo, err := os.Stat(testdataDirPath)
	if err != nil {
		t.Fatal(err)
	}
	testfilePath := os.Getenv("GOPATH") + "/src/github.com/gopherdojo/dojo2/kadai2/test/testdata/image_shallow_A.png"
	testfileInfo, err := os.Stat(testfilePath)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("erro != nil", func(t *testing.T) {
		cli := CLI{}
		input := "abort"
		expected := "abort"
		if actual := cli.delegateFileOperation("", nil, errors.New(input)); actual.Error() != expected {
			t.Errorf("want delegateFileOperation(nil, nil, errors.New(%v)) = %v, got %v", input, expected, actual)
		}
	})
	t.Run("target is dir", func(t *testing.T) {
		cli := CLI{from: "jpeg"}
		if actual := cli.delegateFileOperation(testdataDirPath, testdirInfo, nil); actual != nil {
			t.Errorf("want delegateFileOperation(dirpath, dirfileinfo, nil) = nil, got %v", actual)
		}
	})
	t.Run("target is file, file's ext is not target ext", func(t *testing.T) {
		cli := CLI{from: "jpeg"}
		if actual := cli.delegateFileOperation(testfilePath, testfileInfo, nil); actual != nil {
			t.Errorf("want delegateFileOperation(filepath, fileinfo, nil) = nil, got %v", actual)
		}
	})

}

// 引数の拡張子が許可されているものならばtrue, それ以外なら false を返す
func TestCheckExtPermmited(t *testing.T) {
	cases := []struct {
		name          string
		inputExt      string
		permittedExts []string
		expected      bool
	}{
		{name: "ext is permitted", inputExt: "png", permittedExts: []string{"png", "jpeg"}, expected: true},
		{name: "ext is not permitted", inputExt: "gif", permittedExts: []string{"png", "jpeg"}, expected: false},
		{name: "ext is not permitted", inputExt: "", permittedExts: []string{"png", "jpeg"}, expected: false},
		{name: "ext is not permitted", inputExt: "png", permittedExts: []string{}, expected: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := checkExtPermmited(c.inputExt, c.permittedExts); actual != c.expected {
				t.Errorf("want checkExtPermmited(%s, %v) = %v, got %v", c.inputExt, c.permittedExts, c.expected, actual)
			}
		})
	}
}
