package imageconverter_test

import (
	"os"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/kadai2-2/imageconverter"
)

func TestConverter_Run(t *testing.T) {
	t.Run("jpg to png", func(t *testing.T) {
		testConverter_Run(t,
			imageconverter.FilePath("../sample_dir1/Octocat.jpeg"),
			imageconverter.Format("jpg"),
			imageconverter.Format("png"),
			imageconverter.FilePath("../sample_dir1/Octocat.png"),
			true,
		)
		testConverter_Run(t,
			imageconverter.FilePath("../sample_dir1/dummy_fuga.md"),
			imageconverter.Format("jpg"),
			imageconverter.Format("png"),
			imageconverter.FilePath("../sample_dir1/dummy_fuga.png"),
			false,
		)
		testConverter_Run(t,
			imageconverter.FilePath("../sample_dir1/sample_dir2/Octocat.jpg"),
			imageconverter.Format("jpg"),
			imageconverter.Format("png"),
			imageconverter.FilePath("../sample_dir1/sample_dir2/Octocat.png"),
			true,
		)
	})
	t.Run("png to jpg", func(t *testing.T) {
		testConverter_Run(t,
			imageconverter.FilePath("../sample_dir1/sample_dir2/sample_dir3/Octocat.png"),
			imageconverter.Format("png"),
			imageconverter.Format("jpg"),
			imageconverter.FilePath("../sample_dir1/sample_dir2/sample_dir3/Octocat.jpg"),
			true,
		)
		testConverter_Run(t,
			imageconverter.FilePath("../sample_dir1/dummy_fuga.md"),
			imageconverter.Format("png"),
			imageconverter.Format("jpg"),
			imageconverter.FilePath("../sample_dir1/dummy_fuga.jpg"),
			false,
		)
	})
}

func testConverter_Run(t *testing.T,
	inFP imageconverter.FilePath,
	inF imageconverter.Format,
	outF imageconverter.Format,
	expectedOutFP imageconverter.FilePath,
	expectedGenerated bool,
) {
	t.Helper()
	var c imageconverter.Converter
	fi := imageconverter.FileInfo{Path: inFP}
	c.Run(fi, inF, outF)
	result := fileExists(expectedOutFP)
	if result != expectedGenerated {
		t.Errorf("Converter.Run failed.")
	}
	if result {
		fileClear(expectedOutFP)
	}
}

func fileExists(path imageconverter.FilePath) bool {
	_, err := os.Stat(string(path))
	return !os.IsNotExist(err)
}

func fileClear(path imageconverter.FilePath) {
	os.Remove(string(path))
}
