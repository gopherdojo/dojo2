package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"./download"
	"./option"
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
