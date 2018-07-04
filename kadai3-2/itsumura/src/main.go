package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func geturl() (string, int, error) {
	var url string
	var length int
	var err error

	for {
		fmt.Println("ダウンロードするファイルのURLを入力してください")
		fmt.Scan(&url)

		//200か404かチェック
		resp, err := http.Get(url)
		//fmt.Println(resp.Header)
		length, _ = strconv.Atoi(resp.Header["Content-Length"][0]) //ファイルの長さ取得
		//ステータスコードを表示
		fmt.Println(resp.StatusCode)
		if err != nil || resp.StatusCode != 200 {
			fmt.Println("ファイルにアクセスできません。もう一度入力してください")
			defer resp.Body.Close()
			continue
		} else {
			defer resp.Body.Close()
			break
		}
	}
	return url, length, err
}

func download(url string, length int, limit int) {

}

func main() {
	//URLを入力させる
	url, length, err := geturl()
	fmt.Println(url)
	if err != nil {
		panic(err)
	}

	//ダウンロード
	limit := 10 //ゴールチンの数
	download(url, length, limit)
}
