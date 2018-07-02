package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo2/kadai3-2/Khigashiguchi/sget"
)

func main() {
	cli := sget.New()
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
