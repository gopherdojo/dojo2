package imageconverter_test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

func TestExtValidate(t *testing.T) {
	var cv imageconverter.CommandValidator
	result := cv.ExtValidate(imageconverter.Format("hoge"))
	expected := false
	if result != expected {
		t.Errorf("ExtValidate failed.")
	}
}
