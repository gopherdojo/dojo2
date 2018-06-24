package main

import (
	"testing"

	"./converter"
)

func TestGetFileNameWithoutExt_OnlyFileName(t *testing.T) {
	testdata := "file.png"
	actual := converter.GetFileNameWithoutExt(testdata)
	expected := "file"
	if len(actual) != len(expected) {
		t.Errorf("queue size is different. got %v but want %v", len(actual), len(expected))
	}
}

func TestGetFileNameWithoutExt_WithFolderName(t *testing.T) {
	testdata := "hoge/fuga/file.png"
	actual := converter.GetFileNameWithoutExt(testdata)
	expected := "file"
	if len(actual) != len(expected) {
		t.Errorf("queue size is different. got %v but want %v", len(actual), len(expected))
	}
}
