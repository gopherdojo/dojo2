package loader

import (
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"context"
)

func Download(ctx context.Context, url string, split int) error {
	fileSize, _, e := getFileSize(ctx, url)
	if e != nil {
		return e
	}
	eg, ctx := errgroup.WithContext(ctx)
	fileRanges := splitRange(fileSize, split)
	bodies := make([][]byte, len(fileRanges))
	for i, v := range fileRanges {
		loader := partialLoader{Url: url, Index: i, FileRange: v}
		eg.Go(func() error {
			return loader.Load(ctx, bodies)
		})
	}
	if e := eg.Wait(); e != nil {
		return e
	}
	var result []byte
	for _, v := range bodies {
		result = append(result, v...)
	}
	ioutil.WriteFile("image.jpeg", result, 0666)
	return nil
}
