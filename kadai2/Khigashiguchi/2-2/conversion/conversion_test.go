package conversion_test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/conversion"
	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/options"
)

func TestReplaceExt(t *testing.T) {
	cases := []struct {
		name     string
		filepath string
		to       string
		expected string
	}{
		{
			name:     "convert from png to jpeg",
			filepath: "sample.png",
			to:       "jpeg",
			expected: "sample.jpeg",
		},
		{
			name:     "convert from jpeg to png",
			filepath: "sample.jpeg",
			to:       "png",
			expected: "sample.png",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// FIXME: Optionsを作るのはテストの本筋ではないのでヘルパーに、さらなる改善が可能かもしれない
			opt := &options.Options{
				Files: []string{"sample.jpg", "sample1.jpg"},
				In:    "png",
				Out:   tt.to,
			}
			decoder, err := opt.Decoder()
			if err != nil {
				panic(err)
			}
			encoder, err := opt.Encoder()
			if err != nil {
				panic(err)
			}
			cv := &conversion.Conversion{
				Decoder: decoder,
				Encoder: encoder,
				ToExt:   tt.to,
			}
			actual := cv.ReplaceExt(tt.filepath)
			if actual != tt.expected {
				t.Errorf(`expected=%s, acutal=%s`, tt.expected, actual)
			}
		})
	}
}
