package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

// Client - 分割DL用クライアント
type Client struct {
	URL                        *url.URL
	HTTPClient                 *http.Client
	ContentLength, SplitNumber int
	Filename                   string
}

// NewClient - Client の初期化
func NewClient(urlString string) (*Client, error) {
	var err error
	client := &Client{}
	client.URL, err = url.Parse(urlString)
	splitedPath := strings.Split(client.URL.Path, "/")
	client.Filename = splitedPath[len(splitedPath)-1]
	if err != nil {
		return nil, err
	}
	client.HTTPClient = &http.Client{}
	if err = client.setContentLengthAndSplitNumber(); err != nil {
		return nil, err
	}
	return client, nil
}

// setContentLengthAndSplitNumber - レスポンスヘッダに Accept-Ranges が含まれている場合は分割, ない場合は分割しない
func (c *Client) setContentLengthAndSplitNumber() error {
	req, err := http.NewRequest("HEAD", c.URL.String(), nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	c.ContentLength, err = strconv.Atoi(res.Header["Content-Length"][0])
	if err != nil {
		return err
	}
	if strings.Contains("bytes", res.Header["Accept-Ranges"][0]) {
		c.SplitNumber = runtime.NumCPU()
	} else {
		c.SplitNumber = 1
	}
	return nil
}

func (c *Client) newRequest(ctx context.Context, index int) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.URL.String(), nil)
	if err != nil {
		return req, nil
	}

	req = req.WithContext(ctx)
	req.Header.Add("User-Agent", "bget")
	var chank int
	chank = c.ContentLength / c.SplitNumber
	low := chank * index
	high := chank*(index+1) - 1
	if index+1 == c.SplitNumber && high+1 != c.ContentLength {
		high = c.ContentLength - 1
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low, high))
	return req, nil
}

func (c *Client) save(ctx context.Context, index int) error {
	req, err := c.newRequest(ctx, index)
	if err != nil {
		return err
	}
	file, err := os.Create(c.Filename + "_" + strconv.Itoa(index))
	if err != nil {
		return err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func (c *Client) merge(ctx context.Context) error {
	file, err := os.Create(c.Filename)
	if err != nil {
		return err
	}
	for i := 0; i < c.SplitNumber; i++ {
		subFileName := c.Filename + "_" + strconv.Itoa(i)
		subFile, err := os.Open(subFileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(file, subFile)
		subFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Download - 分割DLを実行
func (c *Client) Download(ctx context.Context) error {
	eg := errgroup.Group{}
	for i := 0; i < c.SplitNumber; i++ {
		index := i
		eg.Go(func() error {
			return c.save(ctx, index)
		})
	}
	if err := eg.Wait(); err != nil {
		c.DeleteFiles()
		return err
	}
	err := c.merge(ctx)
	c.DeleteFiles()
	if err != nil {
		return err
	}

	return nil
}

// DeleteFiles - キャンセルなどがあった際に呼ばれる. 分割ファイルの削除を実行.
func (c *Client) DeleteFiles() error {
	for i := 0; i < c.SplitNumber; i++ {
		subFileName := c.Filename + "_" + strconv.Itoa(i)
		if fi, _ := os.Stat(subFileName); fi != nil {
			if err := os.Remove(subFileName); err != nil {
				return err
			}
		}
	}
	return nil
}
