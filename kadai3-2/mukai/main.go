package main

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"io/ioutil"
)

type PartialData struct {
	index int
	data []byte
}

func main() {
	const Split = 1
	url := "https://www.noao.edu/image_gallery/images/d7/cygloop-8000.jpg"
	fileSize, _, e := GetFileSize(url)
	if e != nil {
		os.Exit(1)
	}
	fileRanges := splitRange(fileSize, Split)
	bodies := make([][]byte, len(fileRanges))
	ch := make(chan PartialData)
	for i, v := range fileRanges {
		go storePartialData(ch, url, i, v)
	}
	count := 0
	for ;; {
		data := <- ch
		bodies[data.index] = data.data
		count = count + 1
		if count == Split {
			close(ch)
			break
		}
	}
	var result []byte
	for _, v := range bodies {
		result = append(result, v...)
	}
	ioutil.WriteFile("image.jpeg", result, 0666)
}

func storePartialData(ch chan <- PartialData, url string, index int, fileRange string) {
	body, _, err := RangeLoad(url, fileRange)
	if err != nil {
		os.Exit(1)
	}
	ch <- PartialData{index: index, data: body}
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
