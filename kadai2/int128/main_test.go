package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestMain performs the integration test for JPEG-PNG conversion.
func TestMain(t *testing.T) {
	dir, err := ioutil.TempDir("", "main")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	subdir := filepath.Join(dir, "subdir")
	if err := os.Mkdir(subdir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create test fixtures
	createJPEG(t, filepath.Join(dir, "image1.jpg"), 100, 200)
	createJPEG(t, filepath.Join(subdir, "image2.jpg"), 300, 400)
	if err := ioutil.WriteFile(filepath.Join(subdir, "dummy.txt"), []byte("dummy"), 0644); err != nil {
		t.Fatal(err)
	}

	// Run main
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"main", dir}
	main()

	// Assert that destination contains PNG files
	assertFilesIn(t, dir, []string{"image1.jpg", "image1.png"})
	assertFilesIn(t, subdir, []string{"dummy.txt", "image2.jpg", "image2.png"})

	// Assert that PNG files are valid
	assertPNG(t, filepath.Join(dir, "image1.png"), 100, 200)
	assertPNG(t, filepath.Join(subdir, "image2.png"), 300, 400)
}

func assertFilesIn(t *testing.T, dir string, expectedFiles []string) {
	t.Helper()
	children, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	files := make([]string, 0)
	for _, child := range children {
		if !child.IsDir() {
			files = append(files, child.Name())
		}
	}
	if !reflect.DeepEqual(expectedFiles, files) {
		t.Errorf("Directory %s wants %v but %v", dir, expectedFiles, files)
	}
}

func createJPEG(t *testing.T, name string, width int, height int) {
	t.Helper()
	r, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if err := jpeg.Encode(r, img, nil); err != nil {
		t.Fatal(err)
	}
}

func assertPNG(t *testing.T, name string, width int, height int) {
	t.Helper()
	r, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	c, err := png.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	if c.Width != width || c.Height != height {
		t.Errorf("PNG %s wants %dx%d but %dx%d", name, width, height, c.Width, c.Height)
	}
}
