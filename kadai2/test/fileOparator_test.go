package test

import (
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/tokunaga"
)

// 完全パスから拡張子を除いた文字列を返されることを確認
func TestFullBasename(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "/AA/BB/C.go", expected: "/AA/BB/C"},
		{input: "AA/BB/C.go", expected: "AA/BB/C"},
		{input: "/AA.go", expected: "/AA"},
		{input: "../C.go", expected: "../C"},
		{input: "./C.go", expected: "./C"},
		{input: "C.go", expected: "C"},
		{input: "/AA/BB/C", expected: "/AA/BB/C"},
		{input: "", expected: ""},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			if actual := tokunaga.FullBasename(c.input); actual != c.expected {
				t.Errorf("want FullBasename(%s) = %s, got %s", c.input, c.expected, actual)
			}
		})
	}
}
