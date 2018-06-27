package main

import (
	"fmt"
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
	if err := createJPEG(filepath.Join(dir, "image1.jpg"), 100, 200); err != nil {
		t.Fatal(err)
	}
	if err := createJPEG(filepath.Join(subdir, "image2.jpg"), 300, 400); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(subdir, "dummy.txt"), []byte("dummy"), 0644); err != nil {
		t.Fatal(err)
	}

	// Run main
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"main", dir}
	main()

	// Assert that destination contains PNG files
	if err := assertFilesIn(dir, []string{"image1.jpg", "image1.png"}); err != nil {
		t.Error(err)
	}
	if err := assertFilesIn(subdir, []string{"dummy.txt", "image2.jpg", "image2.png"}); err != nil {
		t.Error(err)
	}

	// Assert that PNG files are valid
	if err := assertPNG(filepath.Join(dir, "image1.png"), 100, 200); err != nil {
		t.Error(err)
	}
	if err := assertPNG(filepath.Join(subdir, "image2.png"), 300, 400); err != nil {
		t.Error(err)
	}
}

func assertFilesIn(dir string, expectedFiles []string) error {
	children, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	files := make([]string, 0)
	for _, child := range children {
		if !child.IsDir() {
			files = append(files, child.Name())
		}
	}
	if !reflect.DeepEqual(expectedFiles, files) {
		return fmt.Errorf("Directory %s wants %v but %v", dir, expectedFiles, files)
	}
	return nil
}

func createJPEG(name string, width int, height int) error {
	r, err := os.Create(name)
	if err != nil {
		return err
	}
	defer r.Close()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	return jpeg.Encode(r, img, nil)
}

func assertPNG(name string, width int, height int) error {
	r, err := os.Open(name)
	if err != nil {
		return err
	}
	defer r.Close()
	c, err := png.DecodeConfig(r)
	if err != nil {
		return err
	}
	if c.Width != width || c.Height != height {
		return fmt.Errorf("PNG %s wants %dx%d but %dx%d", name, width, height, c.Width, c.Height)
	}
	return nil
}
