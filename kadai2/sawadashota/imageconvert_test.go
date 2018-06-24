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

	cases := []struct {
		imgPath   string
		dest      string
		expectExt string
	}{
		{
			imgPath: path.Join(dir, "testdata/gopher.jpg"),
			dest:    path.Join(dir, "testdata/test.png"),
		},
		{
			imgPath: path.Join(dir, "testdata/subdirectory/gopher.gif"),
			dest:    path.Join(dir, "testdata/subdirectory/test.png"),
		},
	}

	for _, c := range cases {
		if _, err := os.Stat(c.imgPath); err != nil {
			t.Fatalf("%s", err)
		}

		i, err := New(c.imgPath)

		if err != nil {
			t.Errorf("%s", err)
		}

		testImage(t, c.imgPath, c.dest, "png", i.ToPng)
	}
}

func TestImageToJpeg(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	cases := []struct {
		imgPath   string
		dest      string
		expectExt string
	}{
		{
			imgPath: path.Join(dir, "testdata/gopher.png"),
			dest:    path.Join(dir, "testdata/test.jpg"),
		},
		{
			imgPath: path.Join(dir, "testdata/subdirectory/gopher.gif"),
			dest:    path.Join(dir, "testdata/subdirectory/test.jpg"),
		},
	}

	for _, c := range cases {
		if _, err := os.Stat(c.imgPath); err != nil {
			t.Fatalf("%s", err)
		}

		i, err := New(c.imgPath)

		if err != nil {
			t.Errorf("%s", err)
		}

		testImage(t, c.imgPath, c.dest, "jpeg", i.ToJpeg)
	}
}

func TestImageToGif(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	cases := []struct {
		imgPath   string
		dest      string
		expectExt string
	}{
		{
			imgPath: path.Join(dir, "testdata/gopher.png"),
			dest:    path.Join(dir, "testdata/test.gif"),
		},
		{
			imgPath: path.Join(dir, "testdata/subdirectory/gopher.jpg"),
			dest:    path.Join(dir, "testdata/subdirectory/test.gif"),
		},
	}

	for _, c := range cases {
		if _, err := os.Stat(c.imgPath); err != nil {
			t.Errorf("%s", err)
		}

		i, err := New(c.imgPath)

		if err != nil {
			t.Errorf("%s", err)
		}

		testImage(t, c.imgPath, c.dest, "gif", i.ToGif)
	}
}

func testImage(t *testing.T, imgPath, dest, expectExt string, convert func(dest string) error) {
	t.Helper()

	if _, err := os.Stat(imgPath); err != nil {
		t.Fatalf("%s", err)
	}

	if err := convert(dest); err != nil {
		t.Errorf("%s", err)
	}

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

	if format != expectExt {
		t.Errorf("expect %s but %s", expectExt, format)
	}

	if os.Remove(dest) != nil {
		t.Fatalf("%s", err)
	}
}
