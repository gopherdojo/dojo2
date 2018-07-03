package download

import (
	"fmt"
)

// ContentRange represents a HTTP Content-Range header.
// See also RFC7233.
type ContentRange struct {
	Partial  Range  // Received range of the content
	Complete *Range // Complete range of the content (may be nil)
}

// ParseContentRange parses a Content-Range header and returns the Range.
func ParseContentRange(header string) (*ContentRange, error) {
	rng := Range{}
	if _, err := fmt.Sscanf(header, "bytes %d-%d/*", &rng.Start, &rng.End); err == nil {
		return &ContentRange{rng, nil}, nil
	}
	var length int64
	if _, err := fmt.Sscanf(header, "bytes %d-%d/%d", &rng.Start, &rng.End, &length); err == nil {
		return &ContentRange{rng, &Range{0, length - 1}}, nil
	}
	return nil, fmt.Errorf("Invalid Content-Range header: %s", header)
}

// Range represents range for [start, end].
// See also RFC7233.
type Range struct {
	Start int64 // Start position
	End   int64 // End position
}

// HeaderValue returns a value of Range header, e.g. bytes=100-200.
func (r *Range) HeaderValue() string {
	return fmt.Sprintf("bytes=%d-%d", r.Start, r.End)
}

// Length returns length.
func (r *Range) Length() int64 {
	return r.End - r.Start + 1
}

// Split splits the range into chunks.
// If count is 0, it returns an empty slice.
// If count is too big, it returns a possible longest slice.
func (r *Range) Split(count int) []Range {
	if count < 1 {
		return []Range{}
	}
	unit := divCeil(r.Length(), int64(count))
	chunks := make([]Range, 0, count)
	for p := r.Start; p <= r.End; p += unit {
		rng := Range{
			Start: p,
			End:   min(p+unit-1, r.End),
		}
		chunks = append(chunks, rng)
	}
	return chunks
}

// divCeil(a, b) = ceil(a / b)
func divCeil(a int64, b int64) int64 {
	if a%b > 0 {
		return a/b + 1
	}
	return a / b
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
