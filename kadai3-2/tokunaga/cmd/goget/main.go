package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gopherdojo/dojo2/kadai3-2/tokunaga"
)

const (
	ExitOK = iota
	ExitAbort
)
const timeLayout = "2006-01-02 15:04:05"

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "You can only download one file at a time")
		os.Exit(ExitAbort)
	}
	downloadPath := flag.Args()[0]

	t := time.Now()
	fmt.Printf("--%s-- %s\n", t.Format(timeLayout), downloadPath)
	fmt.Print("HTTP による接続要求を送信しました、応答を待っています...")

	// head request
	responseHead, err := http.Head(downloadPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitAbort)
	}
	// response info
	acceptRanges := responseHead.Header.Get("Accept-Ranges")
	statusCode := responseHead.StatusCode
	fileSize := responseHead.ContentLength
	contentType := responseHead.Header.Get("Content-Type")

	downloadFile := tokunaga.File{Uri: downloadPath, FileSize: fileSize}
	fmt.Printf("%d %s\n", statusCode, http.StatusText(statusCode))
	if responseHead.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, http.StatusText(responseHead.StatusCode))
		os.Exit(ExitAbort)
	}
	fmt.Printf("長さ: %d [%s]\n", fileSize, contentType)
	fmt.Printf("`%s' に保存中\n", downloadFile.Filename())
	fmt.Println("")

	if err := downloadFile.Download(acceptRanges); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitAbort)
	}

	fmt.Printf("`%s' に保存完了\n", downloadFile.Filename())
	os.Exit(ExitOK)
}
