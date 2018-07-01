package tokunaga

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

type File struct {
	Uri      string
	FileSize int64
}

// ダウンロードファイルのUriからファイル名を取得
func (f File) Filename() string {
	return path.Base(f.Uri)
}

func (f File) Download(acceptRanges string) error {
	if acceptRanges == "" {
		return f.singleDownload()
	} else {
		return f.splitDownload()
	}
}

// ファイルを一括ダウンロード
func (f File) singleDownload() error {
	responseGet, err := http.Get(f.Uri)
	if err != nil {
		return err
	}
	defer responseGet.Body.Close()

	file, err := os.Create(f.Filename())
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, responseGet.Body)
	if err != nil {
		return err
	}
	return nil
}

func (f File) splitDownload() error {
	splitNum := runtime.NumCPU()
	splitBytes := f.SplitByteSize(int64(splitNum))
	ranges := formatRange(splitBytes)
	createFileMap := map[int]string{}
	for no, rangeValue := range ranges {
		f.rangeDownload(no, rangeValue, createFileMap)
	}
	if err := f.joinSplitFiles(createFileMap); err != nil {
		return err
	}
	return nil
}

func (f File) rangeDownload(fileNo int, rangeValue string, createFiles map[int]string) error {
	req, _ := http.NewRequest("GET", f.Uri, nil)
	req.Header.Set("RANGE", rangeValue)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	createFileName := fmt.Sprintf("%s_%s", f.Filename(), rangeValue)
	file, err := os.Create(createFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	createFiles[fileNo] = createFileName
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (f File) joinSplitFiles(createFiles map[int]string) error {
	originFile, err := os.Create(f.Filename())
	if err != nil {
		return err
	}
	defer originFile.Close()
	for i := 0; i < len(createFiles); i++ {
		splitFile, err := os.Open(createFiles[i])
		if err != nil {
			return err
		}
		_, err = io.Copy(originFile, splitFile)
		if err != nil {
			return err
		}
		splitFile.Close()
		if err := os.Remove(createFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

// ファイルのバイト数と分割数から、分割ダウンロードするファイルの各バイト数の配列を返す
// SplitByteSize(1002, 8) -> [125, 125, 125, 125, 125, 125, 125, 127]
func (f File) SplitByteSize(splitNum int64) []int64 {
	var response = make([]int64, splitNum)
	rest := f.FileSize % splitNum               // ファイルのサイズを分割数で割った余り
	splitUnit := (f.FileSize - rest) / splitNum // 分割したファイルのサイズ
	for i := int64(0); i < splitNum-1; i++ {
		response[i] = splitUnit
	}
	response[splitNum-1] = splitUnit + rest
	return response
}
