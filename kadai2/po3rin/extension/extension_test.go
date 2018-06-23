package extension_test

import (
	"dojo2/kadai2/po3rin/extension"
	"strconv"
	"testing"
)

func TestArg_Convert(t *testing.T) {
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
			From: "pndg",
			To:   "jpg",
			Path: "../images/gopher.png",
		}, {
			From: "png",
			To:   "gif",
			Path: "../images/gopher.png",
		}, {
			From: "gidf",
			To:   "jpg",
			Path: "../images/gopher.gif",
		}, {
			From: "gif",
			To:   "png",
			Path: "../images/gopher.gif",
		},
	}
	for i, v := range testarg {
		t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
			helperConvert(v, t)
		})
	}
}

func helperConvert(v extension.Arg, t *testing.T) {
	t.Helper()
	if err := v.Convert(); err != nil {
		t.Errorf("failed test %#v", err)
	}
}
