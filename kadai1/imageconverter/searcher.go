package imageconverter

import (
	"io/ioutil"
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
	files, _ := ioutil.ReadDir(string(target.Path))
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
