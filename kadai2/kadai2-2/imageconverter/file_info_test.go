package imageconverter_test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

func TestFileInfoExt(t *testing.T) {
	fi := imageconverter.FileInfo{Path: imageconverter.FilePath("/path/to/hoge.txt")}
	result := fi.Ext()
	expected := ".txt"
	if result != expected {
		t.Errorf("FileInfo.Ext failed.  expect:%s, actual:%s", expected, result)
	}
}

func TestFormat(t *testing.T) {
	fi := imageconverter.FileInfo{Path: imageconverter.FilePath("/path/to/hoge.txt")}
	result := fi.Format()
	expected := imageconverter.Format("txt")
	if result != expected {
		t.Errorf("Format failed.  expect:%s, actual:%s", expected, result)
	}
}
