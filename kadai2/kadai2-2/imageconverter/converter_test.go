package imageconverter_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

func TestConverter_Run(t *testing.T) {
	var c imageconverter.Converter
	fi := imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/Octocat.jpeg")}
	c.Run(fi, imageconverter.Format("jpg"), imageconverter.Format("png"))
	generatedFilePath := imageconverter.FilePath("../sample_dir1/Octocat.png")
	result := fileExists(generatedFilePath)
	expected := true
	if result != expected {
		t.Errorf("Converter.Run failed.")
	}
	fileClear(generatedFilePath)
}

func fileExists(path imageconverter.FilePath) bool {
	_, err := os.Stat(string(path))
	return !os.IsNotExist(err)
}

func fileClear(path imageconverter.FilePath) {
	if err := os.Remove(string(path)); err != nil {
		fmt.Println(err)
	}
}
