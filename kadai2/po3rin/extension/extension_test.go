package extension_test

import (
	"dojo2/kadai1/po3rin/extension"
	"testing"
)

func TestConvert(t *testing.T) {
	testarg := []extension.Arg{
		{
			From: "jpg",
			To:   "png",
			Path: "../images/gopher.jpg",
		}, {
			From: "jpg",
			To:   "gif",
			Path: "../images/gopher.jpg",
		}, {
			From: "png",
			To:   "jpg",
			Path: "../images/gopher.png",
		}, {
			From: "png",
			To:   "gif",
			Path: "../images/gopher.png",
		}, {
			From: "gif",
			To:   "jpg",
			Path: "../images/gopher.gif",
		}, {
			From: "gif",
			To:   "png",
			Path: "../images/gopher.gif",
		},
	}
	for _, v := range testarg {
		err := v.Convert()
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}
