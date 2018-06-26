package imageconverter_test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

var (
	calledCountSearcherRun  int
	calledCountConverterRun int
)

type SearcherMock struct{}

func (SearcherMock) Run(target imageconverter.FileInfo) []imageconverter.FileInfo {
	calledCountSearcherRun += 1
	fis := []imageconverter.FileInfo{
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/Octocat.jpeg")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/dummy_fuga.md")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/dummy_hoge.txt")},
	}
	return fis
}

type ConverterMock struct{}

func (ConverterMock) Run(f imageconverter.FileInfo, in, out imageconverter.Format) {
	calledCountConverterRun += 1
	return
}

func TestFacade_Run(t *testing.T) {
	var searcher SearcherMock
	var converter ConverterMock

	targetPath := imageconverter.FilePath("../sample_dir1")
	inputFormat := imageconverter.Format("jpg")
	outputFormat := imageconverter.Format("png")

	facade := imageconverter.Facade{Searcher: searcher, Converter: converter}
	facade.Run(targetPath, inputFormat, outputFormat)

	if calledCountSearcherRun != 1 {
		t.Errorf("Facade Run called count is wrong.")
	}
	if calledCountConverterRun != 3 {
		t.Errorf("Facade Run called count is wrong.")
	}

}
