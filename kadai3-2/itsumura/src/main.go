package main

import (
	"fmt"
	"net/http"
	//"net/url"
	"os"
	//"time"
	"strconv"
	"strings"
	"io/ioutil"
	"sync"
)

func geturl() (string, int, error) {
	var fileurl string
	var length int
	var err error

	for {
		fmt.Println("ダウンロードするファイルのURLを入力してください")
		fmt.Scan(&fileurl)
		
		//tr := http.Transport{}

		//環境変数からプロキシ取得
		proxy := os.Getenv("http_proxy")
		fmt.Println(proxy)

		//parseProxyUrl, _ := url.Parse(proxy)
		//httpClientを作成
		//http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(parseProxyUrl)}
		resp, err := http.Get(fileurl)
		
		/*
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
		tr.Proxy = http.ProxyURL(parseProxyUrl)
		client := &http.Client{Transport:tr}
		req,_ := http.NewRequest("GET", fileurl, nil)
		req.Header.Add("Proxy-Authorization", "")
		resp,err := client.Do(req)
		*/


		fmt.Println(resp)
		length, _ = strconv.Atoi(resp.Header["Content-Length"][0]) //ファイルの長さ取得
		fmt.Println(length)

		//200か404かチェック
		//ステータスコードを表示
		//fmt.Println(resp.StatusCode)
		if err != nil || resp.StatusCode != 200 {
			fmt.Println(err)
			fmt.Println("ファイルにアクセスできません。もう一度入力してください")
			continue
		} else {
			break
		}
	}
	return fileurl, length, err
}

func download(fileurl string, length int, limit int) {
	diff := length % limit // Get the remaining for the last request
	body := make([]string, 11)
	len_sub := length / limit
	//ファイル名
	filemnameIndex :=  len(strings.Split(fileurl, "/"))//  strings.LastIndex(fileurl, "/")
	fmt.Println(filemnameIndex)
	filename := strings.Split(fileurl, "/")[filemnameIndex - 1]
	fmt.Println(filename)

	var wg sync.WaitGroup
	for i := 0; i < limit ; i++ {
        wg.Add(1)

        min := len_sub * i // Min range
        max := len_sub * (i + 1) // Max range

        if (i == limit - 1) {
            max += diff // Add the remaining bytes in the last request
        }

        go func(min int, max int, i int) {
            client := &http.Client {}
            req, _ := http.NewRequest("GET", "fileurl", nil)  
            range_header := "bytes=" + strconv.Itoa(min) +"-" + strconv.Itoa(max-1) // Add the data for the Range header of the form "bytes=0-100"
            req.Header.Add("Range", range_header)
            resp,_ := client.Do(req)
            defer resp.Body.Close()
            reader, _ := ioutil.ReadAll(resp.Body)
            body[i] = string(reader)
            ioutil.WriteFile(strconv.Itoa(i), []byte(string(body[i])), 0x777) // Write to the file i as a byte array
            wg.Done()
            //ioutil.WriteFile(filename, []byte(string(body)), 0x777)
        }(min, max, i)
    }
    wg.Wait()
}

func main() {
	//URLを入力させる
	fileurl, length, err := geturl()
	fmt.Println(fileurl)
	if err != nil {
		panic(err)
	}

	//ダウンロード
	limit := 10 //ゴールチンの数
	download(fileurl, length, limit)
}

/*

		client := &http.Client{Transport:&tr}
		//req,_ := http.NewRequest("GET", fileurl, nil)
		//resp, err := http.Get(fileurl)
		//req.Header.Add("Proxy-Authorization", "")
		//resp,err := client.Do(req) //実行


*/