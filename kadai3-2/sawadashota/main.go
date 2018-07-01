package main

import (
	"io"
	"log"
	"os"

	"fmt"
	"time"

	"github.com/gopherdojo/dojo2/kadai3-2/sawadashota/download"
	"github.com/gopherdojo/dojo2/kadai3-2/sawadashota/option"
)

var writer io.Writer

func init() {
	writer = os.Stdout
}

func main() {
	opts, u, err := option.Parse(writer)

	if err != nil {
		log.Fatal(err)
	}

	d := download.New(u, opts)
	start := time.Now()
	d.Run()
	fmt.Fprintf(d.Writer(), "Duration %f seconds\n", time.Since(start).Seconds())
}
