package download

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Range struct {
	low  int
	high int
	proc int
}

func NewRange(filesize uint, procs, proc int) *Range {
	split := int(filesize) / procs
	return &Range{
		low:  split * (proc - 1),
		high: split * proc,
	}
}

// Low getter
func (r *Range) Low() int {
	return r.low
}

// High getter
func (r *Range) High() int {
	return r.high
}

// FetchFileSize fetch content length and set filesize
func (d *Data) FetchFileSize(URL string) (uint, error) {
	resp, err := http.Head(URL)

	if err != nil {
		return 0, err
	}

	return uint(resp.ContentLength), nil
}

// Get contents concurrently
func (d *Downloader) get(ctx context.Context) *errgroup.Group {
	eg, ctx := errgroup.WithContext(ctx)

	for i := 0; i < d.Proc(); i++ {
		i := i
		r := NewRange(d.filesize, d.Proc(), i+1)

		eg.Go(func() error {
			req, err := http.NewRequest(http.MethodGet, d.URL().String(), nil)
			if err != nil {
				return err
			}

			req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.Low(), r.High()))
			req = req.WithContext(ctx)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}

			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return err
			}

			d.data[i] = body
			return nil
		})
	}

	return eg
}
