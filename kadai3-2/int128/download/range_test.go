package download

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseContentRange(t *testing.T) {
	matrix := []struct {
		header   string
		expected *ContentRange
	}{
		{"bytes 0-0/100", &ContentRange{Range{0, 0}, &Range{0, 99}}},
		{"bytes 0-0/*", &ContentRange{Range{0, 0}, nil}},
		{"bytes 12345-67890/123456", &ContentRange{Range{12345, 67890}, &Range{0, 123455}}},
		{"bytes 12345-67890/*", &ContentRange{Range{12345, 67890}, nil}},
	}
	for _, m := range matrix {
		t.Run(m.header, func(t *testing.T) {
			rng, err := ParseContentRange(m.header)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(m.expected, rng) {
				t.Errorf("range wants %+v but %+v", m.expected, rng)
			}
		})
	}
}

func TestParseWrongContentRange(t *testing.T) {
	rng, err := ParseContentRange("foo")
	if err == nil {
		t.Errorf("wants error but %+v", rng)
	}
}

func TestRangeSplit(t *testing.T) {
	matrix := []struct {
		r        Range
		count    int
		expected []Range
	}{
		{Range{1, 1}, 0, []Range{}},
		{Range{1, 1}, 1, []Range{Range{1, 1}}},

		{Range{1, 2}, 1, []Range{Range{1, 2}}},
		{Range{1, 2}, 2, []Range{Range{1, 1}, Range{2, 2}}},
		{Range{1, 2}, 3, []Range{Range{1, 1}, Range{2, 2}}},

		{Range{1, 3}, 1, []Range{Range{1, 3}}},
		{Range{1, 3}, 2, []Range{Range{1, 2}, Range{3, 3}}},
		{Range{1, 3}, 3, []Range{Range{1, 1}, Range{2, 2}, Range{3, 3}}},
		{Range{1, 3}, 4, []Range{Range{1, 1}, Range{2, 2}, Range{3, 3}}},
	}
	for _, m := range matrix {
		t.Run(fmt.Sprintf("%+v/Count:%d", m.r, m.count), func(t *testing.T) {
			chunks := m.r.Split(m.count)
			if !reflect.DeepEqual(m.expected, chunks) {
				t.Errorf("chunks wants %+v but %+v", m.expected, chunks)
			}
		})
	}
}
