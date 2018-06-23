package sawadashota

import (
	"image"
	"os"
	"path"
	"testing"
)

func TestImageToPng(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	imgPath := path.Join(dir, "testdata/gopher.jpg")

	if _, err := os.Stat(imgPath); err != nil {
		t.Fatalf("%s", err)
	}

	i, err := New(imgPath)

	if err != nil {
		t.Errorf("%s", err)
	}

	dest := path.Join(dir, "testdata/test.jpg")

	i.ToPng(dest)

	if _, err := os.Stat(dest); err != nil {
		t.Errorf("%s", err)
	}

	file, err := os.Open(dest)

	if err != nil {
		t.Errorf("%s", err)
	}
	defer file.Close()

	_, format, err := image.Decode(file)

	if err != nil {
		t.Errorf("%s", err)
	}

	if format != "png" {
		t.Errorf("expect png but %s", format)
	}

	if os.Remove(dest) != nil {
		t.Fatalf("%s", err)
	}
}

func TestImageToJpeg(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	imgPath := path.Join(dir, "testdata/gopher.png")

	if _, err := os.Stat(imgPath); err != nil {
		t.Fatalf("%s", err)
	}

	i, err := New(imgPath)

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	dest := path.Join(dir, "testdata/test.jpg")

	i.ToJpeg(dest)

	if _, err := os.Stat(dest); err != nil {
		t.Errorf("%s", err)
	}

	file, err := os.Open(dest)

	if err != nil {
		t.Errorf("%s", err)
	}
	defer file.Close()

	_, format, err := image.Decode(file)

	if err != nil {
		t.Errorf("%s", err)
	}

	if format != "jpeg" {
		t.Errorf("expect jpeg but %s", format)
	}

	if os.Remove(dest) != nil {
		t.Fatalf("%s", err)
	}
}

func TestImageToGif(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	imgPath := path.Join(dir, "testdata/gopher.png")

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	i, err := New(imgPath)

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	dest := path.Join(dir, "testdata/test.gif")

	i.ToGif(dest)

	if _, err := os.Stat(dest); err != nil {
		t.Errorf("%s", err)
	}

	file, err := os.Open(dest)

	if err != nil {
		t.Errorf("%s", err)
	}
	defer file.Close()

	_, format, err := image.Decode(file)

	if err != nil {
		t.Errorf("%s", err)
	}

	if format != "gif" {
		t.Errorf("expect jpeg but %s", format)
	}

	if os.Remove(dest) != nil {
		t.Fatalf("%s", err)
	}
}
