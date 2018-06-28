package sawadashota_test

import (
	"image"
	"os"
	"path"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/sawadashota"
)

func TestImageToPng(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	cases := []struct {
		name      string
		imgPath   string
		dest      string
		expectExt string
	}{
		{
			name:    "jpg_to_png",
			imgPath: path.Join(dir, "testdata/gopher.jpg"),
			dest:    path.Join(dir, "testdata/test.png"),
		},
		{
			name:    "gif_to_png",
			imgPath: path.Join(dir, "testdata/subdirectory/gopher.gif"),
			dest:    path.Join(dir, "testdata/subdirectory/test.png"),
		},
	}

	for _, c := range cases {
		if _, err := os.Stat(c.imgPath); err != nil {
			t.Fatalf("%s", err)
		}

		i, err := sawadashota.New(c.imgPath)

		if err != nil {
			t.Errorf("%s", err)
		}

		t.Run(c.name, func(t *testing.T) {
			testImage(t, c.imgPath, c.dest, "gif", i.ToGif)
		})
	}
}

func TestImageToJpeg(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	cases := []struct {
		name      string
		imgPath   string
		dest      string
		expectExt string
	}{
		{
			name:    "png_to_jpg",
			imgPath: path.Join(dir, "testdata/gopher.png"),
			dest:    path.Join(dir, "testdata/test.jpg"),
		},
		{
			name:    "gif_to_jpg",
			imgPath: path.Join(dir, "testdata/subdirectory/gopher.gif"),
			dest:    path.Join(dir, "testdata/subdirectory/test.jpg"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := os.Stat(c.imgPath); err != nil {
				t.Fatalf("%s", err)
			}

			i, err := sawadashota.New(c.imgPath)

			if err != nil {
				t.Errorf("%s", err)
			}

			testImage(t, c.imgPath, c.dest, "gif", i.ToGif)
		})
	}
}

func TestImageToGif(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("%s", err)
	}

	cases := []struct {
		name      string
		imgPath   string
		dest      string
		expectExt string
	}{
		{
			name:    "png_to_gif",
			imgPath: path.Join(dir, "testdata/gopher.png"),
			dest:    path.Join(dir, "testdata/test.gif"),
		},
		{
			name:    "jpg_to_gif",
			imgPath: path.Join(dir, "testdata/subdirectory/gopher.jpg"),
			dest:    path.Join(dir, "testdata/subdirectory/test.gif"),
		},
	}

	for _, c := range cases {
		if _, err := os.Stat(c.imgPath); err != nil {
			t.Fatalf("%s", err)
		}

		i, err := sawadashota.New(c.imgPath)

		if err != nil {
			t.Fatalf("%s", err)
		}

		t.Run(c.name, func(t *testing.T) {
			testImage(t, c.imgPath, c.dest, "gif", i.ToGif)
		})
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
		t.Fatalf("%s", err)
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
