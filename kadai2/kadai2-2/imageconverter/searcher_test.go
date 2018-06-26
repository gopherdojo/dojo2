package imageconverter_test

import (
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

func TestSearcher_Run(t *testing.T) {
	var searcher imageconverter.Searcher
	result := searcher.Run(imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1")})
	expected := []imageconverter.FileInfo{
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/Octocat.jpeg")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/dummy_fuga.md")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/dummy_hoge.txt")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/sample_dir2/Octocat.jpg")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/sample_dir2/dummy_fuga.md")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/sample_dir2/dummy_hoge.txt")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/sample_dir2/sample_dir3/Octocat.png")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/sample_dir2/sample_dir3/dummy_fuga.md")},
		imageconverter.FileInfo{Path: imageconverter.FilePath("../sample_dir1/sample_dir2/sample_dir3/dummy_hoge.txt")},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Run(../sample_dir1) failed.  expect:%s, actual:%s", expected, result)
	}

}
