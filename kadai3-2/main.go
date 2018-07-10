package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context/ctxhttp"
	"strconv"
)

const(
	tmpdir = "/tmp/kaznishi_pdownload"
)

func main() {
	url := "http://abehiroshi.la.coocan.jp/abe-top2-4.jpg"
	dirname := "/tmp/pdownload"
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	fmt.Println(url)
	res, _ := ctxhttp.Head(ctx, http.DefaultClient, url)
	fmt.Println(res.Header.Get("Accept-Ranges"))
	fmt.Println(res.Header.Get("Content-Length"))
	os.MkdirAll(dirname, 0755)
	// req, _ := http.NewRequest("GET", url, nil)
	// low := 0
	// high := 65172
	// req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low, high))
	// res, _ = http.DefaultClient.Do(req)

	// ファイルを作って画像を保存する
	// file, _ := os.OpenFile(dirname+"/abehiroshi.jpg", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// defer file.Close()
	// io.Copy(file, res.Body)

	// 2分割でファイルをダウンロードして、それぞれダウンロードし終わった後にマージファイルを作る
	req1, _ := http.NewRequest("GET", url, nil)
	low1 := 0
	high1 := 39999
	req1.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low1, high1))
	res1, _ := http.DefaultClient.Do(req1)
	file1, _ := os.OpenFile(dirname+"/abehiroshi.jpg-p1", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	io.Copy(file1, res1.Body)
	file1.Close()

	req2, _ := http.NewRequest("GET", url, nil)
	low2 := 40000
	high2 := 65172
	req2.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low2, high2))
	res2, _ := http.DefaultClient.Do(req2)
	file2, _ := os.OpenFile(dirname+"/abehiroshi.jpg-p2", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	io.Copy(file2, res2.Body)
	file2.Close()

	file1, _ = os.Open(dirname + "/abehiroshi.jpg-p1")
	file2, _ = os.Open(dirname + "/abehiroshi.jpg-p2")

	file, _ := os.OpenFile(dirname+"/abehiroshi.jpg", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	io.Copy(file, file1)
	io.Copy(file, file2)
	file.Close()

	// 提出用に整えていく
	//// チェック処理
	//// ファイルサイズ分割処理
	//// 分割ファイル名
	//// 分割ダウンロード処理
	//// 分割ファイルマージ処理
	//// 分割ファイルクリア処理
	//// goroutine使って書き換え
	//// キャンセル時処理書く
	//// テスト書く

	//

	return
}

func prepare () {
	os.MkdirAll(tmpdir, 0755)
}

func sizeCheck(url string) (int, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	res, err := ctxhttp.Head(ctx, http.DefaultClient, url)
	if err != nil {
		return 0, err
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, err // ここは自分でエラーを定義するように直す
	}

	l, err := strconv.Atoi(res.Header.Get("Content-Length"))
	return l, err
}


type part struct{
	Low int
	High int
	Filename string
}

func split (pcount int, fullsize int, filename string) [...]part {
	var result [pcount]part

	var low, high int
	for i := 0; i < pcount; i++ {
		if i == 0 {
			low = 0
		} else {
			low = high + 1
		}
		if i == pcount - 1 {
			high = fullsize
		} else {
			high = int(fullsize * (i+1) / pcount)
		}
		fn := filename + "_" + strconv.Itoa(i)
		p := part{Low: low, High: high, Filename: fn}
		result[i] = p
	}
	return result
}

func download(p part, url string) {
	req, _ := http.NewRequest("GET", url, nil)
	low := p.Low
	high := p.High
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low, high))
	res, _ := http.DefaultClient.Do(req)
	file, _ := os.OpenFile(tmpdir+"/"+p.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	io.Copy(file, res.Body)
	file.Close()
	return
}

func merge (parts [...]part, newFilePath string) {
	newFile, _ := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	for _, p := range parts {
		pf, _ := os.Open(tmpdir + "/" + p.Filename)
		io.Copy(newFile, pf)
		pf.Close()
	}
	newFile.Close()
}

func clearPartFiles (parts [...]part) {
	for _, p := range parts {
		os.Remove(tmpdir + "/" + p.Filename)
	}
}