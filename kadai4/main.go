package main

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"io/ioutil"
)

func main() {
	url := "https://kaboompics.com/cache/6/e/e/8/1/6ee81e1477ee1a9610149d0fe7fbd952213ba11d.jpeg?version=v3"
	fileSize, _, e := GetFileSize(url)
	up := fileSize / 2
	if e != nil {
		os.Exit(1)
	}
	body1, _, err := RangeLoad(url, 0, up)
	if err != nil {
		os.Exit(1)
	}
	body2, _, err := RangeLoad(url, up + 1, fileSize)
	if err != nil {
		os.Exit(1)
	}
	bytes := append(body1, body2...)
	ioutil.WriteFile("image.jpeg", bytes, 0666)
}

func GetFileSize(url string) (max int, high int, err error) {
	_, contentRange, err := RangeLoad(url, 0, 1)
	return parseContentRange(contentRange)
}

func RangeLoad(url string, from int, to int) (body []byte, contentRange string, err error) {
	errReturn := func(err error) (byte, string, error) {
		return 0, "", err
	}
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		errReturn(err)
	}
	req.Header.Add("Range", "bytes=" + strconv.Itoa(from) + "-" + strconv.Itoa(to))
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		errReturn(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		errReturn(err)
	}
	contentRange = resp.Header.Get("content-range")
	return body, contentRange, nil
}

func parseContentRange(contentRange string) (max int, high int, err error) {
	compile := regexp.MustCompile("\\d+-(\\d+)/(\\d+)")
	if err != nil {
		return -1, -1, err
	}
	match := compile.FindSubmatch([]byte(contentRange))
	high, err = strconv.Atoi(string(match[1]))
	max, err = strconv.Atoi(string(match[2]))
	if err != nil {
		return -1, -1, err
	}
	return max, high, nil
}
