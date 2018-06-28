package main

import (
	"testing"
)

func TestExtension(t *testing.T) {
	var f fileName = "aaa.txt"
	ex := f.Extension()
	if ex != "txt" {
		t.Fatal(results)
	}
}

func TestSearchDir(t *testing.T) {
	dir := "../pic"
	results, err := searchDir(dir)
	if err != nil {
		t.Fatal("failed test")
		t.Fatal(results)
	}

	dir = "aaa"
	results, err = searchDir(dir)
	if err != nil {
		t.Fatal("failed test")
		t.Fatal(results)
	}
}
