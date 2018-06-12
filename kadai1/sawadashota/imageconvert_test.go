package sawadashota

import (
	"image"
	"os"
	"path"
	"testing"
)

func TestImageToPng(t *testing.T) {
	dir, _ := os.Getwd()
	imgPath := path.Join(dir, "test-data/gopher.jpg")

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	i, err := New(imgPath)

	if err != nil {
		t.Errorf("%s", err)
	}

	dest := path.Join(dir, "test-data/test.jpg")

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

	os.Remove(dest)
}

func TestImageToJpeg(t *testing.T) {
	dir, _ := os.Getwd()
	imgPath := path.Join(dir, "test-data/gopher.png")

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	i, err := New(imgPath)

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	dest := path.Join(dir, "test-data/test.jpg")

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

	os.Remove(dest)
}

func TestImageToGif(t *testing.T) {
	dir, _ := os.Getwd()
	imgPath := path.Join(dir, "test-data/gopher.png")

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	i, err := New(imgPath)

	if _, err := os.Stat(imgPath); err != nil {
		t.Errorf("%s", err)
	}

	dest := path.Join(dir, "test-data/test.gif")

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

	os.Remove(dest)
}
