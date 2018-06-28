package main

import (
	"testing"
)

func TestExtension(t *testing.T) {
	var f fileName = "aaa.txt"
	ex := f.Extension()
	if ex != "txt" {
		t.Fatal(ex)
	}
}

func TestSearchDir(t *testing.T) {
	dir := "../pic"
	results, err := searchDir(dir)
	if err != nil {
		t.Fatal("failed test")
		t.Fatal(results)
	}

	//エラー
	dir = "aaa"
	results, err = searchDir(dir)
	if err != nil {
		t.Fatal("failed test")
		t.Fatal(results)
	}

	//テーブル駆動型
	cases := []struct{
		dir string
	}{
		{dir: "../pic"},
		{dir : "aaa"}, //エラー
	}

	for _, c := range cases{
		results, err = searchDir(c.dir)
		if err != nil {
			t.Fatal("failed test")
			t.Fatal(results)
		}
	}
}
