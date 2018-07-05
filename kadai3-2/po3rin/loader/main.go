package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// get contents length
	res, err := http.Head("http://localhost:8080/5mqx.cif")
	if err != nil {
		fmt.Println(err)
	}
	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		fmt.Println(err)
	}

	// prepare val of split downloaf
	limit := 10
	leng := length / limit
	diff := length % limit
	body := make(map[int][]byte)

	// exec download file
	for i := 0; i < limit; i++ {
		wg.Add(1)
		min := leng * i
		max := leng * (i + 1)

		if i == limit-1 {
			max += diff
		}

		go func(min int, max int, i int) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://localhost:8080/5mqx.cif", nil)
			if err != nil {
				fmt.Println(err)
			}
			rh := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
			req.Header.Add("Range", rh)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			reader, err := ioutil.ReadAll(resp.Body)
			if err != err {
				fmt.Println(err)
			}
			body[i] = reader
			wg.Done()
		}(min, max, i)
	}
	wg.Wait()

	// create new file
	buf := make([]byte, 0)
	for i := 0; i < limit; i++ {
		buf = append(buf, body[i]...)
	}
	file, err := os.Create(`../5mqx.cif`)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write(([]byte)(buf))
}
