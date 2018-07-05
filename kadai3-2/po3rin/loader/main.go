package main

import (
	"context"
	"dojo2/kadai3-2/po3rin/loader/length"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/errgroup"
)

// var wg sync.WaitGroup

func main() {
	length := length.Calc()

	// prepare val of split downloaf
	limit := 10
	leng := length / limit
	diff := length % limit
	body := make(map[int][]byte)

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// exec download file
	for i := 0; i < limit; i++ {
		i := i

		// wg.Add(1)
		min := leng * i
		max := leng * (i + 1)

		if i == limit-1 {
			max += diff
		}

		eg.Go(func() error {
			client := &http.Client{}
			select {
			case <-ctx.Done():
				return errors.New("Error occurred")
			default:
				req, err := http.NewRequest("GET", "http://localhost:8080/5mqx.cif", nil)
				if err != nil {
					cancel()
				}
				rh := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
				req.Header.Add("Range", rh)
				resp, err := client.Do(req)
				if err != nil {
					cancel()
				}
				defer resp.Body.Close()
				reader, err := ioutil.ReadAll(resp.Body)
				if err != err {
					cancel()
				}
				body[i] = reader
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

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
