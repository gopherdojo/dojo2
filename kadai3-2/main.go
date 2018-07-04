package main

import (
	"fmt"
	"time"
	"context"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
	"os"
)

func main() {
	url := "http://example.com"
	dirname := "/tmp/pdownload"
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	fmt.Println(url)
	res, _ := ctxhttp.Head(ctx, http.DefaultClient, url)
	fmt.Println(res.Header.Get("Accept-Ranges"))
	fmt.Println(res.Header.Get("Content-Length"))
	os.MkdirAll(dirname, 0755)
	req, _ := http.NewRequest("GET", url, nil)
	low := 0
	high := 605
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low, high))
	res, _ = http.DefaultClient.Do(req)


	return
}
