package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/options"
)

func main() {
	opts, err := options.Parse(os.Args)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(opts)
}
