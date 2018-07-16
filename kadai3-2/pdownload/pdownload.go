package pdownload

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

// Option はプログラムに与えるオプションをまとめた構造体です
type Option struct {
	TargetURL string // ダウンロードの対象URL
	PCount    int    // 分割数
	OutputDir string // 結合後のファイルの格納場所
	TmpDir    string // 分割ファイルの一時格納場所
}

// Init は新しく生成したオブジェクトにデフォルト値を設定するための関数です
func (o *Option) Init() {
	o.PCount = 5
	o.OutputDir = "."
	o.TmpDir = "/tmp/kaznishi_pdownload"
}

var (
	tmpDir = "/tmp/kaznishi_pdownload" // 分割ファイルを一時的に格納するディレクトリ
)

// Run はpdownloadの処理を実行します
func Run(ctx context.Context, doneCh chan<- int, option Option) error {
	errCh := make(chan error, 1)
	eg, _ := errgroup.WithContext(ctx)

	// 一時保存ディレクトリの作成
	setTmpDir(option.TmpDir)
	if err := mkTmpDir(); err != nil {
		doneCh <- 1
		return err
	}
	//// チェック処理
	fullSize, err := sizeCheck(option.TargetURL)
	if err != nil {
		doneCh <- 1
		return err
	}
	//// ファイルサイズ分割処理
	fileName := path.Base(option.TargetURL)
	parts := split(option.PCount, fullSize, fileName)

	go func() {
		//// 分割ダウンロード処理
		for _, p := range parts {
			fmt.Println("Downloding Part File Started. :" + p.FileName)
			p := p
			eg.Go(func() error {
				return download(p, option.TargetURL)
			})
		}
		if err := eg.Wait(); err != nil {
			errCh <- err
		} else {
			fmt.Println("Downloading Part Files completed.")
		}

		//// 分割ファイルマージ処理
		if err := merge(parts, getNewFilePath(option.OutputDir, fileName)); err != nil {
			errCh <- err
		} else {
			fmt.Println("Combining Part Files completed.")
		}

		//// 分割ファイルクリア処理
		if err = clearPartFiles(parts); err != nil {
			errCh <- err
		}

		errCh <- nil
	}()

	for {
		select {
		case err := <-errCh:
			if err != nil {
				clearWhenCancel(parts, getNewFilePath(option.OutputDir, fileName))
				doneCh <- 1
				return err
			}
			doneCh <- 0
			return nil
		case <-ctx.Done():
			clearWhenCancel(parts, getNewFilePath(option.OutputDir, fileName))
			doneCh <- 0
			return nil
		}
	}
}

///////////////////////////////////////////////////////////////////////////

func getNewFilePath(outputDir, fileName string) string {
	return outputDir + "/" + fileName
}

func setTmpDir(dirPath string) {
	if dirPath != "" {
		tmpDir = dirPath
	}
}

func mkTmpDir() error {
	return os.MkdirAll(tmpDir, 0755)
}

func sizeCheck(url string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	res, err := ctxhttp.Head(ctx, http.DefaultClient, url)
	if err != nil {
		return 0, err
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		err = fmt.Errorf("Accept-Ranges = bytesではありません")
		return 0, err
	}

	l, err := strconv.Atoi(res.Header.Get("Content-Length"))
	return l, err
}

type part struct {
	Low      int
	High     int
	FileName string
}

func (p part) getFilePath() string {
	return tmpDir + "/" + p.FileName
}

func split(pCount int, fullSize int, fileName string) []part {
	result := make([]part, pCount)

	var low, high int
	for i := 0; i < pCount; i++ {
		if i == 0 {
			low = 0
		} else {
			low = high + 1
		}
		if i == pCount-1 {
			high = fullSize - 1
		} else {
			high = int(fullSize * (i + 1) / pCount)
		}
		fn := fileName + "_" + strconv.Itoa(i)
		p := part{Low: low, High: high, FileName: fn}
		result[i] = p
	}
	return result
}

func download(p part, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	low := p.Low
	high := p.High
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low, high))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	file, err := os.Create(p.getFilePath())
	if err != nil {
		return err
	}

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return file.Close()
}

func merge(parts []part, newFilePath string) error {
	newFile, _ := os.Create(newFilePath)
	for _, p := range parts {
		pf, err := os.Open(p.getFilePath())
		if err != nil {
			return err
		}
		io.Copy(newFile, pf)
		pf.Close()
	}
	newFile.Close()
	return nil
}

func clearPartFiles(parts []part) error {
	for _, p := range parts {
		if err := os.Remove(p.getFilePath()); err != nil {
			return err
		}
	}
	return nil
}

func clearWhenCancel(parts []part, newFilePath string) {
	clearPartFiles(parts)
	os.Remove(newFilePath)
}
