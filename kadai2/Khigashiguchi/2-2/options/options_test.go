package options_test

import (
	"errors"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/options"
)

const arg0 = "cli"

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		options  []string
		expected *options.Options
		err      error
	}{
		{
			name:     "noArgs",
			options:  []string{arg0},
			expected: nil,
			err:      errors.New("few argument"),
		},
		{
			name:    "passArgs",
			options: []string{arg0, "sample.jpg", "sample1.jpg"},
			expected: &options.Options{
				Files: []string{"sample.jpg", "sample1.jpg"},
				In:    "jpg",
				Out:   "png",
			},
			err: nil,
		},
		{
			name:    "passArgsAndOpts",
			options: []string{arg0, "--in", "png", "--out", "jpg", "sample.jpg", "sample1.jpg"},
			expected: &options.Options{
				Files: []string{"sample.jpg", "sample1.jpg"},
				In:    "png",
				Out:   "jpg",
			},
			err: nil,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := options.Parse(tt.options)
			if tt.err != nil {
				if err == nil {
					t.Errorf(`expected to cause err %v`, tt.err)
				}
				if err.Error() != tt.err.Error() {
					t.Errorf(`expect err="%v" actual err="%v"`, tt.err, err)
				}
			}
			if tt.expected != nil {
				if actual.In != tt.expected.In {
					t.Errorf(`expect return field in="%v" actual field in="%v"`, tt.expected.In, actual.In)
				}
				if actual.Out != tt.expected.Out {
					t.Errorf(`expect return field out="%v" actual field out="%v"`, tt.expected.Out, actual.Out)
				}
				for i, arg := range tt.expected.Files {
					if actual.Files[i] != arg {
						t.Errorf(`expect return field files="%v" actual field files="%v"`, tt.expected.Files, actual.Files)
					}
				}
			}
		})
	}
}

func TestOptions_Decode(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{
			name: "JPG",
			in:   "jpg",
		},
		{
			name: "PNG",
			in:   "png",
		},
	}
	for _, tt := range cases {
		opts := &options.Options{
			Files: []string{"sample.jpg", "sample1.jpg"},
			In:    tt.in,
			Out:   "png",
		}
		t.Run(tt.name, func(t *testing.T) {
			_, err := opts.Decoder()
			if err != nil {
				t.Errorf(`expect err is nil, actual err is %s`, err)
			}
			// FIXME: どのdecoderが来たのかを確認したい
		})
	}
}

func TestOptions_DecodeUnknown(t *testing.T) {
	opts := &options.Options{
		Files: []string{"sample.jpg", "sample1.jpg"},
		In:    "unknown",
		Out:   "png",
	}
	_, err := opts.Decoder()
	if err == nil {
		t.Error(`expect err is not nil`)
	}
	expected := errors.New("Unknown extention type: unknown")
	if err.Error() != expected.Error() {
		t.Errorf(`expect err is %s, actual err is %s`, err.Error(), expected)
	}
}

func TestOptions_Encode(t *testing.T) {
	cases := []struct {
		name string
		out  string
	}{
		{
			name: "JPG",
			out:  "jpg",
		},
		{
			name: "PNG",
			out:  "png",
		},
	}
	for _, tt := range cases {
		opts := &options.Options{
			Files: []string{"sample.jpg", "sample1.jpg"},
			In:    "jpg",
			Out:   tt.out,
		}
		t.Run(tt.name, func(t *testing.T) {
			_, err := opts.Encoder()
			if err != nil {
				t.Errorf(`expect err is nil, actual err is %s`, err)
			}
			// FIXME: どのdecoderが来たのかを確認したい
		})
	}
}

func TestOptions_EncodeUnknown(t *testing.T) {
	opts := &options.Options{
		Files: []string{"sample.jpg", "sample1.jpg"},
		In:    "png",
		Out:   "unknown",
	}
	_, err := opts.Encoder()
	if err == nil {
		t.Error(`expect err is not nil`)
	}
	expected := errors.New("Unknown extention type: unknown")
	if err.Error() != expected.Error() {
		t.Errorf(`expect err is %s, actual err is %s`, err.Error(), expected)
	}
}
