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
		t.Run(c.name, func(t *testing.T) {
			i := testNewImage(t, c.imgPath)

			if err := i.ToPng(c.dest); err != nil {
				t.Errorf("%s", err)
			}

			if _, err := os.Stat(c.dest); err != nil {
				t.Errorf("%s", err)
			}

			file, err := os.Open(c.dest)

			if err != nil {
				t.Fatalf("%s", err)
			}
			defer file.Close()

			_, format, err := image.Decode(file)

			if err != nil {
				t.Errorf("%s", err)
			}

			if format != "png" {
				t.Errorf("expect %s but %s", "png", format)
			}

			if os.Remove(c.dest) != nil {
				t.Fatalf("%s", err)
			}
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
			i := testNewImage(t, c.imgPath)

			if err := i.ToJpeg(c.dest); err != nil {
				t.Errorf("%s", err)
			}

			if _, err := os.Stat(c.dest); err != nil {
				t.Errorf("%s", err)
			}

			file, err := os.Open(c.dest)

			if err != nil {
				t.Fatalf("%s", err)
			}
			defer file.Close()

			_, format, err := image.Decode(file)

			if err != nil {
				t.Errorf("%s", err)
			}

			if format != "jpeg" {
				t.Errorf("expect %s but %s", "jpeg", format)
			}

			if os.Remove(c.dest) != nil {
				t.Fatalf("%s", err)
			}
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
		t.Run(c.name, func(t *testing.T) {
			i := testNewImage(t, c.imgPath)

			if err := i.ToGif(c.dest); err != nil {
				t.Errorf("%s", err)
			}

			if _, err := os.Stat(c.dest); err != nil {
				t.Errorf("%s", err)
			}

			file, err := os.Open(c.dest)

			if err != nil {
				t.Fatalf("%s", err)
			}
			defer file.Close()

			_, format, err := image.Decode(file)

			if err != nil {
				t.Errorf("%s", err)
			}

			if format != "gif" {
				t.Errorf("expect %s but %s", "gif", format)
			}

			if os.Remove(c.dest) != nil {
				t.Fatalf("%s", err)
			}
		})
	}
}

func testNewImage(t *testing.T, imgPath string) *sawadashota.Image {
	t.Helper()

	if _, err := os.Stat(imgPath); err != nil {
		t.Fatalf("%s", err)
	}

	i, err := sawadashota.New(imgPath)

	if err != nil {
		t.Fatalf("%s", err)
	}

	return i
}
