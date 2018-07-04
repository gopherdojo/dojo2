package loader

import (
	"net/http"
	"strconv"
	"fmt"
	"io/ioutil"
	"regexp"
	"context"
)

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

func getFileSize(ctx context.Context, url string) (max int, high int, err error) {
	_, contentRange, err := rangeLoad(ctx, url, "0-1")
	if err != nil {
		return -1, -1, err
	}
	return parseContentRange(contentRange)
}

func rangeLoad(ctx context.Context, url string, fileRange string) (body []byte, contentRange string, err error) {
	errReturn := func(err error) ([]byte, string, error) {
		return nil, "", err
	}
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return errReturn(err)
	}
	req = req.WithContext(ctx)
	req.Header.Add("Range", "bytes=" + fileRange)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return errReturn(err)
	}
	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("GET not success")
	}
	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return errReturn(err)
	}
	contentRange = resp.Header.Get("content-range")
	return body, contentRange, nil
}

func parseContentRange(contentRange string) (max int, high int, err error) {
	compile, err := regexp.Compile("\\d+-(\\d+)/(\\d+)")
	errFunc := func(e error) (max int, high int, err error) {
		return -1, -1, e
	}
	if err != nil {
		return errFunc(err)
	}
	match := compile.FindSubmatch([]byte(contentRange))
	if len(match) != 3 {
		return errFunc(fmt.Errorf("index out of range"))
	}
	high, err = strconv.Atoi(string(match[1]))
	if err != nil {
		return errFunc(err)
	}
	max, err = strconv.Atoi(string(match[2]))
	if err != nil {
		return errFunc(err)
	}
	return max, high, nil
}
