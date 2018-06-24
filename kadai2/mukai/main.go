package main

import (
	"os"
	"path/filepath"
	"./converter"
	"flag"
	"fmt"
)

//aiueo
func main() {
	s, _ := os.Getwd()
	abs, _ := filepath.Abs(s)
	i := flag.String("i", "", "input image ext")
	o := flag.String("o", "", "output image ext")
	flag.Parse()
	dir := flag.Arg(0)
	if len(*i) == 0 || len(*o) == 0 || len(dir) == 0 {
		fmt.Println(usage())
		return
	}
	if err := converter.Convert(filepath.Join(abs, dir), *i, *o); err != nil {
		println(err.Error())
	}

}

func usage() (string){
	return "usage: imgconv -i jpg -o png dir"
}
