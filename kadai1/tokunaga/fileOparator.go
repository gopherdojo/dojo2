package tokunaga

import (
	"path"
)

// 完全パスから拡張子を除いた文字列を返す 例) /home/hoge/test.png -> /home/hoge/test
func FullBasename(fullFIlePath string) string {
	return fullFIlePath[0 : len(fullFIlePath)-len(path.Ext(fullFIlePath))]
}
