package imageconverter_test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

var cases_CommandValidator_ExtValidate = []struct {
	format   imageconverter.Format
	expected bool
}{
	{imageconverter.Format("jpg"), true},
	{imageconverter.Format("png"), true},
	{imageconverter.Format("hoge"), false},
	{imageconverter.Format("txt"), false},
}

func TestCommandValidator_ExtValidate(t *testing.T) {
	var cv imageconverter.CommandValidator
	for _, c := range cases_CommandValidator_ExtValidate {
		result := cv.ExtValidate(c.format)
		if result != c.expected {
			t.Errorf("ExtValidate failed.")
		}
	}
}
