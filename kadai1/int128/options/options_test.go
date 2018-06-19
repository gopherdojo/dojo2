package options

import (
	"image/jpeg"
	"image/png"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo2/kadai1/int128/images"
)

const arg0 = "kadai1"

func TestNoArg(t *testing.T) {
	_, err := Parse([]string{arg0})
	if err == nil {
		t.Errorf("err wants non-nil but %v", err)
	}
}

func TestUnknownFlag(t *testing.T) {
	_, err := Parse([]string{arg0, "-foo"})
	if err == nil {
		t.Errorf("err wants non-nil but %v", err)
	}
}

func TestDefaultArgs(t *testing.T) {
	opts, err := Parse([]string{arg0, "foo.jpg"})
	if err != nil {
		t.Fatal(err)
	}
	decoder, err := opts.Decoder()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := decoder.(*images.JPEG); !ok {
		t.Errorf("decoder wants JPEG but %+v", decoder)
	}
	encoder, err := opts.Encoder()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := encoder.(*images.PNG); !ok {
		t.Errorf("encoder wants PNG but %+v", encoder)
	}
}

func TestInvalidSourceFormat(t *testing.T) {
	opts, err := Parse([]string{arg0, "-from", "bar", "foo.jpg"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := opts.Encoder(); err != nil {
		t.Fatal(err)
	}
	if _, err := opts.Decoder(); err == nil {
		t.Errorf("err wants non-nil but nil")
	} else if !strings.Contains(err.Error(), "bar") {
		t.Errorf("error message wants bar but %s", err.Error())
	}
}

func TestInvalidDestinationFormat(t *testing.T) {
	opts, err := Parse([]string{arg0, "-to", "bar", "foo.jpg"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := opts.Decoder(); err != nil {
		t.Fatal(err)
	}
	if _, err := opts.Encoder(); err == nil {
		t.Errorf("err wants non-nil but nil")
	} else if !strings.Contains(err.Error(), "bar") {
		t.Errorf("error message wants bar but %s", err.Error())
	}
}

func TestFromPNGToJPEG(t *testing.T) {
	for _, m := range []struct {
		Args    []string
		Quality int
	}{
		{[]string{arg0, "-from", "png", "-to", "jpg", "foo.jpg"}, jpeg.DefaultQuality},
		{[]string{arg0, "-from", "png", "-to", "jpg", "-jpeg-quality", "5", "foo.jpg"}, 5},
	} {
		opts, err := Parse(m.Args)
		if err != nil {
			t.Fatal(err)
		}
		decoder, err := opts.Decoder()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := decoder.(*images.PNG); !ok {
			t.Errorf("decoder wants PNG but %+v", decoder)
		}
		encoder, err := opts.Encoder()
		if err != nil {
			t.Fatal(err)
		}
		if e, ok := encoder.(*images.JPEG); !ok {
			t.Errorf("encoder wants JPEG but %+v", encoder)
		} else if e.Options.Quality != m.Quality {
			t.Errorf("NumColors wants %d but %d", m.Quality, e.Options.Quality)
		}
	}
}

func TestFromJPEGToGIF(t *testing.T) {
	for _, m := range []struct {
		Args      []string
		NumColors int
	}{
		{[]string{arg0, "-to", "gif", "foo.jpg"}, 256},
		{[]string{arg0, "-to", "gif", "-gif-colors", "5", "foo.jpg"}, 5},
	} {
		opts, err := Parse(m.Args)
		if err != nil {
			t.Fatal(err)
		}
		decoder, err := opts.Decoder()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := decoder.(*images.JPEG); !ok {
			t.Errorf("decoder wants JPEG but %+v", decoder)
		}
		encoder, err := opts.Encoder()
		if err != nil {
			t.Fatal(err)
		}
		if e, ok := encoder.(*images.GIF); !ok {
			t.Errorf("encoder wants GIF but %+v", encoder)
		} else if e.Options.NumColors != m.NumColors {
			t.Errorf("NumColors wants %d but %d", m.NumColors, e.Options.NumColors)
		}
	}
}

func TestFromGIFToPNG(t *testing.T) {
	for _, m := range []struct {
		Args             []string
		CompressionLevel png.CompressionLevel
	}{
		{[]string{arg0, "-from", "gif", "foo.jpg"}, png.DefaultCompression},
		{[]string{arg0, "-from", "gif", "-to", "png", "foo.jpg"}, png.DefaultCompression},
		{[]string{arg0, "-from", "gif", "-to", "png", "-png-compression", "no", "foo.jpg"}, png.NoCompression},
		{[]string{arg0, "-from", "gif", "-to", "png", "-png-compression", "best-speed", "foo.jpg"}, png.BestSpeed},
		{[]string{arg0, "-from", "gif", "-to", "png", "-png-compression", "best-compression", "foo.jpg"}, png.BestCompression},
	} {
		opts, err := Parse(m.Args)
		if err != nil {
			t.Fatal(err)
		}
		decoder, err := opts.Decoder()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := decoder.(*images.GIF); !ok {
			t.Errorf("decoder wants GIF but %+v", decoder)
		}
		encoder, err := opts.Encoder()
		if err != nil {
			t.Fatal(err)
		}
		if e, ok := encoder.(*images.PNG); !ok {
			t.Errorf("encoder wants PNG but %+v", encoder)
		} else if e.Options.CompressionLevel != m.CompressionLevel {
			t.Errorf("NumColors wants %v but %v", m.CompressionLevel, e.Options.CompressionLevel)
		}
	}
}

func TestInvalidPNGCompressionLevel(t *testing.T) {
	opts, err := Parse([]string{arg0, "-to", "png", "-png-compression", "zzz", "foo.jpg"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := opts.Decoder(); err != nil {
		t.Fatal(err)
	}
	if _, err := opts.Encoder(); err == nil {
		t.Errorf("err wants non-nil but nil")
	} else if !strings.Contains(err.Error(), "zzz") {
		t.Errorf("error message wants zzz but %s", err.Error())
	}
}

func TestFromAutoToPNG(t *testing.T) {
	for _, m := range []struct {
		Args             []string
		CompressionLevel png.CompressionLevel
	}{
		{[]string{arg0, "-from", "auto", "foo.jpg"}, png.DefaultCompression},
	} {
		opts, err := Parse(m.Args)
		if err != nil {
			t.Fatal(err)
		}
		decoder, err := opts.Decoder()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := decoder.(*images.AutoDetect); !ok {
			t.Errorf("decoder wants AutoDetect but %+v", decoder)
		}
		encoder, err := opts.Encoder()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := encoder.(*images.PNG); !ok {
			t.Errorf("encoder wants PNG but %+v", encoder)
		}
	}
}
