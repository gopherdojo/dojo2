package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
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

	fmt.Printf("%d %s\n", statusCode, http.StatusText(statusCode))
	if responseHead.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, http.StatusText(responseHead.StatusCode))
		os.Exit(ExitAbort)
	}
	fmt.Printf("長さ: %d [%s]\n", fileSize, contentType)
	fmt.Printf("`%s' に保存中\n", path.Base(downloadPath))
	fmt.Println("")

	if acceptRanges == "" {
		singleDownload(downloadPath)
	} else {
		splitNum := runtime.NumCPU()
		splitBytes := splitByteSize(fileSize, int64(splitNum))
		ranges := formatRange(splitBytes)
		createFileMap := map[int]string{}
		for no, rangeValue := range ranges {
			req, _ := http.NewRequest("GET", downloadPath, nil)
			req.Header.Set("RANGE", rangeValue)

			client := new(http.Client)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(ExitAbort)
			}
			//defer resp.Body.Close()
			createFileName := fmt.Sprintf("%s_%s", path.Base(downloadPath), rangeValue)
			file, err := os.Create(createFileName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(ExitAbort)
			}
			//defer file.Close()
			createFileMap[no] = createFileName
			io.Copy(file, resp.Body)
		}
		originFile, err := os.Create(path.Base(downloadPath))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(ExitAbort)
		}
		defer originFile.Close()
		for i := 0; i < splitNum; i++ {
			splitFile, err := os.Open(createFileMap[i])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(ExitAbort)
			}
			io.Copy(originFile, splitFile)
			splitFile.Close()
			if err := os.Remove(createFileMap[i]); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(ExitAbort)
			}
		}
	}
	fmt.Printf("`%s' に保存完了\n", path.Base(downloadPath))
	os.Exit(ExitOK)
}

func singleDownload(downloadPath string) {
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

func splitByteSize(byteSize int64, splitNum int64) []int64 {
	var response = make([]int64, splitNum)
	rest := byteSize % splitNum               // ファイルのサイズを分割数で割った余り
	splitUnit := (byteSize - rest) / splitNum // 分割したファイルのサイズ
	for i := int64(0); i < splitNum-1; i++ {
		response[i] = splitUnit
	}
	response[splitNum-1] = splitUnit + rest
	return response
}

func formatRange(splitBytes []int64) []string {
	var response []string
	var bytePosition int64
	for _, bytes := range splitBytes {
		response = append(response, fmt.Sprintf("bytes=%d-%d", bytePosition, bytePosition+bytes-1))
		bytePosition = bytePosition + bytes
	}
	return response
}
