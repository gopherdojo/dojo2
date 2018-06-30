package main

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"io/ioutil"
)

func main() {
	const Split = 10
	url := "https://kaboompics.com/cache/6/e/e/8/1/6ee81e1477ee1a9610149d0fe7fbd952213ba11d.jpeg?version=v3"
	fileSize, _, e := GetFileSize(url)
	if e != nil {
		os.Exit(1)
	}
	fileRanges := splitRange(fileSize, Split)
	var bodies []byte
	for _, r := range fileRanges {
		body, _, err := RangeLoad(url, r)
		if err != nil {
			os.Exit(1)
		}
		bodies = append(bodies, body...)
	}
	ioutil.WriteFile("image.jpeg", bodies, 0666)
}

// return "0-100"
func splitRange(fileSize int, split int) []string {
	aFileSize := fileSize / split
	var ranges []string
	start := 0
	end := 0
	for i := 0; i < split; i++ {
		if i == 0 {
			start = 0
		} else {
			start = end + 1
		}
		end = end + aFileSize
		ranges = append(ranges,  strconv.Itoa(start) + "-" + strconv.Itoa(end))
	}
	if fileSize % split != 0 {
		ranges = append(ranges, strconv.Itoa(end + 1) + "-" + strconv.Itoa(fileSize))
	}
	return ranges
}

func GetFileSize(url string) (max int, high int, err error) {
	_, contentRange, err := RangeLoad(url, "0-1")
	return parseContentRange(contentRange)
}

func RangeLoad(url string, fileRange string) (body []byte, contentRange string, err error) {
	errReturn := func(err error) (byte, string, error) {
		return 0, "", err
	}
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		errReturn(err)
	}
	req.Header.Add("Range", "bytes=" + fileRange)
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
