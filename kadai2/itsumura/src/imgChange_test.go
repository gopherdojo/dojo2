package main

import (
	"testing"
)

func TestExtension(t *testing.T) {
	var f fileName = "aaa.txt"
	ex := f.Extension()
	println(ex)
}

func TestSearchDir(t *testing.T) {
	dir := "../pic"
	results, err := searchDir(dir)
	if err != nil {
		t.Fatal("failed test")
		t.Fatal(results)
	}
}
