package tokunaga

import "fmt"

// 分割ダウンロードするファイルの各バイト数の配列から、request headerのRANGEに設定する値の配列を返す
// [25 25 25 28]
// -> [bytes=0-25 bytes=26-51 bytes=52-77 bytes=78-103]
func formatRange(splitBytes []int64) []string {
	var response []string
	var bytePosition int64
	for _, bytes := range splitBytes {
		response = append(response, fmt.Sprintf("bytes=%d-%d", bytePosition, bytePosition+bytes-1))
		bytePosition = bytePosition + bytes
	}
	return response
}
