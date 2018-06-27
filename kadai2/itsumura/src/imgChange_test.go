package main

import (
	"testing"
)

func TestSearchDir(t *testing.T) {
	results, err := searchDir("aa")
	if err != nil {
		t.Fatal(results)
		t.Fatal("failed test")
	}
}
