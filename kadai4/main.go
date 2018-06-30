package main

import (
	"net/http"
	"os"
	"fmt"
	"regexp"
)

func main() {
	req, err := http.NewRequest("GET", "https://kaboompics.com/cache/6/e/e/8/1/6ee81e1477ee1a9610149d0fe7fbd952213ba11d.jpeg?version=v3", nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("Range", "bytes=0-1")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
 	contentRange := resp.Header.Get("content-range")
	compile := regexp.MustCompile("\\d-\\d/(\\d)")
	if err != nil {
		os.Exit(1)
	}
	match := compile.FindSubmatch([]byte(contentRange))
	bytes := match[1]
	s := string(bytes)
	fmt.Println(s)
}
