package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
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
	responseHead, err := http.Head(downloadPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitAbort)
	}
	statusCode := responseHead.StatusCode
	fmt.Printf("%d %s\n", statusCode, http.StatusText(statusCode))
	if responseHead.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, http.StatusText(responseHead.StatusCode))
		os.Exit(ExitAbort)
	}
	fmt.Printf("長さ: %d [%s]\n", responseHead.ContentLength, responseHead.Header.Get("Content-Type"))
	fmt.Printf("`%s' に保存中\n", path.Base(downloadPath))
	fmt.Println("")
	responseGet, err := http.Get(downloadPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitAbort)
	}
	defer responseGet.Body.Close()

	file, err := os.Create(path.Base(downloadPath))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitAbort)
	}
	defer file.Close()
	io.Copy(file, responseGet.Body)
}
