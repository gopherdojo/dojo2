package imageconverter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Searcher 対象ディレクトリファイルの検索器
type Searcher struct{}

// Run 対象ディレクトリを再帰的に走査
func (s *Searcher) Run(target FileInfo) []FileInfo {
	return s.recursiveSearch(target)
}

func (s *Searcher) recursiveSearch(target FileInfo) []FileInfo {
	var fis []FileInfo
	files, err := ioutil.ReadDir(string(target.Path))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイルが開けません")
		return fis
	}
	for _, file := range files {
		filePath := FilePath(filepath.Join(string(target.Path), file.Name()))
		fi := FileInfo{Path: filePath}
		if file.IsDir() {
			fis = append(fis, s.recursiveSearch(fi)...)
		} else {
			fis = append(fis, fi)
		}
	}
	return fis
}
