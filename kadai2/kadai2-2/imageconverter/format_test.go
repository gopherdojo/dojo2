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

func TestFormat_NormalizedFormat(t *testing.T) {
	fJpeg := imageconverter.Format("jpeg")
	resultJpeg := fJpeg.NormalizedFormat()
	expectedJpeg := imageconverter.Format("jpg")
	if resultJpeg != expectedJpeg {
		t.Errorf("NormalizedFormat failed.  expect:%s, actual:%s", expectedJpeg, resultJpeg)
	}

	fJpg := imageconverter.Format("jpg")
	resultJpg := fJpg.NormalizedFormat()
	expectedJpg := imageconverter.Format("jpg")
	if resultJpg != expectedJpg {
		t.Errorf("NormalizedFormat failed.  expect:%s, actual:%s", expectedJpeg, resultJpeg)
	}

	fPng := imageconverter.Format("png")
	resultPng := fPng.NormalizedFormat()
	expectedPng := imageconverter.Format("png")
	if resultPng != expectedPng {
		t.Errorf("NormalizedFormat failed.  expect:%s, actual:%s", expectedPng, resultPng)
	}
}
