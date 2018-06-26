package imageconverter

import (
	"fmt"
	"os"
	"path/filepath"
)

// Searcher 対象ディレクトリファイルの検索器
type Searcher struct{}

// Run 対象ディレクトリを再帰的に走査
func (Searcher) Run(target FileInfo) []FileInfo {
	var fis []FileInfo
	err := filepath.Walk(string(target.Path), func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fis = append(fis, FileInfo{Path: FilePath(path)})
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "ディレクトリ検索の途中でエラーが発生しました")
		return fis
	}
	return fis
}
