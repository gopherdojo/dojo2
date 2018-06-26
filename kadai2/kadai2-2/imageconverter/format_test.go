package imageconverter_test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

func TestFormat_Ext(t *testing.T) {
	f := imageconverter.Format("hoge")
	result := f.Ext()
	expected := ".hoge"
	if result != expected {
		t.Errorf("Format.Ext failed.  expect:%s, actual:%s", expected, result)
	}
}

var cases_Format_NormalizedFormat = []struct {
	format   imageconverter.Format
	expected imageconverter.Format
}{
	{imageconverter.Format("jpeg"), imageconverter.Format("jpg")},
	{imageconverter.Format("jpg"), imageconverter.Format("jpg")},
	{imageconverter.Format("png"), imageconverter.Format("png")},
}

func TestFormat_NormalizedFormat(t *testing.T) {
	for _, c := range cases_Format_NormalizedFormat {
		result := c.format.NormalizedFormat()
		if result != c.expected {
			t.Errorf("NormalizedFormat failed.  expect:%s, actual:%s", c.expected, result)
		}
	}
}
