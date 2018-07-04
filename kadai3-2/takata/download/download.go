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
	size  uint
	bytes [][]byte
}

type Option interface {
	Proc() int
	Timeout() time.Duration
	Writer() io.Writer
	Output() string
}

func New(u *url.URL, opts Option) *Downloader {
	return &Downloader{
		url:    u,
		Option: opts,
		Data:   Data{},
	}
}

func (d *Downloader) URL() *url.URL {
	return d.url
}

func (d *Downloader) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filesize, err := d.FetchFileSize(d.url.String())
	if err != nil {
		return err
	}

	d.SetFileSize(filesize)
	d.bytes = make([][]byte, d.Proc())

	errorgroup := d.get(ctx)
	if err := errorgroup.Wait(); err != nil {
		return err
	}

	body := d.Merge()

	f, err := os.Create(d.Output())
	defer f.Close()

	f.Write(body)

	return nil

}

// SetFileSize is to set Data file size
func (d *Data) SetFileSize(size uint) {
	d.size = size
}

// Merge data
func (d *Data) Merge() []byte {
	return bytes.Join(d.bytes, []byte(""))
}

// return String merged data
func (d *Data) String() string {
	return string(d.Merge())
}
