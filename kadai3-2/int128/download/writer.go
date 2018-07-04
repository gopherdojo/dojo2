package download

import (
	"fmt"
	"io"
)

// Position represents relative position in the range.
type Position struct {
	Range  Range
	Offset int64 // offset in the range
}

// Absolute returns absolute position.
func (p *Position) Absolute() int64 {
	return p.Range.Start + p.Offset
}

// Forward forwards the position.
func (p *Position) Forward(n int64) {
	p.Offset += n
}

// CanForward returns true if incremented position is in the range.
func (p *Position) CanForward(n int64) bool {
	return p.Absolute()+n-1 <= p.Range.End
}

// RangeWriter supports partial write.
type RangeWriter struct {
	io.WriterAt
	position Position // relative position in the range
}

// NewRangeWriter returns a new RangeWriter.
func NewRangeWriter(w io.WriterAt, r Range) *RangeWriter {
	return &RangeWriter{w, Position{r, 0}}
}

func (w *RangeWriter) Write(p []byte) (int, error) {
	if !w.position.CanForward(int64(len(p))) {
		return 0, fmt.Errorf("Write position exceeds the range: len(p)=%d, position=%+v", len(p), w.position)
	}
	n, err := w.WriteAt(p, w.position.Absolute())
	w.position.Forward(int64(n))
	return n, err
}
