package download

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"os"
	"time"
)

type Downloader struct {
	url *url.URL
	Option
	Data
}

type Data struct {
	filesize uint
	data     [][]byte
}

type Option interface {
	Proc() int
	Timeout() time.Duration
	Writer() io.Writer
	Output() string
}

// New Downloader
func New(u *url.URL, opts Option) *Downloader {
	return &Downloader{
		url:    u,
		Option: opts,
		Data:   Data{},
	}
}

// URL getter
func (d *Downloader) URL() *url.URL {
	return d.url
}

// Run download
func (d *Downloader) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filesize, err := d.FetchFileSize(d.url.String())
	if err != nil {
		return err
	}

	d.SetFileSize(filesize)
	d.data = make([][]byte, d.Proc())

	eg := d.get(ctx)
	if err := eg.Wait(); err != nil {
		return err
	}

	body := d.Merge()

	f, err := os.Create(d.Output())
	defer f.Close()

	f.Write(body)

	return nil
}

// SetFileSize is filesize setter
func (d *Data) SetFileSize(size uint) {
	d.filesize = size
}

// Merge data slice
func (d *Data) Merge() []byte {
	return bytes.Join(d.data, []byte(""))
}

// String return merged data
func (d *Data) String() string {
	return string(d.Merge())
}
