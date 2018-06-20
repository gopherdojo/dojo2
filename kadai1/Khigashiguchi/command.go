package main

import (
	"flag"
	"log"
	"os"

	"github.com/gopherdojo/dojo2/kadai1/Khigashiguchi/convert"
	"github.com/gopherdojo/dojo2/kadai1/Khigashiguchi/ext"
)

var in string
var out string

func init() {
	flag.StringVar(&in, "in", "jpg", "extension name of target convert file")
	flag.StringVar(&out, "out", "png", "extension name of converted file")
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("You need to specify a target directory.")
	}
	dirpath := flag.Args()[0]
	if _, err := os.Stat(dirpath); err != nil {
		log.Fatal("Your specified directory does not exists.")
	}

	in = ext.Format(in)
	out = ext.Format(out)
	if ext.Validate(in) == false || ext.Validate(out) == false {
		log.Fatalf("Your specified extentision is not supported. You must specify jpg or png, given in %s out %s", in, out)
	}

	converter := convert.Converter{
		Dir:    dirpath,
		InExt:  in,
		OutExt: out,
	}
	converter.Exec()
}
