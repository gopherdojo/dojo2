package loader

import (
	"context"
)

type partialLoader struct {
	Url       string
	Index     int
	FileRange string
}

func (p partialLoader) Load(ctx context.Context, bodies [][]byte) error {
	body, _, err := rangeLoad(ctx, p.Url, p.FileRange)
	if err != nil {
		return err
	}
	bodies[p.Index] = body
	return nil
}
